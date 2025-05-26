---
layout: post
title: Restricting Linux Logins to Specified Group
date: 2012-03-21 11:36:45.000000000 +00:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Systems (Administration)
tags:
- CentOS
- Identity management
- Lightweight Directory Access Protocol
- Linux
- Ubuntu
meta:
  _edit_last: '48492462'
  oc_commit_id: http://drone-ah.com/2012/03/21/restricting-linux-logins-to-specified-group/1332329808
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {'identity-management': {\"text\":\"Identity
    management\",\"slug\":\"identity-management\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/d4d701e1-8133-356a-87d5-37ba4b72dbf4/SocialTag/3\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Identity
    management\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'lightweight-directory-access-protocol': {\"text\":\"Lightweight Directory Access
    Protocol\",\"slug\":\"lightweight-directory-access-protocol\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/d4d701e1-8133-356a-87d5-37ba4b72dbf4/SocialTag/4\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Lightweight
    Directory Access Protocol\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'ubuntu': {\"text\":\"Ubuntu\",\"slug\":\"ubuntu\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/d4d701e1-8133-356a-87d5-37ba4b72dbf4/SocialTag/5\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Ubuntu\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'centos': {\"text\":\"CentOS\",\"slug\":\"centos\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/d4d701e1-8133-356a-87d5-37ba4b72dbf4/SocialTag/6\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"CentOS\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'linux': {\"text\":\"Linux\",\"slug\":\"linux\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/d4d701e1-8133-356a-87d5-37ba4b72dbf4/SocialTag/8\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Linux\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: '783'
  _wp_old_slug: '783'
  _publicize_job_id: '5185288577'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:54'
permalink: "/2012/03/21/restricting-linux-logins-to-specified-group/"
---

If you have linux boxes that authenticate over ldap but want logins for
specific boxes to be restricted to a particular group, there is a simple
way to achieve this.

Firstly, create a new file called /etc/group.login.allow (it can be
called anything - you just need to update the line below to reflect the
name)

In this file, pop in all the groups that should be able to login

    admin
    group1
    group2

Edit /etc/pam.d/common-auth (in ubuntu), it might be
called /etc/pam.d/system-auth or something else very similar. At the top
of the file (or at least above other entries, add the following line:

    auth required pam_listfile.so onerr=fail item=group sense=allow file=/etc/group.login.allow

For the record, found this little tidbit [over at the centos
forums](https://www.centos.org/modules/newbb/viewtopic.php?topic_id=25940 "Allow Only Specific LDAP Group Access (CentOS Forums)")\
