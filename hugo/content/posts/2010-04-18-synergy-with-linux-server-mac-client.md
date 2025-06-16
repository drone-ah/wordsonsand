---
categories:
- Systems (Administration)
date: "2010-04-18T15:41:03Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:39"
  _publicize_pending: "1"
  _wp_old_slug: "298"
  oc_commit_id: http://drone-ah.com/2010/04/18/synergy-with-linux-server-mac-client/1271605270
  original_post_id: "298"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- computing
- Control key
- kvm
- Linux
- Macintosh
- synergy
- synergykm
- Ubuntu
title: Synergy with Linux Server &amp; Mac Client
url: /2010/04/18/synergy-with-linux-server-mac-client/
---

I  borrowed a mac to try and play
with [iPhone](http://en.wikipedia.org/wiki/IPhone) development. I already have a
linux box (running Ubuntu 9.10). Anyone who has used two computers
simultaneously know how annoying it is to have two keyboards/mice plugged. I
originally anticipated just using X11 forwarding. However, it is
an [iMac](http://en.wikipedia.org/wiki/IMac){with a big beautiful screen. It
would be an absolute waste to not use it.

I installed [synergy](http://en.wikipedia.org/wiki/Synergy%20%28software%29) on
both ends, with the linux one as the server

```bash
$ sudo aptitude install synergy
```

and the mac as the client

http://sourceforge.net/projects/synergykm/

and it worked.

There was just one very very annoying problem. The Ctrl key and Cmd keys were
different. This really messed with my muscle memory. After some hunting around,
I just had to update my .synergy.conf file in linux. Here is the relevant
section

```
section: screens
    linux-desktop:
    imac:
    ctrl=alt
    alt=ctrl
    meta=alt
end
```

et voila. It now works a charm. I  have neglected the configuration of the
synergykm and synergys but these can be figured out easily ;-)
