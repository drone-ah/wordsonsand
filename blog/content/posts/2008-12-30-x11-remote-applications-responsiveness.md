---
categories:
  - Systems (Administration)
date: "2008-12-30T15:01:43Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:35"
  _publicize_job_id: "5181540574"
  _wp_old_slug: "123"
  original_post_id: "123"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Systems (Administration)
  - Linux
  - Remote X11
  - ssh
  - X11
title: X11 Remote Applications Responsiveness
url: /2008/12/30/x11-remote-applications-responsiveness/
---

As a developer, I use eclipse a lot... We have a powerful server that off which
eclipse is run which allows us to keep the desktops at a much lower spec. In
general, this works well for us.

However, recently, I have been niggled by the amount of time it takes to switch
perspectives on eclipse. It takes a good 4 seconds to switch between
perspectives.There is also a noticeable lag when performing some operations.

To resolve this, I spent a lot of time looking at the linux real-time and
low-latency patches. I had expected that running X11 applications remotely would
not cause a bottleneck over a gigabit link. Turns out that I was wrong.

To test this, I ran a vnc server on the application server and found that
switching perspectives on there was super fast.

To be able to resolve this, the first thing to do was to remove any latency put
on the X->X communication by ssh.

We use gdm, so I had to enable to TCP on there first. Do this using the
following config line in `/etc/gdm/gdm.conf`

```
DisallowTCP=false
```

Restart gdm

on the remote host, export DISPLAY

```
export DISPLAY=<yourhost>:0
```

and run your application.

I found the application to be a lot more responsive after this. I didn't have to
worry about X auth since we have nfs mounted home. If you don't, check
[this mini howto](http://www.xs4all.nl/~zweije/xauth.html "Remote X Apps Mini HowTo")
