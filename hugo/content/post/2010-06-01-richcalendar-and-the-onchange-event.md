---
categories:
- Java (EE)
date: "2010-06-01T20:20:20Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:40"
  _publicize_job_id: "5185354074"
  _wp_old_slug: "408"
  oc_commit_id: http://drone-ah.com/2010/06/01/richcalendar-and-the-onchange-event/1275423620
  original_post_id: "408"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
published: true
status: publish
tags:
- calendar
- richfaces
- seam
title: rich:calendar and the onchange event
type: post
url: /2010/06/01/richcalendar-and-the-onchange-event/
---

I wanted to trigger some validation based re-renders with the `rich:calendar`
component. I was scratching my head for a while trying to figure out why it
wasn't working.

Then it happened, its supposed to be onchange**d**. This particular component
requires the extra d at the end... and it worked and everyone lived happily ever
after...
