---
layout: post
title: My Thoughts on OCFS2 / Understanding OCFS2 [1110]
date: 2011-11-24 21:42:09.000000000 +00:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Systems (Administration)
tags:
- Computer storage
- DRBD
- File system
- GFS2
- Global File System
- OCFS
- OCFS2
- ocfs2-tools
- SAN
- Shared disk file system
- Ubuntu
meta:
  _edit_last: '48492462'
  oc_commit_id: http://drone-ah.com/2011/11/24/my-thoughts-on-ocfs2-understanding-ocfs2-1110/1322170932
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {'computer-storage': {\"text\":\"Computer
    storage\",\"slug\":\"computer-storage\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/3f4ea6a8-3fbe-38cd-94e8-2ea913aa69bf/SocialTag/2\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Computer
    storage\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'drbd': {\"text\":\"DRBD\",\"slug\":\"drbd\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/3f4ea6a8-3fbe-38cd-94e8-2ea913aa69bf/SocialTag/3\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"DRBD\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'ocfs': {\"text\":\"OCFS\",\"slug\":\"ocfs\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/3f4ea6a8-3fbe-38cd-94e8-2ea913aa69bf/SocialTag/5\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"OCFS\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'shared-disk-file-system': {\"text\":\"Shared disk file system\",\"slug\":\"shared-disk-file-system\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/3f4ea6a8-3fbe-38cd-94e8-2ea913aa69bf/SocialTag/6\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Shared
    disk file system\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'file-system': {\"text\":\"File system\",\"slug\":\"file-system\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/3f4ea6a8-3fbe-38cd-94e8-2ea913aa69bf/SocialTag/7\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"File
    system\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'global-file-system': {\"text\":\"Global File System\",\"slug\":\"global-file-system\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/3f4ea6a8-3fbe-38cd-94e8-2ea913aa69bf/SocialTag/8\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Global
    File System\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'san': {\"text\":\"SAN\",\"slug\":\"san\",\"source\":{\"url\":\"http://d.opencalais.com/genericHasher-1/b0d5d6d3-78c1-3255-aa1c-834011a6a1ac\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/em/e/Technology\",\"name\":\"Technology\",\"_className\":\"ArtifactType\"},\"name\":\"SAN\",\"_className\":\"Entity\",\"rawRelevance\":0.65,\"normalizedRelevance\":0.65},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'ubuntu': {\"text\":\"Ubuntu\",\"slug\":\"ubuntu\",\"source\":{\"url\":\"http://d.opencalais.com/genericHasher-1/6f67c971-6f56-339f-9365-ba86a70b09b0\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/em/e/OperatingSystem\",\"name\":\"OperatingSystem\",\"_className\":\"ArtifactType\"},\"name\":\"Ubuntu\",\"_className\":\"Entity\",\"rawRelevance\":0.364,\"normalizedRelevance\":0.364},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'ocfs2-tools': {\"text\":\"ocfs2-tools\",\"slug\":\"ocfs2-tools\",\"source\":{\"url\":\"http://d.opencalais.com/genericHasher-1/0e3edcd8-65cd-3510-9e32-db89f0c37a46\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/em/e/IndustryTerm\",\"name\":\"IndustryTerm\",\"_className\":\"ArtifactType\"},\"name\":\"ocfs2-tools\",\"_className\":\"Entity\",\"rawRelevance\":0.348,\"normalizedRelevance\":0.348},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'ocfs2': {\"text\":\"OCFS2\",\"slug\":\"ocfs2\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'gfs2': {\"text\":\"GFS2\",\"slug\":\"gfs2\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: '744'
  _wp_old_slug: '744'
  _publicize_job_id: '5181540161'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:52'
permalink: "/2011/11/24/my-thoughts-on-ocfs2-understanding-ocfs2-1110/"
---

[As mentioned
earlier](http://drone-ah.com/2011/11/24/glusterfs-howto/ "GlusterFS HOWTO [1108]"),
we have been considered networked filesystems instead of NFS to
introduce into a number of complex environments. OCFS2 was one of the
first candidates.

In fact, we also considered GFS2 but looking around on the net, there
seemed to be a general consensus recommending ocfs2 over gfs2.

Ubuntu makes it pretty easy to install and manage ocfs2 clusters. You
just need to install ocfs2-tools and ocfs2console. You can then use the
console to manage the cluster.

What I totally missed in all of my research and understanding, and due
to lack of in depth knowledge on clustered filesystems was that OCFS2
(and GFS2 for that matter) are shared disk file systems.

What does this mean?

[Wikipedia](http://en.wikipedia.org/ "Wikipedia"){target="_blank"} defines
a [shared disk
filesystem](http://en.wikipedia.org/wiki/Shared_disk_file_system "Shared Disk File System"){target="_blank"} as
being \"shared by being
simultaneously [mounted](http://en.wikipedia.org/wiki/Mount_(computing) "Mount (computing)") on
multiple [servers](http://en.wikipedia.org/wiki/Server_(computing) "Server (computing)").\"

This essentially means that the storage medium is mounted on to cluster.
The cluster is a collection of the clients. The storage is traditionally
a SAN mount point. This means that a shared storage space is accessed at
high speeds by a number of clients.

From a simplistic point of view, this is not that different from
mounting the SAN point onto a server and then running an NFS server on
it with all the clients mounting over NFS.

The main difference is that OCFS is distributed. There is no single
point of failure (assuming that the storage medium is redundant which a
SAN would be). The NFS server is a clear single point of failure.

If you do not have access to or want to use a SAN, [you can also use
DRBD.](http://www.drbd.org/users-guide/ch-ocfs2.html "Using OCFS2 with DRBD"){target="_blank"}

NFS is described as a network filesystem and GlusterFS is described as a
NAS file system.

 

 
