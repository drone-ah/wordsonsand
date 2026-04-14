---
title: "Did They Have a Problem That Year?"
date: 2026-04-14T10:24:26+01:00
categories:
  - production
tags:
  - leadership
  - production
  - decision-making
  - kraya
  - megabus
  - edfringe
---

2008 was a heck of a year for kraya, and for me. We were already operating
megabus.com in the UK, USA, and Canada, along with Oxford Tube, the sales
website for coach usa - all for Stagecoach.

We were also working on the fringe website. We integrated the website with the
brand spanking new ticketing system - which cost nearly £900k.

We were also hosting websites for Boots, Kellogg's Food Service and dozens of
other clients.

All of this was held together by three or four developers, two systems
administrators and me.

On the 13 June (incidentally, I got married on the same date years later), as I
was just getting ready for a wild night on the town, a call comes through -
which John answers.

I still remember them laughing and then doing a double take "oh, you're serious?
let me get Shri"

It was the fringe. We'd already known that they were having trouble with their
ticketing system. I'd even pitched in, made suggestions - looked at their code
to try and help, but none of that was enough. I expected an update.

They wanted to know if we could put together an interim booking system for them
over the weekend. I wasn't sure. I told them I'd speak to my team and get back
to them.

I wasn't involved with the work on the fringe up until this point. I knew very
little about it. I was focused on megabus.com. The US version of the site had a
big marketing campaign happening in a few days and that was what I was meant to
be focused on.

By the time I put the phone down, Chris, who had been the lead on the fringe
already had a answer. "We can do it!"

"But how - it's got to take more than a weekend - right?"

The fringe website was already built well and had a clean layer interfacing with
the new ticketing system. In fact, that was the bulk of the work that year.

So.. Chris told me - all we would have to do is to implement the functionality
within that thin layer, fattening it up.

He believed we could do it. I believed him.

I hopped in a cab, headed over to the fringe to talk it through. I didn't
promise them we'd be able to get something ready by Monday, but I promised we'd
do our best.

We were in the office on the weekend, writing code. I remember working on the
basket, sending diffs over email and generally having a good time.

I even had some megabus US fun to keep me entertained in the form of issues with
loading sheets - I was already in the office, so it was one step easier to fix.

One of the bits of functionality which took a surprising amount of time was the
seat allocation. None of us had to worry about that before - it was just
capacity management on megabus. For the fringe though, we had to allocate actual
seats with seat numbers and everything.

With the fringe we had multiple tables which all joined together (thanks
hibernate) to encode a tremendous amount of detail about the seating plans -
including their physical location on a map.

It was too much detail for us, so we had to simplify it all down to get it
working in the timeframe. We kept most of the rest of the structures intact to
keep the data migration easier once the ticketing system was fixed.

By Monday, we had each managed at the most 6 hours of sleep each of the previous
three nights. I still have vivid memories of a suit of armour that we put
together using packing material while we were waiting for bits of data or
details of logic.

I broke the MySQL replication at 01:24, fixed by 01:32

I remember the delirium setting in. Email sent to the client with "fun fun fun
fun fun fun fun" as the subject

There were random emails to my brother "I'm still here!"

I also remember making makeshift beds with bubblewrap to get a wee nap here and
there. We were all so exhausted - pumped up on coffee and nicotine.

Finally at 03:35 on the Tue email to client: "DONE DONE DONE DONE DONE NODE NODE
NODE NODE"

Then at 05:33, requesting a PostgreSQL server rebuild for megabus US for their
marketing campaign.

At 10am on Tuesday, the fringe is finally able to sell tickets. The website
promptly fell over from the load, but we nurse it back and it sells 65k+ tickets
in the first week.

It would be at least two more weeks before the ticketing system is fixed and
brought back in.

For the work we did for them that year, and the previous one, we effectively
only charged about 30% - because that's all they could afford. This year, we
asked if they could put our name on the website.

Over the next two weeks(while megabus US was on their marketing campaign), we
fought many battles. There were 750 duplicate bookings. Numerous customer
complaints (thanks to our name being on the website) - almost all of them
blaming us for the failure of the ticketing system. People did not understand
that we put in the interim one, not the one that failed.

Press releases went out from the fringe - only two credited us. Both misspelt
the company name. Both called us a web design company — which, we were not, had
never been, and had no interest in becoming.

In truth, I wanted to be a hero - I think we all did. What we really wanted was
an acknowledgement of what we had done - which was nowhere to be found. We got
paid though - at least for a part of our effort.

For many years after that, I would tell people with pride - "did you know - I
saved the fringe, back in 2008," which was inevitably met with something like
"oh, did they have a problem that year?"
