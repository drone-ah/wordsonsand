---
layout: post
title: Synergy with Linux Server &amp; Mac Client
date: 2010-04-18 15:41:03.000000000 +01:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Systems (Administration)
tags:
- computing
- Control key
- kvm
- Linux
- Macintosh
- synergy
- synergykm
- Ubuntu
meta:
  _publicize_pending: '1'
  _edit_last: '48492462'
  oc_commit_id: http://drone-ah.com/2010/04/18/synergy-with-linux-server-mac-client/1271605270
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {'control-key': {\"text\":\"Control key\",\"slug\":\"control-key\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/9ba43e9a-57ac-3a13-beb1-24df491fbc13/SocialTag/4\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Control
    key\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'ubuntu': {\"text\":\"Ubuntu\",\"slug\":\"ubuntu\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/9ba43e9a-57ac-3a13-beb1-24df491fbc13/SocialTag/6\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Ubuntu\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'macintosh': {\"text\":\"Macintosh\",\"slug\":\"macintosh\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/9ba43e9a-57ac-3a13-beb1-24df491fbc13/SocialTag/7\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Macintosh\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'linux': {\"text\":\"Linux\",\"slug\":\"linux\",\"source\":{\"_className\":\"Entity\",\"url\":\"http://d.opencalais.com/genericHasher-1/8e32c2ad-38a7-3069-96e3-577198801f0a\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/em/e/Technology\",\"name\":\"Technology\"},\"name\":\"Linux\",\"rawRelevance\":0.845,\"normalizedRelevance\":0.845},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'synergy': {\"text\":\"synergy\",\"slug\":\"synergy\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'synergykm': {\"text\":\"synergykm\",\"slug\":\"synergykm\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'kvm': {\"text\":\"kvm\",\"slug\":\"kvm\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'computing': {\"text\":\"computing\",\"slug\":\"computing\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: '298'
  _wp_old_slug: '298'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:39'
permalink: "/2010/04/18/synergy-with-linux-server-mac-client/"
---

I  borrowed a mac to try and play
with [iPhone](http://en.wikipedia.org/wiki/IPhone){#aptureLink_ihNLsBYMct}
development. I already have a linux box (running Ubuntu 9.10). Anyone
who has used two computers simultaneously know how annoying it is to
have two keyboards/mice plugged. I originally anticipated just using X11
forwarding. However, it is
an [iMac](http://en.wikipedia.org/wiki/IMac){#aptureLink_loEUxPLQKw}
with a big beautiful screen. It would be an absolute waste to not use
it.

I
installed [synergy](http://en.wikipedia.org/wiki/Synergy%20%28software%29){#aptureLink_DZ7FSFnNJL}
on both ends, with the linux one as the server

    $ sudo aptitude install synergy

and the mac as the client

<http://sourceforge.net/projects/synergykm/>

and it worked.

There was just one very very annoying problem. The Ctrl key and Cmd keys
were different. This really messed with my muscle memory. After some
hunting around, I just had to update my .synergy.conf file in linux.
Here is the relevant section

    section: screens
        linux-desktop:
        imac:
        ctrl=alt
        alt=ctrl
        meta=alt
    end

et voila. It now works a charm. I  have neglected the configuration of
the synergykm and synergys but these can be figured out easily ;-)
