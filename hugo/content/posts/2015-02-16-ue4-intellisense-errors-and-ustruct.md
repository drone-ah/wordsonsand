---
categories:
- Game Development
date: "2015-02-16T16:55:32Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:00"
  _publicize_job_id: "5186102430"
  _rest_api_client_id: "-1"
  _rest_api_published: "1"
parent_id: "0"
password: ""
status: publish
tags:
- UE4
title: '[UE4] Intellisense Errors and USTRUCT'
type: post
url: /2015/02/16/ue4-intellisense-errors-and-ustruct/
---

The Intellisense feature of Visual Studio, while useful is quite different from
the compiler and on occassion produces false positive errors. According to the
document about
[Setting Up Visual Studio for UE4](https://docs.unrealengine.com/latest/INT/Programming/Development/VisualStudioSetup/index.html),
this can happen for few possible reason

> - The IntelliSense compiler (EDG) is more strict than the MSVC compiler.
> - Some #defines are setup differently for IntelliSense than when building
>   normally.
> - C++ compiled by IntelliSense is always treated as 32-bit.

However, there is also the case where the Intellisense is just plain wrong.
USTRUCT() structures are just such a case. You will find that Intellisense
complains about definitions with the structures defined as USTRUCT.

This may not be a big deal, except for the fact these fales positives makes it
difficult to spot actual errors and I like to see a clean workspace with no
errors (not even false positives). Fortunately, there is a workaround. You can
include the USTRUCT definitions only if it is not being compiled by
Intellisense.

Fortunately, it is just the GENERATED_USTRUCT_BODY() line that is the culprit.
To get these errors to go away, instead of that single line, use the following

```c++
#ifndef __INTELLISENSE__ /* Eliminating erroneous Intellisense Squigglies
*/
GENERATED_USTRUCT_BODY()
#endif

```

That'll make those pesky squiggly lines to disappear and give you a nice clean
workspace :-D
