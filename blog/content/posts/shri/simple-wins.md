---
title: "I Know People Like You"
date: 2026-03-31T10:41:29+01:00
categories:
  - engineering
tags:
  - leadership
  - production
  - decision-making
  - accountability
---

A few years ago, I was interviewed for a role. I was talking about a ticketing
system I'd built - originally in Spring, then rewritten to use EJB 3.2. The
interviewer didn't look impressed.

The team had already written a lot of stuff in Spring - but I really did not
like it. There was all this XML all over the place which was annoying, but what
I really didn't like was that the code and configuration for each component was
spread out all over the place. It meant that to understand how something worked,
I had to go hunting. Eventually, I got sick of it, and ported it to EJB myself.

Later in the same interview, he said: "I know people like you - you come in,
shake things up and get things done - but that's not what I'm looking for."

He was right. I understand that. But I've been thinking about what he was
actually describing.

When megabus.com was still a PHP site, search was the problem. It returned
quickly when the database was healthy and crawled - and slowed down when the
database was struggling. The load came in spikes. Even within a minute, there
were peaks and troughs.

My fix was simple. Before the search query ran, I added a small SQL check: how
many queries are currently active on the database server? If too many, wait a
second and try again. A few retries, then send it anyway.

A rate limiter baked into a search algorithm, written live on a production
server.

There were edge cases to consider, not to mention the load the rate limiter
would add to the database server. I knew though that if it broke, I could fix
it - I could just remove it - live, if needed. Not having the rate limiter was
at the time, more expensive than having it.

It worked. It got us through more than one hump.

The database was still the ceiling. We were on PostgreSQL 7 - no replication
support. Getting a more powerful server was possible but disproportionately
expensive. So I built something.

Two database servers. All writes went to both. Reads were distributed randomly
between them. Everything funnelled through one section of code.

I didn't do this live. I tested it. I knew what failure looked like: if the
servers diverged badly enough, I'd pick a primary and reset the other. That was
the contingency. It wasn't a safety net someone else would pull - it was mine.

The data integrity held better than I expected. Under very high load there were
edge cases - ticket IDs for the same customer could be in a different order
across the two servers on a return purchase - but because the IDs were
consistent within each server, it never caused a real problem. It held the fort
until I could replace it with something better.

I occasionally lie awake at night imagining the databases diverging and figuring
out how I would fix it.

I picked PostgreSQL over MySQL when MySQL was the obvious choice. Under heavy
load it stays up — it slows to a crawl, but it keeps going. And it had
transactions. I was building an ecommerce site; I needed transaction support.
MySQL was fast and popular. It also had a habit of giving up under sustained
load. I still pick PostgreSQL - but nowadays, so do most other people.

The thing these decisions had in common was that I was the person who'd be
fixing them at 3am if they went wrong. When you're personally accountable for
the consequences, the risk calculus changes. You think harder about what failure
looks like. You build the contingency before you go live. You know which
direction to pull if it goes sideways.

Caution that's never personally tested isn't rigour. It's consequence-avoidance
dressed up as responsibility.

"I know people like you - you come in, shake things up and get things done - but
that's not what I'm looking for."
