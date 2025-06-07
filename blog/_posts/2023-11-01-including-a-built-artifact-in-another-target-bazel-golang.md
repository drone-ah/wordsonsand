---
layout: post
title: Including a built artifact in another target (Bazel, golang)
date: 2023-11-01 19:42:30.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories: []
tags:
  - aws-lambda
  - bazel
  - golang
  - pulumi
meta:
  _last_editor_used_jetpack: block-editor
  wpcom_is_first_post: "1"
  _publicize_job_id: "88995902253"
  timeline_notification: "1698867751"
  wordads_ufa: s:wpcom-ufa-v4:1698868013
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:04"
permalink: "/2023/11/01/including-a-built-artifact-in-another-target-bazel-golang/"
---

We use pulumi to do IaC and we use a monorepo with Bazel as the build tool. We
have out modules set out as following

One of the requirements we have is to build a lambda module and then deploy it.
The lambda module is a target being built by Bazel (golang, but shouldn\'t
matter):

```go
go_binary(
    name = "lambda_module",
    visibility = ["//visibility:public"],
)
```

<!-- more -->

We then have the iac module, which should get the built version of the above
module, so that it can then upload it into lambda

```go
go_binary(
    name = "iac",
    args = [
        "-lambda_module",
        "$(location //products/productA/module/lambda_module)",
    ],
    data = ["//products/productA/module/lambda_module"],
    visibility = ["//visibility:public"],
)
```

There are two key parameters here to note:

- `args`: We generate the path to the target module using
  ` //products/productA/module/lambda_module)`
- `data`: We use the data tag to ensure that the built output is included when
  building/running this target

We then need to use runfiles support within golang to be ablet to identify the
correct location for the built binary. The reason this part is complex is to be
able to support multiple operating systems. I should caveat that I have only got
this working on Linux, but Mac/Win shouldn\'t be too different.

```go
package main

import (
    _ "embed"
    "flag"
    "fmt"
    "github.com/bazelbuild/rules_go/go/runfiles"
    "path/filepath"
)

func main() {

    var webhookAuth = flag.String("webhook_auth", "", "bin for webhook_auth")
    flag.Parse()
    fmt.Printf("param : %s \n", *webhookAuth)

    path, err := runfiles.Rlocation(fmt.Sprintf("workspace_name/%s", *webhookAuth))
    fmt.Printf("rLoc path: %s, err: %v \n", path, err)

    symlinks, err := filepath.EvalSymlinks(path)
    fmt.Printf("evaluated path: %s, err: %v \n", symlinks, err)

}
```

We use the flag module to retrieve the path passed in as a runtime parameter

We use `runfiles.Rlocation` to pick up the \"real\" path to the file, prepending
the workspace name to the start. You can
[define the workspace name](https://bazel.build/rules/lib/globals/workspace#workspace)
in the root WORKSPACE file with `workspace(name = "workspace_name")`

Finally, resolve the Symlink to get the actual file path

## References {#references}

There are similar mechanisms to find the rLocation in other languages, a couple
of which are described in
[its design document](https://docs.google.com/document/d/e/2PACX-1vSDIrFnFvEYhKsCMdGdD40wZRBX3m3aZ5HhVj4CtHPmiXKDCxioTUbYsDydjKtFDAzER5eg7OjJWs3V/pub)

There is some documentation in `rules_go` around
[accessing `go_binary` from `go_test`](https://github.com/bazelbuild/rules_go#how-do-i-access-go_binary-executables-from-go_test)
which I referenced and updated to get the above example

I found the above link from
[a stackoverflow post about feeding bazel output to another bazel rule](https://stackoverflow.com/questions/70193581/feed-bazel-output-to-another-bazel-rule)
