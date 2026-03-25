---
title: "It Gets Everywhere"
date: 2026-03-24T20:52:21Z
categories:
  - leadership
tags:
  - leadership
  - ways-of-working
  - production
  - decision-making
  - way-of-the-tortoise
---

In 1999, I was building websites in ASP (before there was .NET) and MSSQL
Server. We had a Windows NT server that I had to restart every week — not
because of updates, because it would get slower and slower until a restart was
the only thing that would fix it.

We had one ADSL connection coming into the office and three of us. I wanted to
share the internet. Windows NT didn't support it cleanly — it had a way, but it
was clunky enough that no internet was arguably better. We'd paid hundreds of
pounds for it.

I'd heard about Linux. Downloaded Red Hat, installed it, configured it for NAT.
It worked — it was like magic. I'm pretty sure I had to recompile the kernel to
get some bits working, but there were instructions and they were honest. It did
what it said.

Here was software that was completely free — free enough that I could read the
source code, make changes, run it however I wanted. It did more than the
hundreds of pounds worth of garbage sitting on the desk. And once I set it up, I
never had to restart it. Never. Compared to once a week on the NT box.

The difference, in my mind, was simple. Linux was built responsibly. NT was
built as a money-making enterprise.

That held for a long time. I moved to Debian, then celebrated when Ubuntu
arrived and made things more accessible. I've recently been able to abandon
Windows altogether — gaming on Linux is finally viable. I came back full time
and felt mostly at home.

But there were minor niggles. Things that felt slightly off but that I couldn't
quite name.

Then I started digging into systemd.

I remembered feeling odd about having to run specific commands to read logs. Odd
about one tool doing many different things — which ran contrary to the Unix
philosophy that had made Linux what it was. When I looked into the history of
the opposition to systemd, it was revelatory.

systemd becoming process 1 is, in a word, irresponsible. It makes everything
easier and more accessible, which is why it won. But unlike the Linux of old,
the tradeoff isn't visible upfront, and there's no real choice. The responsible
option isn't the default anymore — it's the thing you have to go looking for.

I thought I had already done the work. I thought I had found the alternative.

While I was celebrating Linux becoming mainstream, I hadn't considered what it
would cost.

The Linux ecosystem had started optimising for mainstream at the expense of
responsibility. It works now, for far more people. But it's a different thing
than it was. When linux was really taking off, there was a joke going around
(before memes were called memes) about Microsoft Linux. Turns out the joke was
on us!

It is always a tradeoff between security and convenience — something convenient
is rarely secure, and vice versa. I think something similar applies to
responsibility. The more accessible you make something, the harder it becomes to
hold the line on what it was built to do.

There was a time when software going wrong meant losing your work. Now it means
losing your money, your reputation, or — in a car, in a hospital — your life.

The context has changed. The attitudes haven't. And the places that once had
better attitudes — the ones built on responsibility, on craft, on caring about
the thing itself — are being pulled in the same direction. _It gets everywhere._

Do you want your car running Windows? What about systemd?
