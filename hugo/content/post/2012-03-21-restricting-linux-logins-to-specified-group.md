---
categories:
- Systems (Administration)
date: "2012-03-21T11:36:45Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:54"
  _publicize_job_id: "5185288577"
  _wp_old_slug: "783"
  original_post_id: "783"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
published: true
status: publish
tags:
- CentOS
- Identity management
- Lightweight Directory Access Protocol
- Linux
- Ubuntu
title: Restricting Linux Logins to Specified Group
type: post
url: /2012/03/21/restricting-linux-logins-to-specified-group/
---

If you have linux boxes that authenticate over ldap but want logins for specific
boxes to be restricted to a particular group, there is a simple way to achieve
this.

Firstly, create a new file called `/etc/group.login.allow` (it can be called
anything - you just need to update the line below to reflect the name)

In this file, pop in all the groups that should be able to login

```
admin
group1
group2
```

Edit `/etc/pam.d/common-auth` (in ubuntu), it might be
called `/etc/pam.d/system-auth` or something else very similar. At the top of
the file (or at least above other entries, add the following line:

```
auth required pam_listfile.so onerr=fail item=group sense=allow file=/etc/group.login.allow
```

For the record, found this little tidbit
[over at the centos forums](https://www.centos.org/modules/newbb/viewtopic.php?topic_id=25940 "Allow Only Specific LDAP Group Access (CentOS Forums)")\
