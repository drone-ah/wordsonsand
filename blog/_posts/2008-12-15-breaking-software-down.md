---
layout: post
title: Breaking Software Down
date: 2008-12-15 15:53:43.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Software
tags:
  - "*nix"
  - Brave New World
  - Coding Horror
  - collaboration
  - COM
  - Contact Management
  - CORBA
  - find
  - Firefox
  - GNOME Evolution
  - grep
  - Integration
  - Interoperability
  - Jeff Atwood
  - Mylyn
  - OLE
  - REST
  - RPC
  - Safari
  - sed
  - SOAP
  - Software Component
  - Software Garden
  - Subclipse
  - Subversive
  - Thunderbird
meta:
  _edit_last: "48492462"
  restapi_import_id: 591d994f7aad5
  original_post_id: "56"
  _wp_old_slug: "56"
  geo_public: "0"
  _publicize_job_id: "15641121183"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:34"
permalink: "/2008/12/15/breaking-software-down/"
---

[Jeff Atwood likens software development to tending a garden.](http://www.codinghorror.com/blog/archives/000987.html "Tending Your Software Garden"){target="\_blank"
rel="noopener"} I can relate to this. In fact, I would like to ask, if you have
a nice plant in one of your gardens, how complicated is it to \"copy\" that
across to another one?

I realise that I am moving away from the analogy here but there is an important
concept here. Libraries were born out of the desire to share and distribute code
to be re-used.

The idea for
[Remote Procedure Calls](http://en.wikipedia.org/wiki/Remote_procedure_call "Remote Procedure Call"){target="\_blank"
rel="noopener"} dates as far back as 1976. Microsoft brought along
[OLE](http://en.wikipedia.org/wiki/Object_Linking_and_Embedding "Object Linking and Embedding"){target="\_blank"
rel="noopener"} and then
[COM](http://en.wikipedia.org/wiki/Component_Object_Model "Component Object Model"){target="\_blank"
rel="noopener"} made this more generic and better.

RPC is widely in use these days and there are several other mechanisms for inter
process communication including
[CORBA](http://en.wikipedia.org/wiki/CORBA "Common Object Request Broker Architecture"){target="_blank"
rel="noopener"},
[REST](http://en.wikipedia.org/wiki/REST "Representational State Transfer"){target="\_blank"
rel="noopener"} & [SOAP](http://en.wikipedia.org/wiki/SOAP_(protocol) "Simple
Object Access Protocol"){target="\_blank" rel="noopener"}.

I don\'t think software is broken down into small enough components. \*nix is
great in that you can tag a whole bunch of commands together on the command line
to do some amazing things. I have personally piped data through a dozen or so
commands and scripts to do some interesting things.

If we could break everything down into individual components that could be
linked together, we would have a massive arsenal of interoporable tools that
each user can pick and choose to put together very powerful solutions.

How many times have you found a piece of software that does one thing really
well, but fails in something else. Then found another piece of software that
does the other thing really well.

For example, the extensibility of
[Firefox](http://www.mozilla.org/firefox "Firefox"){target="\_blank"
rel="noopener"} is fantastic but I love the rendering of
[Safari](http://www.apple.com/safari/ "Safari Web Browser"){target="\_blank"
rel="noopener"}. I love the Contact Management within
[Evolution](http://projects.gnome.org/evolution/ "Evolution"){target="\_blank"
rel="noopener"} and the Mail capabilities of
[Thunderbird](http://www.mozilla.org/thunderbird "Thunderbird"){target="\_blank"
rel="noopener"}.

Why don\'t we break each software down into each of it\'s individual components
(and I am not talking about libraries here) and allow them to be deployed as
services usable by other pieces of software.

In other words, release the contact management capabilities of Evolution as a
product of it\'s own right with a pre-defined API that any application can link
into (including perhaps a web interface). Release the Mail management component
of Thunderbird as a service, Release GUI\'s as a component. Then we can pick any
GUI we want, link into a specific mail component and another addressbook
component.

Do one thing and do it well. In fact, let\'s take it one step further and
release a public API for each software component - an API for Mail, one for
Contact Management and so on.

Each software component can then be a black box that delivers this API.

Choice can be a bad thing if it makes it difficult to choose -
[Subclipse](http://subclipse.tigris.org "Subclipse"){target="\_blank"
rel="noopener"} vs
[Subversive](http://www.eclipse.org/subversive "Subversive"){target="\_blank"
rel="noopener"} is a good example of this. Let us however, not confuse choice
with flexibility.

Let\'s say that you want to find all the files within a folder modified within
the last 3 days containing the text \"abracadabra\" and then replace all
occurences in those files of the world \"super\"  with \"hyper\".

To do this in linux, all you would do is chain find (to identify files modified
in the last 3 days), grep (to identify only the files that contain
\"abracadabra\") and sed (to do the replacement).

If you know these commands well enough, you could chain something together in
half a minute or so. You could probably figure out how to do this with the
search tools in Windows within a minute or so but where this really shines is if
there are thousands of files that needs to be processed. With other search
tools, you would have to wait for the original search results to be returned
before running to replace operation. This takes up the users time.

With the chaining of commands, I have run it and worked on something else while
it completes.

Let me visualise a brave new world:

In this world, all software would be interoperable components. For example,
there would be components for:

- Mail account management (Perhaps genericised into configuration management)
- Text composition (usage for mail, documents, plain text et al)
- Text reading (again, usable for mail, documents, plain text et al)
- Spam Filtering (already available to some extent)
- Contact Management (optionally linked into organisation\'s LDAP server)
- Task Management (Standalone
  [Mylyn](http://www.eclipse.org/mylyn/ "Eclipse - Mylyn"){target="\_blank"
  rel="noopener"} if you know the product)
- Scheduling (or calendering if you prefer that term)

If all of these components were interoperable, then there would a
[GUI](http://en.wikipedia.org/wiki/Graphical_user_interface "Graphical User Interface"){target="\_blank"
rel="noopener"} that is generic and could bring all of these together. In this
way, the people working on each of the components could concentrate on doing one
thing and one thing well.

If we then start working on public API\'s in a collaborative fashion, each of
the component could be fleshed out to be as flexible and complete as necessary
to gain maximum benefit.

If these components provided the services as a network based API, it would also
allow for the components to be distributed across a network providing redundancy
and efficiency. This makes it easier to turn each desktop into more of dump
terminal concentrating purely on user interaction and getting closer to the
[invisible interface.](http://drone-ah.com/2008/12/12/invisible-interface/ "Invisible Interface"){target="\_blank"
rel="noopener"}

Software as a service has taken a step in the right direction. Can we take a
leap and have software component as a service\...
