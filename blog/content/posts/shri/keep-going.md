---
title: "I Chose to Keep Going"
date: 2026-04-14T20:52:10+01:00
categories:
  - production
tags:
  - production
  - founder
  - decision-making
  - kraya
  - stagecoach
  - megabus
  - burnout
---

In 2008, we all watched Pivotal crash and burn. They'd taken a year and nearly
£900k to build a new ticketing system for the fringe. On launch, they realised
that it could serve only one customer at a time.
[We built an interim ticketing system for them over the weekend](./saving-the-fringe.md).

It was time for kraya to take a leap. We should take the megabus.com ticketing
system to the next level and build a distributed ticketing system that could
potentially scale infinitely.

I'd picked JBoss because it was backed by Red Hat - it had all these features
and capabilities - a lot of which we needed. The other option we considered was
Glassfish, but it just didn't have the features we needed. There were other
options but that involved prohibitive licensing fees.

We costed it out at £650k and a year. It was unrealistic but I could not imagine
Stagecoach paying more, or giving us more time. They wanted it for £500k and in
six months. I should've pushed back, but we'd built a booking system over the
weekend, only a few months back - this should be possible, right?

It wasn't.

We got a year in the end — we didn't ask for it, and we didn't know until we
were most of the way down the path. When I heard about the delay it was a mix of
relief and regret. We'd burned through most of the budget on getting in
contractors who were leaving imminently. We could have got fewer people on
board, but as permanent staff.

We launched to Canada first, and that wasn't too bad. Then it was the US, and
the nightmare started. The UK launched last on my birthday in 2010.

In the intervening years, we had a budget shortfall of £150k and because we'd
rushed to build the product, we'd cut corners, and all the contractors had left.
We even had to let go of some of the permanent staff. I asked Stagecoach if
they'd fund us the extra £150k as we'd asked at the start. They said no. They
renegotiated the contract. They demanded more oversight - and increased our
reporting obligations.

Load testing was already a part of our process and we tested the new system
under load. We identified issues and fixed them. On paper, it looked good.

It wasn't.

The problem wasn't load per se. It was the stability of the system. It struggled
to stay up for extended periods of time. The more nodes there were, the worse it
was.

It was exactly the kind of problem that was hard to replicate and hard to test.
The main option we had was to think through what all it could be and to take
stabs in the dark. I put together a hit list and worked through it methodically.
Convincing Stagecoach to spend the money was sometimes harder than solving the
problem.

Ultimately, the one thing that pushed us over the line in terms of stability was
staggered nightly automated restart of each node.

I thought I was done working with
[software that needed regular restarts to stay functional](./it-gets-everywhere.md)
and at first resisted this. We are running enterprise level software, and that
too on Linux. If I was comfortable with such shenanigans, I might have stayed
with Windows. But we were running out of options. It was a hail Mary - it
worked.

I had been working with Linux and related software for years by that point. I,
in fact had servers that had not been restarted for literal years at that point,
with services that were running just as long. A lot of these services didn't
even need a restart on config change - just a reload.

None of these services even had a paid tier. That should have been the clue.

I cannot imagine having to restart PostgreSQL nightly or even Apache. I still do
not understand how something slated for the enterprise market can have leaks
that would warrant a regular restart to keep it working.

Years later, when I looked up issues around JBoss, I realised that it was
notorious for a whole slew of problems with JGroups and clustering. I remember
scouring the internet for details on any kind of issues and coming up empty.

I'd fallen into a marketing trap. JBoss was no Apache — it was the commercial
product of Red Hat. I thought all enterprise open source software would be of
the same calibre. It wasn't.

From what I understand, Stagecoach spent millions and maybe two years building
the whole ticketing system inhouse. kraya limped along for a few more years
before being shuttered.

The people involved, though, fortunately seem to have gotten through it largely
unscathed. It gives me a great deal of joy to see so many of the juniors I'd
hired now CTO's, VP's, Directors.

I believed that every victory was ours but when something went wrong, it
belonged to me. After all, I made every choice. I chose to pursue the Java EE
ticketing system. I chose JBoss. I chose to keep going.

I was in a narrowing path, with fewer and fewer options.

I knew I was the only one holding that line - I didn't know that there was
another option.

Bit by bit, I lost all sense of what I was paying to fix these mistakes.
