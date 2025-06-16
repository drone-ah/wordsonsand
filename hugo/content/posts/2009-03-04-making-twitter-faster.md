---
categories:
- Systems (Administration)
date: "2009-03-04T17:36:35Z"
excerpt: An idea for making twitter perform better, be more stable and scalable.
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:37"
  _publicize_job_id: "5181540479"
  _wp_old_slug: "185"
  geo_public: "0"
  oc_commit_id: http://drone-ah.com/2009/03/04/making-twitter-faster/1236792306
  original_post_id: "185"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- API
- caching
- Challenge
- Hadoop
- HBase
- Ideas
- Linux
- megabus
- relational database
- search mechanism
- search time
- Twitter
title: Making Twitter Faster
url: /2009/03/04/making-twitter-faster/
---

From my perspective, Twitter has a really really interesting technical problem
to solve. How to store and retrieve a large amount of data really really
quickly.

I am making some assumptions based on how I see twitter working. I have little
information about how it is architected apart from some posts that suggests that
it is running ruby on rails with MySQL?

Twitter is in the rare category where there is a very large number of data being
added. There should be no updates (except to user information but there should
be relatively very small amount of that). There is no need for transactionality.
If I guess right, it should be a large amount of inserts and selects.

While a relational database is probably the only viable choice for the time
being, I think that twitter can scale and perform better if all the extra bits
of a relational database system was removed.

I love challenges like this. Technical ones are easier ;-)

If I didn't have a lifetime job, I would prototype this in a bit more depth.
[Garry](http://garry.blog.kraya.co.uk "Garry's Blog") pointed me in the
direction of [Hadoop](//hadoop.apache.org/ "Hadoop"). Having had a quick look at
it, it can take care of the infrastructure, clustering and massive horizontal
scaling requirements.

<!--more-->

Now for the data layer on top. How to store and retrieve the data.
[HBase](http://hadoop.apache.org/hbase/ "HBase - a scalable distributed database")
is probably a good option but doing it manually should be fairly straightforward
too.

From my limited understanding of twitter, there are two key pieces of
functionality, the timelines and search.

The timelines can be solved by storing each tweet as a file within a directory
structure. My tweets would go into

/w/o/r/d/s/o/n/s/a/n/d/<tweet-filename>

The filename would be <username>-<timestamp>

For the public timeline, you just have a similar folder structure, but with the
timestamp, for example, the timestamp 1236158897 would go into the following
structure as a symlink

/1/2/3/6/1/5/8/8/9/7/<username>

For search, pick up each word in the tweet and pop the tweet as a symlink into
that folder. You could have a folder per word or follow the structure above.

/t/w/i/t/t/e/r/<username>-<timestamp> OR

twitter/<username>-<timestamp>

You would then have an application running on top with a distributed cache with
an API to ease access into the data easier than direct file access. Running on
Linux, the kernel will take care of the large part of the automatic caching and
buffering as long as there is enough RAM on the box.

This can in theory be done without Hadoop in between and separating the
directory structures across multiple servers but that can have complications of
its own, especially with adding and removing boxes for scalability.

You are also likely to run into issues with the number of files /
sub-directories limits but they can be solved by 'archiving' - multiple options
for that too...

Thinking about this problem brought me back to the good old days of working on
the search mechanism within megabus.com. We needed the site to deal with a large
number of searches on limited hardware when the project was still classified as
a pilot.

With some hard work and experimentation, we were able to reduce the search time
to a tenth of the original time.

I'll admit that I don't know the details or the intricacies of the requirements
that twitter has. I have probably over-simplified the problem but it was still
fun to think about. If you can think of problems with this - let me know; I
wanna turn them into opportunities ;-)
