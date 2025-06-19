---
categories:
  - Systems (Administration)
date: "2011-11-02T23:57:22Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:51"
  _publicize_pending: "1"
  _wp_old_slug: "694"
  oc_commit_id: http://drone-ah.com/2011/11/02/saving-your-workspace-window-configuration-in-linux-1102/1320278245
  original_post_id: "694"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Systems (Administration)
  - Devil's Pie
  - GNOME
  - Linux
  - nifty tool
  - System software
  - Technology/Internet
  - Workspace
  - X11
title: Saving your workspace window configuration in Linux [1102]
url: /2011/11/02/saving-your-workspace-window-configuration-in-linux-1102/
---

I am usually working on a good half a dozen things at any given time and this
means that I usually have a good ten or twenty windows open. My chromium
currently has a 134 tabs and this is after I  cleaned up and closed all the tabs
I no longer need.

Luckily, working in Linux means that I can spread each stream of work into the
various workspaces.

Now GNOME 3 makes things a little more complicated with the dynamic workspaces
but I'm learning to use it to my advantage

However, with Ubuntu 11.10 Oneiric Ocelot and GNOME 3, I seem to be running into
an issue regularly...If I leave my computer for a while, it doesn't unlock
correctly. The screen remains black and I can't move the mouse to my second
screen and the unlock screen doesn't show up.

<!--more-->

Thinking about it, it seems like there might be two screen savers being started
but I shall investigate that tomorrow. I have the same issue at both work and
home so it is more likely to be related to Ubuntu + GNOME 3 or something about
the way I set things up.

I  usually resolve this by logging into the console and here a neat trick for
killing all our processes in one fell swoop.

```bash
$ kill -9 -1
```

Another thing I have been doing a bit more of recently is gaming which involves
rebooting in Windows.

Both of the above leaves me with a restarted workspace. Starting up the
applications pops them all into the same workspace. Chrome is especially a
nightmare. I might have 135 open tabs but they are in about 6 windows spread
across four workspaces.

It is annoying to have to distribute these things out each time.

After having done much research, I have not been able to find a clean automated
solution.

There are two half solution that I have found however.

The first one is [Devil's Pie](http://live.gnome.org/DevilsPie "Devil's Pie")
and for a graphical interface
[gdevilspie](http://code.google.com/p/gdevilspie/ "gdevilspie"). According the
website for Devil's Pie, it is "A totally crack-ridden program for freaks and
weirdos who want precise control over what windows do when they appear. If you
want all XChat windows to be on desktop 3, in the lower-left, at 40%
transparency, you can do it."

Unfortunately, that is exactly what it is. If you pre-determine where you want
your windows to be, you can use this very useful application. However, that is
not quite what I want. I want the current configuration to be remember. Exactly
like how Chromium remembers which tabs are in which order in which windows and
their position on the workspace, but for multiple workspaces.

Unfortunately, I couldn't find any way to save the current state.

There is however, another tool
[I found scouring the web.](http://thialfihar.org/projects/window_position_session/ "Window Position Session")

[](http://thialfihar.org/projects/window_position_session/ "Window Position Session")libwnck-3-dev
is what I installed on my Ubuntu box. There are two key commands here

```bash
$ wnckprop --list
```

This will list all the windows across all the workspaces. To get more
information on a specific Window,

```bash
wnckprop --xid [XID]
```

The XID is the number returned next to each window from the first command. The
post that I  mentioned above has a nifty tool attached that saves the window
positions and can also restore them using wnckprop.

However, it saves them based on the Window title. This of course doesn't work
for Chromium or such Windows that changes the title each time you change the
tab.

However, if the save is the last command you run and the restore is the first
command you run after opening up the windows, it can restore the windows into
the correct workspaces.

With the idea of the dynamic workspaces in GNOME 3, you might have to initialise
the workspaces first but it is better than spending five minutes after logging
in each time re-arranging windows...
