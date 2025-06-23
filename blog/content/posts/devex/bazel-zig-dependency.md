---
title: Add zig dependency in bazel
date: 2025-06-11T20:44:01+01:00
categories:
  - devex
tags:
  - devex
  - bazel
  - rules_zig
  - build-systems
  - howto
---

`[rules-zig][https://github.com/aherrmann/rules_zig]` doesn't use `zon`
files(yet) and addind dependencies involves adding it into `MODULES.bazel`. If
you're still using workspaces on bazel,
[this code sample mentioned in the related issue might help](https://github.com/aherrmann/rules_zig/issues/231#issuecomment-2723684385).
I started with that code sample.

I was adding in a dependency to [zig-toml](https://github.com/sam701/zig-toml)
and this is the
[rule that I used](https://github.com/drone-ah/wordsonsand/blob/9dd185e5f330403cad9afe2ad67cc511565e1325/MODULE.bazel#L43):

```starlark
zig_toml_archive(
    name = "zig_toml",
    urls = ["https://github.com/sam701/zig-toml/archive/refs/heads/main.zip"],
    sha256 = "b1f0846c3b5e9696892b0d96f8add4878c057c728623b03d6bfbd508e4af48d5",
    strip_prefix = "zig-toml-main",
    build_file_content = """
load("@rules_zig//zig:defs.bzl", "zig_module")
zig_module(
    name = "toml",
    main = "src/root.zig",
    srcs = glob(["src/**/*.zig"]),
    visibility = ["//visibility:public"],
)
""",
)

```

Let's break it down:

- `name`: Pretty flexible, but you probably want to use the lib name, or the
  name you would use in the zon file.
- `urls`: You need the url to the zip file. If you want main - you can click on
  the button that says `Code` and at the use the link at the bottom that says
  "Download ZIP".
  [GitHub documentation on source code archives](https://docs.github.com/en/repositories/working-with-files/using-files/downloading-source-code-archives)
  has more information as well as details on how to get he url for specific
  commits etc.
- `sha256`: When running for the first time, bazel will tell you that
  `a canonical reproducible form can be obtained by modifying arguments integrity = "sha256-<hash>"`.
  However, for whatever reason, it uses a different format from what you need to
  use in the `MODULES.bazel` file. To convert it, you can use:

```bash
echo '<hash>' | base64 -d | xxd -p -c 256
```

- `strip_prefix`: The zip file will extract the files out into a directory. This
  parameter is that directory name. I just unzipped the zip file and put that
  here. I'm sure there are simpler ways to get at it (please let me know below
  if you know of a way)
- `build_file_content`: This parameter lets you define what the contents of the
  build file will be when building that library. It might be nice to be able to
  reference a file rather than include it here - I think I've done that before
  but I can't remember how. Please let me know below if you know how it's done
  and I'll update it here.
