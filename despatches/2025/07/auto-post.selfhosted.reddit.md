---
publishDate: 2025-07-02 13:35:00+01:00
subreddit: selfhosted
title:
  I automated sharing my Hugo posts to Reddit and Bluesky using frontmatter,
  GitHub Actions, and Python
type: reddit
---

I run a Hugo-powered site and post devlogs, code, and videos regularly. Every
time I finish something, I want to share it everywhere — YouTube, Reddit,
Bluesky — while I'm still in the zone.

But that’s the worst time to share anything for visibility. I wanted to automate
it. I figured this would be a two-hour job.

It took two days.

Challenges included:

- fighting OAuth with Reddit
- hallucinated libraries for Bluesky
- facet formatting madness
- giving up on `rules_python`
- bouncing between Python and Go
- coming back to Poetry with tail between legs

Eventually I wired together:

- Hugo's `publishDate`
- a despatch script in Python
- a GitHub Action that runs every 30 minutes on weekdays
- a commit that marks posts as `publishedAt` to avoid reposts

I wrote up everything I learned, and shared code, curl samples, and a few
mistakes here:

https://drone-ah.com/2025/07/01/automated-posting-to-bluesky-reddit/

Would you do anything differently? Is there anything I didn't think about?

