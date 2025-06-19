---
categories:
  - devex
date: "2025-06-09T20:57:53Z"
meta: null
parent_id: "0"
password: ""
slug: use-protoc-bin-in-bazel
status: publish
tags:
  - devex
  - bazel
  - pulumi
  - pulumi automation api
  - protoc
title: Use `protoc` bin instead of building from source in `bazel`
---

I work on a project that uses `pulumi` automation api, which in turn uses
`protobuf`. I think there are other bits in `bazel` that also uses it.

For a while, it was fine - except that some `CI` runs would take 15 minutes and
we couldn't quite figure out why.

I finally had to update `bazel` from 7.x to 8 and in that process (which
honestly could have been easier, but oh well), I ran into a problem with
compiling protobuf.

Namely, I kept running into this error:

```
error while loading shared libraries: libstdc++.so.6: cannot open shared object file: No such file or directory
```

I tried many things but was not able to get past this error. In the end, I was
able to narrow it down to devbox/nix and how it messes with the environment that
doesn't quite agree with `bazel`.

## What didn't work

- installing `gcc` from devbox
- installing `stdenv.cc.cc.lib`
- Setting `LD_LIBRARY_PATH` (to
  `<workspace-dir>/.devbox/nix/profile/default/lib`, which is where the devbo
  version of `libstdc++.so.6` is installed by `stdenv.cc.cc.lib`)
- Using the `--action-env=`
- [rules-nixpkgs](https://github.com/tweag/rules_nixpkgs)

None of which really worked.

Disabling `devbox` **did** work, which was at least partly encouraging. I also
found a bunch of evidence online around this issue
([1](https://github.com/bazelbuild/bazel/issues/12978),
[2](https://github.com/jetify-com/devbox/issues/1100),
[3](https://github.com/jetify-com/devbox/issues/1596),
[4](https://github.com/jetify-com/devbox/issues/710),
[5](https://github.com/tweag/rules_nixpkgs/issues/573)). Bazel seems to
[generally dislike non-system gcc](https://discuss.ray.io/t/bazel-protobuf-build-errors-libstdc-with-non-system-gcc/3329).

<!--more-->

## A ray of hope

I was just about ready to throw in the towel when ChatGPT, in passing suggested
using the `protobuf` binary directly. I wasn't even using this binary - it was
being pulled in as a dependency and I just needed it to work. I did not need it
to be built from source.

Furthermore, I found, from my research into trying to solve this that building
`protobuf` was the likely culprit in the `CI` taking 15 minutes.

I also remembered that someone else on the team had issues trying to get it to
build on a mac at some point.

All in all, it was a problematic piece of software and I was seriously
considering switching to `opentofu` - but they might be using it too.

## bin `protoc`

It was difficult to find decent documentation about how to achieve this though,
apart from a partially answered on:

- [stackoverflow question](https://stackoverflow.com/questions/68918369/is-it-possible-to-use-bazel-without-compiling-protobuf-compiler),
- [google groups for bazel-discuss](https://groups.google.com/g/bazel-discuss/c/3Q_GEqNZrC0)
  - which also led me to
    [a gitlab repo with some code](https://gitlab.com/mvfwd/issue-bazel-protobuf-compile/-/tree/main)

These resources gave me just enough to be able to cobble together a working
solution, which needs:

### Install `protoc`

The first thing we need is a locally installed `protoc` bin (I'll used devbox to
install it for consistency across dev environments)

```bash
devbox add protobuf
```

You can add a version specifier for better reproducibility.

### Runnable Target

We also need a shell target that will execute this correctly. The `bazel` flag
to override `protoc` takes a local target, not a bin.

`third_party/tools/BUILD.bazel`

```python
package(default_visibility = ["//visibility:public"])

sh_binary(
    name = "protoc",
    srcs = ["protoc.sh"],
)

# https://github.com/protocolbuffers/protobuf/blob/b4b0e304be5a68de3d0ee1af9b286f958750f5e4/BUILD#L773
proto_lang_toolchain(
    name = "cc_toolchain",
    command_line = "--cpp_out=$(OUT)",
    runtime = ":protoc",
    visibility = ["//visibility:public"],
)
```

`third_party/tools/protoc.sh`

```bash
#!/bin/env bash
protoc "$@"

```

### Override protoc

You should now be able to override `protoc` with:

```bash
bazel build --proto_compiler=//third_party/tools:protoc ...
```

Having to pass in a flag each time is annoying though, and you can add it to
your `.bazelrc`

```
build --proto_compiler=//third_party/tools:protoc
```

### Sample code

You can find the sample code in
[the wordsonsand repo](https://github.com/drone-ah/wordsonsand) which uses this
exact solution to override `protoc` and get the `pulumi` sample code to work ;)

PS: It also includes a fully migrated `MODULES.bazel` with support for `golang`.
You will also want to check out `BUILD` in the root.
