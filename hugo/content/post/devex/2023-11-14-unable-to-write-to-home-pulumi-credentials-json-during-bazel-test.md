---
categories:
- devex
date: "2023-11-14T13:46:46Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:05"
  _last_editor_used_jetpack: block-editor
  _publicize_job_id: "89357332809"
  timeline_notification: "1699969607"
  wordads_ufa: s:wpcom-ufa-v4:1699969789
parent_id: "0"
password: ""
published: true
status: publish
tags:
- bazel
- testing
title: Unable to write to $HOME/.pulumi/credentials.json during bazel test
type: post
url: /2023/11/14/unable-to-write-to-home-pulumi-credentials-json-during-bazel-test/
---

# The Problem {#the-problem}

You can add it to your \~/.bazelrc (it needs the path to be absolute)

From our integration tests, we run `pulumi stack output` (or in some cases
`pulumi up`) through the automation API before we run the tests so that we can

- Confirm that the stack is up
- Get the relevant parameters (actual names of lambdas / dynamo db tables etc.)

However, since we use bazel for our tests, we ran into a small problem in that
Bazel (rightly) prevents the tests from writing to anything outside the sandbox.
This restrictions results in this error

```
error: open /home/<username>/.pulumi/credentials.json: read-only file system
```

<!--more-->

# The Solution {#the-solution}

The easiest way to solve this is to ask `bazel` to allow writing to this
location, which you can do with:

```bash
bazel test ... --sandbox_writable_path=$HOME/.pulumi
```

`bazel` needs to the path to be absolute, so `~/.pulumi` won\'t work.

# Automation {#automation}

It is annoying to add this flag into all the tests, but there is an way to
automatically add it to all tests. You can add it to `.bazelrc`. Due to the
aforementioned requirement for the path to be absolute, it is not possible to
put it into the git repo root. However, you can put it into your home directory
rool `.bazelrc`

`$HOME/.bazelrc`

```
test --sandbox_writable_path=/home/<your-username>/.pulumi
```
