---
layout: post
title: Register / Attach Service to JMX
date: 2011-06-21 14:40:45.000000000 +01:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Java (EE)
tags: []
meta:
  _publicize_job_id: "5181426973"
  oc_commit_id: http://drone-ah.com/2011/06/21/register-attach-service-to-jmx/1308663651
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {}\t}"
  _edit_last: "48492462"
  restapi_import_id: 591d994f7aad5
  original_post_id: "681"
  _wp_old_slug: "681"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:50"
permalink: "/2011/06/21/register-attach-service-to-jmx/"
---

Registering a Bean within JMX (at least in JBoss) is very straightforward. It
requires an interface with attributes (getters and setters) and operations.

```java
MBeanServer server = org.jboss.mx.util.MBeanServerLocator.locateJBoss();
ObjectName objectName = new ObjectName("jboss.cache:service=TcpCacheServer");
server.registerMBean(objectToAttach, objectName);
```

objectToAttach is an object with a JMX\'able interface.
