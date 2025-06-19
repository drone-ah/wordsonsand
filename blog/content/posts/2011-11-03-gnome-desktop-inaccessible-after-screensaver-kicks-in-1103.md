---
categories:
  - Software
date: "2011-11-03T11:11:14Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:51"
  _publicize_job_id: "5185240655"
  _wp_old_slug: "698"
  oc_commit_id: http://drone-ah.com/2011/11/03/gnome-desktop-inaccessible-after-screensaver-kicks-in-1103/1320318677
  original_post_id: "698"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Software
  - GNOME
  - Ubuntu
  - Workspace
  - X Window System
title: Gnome Desktop Inaccessible After Screensaver Kicks in [1103]
url: /2011/11/03/gnome-desktop-inaccessible-after-screensaver-kicks-in-1103/
---

Yesterday, I
[mentioned a problem that I\'ve been having](http://drone-ah.com/2011/11/02/saving-your-workspace-window-configuration-in-linux-1102/ "Saving your workspace window configuration in Linux [1102]")
with GNOME 3 on Ubuntu 11.10.

Essentially what happens is that when I leave my desktop for a while, under
specific circumstances, and often, on returning and moving the mouse or using
the keyboard, the pointer would come back  on screen. However, this only works
on one of my two screens.

The unlock dialog does not show up and it seems that there is no way to get back
in.

In the past, I would log into the terminal (Ctrl-Alt-F1 or any function key
through to F5 or so) and

```bash
$ kill -9 -1
```

<!--more-->

This would of course kill all processes owned by me and is therefore unpleasant
at best and have you losing a bunch of work at worst.

After a brainwave yesterday (as detailed in the aforementioned post), I decided
to check the status of the screensaver and killed just those processes. Happily,
this gives me my desktop back. However, my gnome-shell had given up which I had
to restart

```bash
$ gnome-shell --replace
```

Unfortunately, I did not get the windows into the original workspaces since
everything just got dumped into the one workspace but it is better than having
to kill everything off.

EDIT: I just realised that the screen saver of course no longer kicks in and I
had to restart it

```bash
$ gnome-screensaver --no-daemon
```
