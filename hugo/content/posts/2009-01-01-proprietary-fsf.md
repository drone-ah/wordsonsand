---
categories:
- Philosophy
date: "2009-01-01T20:24:16Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:35"
  _publicize_pending: "1"
  _wp_old_slug: "126"
  oc_commit_id: http://drone-ah.com/2009/01/01/proprietary-fsf/1236796134
  original_post_id: "126"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- Freedom
- FSF
- GNU
- GPL
- GPLv3
- LGPL
- Open Source
- Open Source library
- Software
- software developers
- Technology
title: Proprietary FSF
url: /2009/01/01/proprietary-fsf/
---

I have always a big fan and proponent of the FSF and having recently been
interested in researching for a project came across a document covering
[Why you shouldn\'t use the Lesser GPL for your next library](http://www.gnu.org/licenses/why-not-lgpl.html "Why Not LGPL")

What the document basically suggests is to limit what proprietary software
developers can do by licensing libraries as GPL instead of LGPL.

This is no longer free(as in speech, not beer) software. Why?

Freedom means the ability to use something without restriction. If I cannot use
a library in a proprietary product, that is removing an important freedom.

This attitude is likely to alienate the \"commercial\" or proprietary developers
further from FSF/GNU.

<!--more-->

In fact, doing this is just not fair and not in line with how I view is the
concept behind the FSF. The point is to write software / libraries and share
that with the world so others may build upon what you have done. Stand on the
shoulders of giants in a way\...

It makes perfect sense for software to be GPL since you don\'t want somebody to
pick up a GPL software, build something on top, and sell it without source.

However, if libraries are released under the GPL instead of LGPL, it means that
I can not link against that library to write a non-GPL compatible application.

The [GNU Website](http://www.gnu.org/ "The GNU Operating Sytem") states

> \"Free software is a matter of the users\' freedom to run, copy, distribute,
> study, change and improve\"

Additionally, the
[Quick Guide to GPLv3](http://www.gnu.org/licenses/quick-guide-gplv3.html "A Quick Guide To GPLv3")
states that

> Nobody should be restricted by the software they use. There are four freedoms
> that every user should have:

- the freedom to use the software for any purpose,
- the freedom to change the software to suit your needs,
- the freedom to share the software with your friends and neighbors, and
- the freedom to share the changes you make.

This has always been my impressing of the purpose of GPL. Now, how does this
work with Libraries? A little differently\... :-(

From my perspective, if I have the freedom to use the \[library\] for any
purpose, that means that I can write an application that **_uses_** that library
without having to worry about licensing issue.

However, this is not the case. There is a clause that states that the software
cannot be used in a larger software project that has a license incompatible with
the GPL. This includes linking the library into another software application.

Therefore, I do not have the freedom to use the software **_for any purpose_**.

Freedom cannot be uni-directional. If GNU/FSF are trying to muscle out
developers of proprietary software, all they are doing is alienating themselves
further\...

I run a technology firm that uses a heck of a lot of open source software. In
fact, I am posting this from an ubuntu desktop running firefox from a VServer. I
am probably using a dozen open source applications to do this simple
straightforward act.

There is in fact, not a simple closed source application at any point through
this.

The main problem that I see with this is that it makes Open Source so much more
zealot(ous) and FSF, GNU and OSS becomes fundamentalists. The attitude is not
one of freedom and inclusion but of exclusivity and marginalisation.

The worst part is the price that is asked of developers who want to use an Open
Source library. The price is the acceptance and propogation of an idea (Freedom
or else).

Compared to the cost of conversion to another idealogy (Free Software Idealogy),
the cost of a few hundred, thousand, or even millions of dollars / pounds for a
piece of software seems dirt cheap.

I understand that each developer has the freedom to choose which license to use
for their products/libraries. My question is how can an organisation that claims
to be a proponent of freedom encourage the removal of freedoms?

I would like to ask how this shift is any different from religious fanatics who
tell you that their god is the one true god and there is nothing else.
