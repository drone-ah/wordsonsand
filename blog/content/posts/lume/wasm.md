---
title:
  "Using `locateFile` to have js and wasm in different locations with emscripten"
date: 2025-09-18T10:15:42+01:00
tags:
  - lume
  - webassembly
---

As part of building [shine](../../excursions/shine.md), I am using
[lume](https://lume.land) and webassembly with zig.

zig, through emscripten generates both a js and wasm file, which, by default are
expected to be in the same directory.

I wanted to put them in different places, and struggled to get that working with
lume for a bit. I did eventually solve it though:

## Include emscripten js file

Firstly, the js file output from emscripten should be included in you page
manually. I had used `site.add` which meant that it was loaded _before_ we could
put the override for `locateFile` in `window.Module`.

I am using the
[simple-blog theme](https://github.com/lumeland/theme-simple-blog), so I add the
following line to the top of `src/index.vto`

```html
<script defer src="/js/shine.js"></script>
```

## Override `Module`

You can override how emscripten finds the wasm file in the
[Module Object](https://emscripten.org/docs/api_reference/module.html#Module.locateFile).

This can be done in the html file in a script block, or even better, in a js
file that is included automatically.

`js/main.js` felt like a good place.

```javascript
window.Module = {
  locateFile: function (path, scriptDirectory) {
    if (path === "shine.wasm") {
      return "/static/shine/shine.wasm";
    } else {
      return scriptDirectory + path;
    }
  },
};
```

## Closing

With this setup, I can not only have the js and wasm files in different
locations, it's also easy to modify / augment the `Module` object.
