---
layout: post
title: Automating Code Analysis with Hudson [1115]
date: 2011-12-28 15:34:52.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software Development
tags:
  - Continuous integration
  - FindBugs
  - Hudson
  - PMD
  - Static code analysis
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2011/12/28/automating-code-analysis-with-hudson-1115/1325086495
  oc_metadata:
    "{\t\tversion:'1.1',\t\ttags: {'static-code-analysis': {\"text\":\"Static
    code
    analysis\",\"slug\":\"static-code-analysis\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/31a0e4fb-31a6-359b-b950-17b1d47017c2/SocialTag/1\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Static
    code
    analysis\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'findbugs':
    {\"text\":\"FindBugs\",\"slug\":\"findbugs\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/31a0e4fb-31a6-359b-b950-17b1d47017c2/SocialTag/2\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"FindBugs\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'pmd':
    {\"text\":\"PMD\",\"slug\":\"pmd\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/31a0e4fb-31a6-359b-b950-17b1d47017c2/SocialTag/3\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"PMD\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'continuous-integration': {\"text\":\"Continuous
    integration\",\"slug\":\"continuous-integration\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/31a0e4fb-31a6-359b-b950-17b1d47017c2/SocialTag/4\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Continuous
    integration\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'hudson':
    {\"text\":\"Hudson\",\"slug\":\"hudson\",\"source\":{\"url\":\"http://d.opencalais.com/pershash-1/07845f25-13cd-3937-8b42-24eea2bd187c\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/em/e/Person\",\"name\":\"Person\",\"_className\":\"ArtifactType\"},\"name\":\"Hudson\",\"_className\":\"Entity\",\"rawRelevance\":0.857,\"normalizedRelevance\":0.857},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: "759"
  _wp_old_slug: "759"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:53"
permalink: "/2011/12/28/automating-code-analysis-with-hudson-1115/"
---

As part of
[setting up continuous integration and automated builds and source analysis](http://drone-ah.com/2011/12/28/hudson-jenkins-and-continuous-integration-1114/ "Hudson / Jenkins and Continuous Integration [1114]"),
the next step is to integrate in the source analysis parts.

To this end, I installed the following plugins:

Task Scanner\
FindBugs\
CheckStyle\
PMD

After restarting Hudson, there are a few additional configuration bits to
complete.

I added an additional build step and set the goal as

checkstyle:checkstyle findbugs:findbugs pmd:pmd

I then enabled the four plugins, saved, ran a build and et voila\... Just like
magic

 
