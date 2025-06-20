---
categories:
  - Java (EE)
date: "2011-04-25T15:58:00Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:42"
  _publicize_job_id: "5181432565"
  _wp_old_slug: "445"
  original_post_id: "445"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - java-ee
  - caching
  - garbage-collection
  - hibernate
  - java
  - java-memory-allocation
  - java-object-size
  - jboss
title: Java Object Size In Memory
url: /2011/04/25/java-object-size-in-memory/
---

Anyone who has worked with java in a high end application will be well aware of
the double edged sword that is java garbage collection. When it works - it is
awesome but when it doesn't - it is an absolute nightmare. We work on a
ticketing system where it is imperative that the system is as near real-time as
possible. The biggest issue that we have found is the running of memory in the
JVM which causes a stop the world garbage collection. This then results in
cluster failures since an individual node is inaccessible for long enough that
it is kicked out of the cluster.

There are various ways to combat this issue and the first instinct would be
suggest that there is a memory leak. After eliminating this as a possibility,
the next challenge was to identify where the memory was being taken up. This
took some time and effort and the hibernate second level cache was identified.
We were storing far too much in the second level cache.

This is another double edged sword. The hibernate second level cache is
absolutely imperative to a high performance system. It does however, come with a
price. The cache needs to be managed carefully to ensure that balance between
performance and memory requirements.

<!--more-->

To this end, it was important to be able to identify what was taking up all the
memory in the cache. Each object might only take a couple of hundred bytes, but
with our second level cache set to store hundreds of thousands of items, this
quickly takes up hundreds of megabytes. With the metadata of the cache, this
could easily hike it up near a gigabyte of memory usage. This gets substantially
worse with cache evictions and the adding of new items into the cache.

The correct way to resolve this is to identify specific object types that
"overload" the cache. i.e. items that have an large number of instances stored
in the cache. Identifying classes that store a large number of items is easy
enough - we just traverse the cache and count up the number of items. However,
there might be a class that stores a smaller number of items but take a sizeable
amount of memory. For this reason, it is important to understand the object
sizes in memory as well.

If you have ever tried to find a way to identify object sizes, you will know
that this is no easy task. You can calculate to some degree of accuracy the size
of an object based on the data it stores but this is a manual process.

The only real way to get this information is to use a java agent and use that to
calculate a more accurate memory usage. For this purpose, we used the
[classmexer agent](http://www.javamex.com/classmexer/ "ClassMexer Java Profiling Agent")
which requires a simple installation step of adding the following parameter to
java `-javaagent:classmexer.jar`. You can then figure out the memory utilisation
of an object by calling

```java
MemoryUtil.deepMemoryUsageOf(objectInstance)
```

You can also pass in a collection of objects:

```java
MemoryUtil.deepMemoryUsageOfAll(objectInstanceCollection)
```

This was the simple part.

Traversing the node structure of jboss cache and collating a collection
statistics with regards to the number of each type of object and its memory
utilisation was a little more interesting.

I will cover this separately
