---
layout: post
title: Looping from the bash commandline [1113]
date: 2011-12-20 13:32:08.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Systems (Administration)
tags:
  - bash
  - Infinite loop
  - Scripting languages
meta:
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2011/12/20/looping-from-the-bask-commandline-1113/1324387936
  oc_metadata:
    "{\t\tversion:'1.1',\t\ttags: {'scripting-languages': {\"text\":\"Scripting
    languages\",\"slug\":\"scripting-languages\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/98b0861c-aea2-3a5d-8467-ccaa920b7a95/SocialTag/5\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Scripting
    languages\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'infinite-loop': {\"text\":\"Infinite
    loop\",\"slug\":\"infinite-loop\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/98b0861c-aea2-3a5d-8467-ccaa920b7a95/SocialTag/6\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Infinite
    loop\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'bash':
    {\"text\":\"bash\",\"slug\":\"bash\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: "749"
  _wp_old_slug: "749"
  _publicize_job_id: "5185288624"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:53"
permalink: "/2011/12/20/looping-from-the-bask-commandline-1113/"
---

I figured this out the other day from idle curiosity. There is occassionally the
need to have a never ending loop to be executed directly from the bash
commandline instead of writing a script.

I used this to run sl (yes sl, not ls - try it - I love it) repeatedly.

    $ while true; do ; done

for example

    $ while true; do sl; done

Bear in mind that this loop is infinite and there is no way to cancel out of it
except to kill of the terminal.
