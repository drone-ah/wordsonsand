---
layout: post
title:
  Bazel + Pulumi Automation API. Deploying an inline stack along with
  Pulumi.[stack].yaml
date: 2023-11-07 15:57:52.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories: []
tags:
  - bazel
  - golang
  - pulumi
  - pulumi-automation-api
meta:
  _last_editor_used_jetpack: block-editor
  wordads_ufa: s:wpcom-ufa-v4:1699373839
  _publicize_job_id: "89156493747"
  timeline_notification: "1699372673"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:05"
permalink: "/2023/11/07/bazel-pulumi-automation-api-deploying-an-inline-stack-along-with-pulumi-yaml/"
---

I believe in CI/CD/CD as in Continuous, integration, delivery and deployment. As
part of this, I am setting up a workflow where on merge to develop (or
main/trunk), the deployment is triggered automatically. Pulumi deploys the
current state of code and infrastructure through GitHub actions and
[OpenID Connect(OIDC)](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services)
as part of the GitHub Action.

I used to configure Pulumi to be triggered directly from the build process, but
bazel (as far I know), does not support Pulumi. When I used pants, there was a
custom module, developed by one of the community members which did support
pulumi (You might have to ask in the slack channel if you\'re interested), but
they stopped maintaining it as they moved to the Pulumi Automation API.

I am using Automation API from the start, and configuring a \"deployer\" per
product/project within the monorepo. The intention is for the deployer to be as
smart as possible - eventually `up`-ing only the stacks that have changes since
the last time- but that\'s a way down the line.

Another benefit from the Automation API is to pick up the stack outputs
automatically when running integration/e2e tests, making the test configuration
smoother.

The first step is to be able to define a stack within a product, hook it into
the main iac executable and have it working. My directory structure is roughly
as below: While I am using golang as my language of choice, it\'s probably not
hugely different in other languages.

- products
  - productA
    - auth
      - iac_auth
        - BUILD
        - deploy.go
        - Pulumi.dev.yaml
        - Pulumi.yaml
      - OtherAuthModules
    - iac
      - BUILD
      - main.go

# `iac_auth/BUILD` {#iac_authbuild}

```wp-block-syntaxhighlighter-code
# load etc.
go_library(
    name = "iac_auth",
    srcs = ["deploy.go"],
    data = glob([
        "Pulumi.*.yaml",
    ]) + [
        "Pulumi.yaml",
    ],
    visibility = ["//visibility:public"],
    deps = [
       # dependencies automated with gazelle
    ],
)
```

# `iac_auth/deploy.go` {#iac_authdeploy.go}

```wp-block-syntaxhighlighter-code
package iac_sync

import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const moduleName = "auth"

func DeployGithubLambda(ctx *pulumi.Context) error {
    fmt.Println("Deploy")

    conf := config.Require(ctx, "key")
    fmt.Printf("conf: %s\n", conf) // Will successfully pick up the value if key is in Pulumi.<stack>.yaml

    // Actual deployment code here

    return nil
}
```

# `iac/BUILD` {#iacbuild}

```wp-block-syntaxhighlighter-code
# other config including standard config of :iac_lib by gazelle

go_binary(
    name = "iac",
    args = [
        "-iac_auth",
        "$(location //products/productA/auth/iac_auth)",
    ],
    data = [
        "//products/gitolink/productA/iac_auth",
    ],
    embed = [":iac_lib"],
    visibility = ["//visibility:public"],
)
```

Worth noting the `args` bit, which is what we use to identify the path where the
`Pulumi.yaml` and `Pulumi.<stack>.yaml` files are:

# `iac/main.go` {#iacmain.go}

```wp-block-syntaxhighlighter-code
package main

import (
    "context"
    "flag"
    "fmt"
    iacAuth "github.com/drone-ah/monorepo/products/gitolink/auth/iac_auth"
    "github.com/pulumi/pulumi/sdk/v3/go/auto"
    "path/filepath"
)

func main() {

    // Pick up the path from args
    var iacAuthPath = flag.String("iac_auth", "", "bin for iac_auth")
    flag.Parse()
    fmt.Printf("param iac: %s \n", *iacAuthPath)

    ctx := context.Background()

    projectName := "gitolink"
    stackName := "dev"

    stack, err := auto.NewStackInlineSource(ctx,
        stackName,
        projectName,
        iacAuth.DeployGithubLambda,
        // Define workdir for locatin of the Pulumi yaml files
        auto.WorkDir(filepath.Dir(*iacAuthPath)))

    preview, err := stack.Preview(ctx)

    fmt.Println(err)
    fmt.Println(preview)
}
```

You might also have to use `RLocation` (see my previous post about
[including an artifact in another target](https://drone-ah.com/2023/11/01/including-a-built-artifact-in-another-target-bazel-golang/)
for an example of this), though when I tried it, it was missing one of the yaml
files and I didn\'t investigate further.
