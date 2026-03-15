---
title: "Priced In"
date: 2026-04-07T09:55:24+01:00
categories: microblogging
tags:
  - microblogging
  - production
  - decision-making
---

Two ticketing systems. Same client. Same payment provider. We were moving fast —
that was the explicit choice, theirs and mine. The kind of fast where you know
something will go wrong eventually and you price it in rather than try to
prevent it.

We'd done the sensible thing and shared the payment code between them — DRY,
less surface area for error, obvious call.

Then the larger system needed PostAuth. We updated the code, added a scheduled
task to catch anything the non-deterministic bits missed, moved on.

A few months later: why has no money come through on the smaller system?

We'd ported the PostAuth flow across when we updated the shared code. We hadn't
added the scheduled task. The payment provider, chosen for cheap and cheerful
rather than reliability, failed silently rather than erroring. The accounting
department, running at the same pace as everyone else, hadn't caught the gap.

Four separate things had to go wrong simultaneously. Any one of them holding
would have meant no loss at all.

The client lost money. Not a catastrophic amount, but real money. I braced for
the call.

> Try and let me know the next time you decide to run a sale.

He already knew the cost. He'd known before the mistake happened.
