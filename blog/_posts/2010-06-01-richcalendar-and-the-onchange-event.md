---
layout: post
title: rich:calendar and the onchange event
date: 2010-06-01 20:20:20.000000000 +01:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Java (EE)
tags:
  - calendar
  - richfaces
  - seam
meta:
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2010/06/01/richcalendar-and-the-onchange-event/1275423620
  oc_metadata:
    "{\t\tversion:'1.1',\t\ttags: {'seam':
    {\"text\":\"seam\",\"slug\":\"seam\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'richfaces':
    {\"text\":\"richfaces\",\"slug\":\"richfaces\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'calendar':
    {\"text\":\"calendar\",\"slug\":\"calendar\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: "408"
  _wp_old_slug: "408"
  _publicize_job_id: "5185354074"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:40"
permalink: "/2010/06/01/richcalendar-and-the-onchange-event/"
---

I wanted to trigger some validation based re-renders with the rich:calendar
component. I was scratching my head for a while trying to figure out why it
wasn\'t working.

Then it happened, its supposed to be onchange**d**. This particular component
requires the extra d at the end\... and it worked and everyone lived happily
ever after\...
