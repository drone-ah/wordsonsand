---
title: "Auto reload WASM with zig+lume"
date: 2025-09-16T15:56:42+01:00
tags:
  - learning-zig
  - zig
  - wasm
  - mobile-dev
  - supabase
  - imgui
  - lume
---

I’ve been taking some time off to rest and recover from health issues that made
it hard to focus. To ease back in, I’ve started a small project:
[shine](../../excursions/shine.md)

The project would do well to be multiplatform - mobile and web. The obvious
choice was [flutter](/tags/flutter) and I have enjoyed working with it before.

However, as I'm currently in love with [zig](/tags/zig), I wanted to work with
that instead.

## Libraries

### Graphics

I've been playing with [raylib](/tags/raylib) and that was my initial instinct.
However, raylib
[does not support iOs](https://github.com/raysan5/raylib/discussions/2681) and
[has issues with wasm](https://github.com/raysan5/raylib/discussions/3626)

I considered a few options, including

- [sdl](https://www.libsdl.org/), which looked great but was perhaps a little
  too low level for me.
- [jok](https://github.com/Jack-Ji/jok) - does not support mobile and possibly
  has a little more than I needed.

In the end, I decided to go with [sokol](https://github.com/floooh/sokol) and
[sokol-zig](https://github.com/floooh/sokol-zig).

While a little more lower level than raylib, it has:

- a modern clean api
- first class mobile support
- first class wasm support

### UI

I've been working with [dvui](/tags/dvui) a lot recently. Unfortunately, it
doesn't support sokol. [imgui](https://github.com/SpexGuy/Zig-ImGui) is a better
option.

There is even a
[template project that I could start from](https://github.com/floooh/sokol-zig-imgui-sample).

## WASM first

To keep things straightforward, I decided to start with wasm. If I make the site
mobile friendly, I could see how it goes and see if it needs a mobile version.

I will need some shared data storage and have been considering
[supabase](https://supabase.com/), which has [javascript](/tags/javascript)
libs. By using [wasm](/tags/wasm), I can effectively shim in js functions to
handle that instead of having to write bare rest calls from zig.

### Structuring the project

Is this a zig project with a web component, vice versa or indeed two independent
parts that work together.

I did a fair amount of web searching to see if there was some guidance I could I
find for a good way to structure a relatively straightforward zig+js project.

I could not find one. In the end, I decided to keep it fairly straightforward.

```
- shine/
  - src/ # zig code
  - web/ # all the web stuff
```

### Frontend

After a bit of research, and realising that I will probably need a little bit of
supporting content around [shine](../../excursions/shine.md), I decided to go
with [lume](https://lume.land/)

I used the [simple-blog theme](https://github.com/lumeland/theme-simple-blog) as
a template to start from. I could have just pulled the template in but I wanted
a custom homepage.

When I tried to add an `index.md`, it complained about two files wanting to
write `index.html`. From what I could find, the easiest way to override the
homepage was to just pick up the theme and edit it - which was easy enough.

### WASM => frontend

I didn't want to copy over the wasm and the js file every time, so I added a
couple of steps to `build.zig` right after the `link_step` (also included below)

```zig
// build.zig
// create a build step which invokes the Emscripten linker
const link_step = try sokol.emLinkStep(b, .{
    .lib_main = shine,
    .target = opts.mod_main.resolved_target.?,
    .optimize = opts.mod_main.optimize.?,
    .emsdk = dep_emsdk,
    .use_webgl2 = true,
    .use_emmalloc = true,
    .use_filesystem = false,
    .shell_file_path = opts.dep_sokol.path("src/sokol/web/shell.html"),
});
// attach to default target
b.getInstallStep().dependOn(&link_step.step);

// Copy shine.js from default emscripten output
const js_install = b.addInstallFileWithDir(
    b.path("zig-out/web/shine.js"),
    .{ .custom = "../web/src/static/shine" },
    "shine.js",
);
js_install.step.dependOn(&link_step.step);
b.getInstallStep().dependOn(&js_install.step);

// Copy shine.wasm from default emscripten output
const wasm_install = b.addInstallFileWithDir(
    b.path("zig-out/web/shine.wasm"),
    .{ .custom = "../web/src/static/shine" },
    "shine.wasm",
);
wasm_install.step.dependOn(&link_step.step);
```

These steps will copy across the wasm and the js file across to `static/shine`.
I wanted to put the js in `src/js` and the wasm in the static dir. However, the
js file expects the wasm in the same dir. I tried overriding `locateFile` but it
didn't work.

```html
<!-- index.vto -->
<script>
  window.Module = {
    locateFile: (path, prefix) => {
      if (path.endsWith(".wasm")) {
        return "/static/shine/shine.wasm";
      }
      return prefix + path;
    },
  };
</script>
<script src="/js/shine.js"></script>
```

I was able to get lume to process the javascript file by `add`ing it.

```javascript
// _config.ts
site.add("static/shine/shine.js");
```

## Conclusion

With all of these set up, I was able to run:

```bash
zig build -Dtarget=wasm32-emscripten --watch
```

```

```

in one window. This command will rebuild wasm and provide it to lume whenever
the zig code changes.

```bash
deno task serve
```

Running this in another window will mean that lume will rebuild on any changes,
including a new wasm file and redeploy. The redeploy will trigger an auto-reload
of the page as well if I have it in a browser.

I now effectively have automated reload with changes if I make changes in either
zig or the frontend.

I don't have _hot_ reload - but this is pretty good for now.
