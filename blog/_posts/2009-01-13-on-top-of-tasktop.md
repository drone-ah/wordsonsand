---
layout: post
title: On top of Tasktop
date: 2009-01-13 13:43:56.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software
tags:
  - Context Management
  - Eclipse
  - Firefox
  - Linux
  - Mylyn
  - Tasktop
  - Time Tracking
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  restapi_import_id: 591d994f7aad5
  original_post_id: "133"
  _wp_old_slug: "133"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:36"
permalink: "/2009/01/13/on-top-of-tasktop/"
---

My post about
[tracking time](http://drone-ah.com/2008/12/13/your-time/ "Your Time [words on sand]")
attracted the attention of [Tasktop](http://tasktop.com/ "Tasktop"). While this
had been mentioned to me before, I was **mistakenly** under the impression that
this was a windows only app.

I was pleased to find out that this was also available for linux. Great\... Lets
try it out.

First stumbling block is the requirement to register on the website before I can
download a trial. I am a firm believer of try before you buy. I should be able
to register but it should be entirely my choice.

I am more comfortable with registering before buying or for the use of a free
piece of software. However, registering for a trial always irritates me. This
was also the case when I wanted to trial InDesign / Illustrator the other day.

<!-- more -->

After registering, there was the irritating wait for the email to arrive. Now,
this is irritating. When I want something, I want it **_NOW_**. I hate waiting.
Adobe did not make me wait for the confirmation email of registration before
downloading the trials. There are two good reasons as to why this irritates me.

1.  Email, as reliable as it is generally, can take time. In theory, this can be
    anywhere from a few seconds to hours. How about if my mail server is
    currently down. Or even more importantly, what if I have shut down my mail
    client so that it does not keep distracting me from something that I am
    trying to do. Opening up my mail client, I now want to find out about the
    other emails that are in my inbox and whether any of them require an
    action\...
2.  I have reluctantly provided details about myself. Confirming my email
    address before I am allowed to download a trial suggests that Tasktop does
    not trust me enough to just let me download the trial. The software has
    started off on the wrong foot. How much of an issue is it really if someone
    gave the wrong details before downloading a trial. Is it really that
    important that you are able to keep bugging them via email to buy the
    product?

I was curious enough to jump through the hoops to download the product. The
first thing I noticed is that there is no 64bit for Linux :-(. More steps
involved in installing this on my 64bit machine. So instead, I installed it one
of my 32bit machines - save time.

Once the download completed, the steps on the website suggested that I needed to
configure it (with ./configureTasktop.sh) and then run Tasktop. The
configuration step required no input from the user and outputted nothing. I have
to ask:

1.  Why is the configuration step not integrated into Tasktop and configured to
    run once? Alternatively,
2.  Why does the configuration step, not start Tasktop right after.
3.  Even better: Make Tasktop a symlink to configureTasktop.sh, which then
    relinks that to the Tasktop Binary with the configureTasktop running Tasktop
    right after. This means that from the users perspective, they are always
    running the same command, and you save any cost associated with run once
    checks.

I finally got Tasktop to run and it asks me if I want to install the firefox
addon to integrate with Tasktop. I want to see how it integrates, so I do. Of
course, this is yet **another** step.

A restart later, I was ready to try out Tasktop - or was I? We use bugzilla to
track tasks and I wanted to integrate that in similar to how I do it in Eclipse.
This was also trickier than I expected.

I went into the partner connectors section which did not cover bugzilla, which I
assumed meant that it came with Bugzilla integration by default. This is true
but how the hell do I get there to configure it. It took me a little while to
find the configuration section (there are no menus). Once I was there, I wanted
to get back to the original layout which was tricky since the \"close
configuration\" button was nicely hidden away up at the top right.

Once I had this working, I tried out the active/deactive mechanisms and this
works just the same as in Eclipse. Except with the Firefox plugin, it adds in
the links that you browse as part of your context - GREAT!

Add in a task to blog about it and went through writing half the document, then
decided to de-activate it before I started working on something else. All the
firefox tabs were closed - again, great\...

The problem is that when you re-activate the context, it just clears the tabs in
firefox and shows you the links you last had open. The page titles for the pages
that I had open were the same for a few, so going through them trial and error
to get to the blog post was tricky. More importantly, the cookie was already
gone and I had to re-login. This might be a timeout issue with Wordpress so wont
tag that against Tasktop.

I haven\'t tried linking folders / files yet but considering that with the above
process taking me more time than I expected due to the sheer number of steps
involved, I shall have to leave that to another day. In all honesty, it might
never happen.

I do like the time logging feature of Tasktop as it tells me which tasks I spent
my time on in different chart formats. This is great. However, I have a problem
in that this is on an individual basis. I see nothing on here about how a team
leader can link in Tasktop used by the team to calculate total time spent on a
project / task. This is a necessary feature for a tool like this in the team
environment.

It is possible that all of this is easier in a windows environment. Possibly
because it was built on there, but more likely because Windows users are used to
taking several steps to achieve something (what is it - 7 clicks to delete a
file in Vista?)

Having ranted on for a while, dont get me wrong. I think that Tasktop is a
fantastic concept and with a bunch of tweaking can be a very intuitive tool to
use. However, at the stage that it is in, it does not do what I need it to do.
It is actually more obtrusive than useful (e.g. by removing all my tabs from
firefox when switching out of a context and not re-instating them on going back
to the context).

Then, it is probably just because I simply expect too much\... :-(
