---
categories:
  - Systems (Administration)
date: "2011-12-20T13:26:55Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:53"
  _publicize_job_id: "5181540093"
  _wp_old_slug: "746"
  oc_commit_id: http://drone-ah.com/2011/12/20/expanding-glusterfs-volumes-1112/1324387621
  original_post_id: "746"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Systems (Administration)
  - computing
  - GlusterFS
  - Logical volume management
  - Storage
title: Expanding glusterfs volumes [1112]
url: /2011/12/20/expanding-glusterfs-volumes-1112/
---

[Once you have set up a glusterfs volume](/2011/11/24/glusterfs-howto/ "GlusterFS HOWTO [1108]"),
you might want to expand the volume to add storage. This is an astoundingly easy
task.

The first thing that you'll want to do is to add in bricks. Bricks are similar
to physical volumes a la LVM. The thing to bear in mind is that depending on
what type of cluster you have (replicated / striped), you will need to add a
certain number of blocks at a time.

Once you have a initialised the nodes, to add in a set of bricks, you need the
following command which adds two more bricks to a cluster which keeps two
replicas.

```bash
$ gluster volume add-brick testvol cserver3:/gdata cserver4:/gdata
```

Once you have done this, you will need to rebalance the cluster, which involves
redistributing the files across all the bricks. There are two steps to this
process, the "fixing" of the layout changes and the rebalancing of the data
itself. You can perform both tasks together.

As a starting point, to view the status of a rebalance, you can use:

```bash
$ gluster volume rebalance testvol status
```

You can also stop / pause a rebalance with

```bash
$ gluster volume rebalance testvol stop
```

To "fix" the layout changes, you need to run:

```bash
$ gluster volume rebalance testvol fix-layout start
Starting rebalance on volume test-volume has been successful
```

Rebalancing the volume to migrate the data is easy and can be done using a
similar command:

```bash
$ gluster volume rebalance testvol migrate-data start
```

To complete both in one command, you just need:

```bash
$ gluster volume rebalance testvol start
```

Easy right?

With this mechanism, you have the ability to have storage that can be expanded
on the fly by using additional hardware. You can also remove existing bricks
using:

```bash
$ gluster volume remove-brick testvol cserver2:/gdata
```

This means that you can remove a brick with smaller hard drives, upgrade the
harddrives, and re-integrate into the cluster with bigger hard drives. This
means that you have a cloud like storage solution which you can easily grow as
necessary without worrying about resizing underlying filesystems or hotswapping
hardisks or any of that hassle.
