---
layout: post
title: Understanding ZFS Disk Utilisation and available space
date: 2017-08-16 12:02:07.000000000 +01:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Systems (Administration)
tags: []
meta:
  _edit_last: "48492462"
  geo_public: "0"
  _publicize_job_id: "8320377172"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:01"
permalink: "/2017/08/16/understanding-zfs-disk-utilisation-and-available-space/"
---

I am hopeful the following will help someone scratch their head a little less in
trying to understand the info returned by zfs.

I set up a pool using 4 2TB SATA disks.

```bash
$ zpool list -v
NAME SIZE ALLOC FREE EXPANDSZ FRAG CAP DEDUP HEALTH ALTROOT
rpool 7.25T 2.50T 4.75T - 10% 34% 1.00x ONLINE -
raidz2 7.25T 2.50T 4.75T - 10% 34%
sda2 - - - - - -
sdb2 - - - - - -
sdc2 - - - - - -
sdd2 - - - - - -
```

The total size displayed here is the total size of the 4 disks. The maths works
as 4\*2TB = 8TB = ~7.25TiB

RAIDZ2 is like RAID6 and it uses two disks for parity. Thus, I would expect to
have ~4TB or 3.63TiB of available space. I haven't been able to find this number
displayed anywhere.

**However**, you can find the amount of disk space still available using the
following command.

```bash
$# zfs list
NAME USED AVAIL REFER MOUNTPOINT
rpool 1.21T 2.19T 140K /rpool
rpool/ROOT 46.5G 2.19T 140K /rpool/ROOT
rpool/ROOT/pve-1 46.5G 2.19T 46.5G /
rpool/data 1.16T 2.19T 140K /rpool/data
rpool/data/vm-100-disk-1 593M 2.19T 593M -
rpool/data/vm-101-disk-1 87.1G 2.19T 87.1G -
rpool/data/vm-102-disk-1 71.2G 2.19T 71.2G -
rpool/data/vm-103-disk-1 2.26G 2.19T 2.26G -
rpool/data/vm-103-disk-2 13.2M 2.19T 13.2M -
rpool/data/vm-103-disk-3 13.2M 2.19T 13.2M -
rpool/data/vm-103-disk-4 93K 2.19T 93K -
rpool/data/vm-103-disk-5 1015G 2.19T 1015G -
rpool/data/vm-104-disk-1 4.73G 2.19T 4.73G -
rpool/data/vm-105-disk-1 4.16G 2.19T 4.16G -
rpool/swap 8.66G 2.19T 8.66G -
```

The value of **2.19T** is the amount of unallocated space available in the pool.
To verify this, you can run

```
# zfs get all rpool
NAME PROPERTY     VALUE                           SOURCE
rpool type        filesystem                       -
rpool creation    Fri Aug 4 20:39 2017             -
rpool used        1.21T                            -
rpool available   2.19T                            -

...

```

If we add the two numbers here, 1.21T + 2.19T = 3.4T.

5% of disk space is reserved, so 3.63 \* 0.95 = 3.4T

et voila
