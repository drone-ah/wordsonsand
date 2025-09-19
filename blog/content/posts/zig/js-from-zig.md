---
title: "Calling Javascript from Zig through WebAssembly"
date: 2025-09-17T15:52:20+01:00
tags:
  - zig
  - clang
  - webassembly
  - javascript
  - typescript
  - emscripten
  - supabase
  - deno
  - lume
---

The next step for [shine](../../excursions/shine.md) is to build a bridge
between [zig](/tags/zig) and [javascript](/tags/javascript).

I am currently planning to using [supabase](https://supabase.com/) for storage.
Unsurprisingly, it does not have a zig sdk. It does, however, have a javascript
sdk.

If I can write basic CRUD operations in javascript and call that from zig
through webassembly, that could make that integration a lot easier.

## Goals

There are a few ideal restrictions for me - mainly because writing javascript is
not fun for me.

- Use the Supabase js/ts library through zig
- Use TypeScript as much as possible. (I don't love TypeScript, but at least
  it's not javascript)
- Keep as much of the supabase related code in the web part so that
  [deno](/tags/deno) and [lume](/tags/lume) can handle any heavy lifting.

## Options

I can't use FFI(Foreign Function Interface):

```zig
extern "env" fn jsLog(ptr: [*]const u8, len: usize) void;
```

```javascript
const wasm = await WebAssembly.instantiateStreaming(fetch("prog.wasm"), {
  env: {
    jsLog: (ptr, len) => {
      /* read from memory and console.log */
    },
  },
});
```

because imgui pulls in emscripten, which means we don't have the ability to call
`instantiateStreaming`.

With emscripten, declaring an external function is easy enough.

```zig
extern fn jsLog(ptr: [*]const u8, len: usize) void;
```

There are a couple of options to wire them up to the javascript:

### `library.js` / `mergeInto`

This option requires javascript files on the zig side. If you want to start with
typescript, you'll need to integrate a transpiler into the build chain as well.

First, you want a javascript file - let's call it `libshine.js`, and pop it into
a `js` dir.

```javascript
// js/libshine.js
// Emscripten will provide these globals at link/runtime
declare var mergeInto: (lib: any, funcs: Record<string, Function>) => void;
declare var LibraryManager: { library: any };
declare function UTF8ToString(ptr: number, len?: number): string;

mergeInto(LibraryManager.library, {
  jsLog: (ptr: number) => {
    const msg = UTF8ToString(ptr);
    console.log("🟢 Zig says:", msg);
  },
});
```

We then need to pass this `js` file into the build step

as part of my sokol build step, I pass it in as `.extra_args`.

```zig
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
    // set the js file here
    .extra_args = &.{
        "--js-library", "js/libshine.js",
    },
});
```

We can then call it from zig, with something like:

```zig
pub fn main() void {
    jsLog("hello from zig");
}

extern fn jsLog(ptr: [*]const u8) void;
```

From my firefox console:

```
Lume live reloading is ready. Listening for changes...     localhost:3000:102:15
🟢 Zig says: hello from zig                                shine.js:3168:11
```

### `EM_JS` / `EM_ASM`

The other option is to use
[`EM_JS`](https://emscripten.org/docs/porting/connecting_cpp_and_javascript/Interacting-with-code.html#interacting-with-code-call-javascript-from-native)
which involves writing a wee bit of `C`, which can embed the `javascript`.

In theory, it's as simple as:

```c
#include <emscripten.h>

EM_JS_DEPS(bla, "$UTF8ToString");

EM_JS(void, jsLog, (const char* s), {
  console.log(UTF8ToString(s));
});

```

and adding it into the build file:

```zig
// build the main file into a library, this is because the WASM 'exe'
// needs to be linked in a separate build step with the Emscripten linker
const shine = b.addLibrary(.{
    .name = "shine",
    .root_module = opts.mod_main,
});

// get the Emscripten SDK dependency from the sokol dependency
const dep_emsdk = opts.dep_sokol.builder.dependency("emsdk", .{});

// need to inject the Emscripten system header include path into
// the cimgui C library otherwise the C/C++ code won't find
// C stdlib headers
const emsdk_incl_path = dep_emsdk.path("upstream/emscripten/cache/sysroot/include");

shine.root_module.addCSourceFile(.{
    .file = b.path("src/libjs.c"),
    .flags = &.{}, // optional extra emcc flags
});
shine.addSystemIncludePath(emsdk_incl_path);
```

The calling code in `main.zig` remains the same:

```zig
pub fn main() void {
    jsLog("hello from zig");
}

extern fn jsLog(ptr: [*]const u8) void;
```

However, this didn't work, and failed with:

```
error: undefined symbol: jsLog (referenced by root reference (e.g. compiled C/C++ code))
warning: To disable errors for undefined symbols use `-sERROR_ON_UNDEFINED_SYMBOLS=0`
warning: _jsLog may need to be added to EXPORTED_FUNCTIONS if it arrives from a system library
Error: Aborting compilation due to previous errors
```

Thanks to
[some help](https://ziggit.dev/t/help-with-getting-a-simple-call-to-js-through-emscripten-working/12090/3)
from [flooh](https://ziggit.dev/u/floooh) (who btw put together the
[sokol](https://github.com/floooh/sokol) and
[sokol-zig](https://github.com/floooh/sokol-zig) packages as well the
[sokol-imgui-sample](https://github.com/floooh/sokol-zig-imgui-sample) template
which I used to kick start this project.), I was able to get it working.

Turns out the c file needs to have a function in it that is used in the zig
file - it doesn't need to do anything.

So, based on the suggestion, `libjs.c` changes to:

```c
#include <emscripten.h>

EM_JS_DEPS(bla, "$UTF8ToString");

EM_JS(void, jsLog, (const char* s), {
  console.log(UTF8ToString(s));
});

void dummy(void) {};
```

and in `main.zig`:

```zig
pub fn main() void {
    dummy();
    jsLog("hello from zig");
}

extern fn jsLog(ptr: [*]const u8) void;

extern fn dummy() void;
```

From my firefox console:

```
Lume live reloading is ready. Listening for changes...     localhost:3000:102:15
🟢 Zig says: hello from zig                                shine.js:3168:11
```

You can see a working example in [my forked repo](

## `EM_JS` directly through `zig` \[unsuccessful\]

Looking at the macro for `EM_JS` and with my good friend ChatGPT, I attempted
translating it to zig and made some progress, but ultimately failed to get it
working. I'll leave the work here in the hopes it might be helpful.

```c
#define _EM_JS(ret, c_name, js_name, params, code)                             \
  _EM_BEGIN_CDECL                                                              \
  ret c_name params EM_IMPORT(js_name);                                        \
  __attribute__((visibility("hidden")))                                        \
  void* __em_js_ref_##c_name = (void*)&c_name;                                 \
  EMSCRIPTEN_KEEPALIVE                                                         \
  __attribute__((section("em_js"), aligned(1))) char __em_js__##js_name[] =    \
    #params "<::>" code;                                                       \
  _EM_END_CDECL
```

The above macro translates to zig roughly (with help from ChatGPT) as:

```zig
extern fn jsLog(ptr: [*]const u8) void;

/// 2. Keep a reference to avoid the linker removing the function.
///    Same role as __em_js_ref_* in the C macro.
pub export const __em_js_ref_jsLog = &jsLog;

/// 3. Embed the JS implementation in a special section called "em_js".
///    Emscripten will scan this and inject the code into the output JS.
export const __em_js__jsLog align(1) linksection("em_js") =
    "(const char* s)<::>{ console.log(UTF8ToString(s)); }\x00";

pub export fn dummy() void {}
```

I added a `pub fn` and called it from main:

```zig
pub fn log(ptr: [*]const u8) void {
    jsLog(ptr);
}
```

Which gave me the familiar error about not being able to find `jsLog`.

comparing the linker sections gave some clues:

```
❯ wasm-objdump --section=linking -x <path/to/libjs.o>

libjs.o:        file format wasm 0x1

Section Details:

Custom:
 - name: "linking"
  - symbol table [count=9]
   - 0: F <dummy> func=1 [ binding=global vis=hidden ]
   - 1: D <__em_js_ref_jsLog> segment=0 offset=0 size=4 [ binding=global vis=hidden ]
   - 2: F <jsLog> func=0 [ undefined explicit_name binding=global vis=default ]
   - 3: D <__em_js__jsLog> segment=1 offset=0 size=53 [ exported no_strip binding=global vis=hidden ]
   - 4: S <.debug_abbrev> section=7 [ binding=local vis=default ]
   - 5: G <env.__stack_pointer> global=0 [ undefined binding=global vis=default ]
   - 6: S <.debug_str> section=9 [ binding=local vis=default ]
   - 7: T <env.__indirect_function_table> table=0 [ undefined exported no_strip binding=global vis=default ]
   - 8: S <.debug_line> section=10 [ binding=local vis=default ]
  - segment info [count=2]
   - 0: .data.__em_js_ref_jsLog p2align=2 [ ]
   - 1: em_js p2align=0 [ RETAIN ]
```

and the zig object:

```
❯ wasm-objdump --section=linking -x js.o

js.o:   file format wasm 0x1

Section Details:

Custom:
 - name: "linking"
  - symbol table [count=6]
   - 0: F <dummy> func=1 [ binding=global vis=default ]
   - 1: D <__em_js_ref_jsLog> segment=0 offset=0 size=4 [ binding=global vis=default ]
   - 2: F <jsLog> func=0 [ undefined explicit_name binding=global vis=default ]
   - 3: D <__em_js__jsLog> segment=1 offset=0 size=4 [ binding=global vis=default ]
   - 4: D <__anon_946> segment=2 offset=0 size=54 [ binding=local vis=default ]
   - 5: T <env.__indirect_function_table> table=0 [ undefined exported no_strip binding=global vis=default ]
  - segment info [count=3]
   - 0: .rodata.__em_js_ref_jsLog p2align=2 [ ]
   - 1: em_js p2align=0 [ ]
   - 2: .rodata.__anon_946 p2align=0 [ ]`
```

From what I could understand (which is little), it looks like `__em_js__jsLog`
in the zig obj is a pointer while from C, it's the full string.

hardcoding it as a static array helped:

```zig
export const __em_js__jsLog align(1) linksection("em_js") = [_]u8{
    '(', 'c','o','n','s','t',' ','c','h','a','r','*',' ','s',')',
    '<',':',':','>','{',' ',
    'c','o','n','s','o','l','e','.','l','o','g','(',
    'U','T','F','8','T','o','S','t','r','i','n','g','(',
    's',')',')',';',' ','}','\x00',
};
```

The output from this is a little more promising

```
❯ wasm-objdump --section=linking -x js.o

js.o:   file format wasm 0x1

Section Details:

Custom:
 - name: "linking"
  - symbol table [count=5]
   - 0: F <dummy> func=1 [ binding=global vis=default ]
   - 1: D <__em_js_ref_jsLog> segment=0 offset=0 size=4 [ binding=global vis=default ]
   - 2: F <jsLog> func=0 [ undefined explicit_name binding=global vis=default ]
   - 3: D <__em_js__jsLog> segment=1 offset=0 size=53 [ binding=global vis=default ]
   - 4: T <env.__indirect_function_table> table=0 [ undefined exported no_strip binding=global vis=default ]
  - segment info [count=2]
   - 0: .rodata.__em_js_ref_jsLog p2align=2 [ ]
   - 1: em_js p2align=0 [ ]
```

Let's look at the two side by side

```
# From C
- 3: D <__em_js__jsLog> segment=1 offset=0 size=53 [ exported no_strip binding=global vis=hidden ]
# From zig
- 3: D <__em_js__jsLog> segment=1 offset=0 size=53 [ binding=global vis=default ]
```

There are some clear differences in how the two are output and I am already
beyond my knowledge level here - so I'll leave it to someone who knows this
stuff better (or wait until I do)

You can
[check out the code in the branch of my forked repo](https://github.com/drone-ah/sokol-zig-imgui-sample/tree/zig_em_js)

## Next steps

My plan is to use `EM_JS` through `C` to implement glue JavaScript functions -
something like:

```c
EM_JS(void, jsLog, (const char* s), {
	Module.jsLog(UTF8ToString(s));
});
```

By doing this, I can have one-line js code in the `.c` file and all the
implementation can go into the web side (and can easily be TypeScript too).

```javascript
window.Module = {
  jsLog: function (msg) {
    console.log("🟢 Zig says:", msg);
  },
};
```
