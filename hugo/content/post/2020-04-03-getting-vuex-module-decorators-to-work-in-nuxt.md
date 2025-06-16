---
categories: []
date: "2020-04-03T10:10:28Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:03"
  _publicize_job_id: "42570371069"
  timeline_notification: "1585908629"
parent_id: "0"
password: ""
published: true
status: publish
tags: []
title: Getting vuex-module-decorators to work in nuxt
type: post
url: /2020/04/03/getting-vuex-module-decorators-to-work-in-nuxt/
---

There are a few caveats to integrating
[vuex-module-decorators](https://github.com/championswimmer/vuex-module-decorators)
with nuxt. The first steps are described in the
[README](https://github.com/championswimmer/vuex-module-decorators#accessing-modules-with-nuxtjs)
(although I missed it because it nuxt was in the small text).

In addition to this, according to
[issue #179](https://github.com/championswimmer/vuex-module-decorators/issues/179),
there are a few other caveats

- [The file name has to be the same as the module name](https://github.com/championswimmer/vuex-module-decorators/issues/179#issuecomment-533853333)
- [The store should go in \~/store, not \~/store/modules](https://github.com/championswimmer/vuex-module-decorators/issues/179#issuecomment-549326864)
- [You have to **export default** your module classes](https://github.com/championswimmer/vuex-module-decorators/issues/179)
