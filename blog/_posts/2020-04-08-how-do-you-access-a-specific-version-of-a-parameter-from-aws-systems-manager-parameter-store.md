---
layout: post
title:
  How do you access a specific version of a parameter from AWS Systems Manager
  Parameter Store
date: 2020-04-08 15:53:50.000000000 +01:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software Development
tags:
  - aws
  - node
  - ssm
meta:
  _publicize_job_id: "42784269381"
  timeline_notification: "1586361231"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:04"
permalink: "/2020/04/08/how-do-you-access-a-specific-version-of-a-parameter-from-aws-systems-manager-parameter-store/"
---

This was a bit tricky to find and doesn\'t seem to be well documented.

```wp-block-code
    // version number below can be a number, i.e. 3 or a label
    let params = {
      Name: "/path/to/parameter:<version-number>",
      WithDecryption: false
    };

    ssm.getParameter(params, function (err, data) {
      if (err) console.error(err, err.stack);
      else console.log(data)
    });
```
