---
layout: post
title: Exporting X11 to Windows [1109]
date: 2011-11-24 21:10:55.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Systems (Administration)
tags:
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
meta:
  _edit_last: "48492462"
  oc_metadata:
    "{\t\tversion:'1.1',\t\ttags: {'cross-platform-software':
    {\"text\":\"Cross-platform
    software\",\"slug\":\"cross-platform-software\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/3\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Cross-platform
    software\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'linux':
    {\"text\":\"Linux\",\"slug\":\"linux\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/4\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Linux\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'x-servers': {\"text\":\"X
    servers\",\"slug\":\"x-servers\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/5\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"X
    servers\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'cygwin':
    {\"text\":\"Cygwin\",\"slug\":\"cygwin\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/7\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Cygwin\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'putty':
    {\"text\":\"PuTTY\",\"slug\":\"putty\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/9\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"PuTTY\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'x-window-system': {\"text\":\"X Window
    System\",\"slug\":\"x-window-system\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/10\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"X
    Window
    System\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'xming':
    {\"text\":\"Xming\",\"slug\":\"xming\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/11\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Xming\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'linux-kernel': {\"text\":\"Linux
    kernel\",\"slug\":\"linux-kernel\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/0888b5c3-fff0-3f6a-a016-8116518b1b97/SocialTag/12\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Linux
    kernel\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'x-windows': {\"text\":\"X
    Windows\",\"slug\":\"x-windows\",\"source\":{\"url\":\"http://d.opencalais.com/genericHasher-1/a8bed2b3-0f9d-3dba-83c9-698d783d28f2\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/em/e/Technology\",\"name\":\"Technology\",\"_className\":\"ArtifactType\"},\"name\":\"X
    Windows\",\"_className\":\"Entity\",\"rawRelevance\":0.333,\"normalizedRelevance\":0.333},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'windows-7': {\"text\":\"Windows
    7\",\"slug\":\"windows-7\",\"source\":{\"url\":\"http://d.opencalais.com/genericHasher-1/1a627e2d-ba9c-3368-aa24-5b51d4ed863f\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/em/e/OperatingSystem\",\"name\":\"OperatingSystem\",\"_className\":\"ArtifactType\"},\"name\":\"Windows
    7\",\"_className\":\"Entity\",\"rawRelevance\":0.418,\"normalizedRelevance\":0.418},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  oc_commit_id: http://drone-ah.com/2011/11/24/exporting-x11-to-windows-1109/1322169058
  restapi_import_id: 591d994f7aad5
  original_post_id: "732"
  _wp_old_slug: "732"
  _publicize_job_id: "5181540244"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:52"
permalink: "/2011/11/24/exporting-x11-to-windows-1109/"
---

Playing Skyrim the last week, sometimes I just missed Linux so terribly that I
wanted a piece of it and not just the command line version. I wanted X Windows
on my Windows 7.

There has been a solution for this for several years and the first time I did
this, I installed [cygwin](http://www.cygwin.com/ "cygwin") with X11 but there
is a far simpler way to accomplish this.

Install
[XMing](http://www.straightrunning.com/XmingNotes/ "XMing"){target="\_blank"}. I
then used putty, which has the forward X11 option. Once logged in, running xeyes
shows the window exported onto my Windows 7. Ah.. so much better.

I actually used this to run terminator to connect to a number of servers. Over
local LAN, the windows didn\'t have any perceptible lag or delay. It was more or
less like running it locally.

It is possible to set up shortcuts to run an application through putty and have
it exported to your desktop. I haven\'t played with this enough to comment
though.

This of course only worked because I have another box which is running Linux. If
that is not the case for you, then you might want to try
[VirtualBox](https://www.virtualbox.org/ "VirtualBox"){target="\_blank"} but
since the linux kernel developers have described the kernel modules as
[tainted crap](http://www.phoronix.com/scan.php?page=news_item&px=OTk5Mw "The VirtualBox Kernel Driver Is Tainted Crap"){target="\_blank"},
you might want to consider
[vmware](http://www.vmware.com "vmware"){target="\_blank"} instead which is an
excellent product.
