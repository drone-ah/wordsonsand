---
layout: post
title: 'Getting Started: Emacs &amp; C++ (w/ cmake) (On the fly syntax highlighting)'
date: 2013-02-24 11:54:03.000000000 +00:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Software
tags:
- CMake
- Compiling tools
- Emacs
- Flymake
- GNU build system
- Make
meta:
  _edit_last: '48492462'
  oc_commit_id: http://drone-ah.com/2013/02/24/getting-started-emacs-c-w-cmake/1361706846
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {'compiling-tools': {\"text\":\"Compiling
    tools\",\"slug\":\"compiling-tools\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/b6e22a64-d169-3a9a-a306-dc95fe140c5b/SocialTag/3\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Compiling
    tools\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'cmake': {\"text\":\"CMake\",\"slug\":\"cmake\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/b6e22a64-d169-3a9a-a306-dc95fe140c5b/SocialTag/4\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"CMake\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'emacs': {\"text\":\"Emacs\",\"slug\":\"emacs\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/b6e22a64-d169-3a9a-a306-dc95fe140c5b/SocialTag/5\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Emacs\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'make': {\"text\":\"Make\",\"slug\":\"make\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/b6e22a64-d169-3a9a-a306-dc95fe140c5b/SocialTag/6\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Make\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'gnu-build-system': {\"text\":\"GNU build system\",\"slug\":\"gnu-build-system\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/b6e22a64-d169-3a9a-a306-dc95fe140c5b/SocialTag/7\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"GNU
    build system\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'flymake': {\"text\":\"Flymake\",\"slug\":\"flymake\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  restapi_import_id: 591d994f7aad5
  original_post_id: '908'
  _wp_old_slug: '908'
  _publicize_job_id: '5185240607'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:57'
permalink: "/2013/02/24/getting-started-emacs-c-w-cmake/"
---

I am a recent convert to emacs. My vast majority of development is in
Java EE and I have not found an easy way to get the functionality in
eclipse into emacs. So I still use eclipse for this.

However, I like to tinker with C++ and I wanted to get some of the CDT
functionality into emacs. In truth, I have used very little CDT so my
expectations from emacs will be set differently. Considering that emacs
has been used for C/C++ development for decades, I am hopeful that it
will be more feature-rich than eclipse or any of the other IDE\'s like
Anjuta, Code::Blocks etc (both I have tried to use).

First things first. In the world of Java, I am a massive fan of maven
which makes build management so easy and simple. Having used it now for
years, it is easy to forget how much of a learning curve it had to get
started.

Autotools are a massive pain to use and has a very steep learning curve.
I have used it in the past to set up build environments and it works
fine. pkg-config is pretty awesome and in a lot of ways, maven does pale
in comparison. i.e. instead of having maven pull in dependencies, you
just use your systems package manager like apt-get or yum and it
installs the libraries for you.

Long story short, I am using [cmake](http://www.cmake.org/ "CMake")
which has the added advantage of being a little more cross-platform
(i.e. supported in Windows as well as \*nix). If you haven\'t used CMake
before, let me tell you - it\'s a heck of a lot easier to get used to
than Autotools. Just go through their
[tutorial](http://www.cmake.org/cmake/help/cmake_tutorial.html "CMake Tutorial")
and you should be off.

The next thing I wanted to sort out was on the fly syntax checking. This
makes life a lot easier and means that you can write and correct syntax
errors etc without having to build manually.

[Flymake](http://www.emacswiki.org/emacs/FlyMake) is what you want to
use for this. The later versions of emacs comes with Flymake so you
won\'t necessary need to install it to get started. However, flymake
doesn\'t (unfortunately) just work magically out of the box and requires
a little configuration.

After hunting around for a bit, and finally from the [EmacsWiki
Flymake](http://www.emacswiki.org/emacs/FlyMake) page, found a couple of
options
[cpputils-make](https://github.com/redguardtoo/cpputils-cmake) and
[cmake-project](https://github.com/alamaison/emacs-cmake-project).
Cmake-project seemed simpler and I opted for that. I also tried
installing cpputils-make and didn\'t have any issues with that either.

There is one thing you need to be aware of though - both these tools
expect you to do out-of-source-builds. This essentially requires you to
create a build folder (called bin for cmake-project and build for
cpputils-make) and generate the Makefile etc. in there.

This is the preferred way with CMake anyway so it\'ll be better to do
build in that way. It\'ll also  make it easier to have different builds
(Debug/Release etc.)

The easiest way to install either of these is through marmalade. If you
don\'t already have it installed - it is so easy - just follow the
instructions on their [homepage](http://marmalade-repo.org/). You can
then install by running

M-x package-install cmake-project

OR

M-x package-install cpputils-cmake

add the following to your .emacs file for cmake project. Instructions
for cpputils-make can be found on [their github
page](https://github.com/redguardtoo/cpputils-cmake)

[(require
\'cmake-project)]{style="font-family:Consolas, Monaco, monospace;font-size:12px;line-height:18px;"}

Do a full build on your sources first by going to the bin or build
directory and generating the makefiles by using cmake (cmake .. or cmake
../src depending on how you set up cmake) and then make.

You can then initialise the mode within emacs

M-x cmake-project-mode    (for cmake

You may have to also enable flymake

M-x flymake-mode
