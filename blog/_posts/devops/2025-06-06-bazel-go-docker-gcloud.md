---
layout: post
title: Build and push container image using bazel
date: 2025-06-06 13:06:14
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - devops
tags:
  - google cloud
  - webapp
  - bazel
  - golang
  - pulumi
  - oci
  - docker
  - iac
meta:
---

I am building a small [golang](/tags/golang) [webapp](/tags/webapp), and I want
to push a [container](/tags/oci) up for it, which can eventually be used in
Google Cloud Run, or elsewhere.

In this post, I want to describe how I got it to push images build locally up to
googles artifact repository. It will include creating the artifac repository
using infrastructure as code

The project is in a [monorepo](/tags/monorepo) that uses [bazel](/tags/bazel).

## The executable to be packaged

```starlark
# //products/example/cmd/server/BUILD.bazel
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "server_lib",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)
```

<!-- more -->

## Build Container Image

### add `rules_oci`

[rules_oci](https://github.com/bazel-contrib/rules_oci) is a good place to start
to integrate container support into your bazel configuration. It's also worth
[reading up a bit on distroless](https://github.com/GoogleContainerTools/distroless)
if you are not aware of it.

Adding `rules_oci` into bazel `MODULES` is straightforward:

```starlark
bazel_dep(name = "rules_oci", version = "2.2.6")
# For testing, we also recommend https://registry.bazel.build/modules/container_structure_test

oci = use_extension("@rules_oci//oci:extensions.bzl", "oci")

# Declare external images you need to pull, for example:
oci.pull(
    name = "distroless_base",
    # 'latest' is not reproducible, but it's convenient.
    # During the build we print a WARNING message that includes recommended 'digest' and 'platforms'
    # values which you can use here in place of 'tag' to pin for reproducibility.
    tag = "latest",
    image = "gcr.io/distroless/base",
    platforms = ["linux/amd64"],
)

# For each oci.pull call, repeat the "name" here to expose them as dependencies.
use_repo(oci, "distroless_base")
```

If you use the `WORKSPACE`, or if you want the latest version, you
[can find details on their releases page](https://github.com/bazel-contrib/rules_oci/releases).

### `tar.bzl`

`oci_rules` uses tar files to build the image. To build tar images, you will
want to use [tar.bzl](https://github.com/bazel-contrib/tar.bzl).

Add the following to you `MODULES`

```starlark
bazel_dep(name = "tar.bzl", version = "0.3.0")
```

Latest version and instructions for `WORKSPACE`
[can be found on their releases page](https://github.com/bazel-contrib/tar.bzl/releases/tag/v0.3.0)

### `BUILD.bazel`

There are
[language specific sample build files](https://github.com/bazel-contrib/rules_oci?tab=readme-ov-file#usage)
that you can start from.

Starting
[from the example go `BUILD.bazel`](https://github.com/aspect-build/bazel-examples/blob/main/oci_go_image/BUILD.bazel),
and simplifying it, I have:

```starlark
# //products/example/deploy/BUILD.bazel
load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")
load("@tar.bzl", "mutate", "tar")

# Put app go_binary into a tar layer.
tar(
    name = "app_layer",
    srcs = ["//products/muster/cmd/server:server"],
    out = "app_layer.tar",
    mutate = mutate(strip_prefix = package_name() + "/app_"),
)


oci_image(
    name = "image",
    # This is defined by an oci.pull() call in /MODULE.bazel
    base = "@distroless_base",
    entrypoint = ["/app"],
    # Link the resulting image back to the repository where the build is defined.
    labels = {
        "org.opencontainers.image.source": "https://github.com/aspect-build/bazel-examples",
    },
    tars = [":app_layer"],
)

```

There is
[more information at the go doc page](https://github.com/bazel-contrib/rules_oci/blob/main/docs/go.md)
including details of migrating from
[`rules_docker`](https://github.com/bazelbuild/rules_docker)

```bash
products/example/deploy $ bazel build ...
INFO: Analyzed 2 targets (0 packages loaded, 0 targets configured).
INFO: Found 2 targets...
INFO: Elapsed time: 0.444s, Critical Path: 0.07s
INFO: 2 processes: 1 internal, 1 linux-sandbox.
INFO: Build completed successfully, 2 total actions
```

Building of the images now completes

## Push Image

Pushing the image can be pretty straightforward:

- Enable the container repository in [Google Cloud](/tags/google-cloud)
- Configure it in the `BUILD.bazel` file
- `bazel run` to push

I would like to do the whole thing through [IaC](/tags/iac) so that it is

- Repeatable, and
- more importantly documented

### IaC Tool Options

[terraform](https://developer.hashicorp.com/terraform) has no bazel support, and
is a questionable choice from an ethics perspective. [opentofu]() has

[rules_tf](https://github.com/yanndegat/rules_tf). However, that
[does not support `apply`](https://github.com/yanndegat/rules_tf/issues/5)

[Google Cloud Infrastucture Manager](https://cloud.google.com/infrastructure-manager/docs)
is vendor lock in.

That leaves us with [pulumi](https://www.pulumi.com/) (If there is another
alternative, please let me know).

I like the stacks concept in pulumi, and while it doesn't have bazel
integration, what it does have is an
[automation api](https://github.com/pulumi/automation-api-examples). While not
he best documented, there is enough documentation out there to be able to piece
it together.

With this, we can build a runnable unit, and then run it too. I've experimented
with tagging specific runnables and then running them through the CI. I won't
get that far in this post. My focus here is to get it working locally.

### Enable Artifact Repository

In pulumi, you can enable the artifact repository with:

```go

func enableArtifactRepository(ctx *pulumi.Context) error {
	// Set your GCP project and region
	projectID := "<projectId>"
	region := "europe-west1"
	repoName := "<repo-name>"

	// Create Artifact Registry repository
	repo, err := artifactregistry.NewRepository(ctx, repoName, &artifactregistry.RepositoryArgs{
		Format:       pulumi.String("DOCKER"),
		Location:     pulumi.String(region),
		RepositoryId: pulumi.String(repoName),
		Description:  pulumi.String("Repository for OCI container images"),
		Project:      pulumi.String(projectID),
	})
	if err != nil {
		return err
	}

	// Export the repository URL (used for pushes)
	ctx.Export("repositoryURL", pulumi.Sprintf("%s-docker.pkg.dev/%s/%s", region, projectID, repo.Name))

	return nil

}

```

### Authenticating against the repo

You need to authenticate against the repository so that the `oci_rule` can pick
it up

```bash
gcloud auth configure-docker europe-west1-docker.pkg.dev
```

### Define the push target

```starlark
oci_push(
    name = "push",
    image = ":image",
    remote_tags = [
        "latest",
        "1h",
    ],
    repository = "europe-west1-docker.pkg.dev/<projectId>/<repo-name>/<image-name>",
)

```

## Pushing

Once all the above is set up, you can then push the image with one of the
following:

```bash
bazel run :push # if you in the directory
bazel run //<target/path>:push # fron anywhere in the repo
```

## Side note

Please do not take my heavy usage of google and its products to be a fan letter
or an endorsement. They _might_ at best, be the lesser evil.
