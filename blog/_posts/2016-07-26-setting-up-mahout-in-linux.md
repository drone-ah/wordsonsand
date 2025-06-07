---
layout: post
title: Setting up Mahout in Linux
date: 2016-07-26 10:06:41.000000000 +01:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Artificial Intelligence
tags:
  - mahout
meta:
  _rest_api_published: "1"
  _rest_api_client_id: "-1"
  _publicize_job_id: "5186353710"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:00"
permalink: "/2016/07/26/setting-up-mahout-in-linux/"
---

A few simple steps to get Mahout running in Linux. This is mostly about the bash
script to get it to run easily

You'll need to install Java first, then download and unpack the mahout
distribution.

I then placed it in `/usr/local/mahout`

To be able to run Mahout from the path, the following bash script was placed in
`/usr/local/bin`

Update the paths as relevant

```bash
 #!/bin/bash
 export MAHOUT_JAVA_HOME=/usr/lib/jvm/java-8-oracle/jre/
export MAHOUT_HOME=/usr/local/mahout
export MAHOUT_HEAPSIZE=4000
export MAHOUT_LOCAL=y
```
