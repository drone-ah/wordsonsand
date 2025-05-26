---
layout: post
title: Proprietary FSF
date: 2009-01-01 20:24:16.000000000 +00:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Philosophy
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
meta:
  _publicize_pending: '1'
  oc_metadata: "{\t\tversion:1.0,\t\ttags: {'freedom': {\t\t\ttext:'Freedom',\t\t\tslug:'freedom',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'fsf':
    {\t\t\ttext:'FSF',\t\t\tslug:'fsf',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'gnu':
    {\t\t\ttext:'GNU',\t\t\tslug:'gnu',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'gpl':
    {\t\t\ttext:'GPL',\t\t\tslug:'gpl',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'gplv3':
    {\t\t\ttext:'GPLv3',\t\t\tslug:'gplv3',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'lgpl':
    {\t\t\ttext:'LGPL',\t\t\tslug:'lgpl',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'open-source':
    {\t\t\ttext:'Open Source',\t\t\tslug:'open-source',\t\t\tsource:null,\t\t\tbucketName:'current'\t\t},'software-developers':
    {\t\t\ttext:'software developers',\t\t\tslug:'software-developers',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/53915970-e599-3269-9036-dadd14e5d068',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/IndustryTerm',\t\t\ticonURL:'',\t\t\tname:'IndustryTerm'\t\t},\t\t\tname:'software
    developers',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t},'technology':
    {\t\t\ttext:'technology',\t\t\tslug:'technology',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/5a253e82-7004-3a68-8797-61b5b7eda895',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/IndustryTerm',\t\t\ticonURL:'',\t\t\tname:'IndustryTerm'\t\t},\t\t\tname:'technology',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t},'software':
    {\t\t\ttext:'software',\t\t\tslug:'software',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/30c49aa7-f102-3988-8215-2e76315c6ed3',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/IndustryTerm',\t\t\ticonURL:'',\t\t\tname:'IndustryTerm'\t\t},\t\t\tname:'software',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t},'open-source-library':
    {\t\t\ttext:'Open Source library',\t\t\tslug:'open-source-library',\t\t\tsource:{\t\t\turl:'http://d.opencalais.com/genericHasher-1/08787c79-222c-356b-b892-12e7feb33809',\t\t\ttype:{\t\t\turl:'http://s.opencalais.com/1/type/em/e/Facility',\t\t\ticonURL:'',\t\t\tname:'Facility'\t\t},\t\t\tname:'Open
    Source library',\t\t\tnInstances:1\t\t},\t\t\tbucketName:'current'\t\t}}\t}"
  oc_commit_id: http://drone-ah.com/2009/01/01/proprietary-fsf/1236796134
  _edit_last: '48492462'
  restapi_import_id: 591d994f7aad5
  original_post_id: '126'
  _wp_old_slug: '126'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:35'
permalink: "/2009/01/01/proprietary-fsf/"
---

I have always a big fan and proponent of the FSF and having recently
been interested in researching for a project came across a document
covering [Why you shouldn\'t use the Lesser GPL for your next
library](http://www.gnu.org/licenses/why-not-lgpl.html "Why Not LGPL"){target="_blank"}

What the document basically suggests is to limit what proprietary
software developers can do by licensing libraries as GPL instead of
LGPL.

This is no longer free(as in speech, not beer) software. Why?

Freedom means the ability to use something without restriction. If I
cannot use a library in a proprietary product, that is removing an
important freedom.

This attitude is likely to alienate the \"commercial\" or proprietary
developers further from FSF/GNU.

In fact, doing this is just not fair and not in line with how I view is
the concept behind the FSF. The point is to write software / libraries
and share that with the world so others may build upon what you have
done. Stand on the shoulders of giants in a way\...

It makes perfect sense for software to be GPL since you don\'t want
somebody to pick up a GPL software, build something on top, and sell it
without source.

However, if libraries are released under the GPL instead of LGPL, it
means that I can not link against that library to write a non-GPL
compatible application.

The [GNU
Website](http://www.gnu.org/ "The GNU Operating Sytem"){target="_blank"}
states

> \"Free software is a matter of the users\' freedom to run, copy,
> distribute, study, change and improve\"

Additionally, the [Quick Guide to
GPLv3](http://www.gnu.org/licenses/quick-guide-gplv3.html "A Quick Guide To GPLv3"){target="_blank"}
states that

> Nobody should be restricted by the software they use. There are four
> freedoms that every user should have:

-   the freedom to use the software for any purpose,
-   the freedom to change the software to suit your needs,
-   the freedom to share the software with your friends and neighbors,
    and
-   the freedom to share the changes you make.

This has always been my impressing of the purpose of GPL. Now, how does
this work with Libraries? A little differently\... :-(

From my perspective, if I have the freedom to use the \[library\] for
any purpose, that means that I can write an application that ***uses***
that library without having to worry about licensing issue.

However, this is not the case. There is a clause that states that the
software cannot be used in a larger software project that has a license
incompatible with the GPL. This includes linking the library into
another software application.

Therefore, I do not have the freedom to use the software ***for any
purpose***.

Freedom cannot be uni-directional. If GNU/FSF are trying to muscle out
developers of proprietary software, all they are doing is alienating
themselves further\...

I run a technology firm that uses a heck of a lot of open source
software. In fact, I am posting this from an ubuntu desktop running
firefox from a VServer. I am probably using a dozen open source
applications to do this simple straightforward act.

There is in fact, not a simple closed source application at any point
through this.

The main problem that I see with this is that it makes Open Source so
much more zealot(ous) and FSF, GNU and OSS becomes fundamentalists. The
attitude is not one of freedom and inclusion but of exclusivity and
marginalisation.

The worst part is the price that is asked of developers who want to use
an Open Source library. The price is the acceptance and propogation of
an idea (Freedom or else).

Compared to the cost of conversion to another idealogy (Free Software
Idealogy), the cost of a few hundred, thousand, or even millions of
dollars / pounds for a piece of software seems dirt cheap.

I understand that each developer has the freedom to choose which license
to use for their products/libraries. My question is how can an
organisation that claims to be a proponent of freedom encourage the
removal of freedoms?

I would like to ask how this shift is any different from religious
fanatics who tell you that their god is the one true god and there is
nothing else.
