---
categories:
- Systems (Administration)
date: "2018-04-12T10:15:02Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:02"
  _publicize_job_id: "16719033684"
  _rest_api_client_id: "-1"
  _rest_api_published: "1"
  timeline_notification: "1523528103"
parent_id: "0"
password: ""
published: true
status: publish
tags:
- SysAdmin
- Technology
- zfs
title: ZFS Deleting files doesn't free up space
type: post
url: /2018/04/12/zfs-deleting-files-doesnt-free-up-space/
---

So I have a proxmox server on which I run a few VMs and the other day it
completely ran out of space. This was because of overprovisioning through thin
volumes.

After much head scratching and metaphorically banging my head against a wall,
here are the things I learnt.

<!--more-->

## Empty Trash

### Local Trash

Make sure that have emptied the trash on the VMs .Ubuntu has this issue and so
might other distributions

### Network Trash

If you have SAMBA enabled on your VMs make sure that the Recycle Bin is not
enabled. I have openmediavault running on a VM and I had to go through and
disable the Recycling Bin. Make sure that the Recycle bin is emptied. They are
hidden folders in the root of your shares.

## Correct Driver & Settings

- When setting up the hard drive for your VM, make sure you use virtio-scsi (or
  just scsi on the web interface).
  - If you disk is already set up using IDE or VirtIO,
    - Delete it. Don\'t worry, it\'s only deleting the link. The disk itself
      will show up in the interface afterwards
    - Double click on the unattached disk and select SCSI and Discard
    - You might have to fix the references to the drive in the OS
- On the Device Screen, make sure discard is selected.

## TRIM

Configure the OS to send TRIM commands to the drive

### Linux

#### On Mount

You can pass the parameter discard to any mountpoint and the correct TRIM
commands will be sent to the disk. **HOWEVER**, this is apparently a big
performance hit.

To do the actual trim, run

```bash
$ fstrim
```

OR to run fstrim on all supported drives

```bash
$ fstrim -a
```

[Digital Ocean has a detailed post about setting TRIM and setting up a schedule etc.](https://www.digitalocean.com/community/tutorials/how-to-configure-periodic-trim-for-ssd-storage-on-linux-servers)

### Windows

My condolences! Also, I don't run Windows on any of my VM's so I have no
experience with it.
