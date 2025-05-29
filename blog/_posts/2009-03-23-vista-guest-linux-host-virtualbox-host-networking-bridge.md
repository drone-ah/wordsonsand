---
layout: post
title: Vista Guest, Linux Host, VirtualBox, Host Networking - Bridge
date: 2009-03-23 15:39:41.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Systems (Administration)
tags:
  - Guest OS
  - host networking
  - Linux
  - Microsoft Vista
  - Ubuntu
  - VirtualBox
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2009/03/23/vista-guest-linux-host-virtualbox-host-networking-bridge/1237822782
  oc_metadata:
    "{\t\tversion:1.0,\t\ttags: {'ubuntu':
    {\t\t\ttext:'Ubuntu',\t\t\tslug:'ubuntu',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/6f67c971-6f56-339f-9365-ba86a70b09b0',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/OperatingSystem',\t\t\ticonURL:'',\t\t\tname:'OperatingSystem'\t\t},\t\t\tname:'Ubuntu',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t},'host-networking':
    {\t\t\ttext:'host
    networking',\t\t\tslug:'host-networking',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/b4cb0156-47a7-36cc-a0e1-8b262208ee21',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/IndustryTerm',\t\t\ticonURL:'',\t\t\tname:'IndustryTerm'\t\t},\t\t\tname:'host
    networking',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t},'microsoft-vista':
    {\t\t\ttext:'Microsoft
    Vista',\t\t\tslug:'microsoft-vista',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/fdd3ca0c-7a56-3474-83f1-b94ca64d21c2',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/OperatingSystem',\t\t\ticonURL:'',\t\t\tname:'OperatingSystem'\t\t},\t\t\tname:'Microsoft
    Vista',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t},'linux':
    {\t\t\ttext:'Linux',\t\t\tslug:'linux',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/8e32c2ad-38a7-3069-96e3-577198801f0a',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/Technology',\t\t\ticonURL:'',\t\t\tname:'Technology'\t\t},\t\t\tname:'Linux',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t},'virtualbox':
    {\t\t\ttext:'VirtualBox',\t\t\tslug:'virtualbox',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'guest-os':
    {\t\t\ttext:'Guest
    OS',\t\t\tslug:'guest-os',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: "242"
  _wp_old_slug: "242"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:38"
permalink: "/2009/03/23/vista-guest-linux-host-virtualbox-host-networking-bridge/"
---

One would think that it would be straightforward, work off the bat, or at least
have some reasonable documentation. Unfortunately, no!

I needed host networking to be able to access network resources (Samba shares
etc.) which does not work if the guest OS is on NAT :-(

Solving it was easy though\... I assume Vista is installed as a guest with the
guest additions and that your user account is a part of the vboxusers group.

On the linux host, first install bridge utils. I run Ubuntu, so it was as easy
as:

\$ sudo aptitude install bridge-utils

Next, you need to set up the bridge; again, easy on Ubuntu:

add the following section to /etc/network/interfaces

auto br0\
iface br0 inet dhcp\
bridge_ports eth1

Add the interfaces to VirtualBox

\$ sudo VBoxAddIF vbox0 \'shri\' br0

Within the VirtualBox Guest settings, choose Host Networking and fo the
interface, choose br0

bring the interface up:

\$ sudo ifup br0

and start your guest os\... et voila, it just works\...
