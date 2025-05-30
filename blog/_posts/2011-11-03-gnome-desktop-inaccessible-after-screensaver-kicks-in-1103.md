---
layout: post
title: Gnome Desktop Inaccessible After Screensaver Kicks in [1103]
date: 2011-11-03 11:11:14.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software
tags:
  - GNOME
  - Ubuntu
  - Workspace
  - X Window System
meta:
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2011/11/03/gnome-desktop-inaccessible-after-screensaver-kicks-in-1103/1320318677
  oc_metadata:
    "{\t\tversion:'1.1',\t\ttags: {'gnome':
    {\"text\":\"GNOME\",\"slug\":\"gnome\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/73728cc7-f8ba-3dc3-8f2e-7b09cc9aa3b9/SocialTag/5\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"GNOME\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'x-window-system': {\"text\":\"X Window
    System\",\"slug\":\"x-window-system\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/73728cc7-f8ba-3dc3-8f2e-7b09cc9aa3b9/SocialTag/7\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"X
    Window
    System\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'workspace':
    {\"text\":\"Workspace\",\"slug\":\"workspace\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/73728cc7-f8ba-3dc3-8f2e-7b09cc9aa3b9/SocialTag/8\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Workspace\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'ubuntu':
    {\"text\":\"Ubuntu\",\"slug\":\"ubuntu\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/73728cc7-f8ba-3dc3-8f2e-7b09cc9aa3b9/SocialTag/9\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Ubuntu\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: "698"
  _wp_old_slug: "698"
  _publicize_job_id: "5185240655"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:51"
permalink: "/2011/11/03/gnome-desktop-inaccessible-after-screensaver-kicks-in-1103/"
---

Yesterday, I
[mentioned a problem that I\'ve been having](http://drone-ah.com/2011/11/02/saving-your-workspace-window-configuration-in-linux-1102/ "Saving your workspace window configuration in Linux [1102]"){target="\_blank"}
with GNOME 3 on Ubuntu 11.10.

Essentially what happens is that when I leave my desktop for a while, under
specific circumstances, and often, on returning and moving the mouse or using
the keyboard, the pointer would come back  on screen. However, this only works
on one of my two screens.

The unlock dialog does not show up and it seems that there is no way to get back
in.

In the past, I would log into the terminal (Ctrl-Alt-F1 or any function key
through to F5 or so) and

    $ kill -9 -1

This would of course kill all processes owned by me and is therefore unpleasant
at best and have you losing a bunch of work at worst.

After a brainwave yesterday (as detailed in the aforementioned post), I decided
to check the status of the screensaver and killed just those processes. Happily,
this gives me my desktop back. However, my gnome-shell had given up which I had
to restart

    $ gnome-shell --replace

Unfortunately, I did not get the windows into the original workspaces since
everything just got dumped into the one workspace but it is better than having
to kill everything off.

EDIT: I just realised that the screen saver of course no longer kicks in and I
had to restart it

    $ gnome-screensaver --no-daemon
