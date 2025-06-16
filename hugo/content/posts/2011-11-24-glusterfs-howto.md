---
categories:
- Systems (Administration)
date: "2011-11-24T20:53:33Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:52"
  _publicize_job_id: "5181540343"
  _wp_old_slug: "738"
  oc_commit_id: http://drone-ah.com/2011/11/24/glusterfs-howto/1322168013
  original_post_id: "738"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- disk cloud
- GlusterFS
- Ubuntu
title: GlusterFS HOWTO [1108]
url: /2011/11/24/glusterfs-howto/
---

So, I  am catching up a bit on the technical documentation. A week taken to play
Skyrim combined with various other bits and pieces made this a little difficult.

On the bright side, there are a few new things that have been worked on so
hopefully plenty of things to cover soon.

We manage a number of servers and all over the place and all of them require to
be backed up. We also have a number of desktops all with mirrored disks also
getting backed up.

I like things to be all nicely efficient and its annoying when one server /
desktop runs out of space when another two (or ten) has plenty of space. We grew
to dislike NFS particularly due to the single point of failure and there were
few other options.

We had tried [glusterfs](http://www.gluster.org/ "GlusterFS") a few years ago
(think it was at version 1.3 or something) and there were various issues
particularly around small files and configuration was an absolute nightmare.

With high hopes that version 3.2 was exactly what we were looking for, we set up
three basic machines for testing

<!--more-->

Previously, [glusterfs](http://www.gluster.org/ "GlusterFS") required all the
configuration to be completed manually and with text files. It also required a
fairly detailed knowledge of what they called translators and a lot of tweaking
and fiddling with parameters.

I am very happy to report that this is no longer the case with 3.2.

The three servers(cserver[1-3]) are running Ubuntu and was updated to Natty
(11.10) to get access to glusterfs 3.2 (11.04 only had 3.0). One thing to bear
in mind is that the glusterfs website seemed to only have the 64 bit version but
Ubuntu 11.10 also has the 32bit version.

Installing the server part of glusterfs was simple and straightfoward

```bash
$ sudo aptitude install glusterfs-server
```

Once this was done all three servers
([terminator](http://www.tenshu.net/p/terminator.html "Terminator") is a godsend
when doing these things across a number of servers), adding the servers into a
"cluster" was easy enough.

```bash
shri@cserver1:~$ gluster peer probe cserver2
Probe successful

shri@cserver1:~$ sudo gluster peer probe cserver3
Probe successful
```

The thing to note is that these probe statements are two way. In other words,
all three servers are now part of the same cluster.

```bash
shri@cserver3:~$ sudo gluster peer status
Number of Peers: 2

Number of Peers: 3

Hostname: cserver1
Uuid: 8fe63300-e227-4aec-81f3-69b33f894330
State: Peer in Cluster (Connected)

Hostname: cserver2
Uuid: 275ce612-2dd8-4e2a-8cc8-3115ad18c594
State: Peer in Cluster (Connected)
```

The thing is that if you type in an incorrect hostname, the probe will keep
trying to connect to it. I haven't left it running long enough to know if it
every returns.

```bash
shri@cserver3:~# sudo gluster peer probe does-not-exist
^C
shri@cserver3:~# sudo gluster peer status
Number of Peers: 3

Hostname: cserver1
Uuid: 8fe63300-e227-4aec-81f3-69b33f894330
State: Peer in Cluster (Connected)

Hostname: cserver2
Uuid: 275ce612-2dd8-4e2a-8cc8-3115ad18c594
State: Peer in Cluster (Connected)

Hostname: does-not-exist
Uuid: 00000000-0000-0000-0000-000000000000
State: Establishing Connection (Disconnected)
```

Thankfully, removing host does-not-exist is simple enough

```bash
shri@cserver3:~# sudo gluster peer detach does-not-exist
Detach successful
shri@cserver3:~# sudo gluster peer status
Number of Peers: 2

Hostname: cserver1
Uuid: 8fe63300-e227-4aec-81f3-69b33f894330
State: Peer in Cluster (Connected)

Hostname: cserver2
Uuid: 275ce612-2dd8-4e2a-8cc8-3115ad18c594
State: Peer in Cluster (Connected)
```

Creating a volume is straightforward. There are a number of different types of
volumes which you can find out from the documentation. In this particular
instance, we are creating a distributed replicated

```bash
shri@cserver3:~$ sudo gluster volume create testvol replica 2 transport tcp cserver1:/gdata cserver2:/gdata
Creation of testvol has been successful
Please start the volume to access data
```

The reason I have not included cserver3 in here is that the volume needs a
multiple of the replica number of bricks. In the case, the there needs to be a
muliple of 2 number of bricks.

Additionally, you could use rdma instead of tcp if you are using infiniband

Starting the volume is simple enough

```bash
shri@cserver3:~$ sudo gluster volume start testvol
```

this volume is now accessible from all the boxes in the cluster

```bash
shri@cserver1:~# sudo gluster volume info

Volume Name: testvo
Type: Distributed-Replicate
Status: Started
Number of Bricks: 2 x 1 = 2
Transport-type: tcp
Bricks:
Brick1: cserver1:/gdata
Brick2: cserver2:/gdata
```

Mounting this from another box is easy

```bash
$ sudo aptitude install glusterfs-client
$ sudo mount -t glusterfs /mnt cserver:/testvol
```

If you get the error of "endpoint not connected" when listing the content of the
mount, it is likely because the volume is not started.

If you are curious, check the gdata folders in the bricks after copying some
files into the mount and you'll find them show up intact and on both bricks in
the above example.
