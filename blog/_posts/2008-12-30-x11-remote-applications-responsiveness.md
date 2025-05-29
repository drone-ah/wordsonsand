---
layout: post
title: X11 Remote Applications Responsiveness
date: 2008-12-30 15:01:43.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Systems (Administration)
tags:
  - Linux
  - Remote X11
  - ssh
  - X11
meta:
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2008/12/30/x11-remote-applications-responsiveness/1236796166
  oc_metadata:
    "{\t\tversion:1.0,\t\ttags: {'remote-x11': {\t\t\ttext:'Remote
    X11',\t\t\tslug:'remote-x11',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'ssh':
    {\t\t\ttext:'ssh',\t\t\tslug:'ssh',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'x11':
    {\t\t\ttext:'X11',\t\t\tslug:'x11',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'linux':
    {\t\t\ttext:'Linux',\t\t\tslug:'linux',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/8e32c2ad-38a7-3069-96e3-577198801f0a',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/Technology',\t\t\ticonURL:'',\t\t\tname:'Technology'\t\t},\t\t\tname:'Linux',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: "123"
  _wp_old_slug: "123"
  _publicize_job_id: "5181540574"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:35"
permalink: "/2008/12/30/x11-remote-applications-responsiveness/"
---

As a developer, I use eclipse a lot\... We have a powerful server that off which
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
on the X-\>X communication by ssh.

We use gdm, so I had to enable to TCP on there first. Do this using the
following config line in /etc/gdm/gdm.com

> DisallowTCP=false

Restart gdm

on the remote host, export DISPLAY

> export DISPLAY=\<yourhost\>:0

and run your application.

I found the application to be a lot more responsive after this. I didn\'t have
to worry about X auth since we have nfs mounted home. If you don\'t, check
[this mini howto](http://www.xs4all.nl/~zweije/xauth.html "Remote X Apps Mini HowTo"){target="\_blank"}\
