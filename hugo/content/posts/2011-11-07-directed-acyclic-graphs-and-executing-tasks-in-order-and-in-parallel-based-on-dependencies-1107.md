---
categories:
- Software Development
date: "2011-11-07T23:44:36Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:52"
  _publicize_pending: "1"
  _wp_old_slug: "723"
  oc_commit_id: http://drone-ah.com/2011/11/07/directed-acyclic-graphs-and-executing-tasks-in-order-and-in-parallel-based-on-dependencies-1107/1320709479
  original_post_id: "723"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- DAG
- Dependency
- Directed Acyclic Graph
- Graphs
- java
- Pointer
- Task
title: Directed Acyclic Graphs and Executing Tasks in Order (and in Parallel) Based
  on Dependencies [1107]
url: /2011/11/07/directed-acyclic-graphs-and-executing-tasks-in-order-and-in-parallel-based-on-dependencies-1107/
---

A little while ago, there was a requirement to write a tool that could take a
number of tasks each with a set of dependencies and execute them in parallel
while taking the dependencies into account.

The tasks themselves were meant for data migration but that is not particularly
relevant. We were writing a number of tasks which all had a set of dependencies
(some of the tasks did not have any dependencies or the process could of course
never start).

It was assumed that there were no cyclic dependencies (which would be error in
this particular case anyway)

Bearing in mind that this was a quick and dirty tool for use three times, some
of the bits in here could do with tidying up.

Each task was defined to implement the following interface

```java
    public interface Task extends Runnable {

        public String getName();

        public Set getDependencies();

    }
```

<!--more-->

It should all be self explanatory. Extending the Runnable interface ensure that
we can pass it into threads and other relevant bits of code. The getDependencies
is expected to return the name of the tasks that it depends on.

The basic task runner which I describe below does not check if the task
described in any list of dependencies actually exist. If an non-existing
dependency is defined, it will likely just throw a Null Pointer Exception. I
wrote this a long time ago, so don't actually remember.

The BasicTaskRunner which we used to run the tasks implemented the TaskRunner
Interface

```java
    public interface TaskRunner {

        public boolean addTask(Task task);

        public boolean prepare();

        public boolean start();

        public void waitToComplete();
    }
```

The `addTask` method simply added it to a map from String -> Task and threw an
exception in the event of a duplicate task being added in.

```java
       @Override
        public synchronized boolean addTask(Task task) {
            LOG.info("Adding task: " + task.getName());

            if (tasks.put(task.getName(), task) != null) {
                throw new RuntimeException("Task with same name already exists: " + task.getName());
            }

            return true;
        }
```

the prepare method just calls a method to buildGraph. This uses the jGrapht
library to build a
[Directed Acyclic Graph](http://en.wikipedia.org/wiki/Directed_acyclic_graph "Directed Acyclic Graph")

```java
       private boolean buildGraph() {

            LOG.info("Building DAG of tasks");
            graph = new SimpleDirectedGraph(DefaultEdge.class);

            LOG.info("Adding tasks");

            for (Task task: tasks.values()) {
                graph.addVertex(task);
            }

            LOG.info("Adding Relationships");

            for (Task task: tasks.values()) {

                if (task.getDependencies() != null) {
                    for (String depend: task.getDependencies()) {

                        Task dependOnTask = tasks.get(depend);

                        LOG.info("Adding relationship between " + task.getName() + " and " + dependOnTask.getName());
                        graph.addEdge(dependOnTask, task);

                    }
                }
            }

            return true;
        }
```

So we create a simple directed graph, loop through the tasks, then each of its
dependencies to create an edge, which we then add to the graph. Simple stuff
really.

the start method, which actually executes the task is as follows:

```java
       public boolean start() {

            int cpus = Runtime.getRuntime().availableProcessors();

            executor = new ThreadExecutor(cpus, 60, new LinkedBlockingQueue());

            numTasks = graph.vertexSet().size();
            LOG.info("Starting... Num Tasks: " + numTasks);
            startTime = System.currentTimeMillis();

            scheduleTasks();

            return true;

        }
```

As a basic algorithm, we pick up the number of available processors and use that
many threads. scheduleTasks is a pseudo-recursive function whose role is to add
the currently executable list of tasks into the executor to execute.

```
       private void scheduleTasks() {
            if (graph.vertexSet().size() == 0) {
                executor.shutdown();
            }

            synchronized (graph ) {
                Iterator iter = new TopologicalOrderIterator(graph);
                Set executing = new HashSet();

                while(iter.hasNext()) {

                    Task task = iter.next();
                    //System.out.println(task.getName());
                    if (graph.incomingEdgesOf(task).size() == 0 && !executing.contains(task)) {
                        executor.execute(task);
                        executing.add(task);
                    }

                }
            }

        }
```

If there are no tasks left to execute, we shut the executor down. All being
well, we add every single task in the graph that has no dependencies to be
executed. The threadpool ensures that any tasks that cannot currently be
executed are queued.

We use a custom version of the threadpool as follows:

```
       private class ThreadExecutor extends ThreadPoolExecutor {

            public ThreadExecutor(int corePoolSize, long keepAliveSeconds, BlockingQueue workQueue) {
                super(corePoolSize, corePoolSize, keepAliveSeconds, TimeUnit.SECONDS, workQueue);
            }

            @Override
            protected void afterExecute(Runnable runTask, Throwable e) {
                super.afterExecute(runTask, e);

                if (e == null) {
                    completed((Task) runTask);
                } else {
                    failed((Task) runTask, e);
                }
            }

        }
```

The main purpose of this is to use the completed and failed callbacks to ensure
that on complete, dependent tasks can be executed. On fail, we ensure that
dependent tasks are not executed. The code currently does not allow for tasks
that are left behind and will hang indefinitely after executing all tasks it
can.

```java
       public void completed(Task t) {
            LOG.info("Completed Task: " + t.getName());

            synchronized (graph) {
                graph.removeVertex(t);
            }

            long timeTaken = (System.currentTimeMillis() - startTime);
            int tasksComplete = numTasks - graph.vertexSet().size();

            long timePerTask = timeTaken/tasksComplete;

            long totalTime = timePerTask * numTasks;
            long timeToComplete = timePerTask * graph.vertexSet().size();

            LOG.info(" ## Tasks left: " + graph.vertexSet().size()
                   + " ## Elapsed: " + timeTaken/1000
                   + " ## Est. Total " + totalTime/1000
                   + " ## E.T.A : " + timeToComplete/1000);

            scheduleTasks();
        }

        public void failed(Task t, Throwable e) {
            LOG.fatal("Failed Task: " + t.getName(), e);
            scheduleTasks();
        }
```

On completion of a task, we simply remove the task from the graph. The frees up
all its dependencies to be executed. We add these tasks into the list by calling
scheduleTasks again. There is nothing more for us to do when a task fails except
to schedule any other tasks that can be executed. In theory, this call is
redundant since any tasks that could be executed before the failure are already
in the queue. Any tasks that can be completed on the completion of another item
will be initiated on the completion of that task.

I hope the above makes sense and has been helpful. The code for the full class
including further logging statements follows. Please bear in mind that this was
hacked together over a couple of hours for something that was to be executed a
grand total of three times.

```java
    public class BasicTaskRunner implements TaskRunner {

        private static final Logger LOG = Logger.getLogger(BasicTaskRunner.class);

        private Map tasks = new HashMap();

        private DirectedGraph graph;

        private ThreadExecutor executor;

        @Override
        public synchronized boolean addTask(Task task) {
            LOG.info("Adding task: " + task.getName());

            if (tasks.put(task.getName(), task) != null) {
                throw new RuntimeException("Task with same name already exists: " + task.getName());

            }

            return true;
        }

        @Override
        public boolean prepare() {

            LOG.info("Preparing task runner. Num Tasks: " + tasks.size());

            buildGraph();

            return false;
        }

        private boolean buildGraph() {

            LOG.info("Building DAG of tasks");
            graph = new SimpleDirectedGraph(DefaultEdge.class);

            LOG.info("Adding tasks");

            for (Task task: tasks.values()) {
                graph.addVertex(task);
            }

            LOG.info("Adding Relationships");

            for (Task task: tasks.values()) {

                if (task.getDependencies() != null) {
                    for (String depend: task.getDependencies()) {

                        Task dependOnTask = tasks.get(depend);

                        LOG.info("Adding relationship between " + task.getName() + " and " + dependOnTask.getName());
                        graph.addEdge(dependOnTask, task);

                    }
                }
            }

            return true;
        }

        public void waitToComplete() {
            try {
                executor.awaitTermination(3, TimeUnit.DAYS);
            } catch (InterruptedException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }
        }

        private long startTime;
        private int numTasks;

        @Override
        public boolean start() {

            int cpus = Runtime.getRuntime().availableProcessors();

            executor = new ThreadExecutor(cpus, 60, new LinkedBlockingQueue());

            numTasks = graph.vertexSet().size();
            LOG.info("Starting... Num Tasks: " + numTasks);
            startTime = System.currentTimeMillis();

            scheduleTasks();

            return true;

        }

        private void scheduleTasks() {
            if (graph.vertexSet().size() == 0) {
                executor.shutdown();
            }

            synchronized (graph ) {
                Iterator iter = new TopologicalOrderIterator(graph);
                Set executing = new HashSet();

                while(iter.hasNext()) {

                    Task task = iter.next();
                    //System.out.println(task.getName());
                    if (graph.incomingEdgesOf(task).size() == 0 && !executing.contains(task)) {
                        executor.execute(task);
                        executing.add(task);
                    }

                }
            }

        }

        public void completed(Task t) {
            LOG.info("Completed Task: " + t.getName());

            synchronized (graph) {
                graph.removeVertex(t);
            }

            long timeTaken = (System.currentTimeMillis() - startTime);
            int tasksComplete = numTasks - graph.vertexSet().size();

            long timePerTask = timeTaken/tasksComplete;

            long totalTime = timePerTask * numTasks;
            long timeToComplete = timePerTask * graph.vertexSet().size();

            LOG.info(" ## Tasks left: " + graph.vertexSet().size()
                   + " ## Elapsed: " + timeTaken/1000
                   + " ## Est. Total " + totalTime/1000
                   + " ## E.T.A : " + timeToComplete/1000);

            scheduleTasks();
        }

        public void failed(Task t, Throwable e) {
            LOG.fatal("Failed Task: " + t.getName(), e);
            scheduleTasks();
        }

        private class ThreadExecutor extends ThreadPoolExecutor {

            public ThreadExecutor(int corePoolSize, long keepAliveSeconds, BlockingQueue workQueue) {
                super(corePoolSize, corePoolSize, keepAliveSeconds, TimeUnit.SECONDS, workQueue);
            }

            @Override
            protected void beforeExecute(Thread thread, Runnable runTask) {
                super.beforeExecute(thread, runTask);

                Task task = (Task) runTask;

                LOG.info("Starting task: " + task.getName());
            }

            @Override
            protected void afterExecute(Runnable runTask, Throwable e) {
                super.afterExecute(runTask, e);

                if (e == null) {
                    completed((Task) runTask);
                } else {
                    failed((Task) runTask, e);
                }
            }

        }

    }
```
