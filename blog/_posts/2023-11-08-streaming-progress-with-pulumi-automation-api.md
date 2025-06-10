---
layout: post
title: Streaming Progress With Pulumi Automation API
date: 2023-11-08 12:38:32.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories: []
tags:
  - golang
  - pulumi
  - pulumi automation api
meta:
  _last_editor_used_jetpack: block-editor
  wordads_ufa: s:wpcom-ufa-v4:1699447355
  _publicize_job_id: "89181397227"
  timeline_notification: "1699447113"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:05"
permalink: "/2023/11/08/streaming-progress-with-pulumi-automation-api/"
---

When using the pulumi automation API, you lose some of the niceties of the
pulumi CLI, like having to set up command line args processing and the output is
not as friendly or pretty as before. It also doesn\'t stream the output - though
this one is easier to fix.

This is lifted straight out of
[their golang example code](https://github.com/pulumi/automation-api-examples/blob/3114b754ea84ebd0cc1e1b67f128df75795bd4c3/go/local_program/automation/main.go#L74C2-L82C3),
so if you\'re working in another language - you should be able to find the
relevant code in the same repo

```go
   // wire up our update to stream progress to stdout
    stdoutStreamer := optup.ProgressStreams(os.Stdout)

    // run the update to deploy our fargate web service
    res, err := stack.Up(ctx, stdoutStreamer)
    if err != nil {
        fmt.Printf("Failed to update stack: %v\n\n", err)
    }
```
