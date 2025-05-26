---
layout: post
title: 'Drupal 7 + Services: Paging &amp; Filtering the index endpoint'
date: 2017-09-03 12:20:54.000000000 +01:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- drupal
tags:
- drupal-7
- drupal-services
meta:
  _edit_last: '48492462'
  geo_public: '0'
  _publicize_job_id: '8909860171'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:55:01'
permalink: "/2017/09/03/drupal-7-services-paging-filtering-the-index-endpoint/"
---

There are a lot of ways to manipulate the data returned by the index
endpoint. In this post, we are going to consider the node index
endpoint. By default, this endpoint returns all nodes sorted in
descending order of last update with 20 items per page.

You access the node index endpoint by going to

\[code\]http://\<domain\>/\<endpoint-path/node.json (or the alias given
to node in the resources section)\[/code\]

You can replace .json with with other extensions to get the same data in
different formats

To access the second page, you can use the page parameter

\[code\]

node.json?page=2

node.json?page=5

\[/code\]

    To change the number of items on each page, you need the "perform unlimited index" queries permission. You use the pagesize parameter to change it

\[code\]\
node.json?pagesize=100\
node.json?pagesize=50\
\[/code\]

To filter a field, you can use the parameters\[property\] where
\'property\' is the field on which you want to filter. It needs to be a
field on the node table, and not a drupal field as it does not do the
joins to pull in field data.

\[code\]

node.json?parameters\[type\]=blog_post

node.json?parameters\[type\]=article

\[/code\]

To apply a different filter than of equality, you can use
options\[parameters_op\]\[property\] where property is the same as
above.

\[code\]

node.json?parameters\[created\]=1431065220&options\[parameters_op\]\[created\]=\<

node.json?parameters\[changed\]=1431065220&options\[parameters_op\]\[changed\]=\>

\[/code\]

To return fewer fields, you can use fields and comma separate the
properties. Once again, you can only specify properties on the entity
(i.e fields on the base table)

\[code\]

node.json?fields=nid,changed

node.json?fields=nid,created,title

\[/code\]

you can sort the results by using
options\[orderby\]\[property\]=\<asc\|desc\>

\[code\]

node.json?options\[orderby\]\[nid\]=asc

node.json?options\[orderby\]\[created\]=desc

\[/code\]

You can also mix and match these separate options

\[code\]

node.json?page=10&pagesize=100&parameters\[type\]=blog_post&options\[parameters_op\]\[type\]=!=&fields=nid,changed

\[/code\]\
