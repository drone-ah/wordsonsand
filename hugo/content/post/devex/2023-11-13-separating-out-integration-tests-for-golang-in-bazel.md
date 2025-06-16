---
categories:
- devex
date: "2023-11-13T13:47:53Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:05"
  _last_editor_used_jetpack: block-editor
  _publicize_job_id: "89325811494"
  timeline_notification: "1699883275"
  wordads_ufa: s:wpcom-ufa-v4:1699886196
parent_id: "0"
password: ""
published: true
status: publish
tags:
- bazel
- golang
- testing
title: Separating out integration tests for golang in Bazel
type: post
url: /2023/11/13/separating-out-integration-tests-for-golang-in-bazel/
---

There are many kinds of automated tests and two main kinds are integration tests
and unit tests.

Unit tests are designed to run as fast as possible, so any slower processes like
databases are mocked out. While super helpful and powerful in terms of providing
confidence in the software, it should be only one part of the testing strategy.

Integration tests, as is implied runs tests of the different part of the
software integrated together. Technically speaking, you can still mock out the
database and other slower layers to keep it running quickly. However, there is
value in including a database or other slower services in the process to test as
them in an automated fashion.

What this does mean though, is that you want to be able to run only the unit
tests or run the integration tests as well. You might also want to have smoke
tests, which are run on your live production environment.

<!--more-->

# How {#how}

You could define a separate target in your `BUILD` file with the unit tests and
let `gazelle` automatically build your default test target with all the tests. I
found this frustrating to use as I had to keep tweaking the dependencies
manually whenever anything changed (which happened often)

## Tagging {#tagging}

The easiest way to achieve this for golang and bazel is to tag your source code
files. You can do this by adding the following to the top of your integration
test files

`something_integration_test.go`

```go
//go:build integration_test

package somepackage
```

You can pick any tag name you want instead of `integration_test` like
`integration`, `smoke_test` etc.

## IDE Support {#ide-support}

You will likely need to add this source file into the IDEs build constraints to
get the IDE to treat it as a source file. In IntelliJ (IDEA/Goland), you will be
warned of this

![idea screenshot, main_test.go is ignored by build tool](/assets/2023/11/image.png)

If you click `Edit settings`, you can add the tag in

![set custom tags in build tags](/assets/2023/11/image-1.png)

When running `gazelle`, you want to include the files with these tags

## Gazelle {#gazelle}

```bash
bazel run //:gazelle -- -build_tags=integration_test
```

If you have multiple tags, you can separate them with commas. This command will
generate a test target with all of the source files and its dependencies

## Bazel Integration on test {#bazel-integration-on-test}

To run only the unit tests, you test as normal:

```bash
bazel test ... # or the specific target, and it'll run only the unit tests
```

To run the integration tests as well, include that tag

```bash
bazel test ... --define  gotags=integration_test # Will run unit & integration tests
```

## Run only Unit Tests {#run-only-unit-tests}

This setup will currently not allow you to run ONLY the integration tests. To be
able to do that you\'ll need to add a `unit_test` tag to the unit test files so
that you can exclude them.

```
something_test.go
```

```go
//go:build integration_test

package somepackage
```

You can then run only the unit tests with

```bash
bazel test … --define gotags=unit_test # Will run unit & integration tests
```

Only the integration tests

```bash
bazel test … --define gotags=integration_test # Will run unit & integration tests
```

Or both:

```bash
bazel test … --define gotags=unit_test,integration_test # Will run unit & integration tests
```

## Simpler gazelle command {#simpler-gazelle-command}

You can enable the tags by default in the `BUILD` file so that you don\'t have
to pass the tags into gazelle each time.

`BUILD`

```starlark
# gazelle:build_tags unit_test,integration_test
```

You can then just run `bazel run //:gazelle` which will run with these tags
enabled.

# Sample Source {#sample-source}

You can find sample source code demonstrating this in
[my github repo](https://github.com/drone-ah/wordsonsand), under
[post/2023/11/separatetests](https://github.com/drone-ah/wordsonsand/tree/main/post/2023/11/separatetests)
