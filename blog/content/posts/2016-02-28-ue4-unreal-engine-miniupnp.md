---
categories:
  - Game Development
date: "2016-02-28T10:11:15Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:00"
  _publicize_job_id: "5186311673"
  _rest_api_client_id: "-1"
  _rest_api_published: "1"
parent_id: "0"
password: ""
status: publish
tags:
  - Game Development
  - UE4
title: "[UE4] Unreal Engine &amp; miniupnp"
url: /2016/02/28/ue4-unreal-engine-miniupnp/
---

This post covers how to integrate pnp into an unreal project using miniupnp.

The first step is to clone the project from
[github](https://github.com/miniupnp/miniupnp/)

The module that we are interested in is the miniupnpc and in that directory,
there is another directory called msvc and this contains the solution file for
Visual Studio. Open this and if you have a more recent version of visual studio
(which very likely do), it will want to upgrade everything. Let it go through
the upgrade process.

Building the project now will most likely fail due to a missing file
miniupnpcstrings.h . This file needs to be generated and the way to do that is
to run a script in that folder called updateminiupnpstrings.sh. You will most
probably need something like [cygwin](https://www.cygwin.com/) to for this
script to work as it a [unix shell](https://en.wikipedia.org/wiki/Unix_shell) to
work

<!--more-->

Once the miniupnpcstrings.h has been generated, we also need to follow some
instructions for Unreal Engine
for [Linking Static Libraries Using The Build System](https://wiki.unrealengine.com/Linking_Static_Libraries_Using_The_Build_System),
particularly the section on customizations for targeting UE4 modules.

From the project properties page, choose configuration manager. From
the **Active Solutions Platform**, select new and type in or select x64 and save
it. You have to do this for only one of the projects.

Building of the static project will fail since it can't find the lib which is
now in x64Release as opposed to just Release. The exe is not required for
integrating with Unreal Engine, but if you want to complete the build, just fix
the path in Project `Properties -> Linker -> Input`.

You should choose the release build instead of the debug build and you should
now be able to build the solution from visual studio. It did pop up some
warnings for me, but the build completed successfully.

The rest of the instructions are from the unreal engine documentation about
integrating static libraries, starting from section about
[Third Party Directory](https://wiki.unrealengine.com/Linking_Static_Libraries_Using_The_Build_System#Third_Party_Directory)
