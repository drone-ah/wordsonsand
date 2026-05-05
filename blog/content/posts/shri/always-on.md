---
title: "Always On"
date: 2026-04-14T19:51:09+01:00
categories:
  - production
tags:
  - production
  - founder
  - decision-making
  - kraya
  - megabus
  - polskibus
  - burnout
---

I knew as soon as my phone rang what it was about.

It was the same every time. I would drag myself up to answer the phone - my
body, my mind screamed at me, but I had gotten good at overriding every instinct
through sheer willpower. I could hear the apologetic tone on the other side, and
I could recognise some of the voices after a while. I mustered up all of my
strength to be and sound as awake as possible. I needed to be professional even
if I was still in my underwear.

I had a glass of water on my bedside table. I'd pick that up and head to my
office in the spare room. The computer was always on and always ready to go —
like me I guess. I'd log on to the servers, and check the logs. If I can
identify which one fell out of the group, I can restart just that one. If I was
too late or if the issue had escalated, I'd have to restart the whole cluster —
shut them all down, give them a few seconds, then bring each one up, while
keeping an eye on them. I could do it half asleep after a while.

Falling back asleep wasn't a breeze either - I was tired - exhausted - but I was
now also wired. Waking up in the morning was harder - the alarm would go off and
my body would be limp. I still remember the sheer power of will to drag myself
into the shower, then carry on with the rest of the day.

Of 266 incidents over about two years, I answered 156.

I remember one particular night, though I do not remember how many times I'd
woken up beforehand. I was already tired.

megabus.com had gone offline. I got an alert. "But someone else is on call
tonight," I told them. "We already tried them twice," came the reply. I had to
deal with this. I had to deal with this.

I remember sitting at my desk at home working on fixing it. At some point,
something was different, though I don't remember what. While I was working on
fixing it, I remember being overcome with an overwhelming impulse to get up from
the chair and walk away — I almost imagined myself walking away. I resisted and
shut down that impulse. I fixed megabus as I had always done. In fixing megabus
though, something broke inside me, somewhere deep, in the very core of my being.
I was never the same again.

I analysed the system top to bottom, inside and out. I even waded through JVM
internals.

It got incrementally better, more stable. I think I rewrote every component that
wasn't the core ticketing system. In the end, what pushed it over the line were
two unexpected changes. Automated nightly restarts of each node in the cluster
and a rate limiter.

On the 10th December 2012, the system, now serving Polskibus, had a sale event.
We had a bank of screens on a wall with all the key stats for the system. It
looked cool, and we felt a bit like we were on a TV show. At peak, nearly 15,000
concurrent sessions — six or seven times the average. Over 30,000 bookings in a
single day, three times more than the normal amounts across all systems.

We watched it closely, all day. Nothing broke. Nothing screamed. Everyone
smiled, but there was no celebration.
