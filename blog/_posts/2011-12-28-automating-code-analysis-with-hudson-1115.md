---
layout: post
title: Automating Code Analysis with Hudson [1115]
date: 2011-12-28 15:34:52.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software Development
tags:
  - Continuous integration
  - FindBugs
  - Hudson
  - PMD
  - Static code analysis
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2011/12/28/automating-code-analysis-with-hudson-1115/1325086495
  restapi_import_id: 591d994f7aad5
  original_post_id: "759"
  _wp_old_slug: "759"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:53"
permalink: "/2011/12/28/automating-code-analysis-with-hudson-1115/"
---

As part of
[setting up continuous integration and automated builds and source analysis](http://drone-ah.com/2011/12/28/hudson-jenkins-and-continuous-integration-1114/ "Hudson / Jenkins and Continuous Integration [1114]"),
the next step is to integrate in the source analysis parts.

To this end, I installed the following plugins:

Task Scanner\
FindBugs\
CheckStyle\
PMD

After restarting Hudson, there are a few additional configuration bits to
complete.

I added an additional build step and set the goal as

```
checkstyle:checkstyle findbugs:findbugs pmd:pmd
```

I then enabled the four plugins, saved, ran a build and et voila\... Just like
magic


