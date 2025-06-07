---
layout: post
title: WARN - Missing artifact descriptor for XXX
date: 2013-01-31 20:38:48.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software Development
tags:
  - Apache Maven
  - Arquillian
  - Maven
  - Software
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2013/01/31/warn-missing-artifact-descriptor-for-xxx/1359664732
  restapi_import_id: 591d994f7aad5
  original_post_id: "901"
  _wp_old_slug: "901"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:57"
permalink: "/2013/01/31/warn-missing-artifact-descriptor-for-xxx/"
---

Working on an Arquillian test deployment which had some library changes
recently, I ran into the following error.

```
WARN - Missing artifact descriptor for org.javassist:javassist:jar:3.16.1-GA
```

The particular library was in the pom.xml dependency hierarchy but it was
resolving to an earlier  version. Maven was switched to offline mode during the
tests and I had never needed this version of the library before. This meant
 that the local version of my maven repository did not have jar and maven emits
it slightly unhelpful error. It would be better if it told us that it could not
find the artifact and since its in offline mode, can\'t go and retrieve it.

There are two quick hacks. Add the library and version into the pom.xml and do a
build. This will pull the library into your local repository and maven will be
able to find it offline. You could also just take maven in the tests online by
removing the goOffline() method call.

As for the cause of the issue, it stems from the way maven resolves dependencies
from Arquillian in comparison to building. I had updated a library version,
which now depends on the newer version of javassist. However, in considering all
the other things within the pom.xml, maven brings it down to an earlier version
when building.

However, the dependency resolution within the maven run through Arquillian
considers a slightly different set of requirements and resolves to a later
version of the lib which is not available.


