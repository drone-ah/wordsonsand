---
title: "What's Circuit Breaker?"
date: 2026-05-06T10:53:27+01:00
categories:
  - production
tags:
  - production
  - founder
  - decision-making
  - kraya
  - megabus
  - circuit-breaker
---

I'd not been interviewed very often. For my first job, sure, I sent around my CV
to a bunch of places and did a bunch of interviews.

I then started my own company and ran that for nearly two decades, everything
largely self-taught. After that, because of the peculiar shape of my CV, and
"falling between a lot of stools" as a friend pointed out, interviews were...
tricky!

I remember one such interview a few years ago. He was good at diffusing any
anxiety and it was a pretty relaxed interview. He wanted to know if I knew
architectural design patterns - a term I hadn't heard before. He gave me
examples, circuit breaker, pub / sub and a third one I can't remember anymore. I
recognised pub / sub and the other one I think - but was not aware of circuit
breaker. He explained it to me - something about reducing the risk of a
downstream service failing.

It made sense - but I couldn't remember any instance where I'd done that and I
said as much.

No big deal! We moved on.

As you may have noticed from my recent flurry of posts, I've been going through
a bit of an excavation of all things kraya, and discovering little nuggets that
I'd forgotten all about.

Claude code went through all of my emails, all the code repositories and put
together interesting things I had done. The purpose was mainly to pick out
things that could be good blog posts, stories, or just reminders.

Claude had found two circuit breakers.

The first one was built by the seat of my pants in 2004. There was an active
marketing campaign on megabus and the system was struggling. I'd plumbed the
depths of the tech stack - the web servers, the database server to get every
last ounce of performance out of the system.

What I really needed was to slow down the deluge of people coming into the
site - just a little bit. I needed to prevent the snowball effect, and I tried a
simple way to achieve it. I wrote the following around my 23rd birthday.

```php
$timestamp = time() - 120;
$sTimestamp = date("d-F-Y H:i");
$sql = "SELECT count(*) FROM tSearches WHERE Script_Start > '$sTimestamp' AND Script_End IS NULL;";

$numSearches = $dbh->getOne($sql);
if ($numSearches > 10)
    sleep(5);

$numSearches = $dbh->getOne($sql);
if ($numSearches > 10)
    sleep(5);

$numSearches = $dbh->getOne($sql);
if ($numSearches > 10)
    sleep(5);
```

(The eagle eyed among you might notice a bug in the above code. It wasn't
resolved for years. Coding live on a production system tends to create bugs. See
if you can find it before reading on.)

(Regular readers may also remember a part of this story from
[I Know People Like You](./simple-wins.md))

The system already tracked searches - so I just needed to check it.

This bit of code checks to see how many of the searches that started in the last
two minutes were incomplete. Except that's not what it did - I forgot to
actually use `$timestamp` so it only checked the current minute's worth.

I considered failing and letting the user retry manually - but I didn't want to
force the user to take action unless absolutely necessary. This one waited up to
15 seconds and then let the search happen anyway.

If the user was still waiting after 15 seconds, might as well try and do a
search and see what happens. In hindsight, it might have been better to fail at
that point. If there were too many searches after 15 seconds, the database
server was likely already snowballing.

This bit of code survived through to the end of the PHP codebase apart from
minor tweaks. The Java version had layers of circuit breakers. Session limits,
rate limiters and threadpool configuration for scaling up. None of it though,
was called a circuit breaker - not by me or by the team.

So, have I built a circuit breaker? Thinking back, what threw me off was the
wording. I had the feeling that the circuit breaker would "protect" another
service.

In my mind, the database wasn't another service - it was a part of the same
service.
