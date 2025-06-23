---
title: Limit Java Memory Globally
date: 2024-07-26T11:18:35+01:00
categories: devex
tags:
  - java
---

Java can be greedy with RAM. By default, it grabs up to half your system memory,
which is fine for servers—but annoying on a dev machine.

You can tame this by setting:

```bash
export JAVA_TOOL_OPTIONS="-Xmx1G"
```

I dropped this into my .bashrc, and it instantly reduced background memory
pressure. One gig is plenty for most compile-and-run tasks during development.
