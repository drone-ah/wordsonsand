---
categories:
- Systems (Administration)
date: "2010-04-21T13:47:26Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:39"
  _publicize_job_id: "5185288651"
  _wp_old_slug: "362"
  oc_commit_id: http://drone-ah.com/2010/04/21/linux-bulk-search-and-replace/1271857655
  original_post_id: "362"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- bash
- bulk search
- computing
- find
- grep
- search and replace
title: Linux bulk search and replace
url: /2010/04/21/linux-bulk-search-and-replace/
---

Doing a bulk search and replace across a set of files is actually surprisingly
easy. sed is the key. It has a flag - i that will modify the files passed to it
in-place.

```bash
$ sed -e 's/TextToFind/Replacement/' -i file1 file2 file3
```

Tie this power with either grep `-l`. (Thanks to Steve for pointing out a
mistake in the following, now corrected)

```bash
$ grep -l TextToFind * |xargs sed -e 's/TextToFind/Replacement/' -i
```

or find

```bash
$ find . -exec sed -e 's/TextToFind/Replacement' -i {} ;
```

If there are multiple changes you want to make, just put them all into a file
and pass it in via the -f flag.

file: replacements.patterns

```sed
s/TextToFind1/Replacement1/
s/TextToFind2/Replacement2/
s/TextToFind3/Replacement3/
```

and the command, using find to iterate through all files in the current
directory and subdirectories.

```bash
find . -exec sed -f replacements.patterns -i {} ;
```

et voila - hope it helps.
