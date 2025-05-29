---
layout: post
title: Eclipse TPTP on Ubuntu (64bit)
date: 2008-12-28 18:15:45.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software
tags:
  - 64bit
  - Dev Env
  - Eclipse
  - Eclipse TPTP
  - java
  - libstdc++.so.5
  - Ubuntu
  - VServer
meta:
  _edit_last: "48492462"
  restapi_import_id: 591d994f7aad5
  original_post_id: "108"
  _wp_old_slug: "108"
  _publicize_job_id: "5181448836"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:35"
permalink: "/2008/12/28/eclipse-tptp-on-ubuntu-64bit/"
---

I run ubuntu 64 bit (technically, I run an ubuntu 64bit vserver which I access
from ubuntu 32 bit but thats not really relevant).

In the open source world, I expect that all things which are accessible as 32bit
are also accessible and 64bit and ubuntu makes it automagic enough that
everything just works. Yes, I run into problems with closed source software like
Flash Player (recently resolved with flash player 10) and the Java Plugin but
that is another story. I use Eclipse and wanted to do some performance analysis
and benchmarking to find a bottleneck and installed the TPTP plugin; and ran
into a problem. It just didn\'t work.

To resolve it, I turned to google\... In this instance, it turned out to be a
distraction and a red-herring. It lead me in the direction of installing
libstdc++2.10-glibc2.2_2.95.4-27_i386.deb which was difficult at best since
there was only a 32bit version of the package and that wasn\'t even in the
standard repository.

In the end, digging deeper, I found that it simply missed the following shared
object libstdc++.so.5.

All I had to do was install libstdc++5:

> sudo aptitude install libstdc++5

and it worked\... :-D

Now, I think that ACServer which Eclipse uses to do TPTP should not link to an
outdated library but that is another issue\...
