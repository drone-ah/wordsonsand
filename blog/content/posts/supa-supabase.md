---
title: "Supa Supabase"
date: 2025-09-23T12:09:57+01:00
tags:
  - postgresql
  - supabase
---

I have always been a fan of [PostgreSQL](/tags/postgresql). I picked it for
megabus.com and used it when the system grew from a few hundred orders a day to
tens of thousands each day.

Back in the late noughties, when MySQL was getting popular, people would often
ask me why I was picking Postgres. MySQL was so popular and so fast and it had
cool things like query caching (with postgres did not have). My answer was
simple - postgresql was a "real database system." I remember being shocked when
I trying to use transactions and it just ignored it. Postgresql was also really
good at failing gracefully under high load while MySQL had a bad habit of just
stopping processing any requests and becoming unresponsive.

This was more than 15 years ago now so things may have changed with MySQL,
though its acquisition by Oracle was certainly a bad sign.

Years on, and it would seem that people are better informed as to the benefits
of postgresql.

I am currently working on a small game-like social experiment app and it
requires persistent storage. I was intially considering firebase but the costs
put me off. In the world of cloud, I didn't think that spinning up a postgresql
server was really an option - or is it?

## Supabase

I ran into Supabase as a more cost-effective option and as a bonus it's open
source and based on PostgreSQL. I'll admit that I was skeptical that it would
embody the postgresql philosophy that I knew and loved. I expected another
capitalistic effort at monetising open source.

I was pleasantly surprised. I am in the very early stages of using it, but so
far, I love it. Not a huge fan its love of javascript, but the whole world seems
to be a big fan (I am not - but then I wasn't a fan of MySQL either ;))

I didn't end up going further with it, but the early impression was good enough
that I'd reach for it again.
