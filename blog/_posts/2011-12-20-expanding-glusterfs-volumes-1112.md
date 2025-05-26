---
layout: post
title: Expanding glusterfs volumes [1112]
date: 2011-12-20 13:26:55.000000000 +00:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Systems (Administration)
tags:
- computing
- GlusterFS
- Logical volume management
- Storage
meta:
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {'glusterfs': {\"text\":\"GlusterFS\",\"slug\":\"glusterfs\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/43fe36c2-a883-33a0-a757-4cada4db79f2/SocialTag/3\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"GlusterFS\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'logical-volume-management': {\"text\":\"Logical volume management\",\"slug\":\"logical-volume-management\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/43fe36c2-a883-33a0-a757-4cada4db79f2/SocialTag/6\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Logical
    volume management\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'storage': {\"text\":\"Storage\",\"slug\":\"storage\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/43fe36c2-a883-33a0-a757-4cada4db79f2/SocialTag/8\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Storage\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'computing': {\"text\":\"computing\",\"slug\":\"computing\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  oc_commit_id: http://drone-ah.com/2011/12/20/expanding-glusterfs-volumes-1112/1324387621
  _edit_last: '48492462'
  restapi_import_id: 591d994f7aad5
  original_post_id: '746'
  _wp_old_slug: '746'
  _publicize_job_id: '5181540093'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:53'
permalink: "/2011/12/20/expanding-glusterfs-volumes-1112/"
---

[Once you have set up a glusterfs
volume](http://drone-ah.com/2011/11/24/glusterfs-howto/ "GlusterFS HOWTO [1108]"),
you might want to expand the volume to add storage. This is an
astoundingly easy task.

The first thing that you\'ll want to do is to add in bricks. Bricks are
similar to physical volumes a la LVM. The thing to bear in mind is that
depending on what type of cluster you have (replicated / striped), you
will need to add a certain number of blocks at a time.

Once you have a initialised the nodes, to add in a set of bricks, you
need the following command which adds two more bricks to a cluster which
keeps two replicas.

\$ gluster volume add-brick testvol cserver3:/gdata cserver4:/gdata

Once you have done this, you will need to rebalance the cluster, which
involves redistributing the files across all the bricks. There are two
steps to this process, the \"fixing\" of the layout changes and the
rebalancing of the data itself. You can perform both tasks together.

As a starting point, to view the status of a rebalance, you can use:

    $ gluster volume rebalance testvol status

You can also stop / pause a rebalance with

    $ gluster volume rebalance testvol stop

To \"fix\" the layout changes, you need to run:

    $ gluster volume rebalance testvol fix-layout start
    Starting rebalance on volume test-volume has been successful

Rebalancing the volume to migrate the data is easy and can be done using
a similar command:

    $ gluster volume rebalance testvol migrate-data start

To complete both in one command, you just need:

    $ gluster volume rebalance testvol start

Easy right?

With this mechanism, you have the ability to have storage that can be
expanded on the fly by using additional hardware. You can also remove
existing bricks using:

    $ gluster volume remove-brick testvol cserver2:/gdata

This means that you can remove a brick with smaller hard drives, upgrade
the harddrives, and re-integrate into the cluster with bigger hard
drives. This means that you have a cloud like storage solution which you
can easily grow as necessary without worrying about resizing underlying
filesystems or hotswapping hardisks or any of that hassle.
