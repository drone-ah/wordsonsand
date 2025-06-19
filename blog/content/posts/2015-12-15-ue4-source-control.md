---
categories:
  - Game Development
date: "2015-12-15T10:35:17Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:00"
  _publicize_job_id: "5185637349"
  _rest_api_client_id: "-1"
  _rest_api_published: "1"
parent_id: "0"
password: ""
status: publish
tags:
  - UE4
title: "[UE4] Source Control"
url: /2015/12/15/ue4-source-control/
---

As with any software project, it is important to use some form of source control
solution. For most software projects, there are a number of good solutions out
there. However, in the world of game development, most of these are not viable
since they won't handle binary files very well, and unreal engine (as most
games) will have a large amount of binary resources.

Perforce is a good tool for small teams since it is free for teams with less
than 20 developers.

Another thing that can be confusing is what files/folders to add into the source
control. Generally, we do not want to include any files which can be
auto-generated (e.g. builds) or are transient (e.g. logs). You should generally
include the following folders into source control

- Config
- Content
- `Intermediate/ProjectFiles/<projectname>.vcxproj\*` (the .vcxproj.user file may
  not be relevant if there are multiple developers)
- Source
- `<projectname>.sln`
- `<projectname>/.uproject`

I found it odd that the project file is in the intermediate folder since one
wouldn't intuitively think to include it in source control
