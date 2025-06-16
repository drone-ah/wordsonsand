---
categories:
- Java (EE)
date: "2011-06-21T14:40:45Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:50"
  _publicize_job_id: "5181426973"
  _wp_old_slug: "681"
  oc_commit_id: http://drone-ah.com/2011/06/21/register-attach-service-to-jmx/1308663651
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {}\t}"
  original_post_id: "681"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
published: true
status: publish
tags: []
title: Register / Attach Service to JMX
type: post
url: /2011/06/21/register-attach-service-to-jmx/
---

Registering a Bean within JMX (at least in JBoss) is very straightforward. It
requires an interface with attributes (getters and setters) and operations.

```java
MBeanServer server = org.jboss.mx.util.MBeanServerLocator.locateJBoss();
ObjectName objectName = new ObjectName("jboss.cache:service=TcpCacheServer");
server.registerMBean(objectToAttach, objectName);
```

objectToAttach is an object with a JMX\'able interface.
