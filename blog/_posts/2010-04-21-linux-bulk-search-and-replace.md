---
layout: post
title: Linux bulk search and replace
date: 2010-04-21 13:47:26.000000000 +01:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Systems (Administration)
tags:
- bash
- bulk search
- computing
- find
- grep
- search and replace
meta:
  _publicize_job_id: '5185288651'
  _edit_last: '48492462'
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {'bulk-search': {\"text\":\"bulk search\",\"slug\":\"bulk-search\",\"source\":{\"_className\":\"Entity\",\"url\":\"http://d.opencalais.com/genericHasher-1/1fa5711a-3115-34ed-baee-388d6ef2e5fc\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/em/e/IndustryTerm\",\"name\":\"IndustryTerm\"},\"name\":\"bulk
    search\",\"rawRelevance\":0.857,\"normalizedRelevance\":0.857},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'search-and-replace': {\"text\":\"search and replace\",\"slug\":\"search-and-replace\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'bash': {\"text\":\"bash\",\"slug\":\"bash\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'grep': {\"text\":\"grep\",\"slug\":\"grep\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'find': {\"text\":\"find\",\"slug\":\"find\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'computing': {\"text\":\"computing\",\"slug\":\"computing\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  oc_commit_id: http://drone-ah.com/2010/04/21/linux-bulk-search-and-replace/1271857655
  restapi_import_id: 591d994f7aad5
  original_post_id: '362'
  _wp_old_slug: '362'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:39'
permalink: "/2010/04/21/linux-bulk-search-and-replace/"
---

Doing a bulk search and replace across a set of files is actually
surprisingly easy. sed is the key. It has a flag - i that will modify
the files passed to it in-place.

    $ sed -e 's/TextToFind/Replacement/' -i file1 file2 file3

Tie this power with either grep -l . \[Thanks to Steve for pointing out
a mistake in the following, now corrected\]

    $ grep -l TextToFind * |xargs sed -e 's/TextToFind/Replacement/' -i

or find

    $ find . -exec sed -e 's/TextToFind/Replacement' -i {} ;

If there are multiple changes you want to make, just put them all into a
file and pass it in via the -f flag.

file: replacements.patterns

    s/TextToFind1/Replacement1/
    s/TextToFind2/Replacement2/
    s/TextToFind3/Replacement3/

and the command, using find to iterate through all files in the current
directory and subdirectories.

    find . -exec sed -f replacements.patterns -i {} ;

et voila - hope it helps.
