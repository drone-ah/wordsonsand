---
categories:
- Systems (Administration)
date: "2011-11-24T21:42:09Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:52"
  _publicize_job_id: "5181540161"
  _wp_old_slug: "744"
  original_post_id: "744"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
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
title: My Thoughts on OCFS2 / Understanding OCFS2 [1110]
type: post
url: /2011/11/24/my-thoughts-on-ocfs2-understanding-ocfs2-1110/
---

[As mentioned earlier](/2011/11/24/glusterfs-howto/ "GlusterFS HOWTO [1108]"),
we have been considered networked filesystems instead of NFS to introduce into a
number of complex environments. OCFS2 was one of the first candidates.

In fact, we also considered GFS2 but looking around on the net, there seemed to
be a general consensus recommending ocfs2 over gfs2.

Ubuntu makes it pretty easy to install and manage ocfs2 clusters. You just need
to install ocfs2-tools and ocfs2console. You can then use the console to manage
the cluster.

What I totally missed in all of my research and understanding, and due to lack
of in depth knowledge on clustered filesystems was that OCFS2 (and GFS2 for that
matter) are shared disk file systems.

What does this mean?

<!--more-->

[Wikipedia](http://en.wikipedia.org/ "Wikipedia") defines a
[shared disk filesystem](http://en.wikipedia.org/wiki/Shared_disk_file_system "Shared Disk File System") as
being "shared by being
simultaneously [mounted](<http://en.wikipedia.org/wiki/Mount_(computing)> "Mount
(computing)") on
multiple [servers](<http://en.wikipedia.org/wiki/Server_(computing)> "Server (computing)")."

This essentially means that the storage medium is mounted on to cluster. The
cluster is a collection of the clients. The storage is traditionally a SAN mount
point. This means that a shared storage space is accessed at high speeds by a
number of clients.

From a simplistic point of view, this is not that different from mounting the
SAN point onto a server and then running an NFS server on it with all the
clients mounting over NFS.

The main difference is that OCFS is distributed. There is no single point of
failure (assuming that the storage medium is redundant which a SAN would be).
The NFS server is a clear single point of failure.

If you do not have access to or want to use a SAN,
[you can also use DRBD.](http://www.drbd.org/users-guide/ch-ocfs2.html "Using OCFS2 with DRBD")

NFS is described as a network filesystem and GlusterFS is described as a NAS
file system.
