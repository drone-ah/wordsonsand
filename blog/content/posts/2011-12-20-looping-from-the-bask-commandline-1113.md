---
categories:
  - Systems (Administration)
date: "2011-12-20T13:32:08Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:53"
  _publicize_job_id: "5185288624"
  _wp_old_slug: "749"
  oc_commit_id: http://drone-ah.com/2011/12/20/looping-from-the-bask-commandline-1113/1324387936
  original_post_id: "749"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - systems-administration
  - bash
  - infinite-loop
  - scripting-languages
title: Looping from the bash commandline [1113]
url: /2011/12/20/looping-from-the-bask-commandline-1113/
---

I figured this out the other day from idle curiosity. There is occassionally the
need to have a never ending loop to be executed directly from the bash
commandline instead of writing a script.

I used this to run sl (yes sl, not ls - try it - I love it) repeatedly.

```
    $ while true; do ; done
```

for example

```
    $ while true; do sl; done
```

Bear in mind that this loop is infinite and there is no way to cancel out of it
except to kill of the terminal.
