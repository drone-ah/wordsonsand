---
categories: []
date: "2019-03-17T15:13:53Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:03"
  _publicize_job_id: "28738414385"
  timeline_notification: "1552835634"
parent_id: "0"
password: ""
status: publish
tags:
- java
title: List of Installed Programs on Windows from java
type: post
url: /2019/03/17/list-of-installed-programs-on-windows-from-java/
---

I was recently in need of a way to pick up the list of installed software on a
windows computer from Java. It was shrouded in a veil a mystery. There did not
seem to be a functional call that I could make, which makes sense since Java is
cross-platform and there is not universal way to pick up all the installed
packages on an OS.

In fact, picking up all the installed packages on Windows seems be a bit
cryptic. There are API calls you can make and I considered JNI. I suspect this
might be a superior solution, but I haven\'t tried it and I read that it may be
slow.

After much research, I came across
[ListPrograms.](https://github.com/mavenlin/ListPrograms) My initial thought was
to link to it using JNI. However, it seemed simple enough to warrant a rewrite
in Java if I could access the registry somehow.

This is where [JNA](https://github.com/java-native-access/jna) and the
[Advapi32Util](https://java-native-access.github.io/jna/4.2.0/com/sun/jna/platform/win32/Advapi32Util.html)
class came in handy.

It didn\'t take me long to put together a quick replica of behaviour. I skipped
out the part about user installed programs because it isn\'t relevant for me
(yet).

I also missed out issues which will be around when running it as part of a 32bit
VM within a 64bit OS.

You can find [JavaListApps at GitHub](https://github.com/drone-ah/JavaListApps)
