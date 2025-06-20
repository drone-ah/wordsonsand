---
categories:
  - Systems (Administration)
date: "2020-06-25T08:47:45Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:04"
  _publicize_job_id: "45837663329"
  timeline_notification: "1593074865"
parent_id: "0"
password: ""
status: publish
tags:
  - systems-administration
  - bash
title: How to fix terminal title after disconnecting from ssh
url: /2020/06/25/how-to-terminal-title-after-disconnecting-from-ssh/
---

For some reason, ssh does not clean up after itself in terms of updating the
terminal title when you disconnect.

Here is a simple solution, a combination of
<https://unix.stackexchange.com/a/341277/25975> and
<https://unix.stackexchange.com/a/28520/25975>

Add the following functions into your `~/.bashrc` It will push the current title
and icon into a stack and pop it afterwards.

```bash
function ssh()
{
    # push current title and icon to stack
    echo -ne '\e[22t'
    # Execute ssh as expected
    /usr/bin/ssh "$@"
    # revert the window title after the ssh command
    echo -ne '\e[23t'
}
```

Restart bash / log out and back in, and it should work.

For security reasons, it is not possible to query the current title of the
terminal. However, with the following command, you can push the current one on
to a stack

```bash
echo -ne '\e[22t'
```

The title can then be set to anything, by ssh for example. You can then pop that
back from the stack using

```bash
echo -ne '\e[23t'
```
