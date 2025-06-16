---
categories:
- yertoob
date: "2018-06-26T11:18:38Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:03"
  _publicize_job_id: "19370182202"
  _rest_api_client_id: "-1"
  _rest_api_published: "1"
  timeline_notification: "1530011919"
parent_id: "0"
password: ""
published: true
status: publish
tags:
- java
- javafx
title: What I learnt developing a small JavaFX App
type: post
url: /2018/06/26/what-i-learnt-developing-a-small-javafx-app-wip/
---

# Introduction

This is a collection of the things I learnt developing a simple JavaFX app over
the last month or two. My background is very much in Java EE with decades of
experience building high end, high-performance ticketing systems. This means
that my expectation from a development environment is relatively high. There are
many optional components in here that I find worthwhile setting up at the start,
but are not necessary

## Tools

### Maven

One of the most useful tools I have found while working with Java is maven. If
maven isn\'t a part of your build, have a look at it and re-evaluate that. I
have no doubt that maven has saved me hundreds, if not thousands of hours over
the last few years.

### JavaFX Scene Builder

While this one has a bunch of issues and serious limitations, it can still be a
helpful tool. It helped me get a handle on the components available and placing
items.

# Libraries

## Testing

I use **junit5,** but there are other options like test-ng which are equally
good. I use **Mockito **for mocking, but there are many other options like
PowerMock, JMockit, EasyMock, etc.

For UI Testing, you can use **TestFX.** I don\'t like UI work, so haven\'t done
much work with this.

        org.junit.platform
        junit-platform-launcher
        1.2.0
        test


        org.junit.jupiter
        junit-jupiter-engine
        5.2.0
        test


        org.junit.vintage
        junit-vintage-engine
        5.2.0
        test



        org.mockito
        mockito-core
        2.18.3
        test

## Logging

I can't live without logging in any application. It can make troubleshooting
much easier, particularly when you've deployed your app. **log4j2** is the main
logging framework out there. You can choose another one if you like, but I
strongly recommend having and using one.

        org.apache.logging.log4j
        log4j-slf4j-impl
        2.11.0


        org.apache.logging.log4j
        log4j-api
        2.11.0


        org.apache.logging.log4j
        log4j-core
        2.11.0

## Dependency Injection

If you have working the Java EE Environment, you have almost certainly come
across
[Inversion of Control](https://en.wikipedia.org/wiki/Inversion_of_control), 
particularly in the form
of [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection). I
love dependency injection. It helps with decoupling components and with testing.
I looked at various frameworks including
[Dagger 2](https://google.github.io/dagger/), [Spring](https://spring.io/),
[Guice](https://github.com/google/guice).

### Dagger

The fully static (compile time) nature of Dagger 2 means that it doesn\'t gel
well with JavaFX which is very dynamic.

### Spring

I had worked with Spring many years ago and didn\'t want to tangle with a
behemoth for a small project. There are many components and loads of
functionality in spring and if you building a large and complex project, it
might be worth it.

### Google Guice

Google Guice is the framework that I ended up going with. It does have some
dependencies like Guava, but as it turned out, I Guava comes in handy for JavaFX
anyway. We don\'t need an entry in the pom.xml for this because of the following
dependency.

## Gluon Ignite

[Gluon Ignite](https://gluonhq.com/labs/ignite/) was released by gluon labs to
integrate Dependency Injection frameworks with JavaFX. In other words, it ties
in the DI framework with the FXMLLoader so that it will load the correct
controller instances. Since I am using Guice, I needed the ignore-guice module.
If you add this into your pom.xml, it will also pull in google guice. Easy eh?
;)

        com.gluonhq
        ignite-guice
        1.0.2

If you don\'t want to add another dependency, you could take a look at the code
in this module. It\'s fairly straightforward to integrate that manually into
your app. It\'s just easier to add in the dependency and let it do the magic

## Event / Publisher / Subscriber Framework

It is likely that you will need an event framework or a publisher/subscriber
framework. The nature of GUI design and work is that it is an easy and simple
solution for a number of problems you will come across. Fortunately, we already
have an
[event framework](https://github.com/google/guava/wiki/EventBusExplained) in
place within [Guava](https://github.com/google/guava) which is a dependency of
Google Guice.

## Project Lombok

Don\'t you love adding in a getter and a setter for each field? When you change
your fields, don\'t you love going in and updating all the getters and setters?
How about defining the long, pita to type types with generics of variables that
you are assigning from a method call? I mean the compiler can\'t possibly figure
that out by itself, right? How about writing out the toString, equals and
hashCode for each class? What do you mean no, you don\'t? You don\'t love these
tedious repetitive tasks of development? Good! You will love
[Lombok.](https://projectlombok.org/setup/eclipse) Unfortunately, with Lombok,
it\'s not as simple as adding it into your pom.xml. Check out the install
instructions on their website for you IDE etc. You also need it in your pom.xml.

        org.projectlombok
        lombok
        1.18.0
        provided

There is some controversy around Lombok. There is a
[good post on stackoverflow that covers some of these things in a reasonable fashion](https://stackoverflow.com/questions/3852091/is-it-safe-to-use-project-lombok).

## Apache Commons

Last but certainly not least, we have
[apache commons](https://commons.apache.org/). This is a collection of libraries
rather than a single one. You know all those bits of code you write over and
over. Chances are that there is something in here that does it better and in a
simpler way.

## Persistence

TBC

# Packaging

TBC
