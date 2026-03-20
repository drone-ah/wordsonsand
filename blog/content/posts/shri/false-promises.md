---
title: "False Promises"
date: 2026-03-18T20:25:00Z
draft: true
---

<!--

**The enterprise software marketing trap**
    — you came from Apache/Linux/PostgreSQL, where the documentation told you what it did and it did that. JBoss/J2EE made promises that didn't survive contact with production. The lesson isn't "Java bad," it's about the difference between software you can trust at face value and software you have to earn understanding of the hard way.

**The nightly restart as a moment of clarity**
    — you resisted it because it felt wrong, then accepted it, then it worked. That moment of "enterprise software didn't mean what I thought it meant" is genuinely interesting. It's about the gap between architectural elegance and what actually keeps a system running.

**The on-call reality**
    — a system that randomly fails and requires a human to restart it is a fundamentally different kind of problem than one with diagnosable bugs. You were essentially absorbing the unreliability personally, for years.

-->

I built the first version of megabus.com in six weeks, by myself, on Apache, PHP
and PostgreSQL.

It launched. It worked. It scaled — not infinitely, but honestly, further than
it had any right to.

When the time came to rebuild it properly, I did what you do when you've
outgrown your tools. I looked at what the enterprise was using. J2EE. JBoss. The
full stack. The documentation was confident. The feature list was long.
Clustering, failover, distributed caching — it had answers for everything we'd
been patching around.

I was used to Apache and PostgreSQL. When those projects told you something
worked a certain way, it worked that way. You could trust the documentation like
a contract.

I assumed J2EE was the same kind of software. Better resourced, if anything.

It wasn't.

\[What actually happened in production — the random failures, no reproducible
cause, no clear fix\]

We brought in experts. We tried \[X\]. We tried \[Y\]. Each thing helped, a
little. None of it solved it.

Eventually someone suggested nightly automatic restarts.

I pushed back. A system that needs restarting isn't fixed — it's managed. That's
not a solution, it's an admission.

We did it anyway.

It worked.

That was the moment I understood that enterprise software didn't mean what I
thought it meant. It didn't mean reliable. It meant complex enough that the
instability becomes somebody's full-time job to absorb.

Apache didn't need me to absorb anything. It just ran.

\[What I took from this — about evaluating software, about marketing vs.
production reality, about what "enterprise" actually signals\]
