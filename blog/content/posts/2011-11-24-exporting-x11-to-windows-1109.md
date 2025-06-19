---
categories:
  - Systems (Administration)
date: "2011-11-24T21:10:55Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:52"
  _publicize_job_id: "5181540244"
  _wp_old_slug: "732"
  oc_commit_id: http://drone-ah.com/2011/11/24/exporting-x11-to-windows-1109/1322169058
  original_post_id: "732"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Systems (Administration)
  - Cross-platform software
  - Cygwin
  - Linux
  - Linux kernel
  - PuTTY
  - Windows 7
  - X servers
  - X Window System
  - X Windows
  - Xming
title: Exporting X11 to Windows [1109]
url: /2011/11/24/exporting-x11-to-windows-1109/
---

Playing Skyrim the last week, sometimes I just missed Linux so terribly that I
wanted a piece of it and not just the command line version. I wanted X Windows
on my Windows 7.

There has been a solution for this for several years and the first time I did
this, I installed [cygwin](http://www.cygwin.com/ "cygwin") with X11 but there
is a far simpler way to accomplish this.

Install [XMing](http://www.straightrunning.com/XmingNotes/ "XMing"). I then used
putty, which has the forward X11 option. Once logged in, running xeyes shows the
window exported onto my Windows 7. Ah.. so much better.

I actually used this to run terminator to connect to a number of servers. Over
local LAN, the windows didn\'t have any perceptible lag or delay. It was more or
less like running it locally.

It is possible to set up shortcuts to run an application through putty and have
it exported to your desktop. I haven\'t played with this enough to comment
though.

This of course only worked because I have another box which is running Linux. If
that is not the case for you, then you might want to try
[VirtualBox](https://www.virtualbox.org/ "VirtualBox") but since the linux
kernel developers have described the kernel modules as
[tainted crap](http://www.phoronix.com/scan.php?page=news_item&px=OTk5Mw "The VirtualBox Kernel Driver Is Tainted Crap"),
you might want to consider [vmware](http://www.vmware.com "vmware") instead
which is an excellent product.
