---
categories:
- drupal
date: "2017-09-03T12:20:54Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:01"
  _publicize_job_id: "8909860171"
  geo_public: "0"
parent_id: "0"
password: ""
published: true
status: publish
tags:
- drupal-7
- drupal-services
title: 'Drupal 7 + Services: Paging &amp; Filtering the index endpoint'
type: post
url: /2017/09/03/drupal-7-services-paging-filtering-the-index-endpoint/
---

There are a lot of ways to manipulate the data returned by the index endpoint.
In this post, we are going to consider the node index endpoint. By default, this
endpoint returns all nodes sorted in descending order of last update with 20
items per page.

You access the node index endpoint by going to

`http://<domain>/<endpoint-pathnode.json` (or the alias given to node in the
resources section)

You can replace .json with with other extensions to get the same data in
different formats

To access the second page, you can use the page parameter

```
node.json?page=2
node.json?page=5
```

To change the number of items on each page, you need the "perform unlimited
index" queries permission. You use the `pagesize` parameter to change it

```
node.json?pagesize=100
node.json?pagesize=50
```

To filter a field, you can use the parameters[property] where 'property' is the
field on which you want to filter. It needs to be a field on the node table, and
not a drupal field as it does not do the joins to pull in field data.

```
node.json?parameters[type]=blog_post
node.json?parameters[type]=article
```

To apply a different filter than of equality, you can use
options[parameters_op][property] where property is the same as above.

```
node.json?parameters[created]=1431065220&options[parameters_op][created]=<
node.json?parameters[changed]=1431065220&options[parameters_op][changed]=>
```

To return fewer fields, you can use fields and comma separate the properties.
Once again, you can only specify properties on the entity (i.e fields on the
base table)

```
node.json?fields=nid,changed
node.json?fields=nid,created,title
```

you can sort the results by using options`[property]=<asc|desc>`

```
node.json?options[orderby][nid]=asc
node.json?options[orderby][created]=desc
```

You can also mix and match these separate options

```
node.json?page=10&pagesize=100&parameters[type]=blog_post&options[parameters_op][type]=!=&fields=nid,changed
```
