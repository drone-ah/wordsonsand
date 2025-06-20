---
categories:
  - Systems (Administration)
date: "2009-03-23T15:39:41Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:38"
  _publicize_pending: "1"
  _wp_old_slug: "242"
  oc_commit_id: http://drone-ah.com/2009/03/23/vista-guest-linux-host-virtualbox-host-networking-bridge/1237822782
  original_post_id: "242"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - sysadmin
  - host-networking
  - linux
  - microsoft-vista
  - ubuntu
  - virtualbox
title: Vista Guest, Linux Host, VirtualBox, Host Networking - Bridge
url: /2009/03/23/vista-guest-linux-host-virtualbox-host-networking-bridge/
---

One would think that it would be straightforward, work off the bat, or at least
have some reasonable documentation. Unfortunately, no!

I needed host networking to be able to access network resources (Samba shares
etc.) which does not work if the guest OS is on NAT :-(

Solving it was easy though... I assume Vista is installed as a guest with the
guest additions and that your user account is a part of the vboxusers group.

On the linux host, first install bridge utils. I run Ubuntu, so it was as easy
as:

```bash
$ sudo aptitude install bridge-utils
```

Next, you need to set up the bridge; again, easy on Ubuntu:

add the following section to /etc/network/interfaces

```
auto br0
iface br0 inet dhcp
bridge_ports eth1
```

Add the interfaces to VirtualBox

```bash
$ sudo VBoxAddIF vbox0 'shri' br0
```

Within the VirtualBox Guest settings, choose Host Networking and fo the
interface, choose br0

bring the interface up:

```
$ sudo ifup br0
```

and start your guest os... et voila, it just works...
