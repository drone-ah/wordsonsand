---
categories: []
date: "2021-08-09T16:22:08Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:04"
  _last_editor_used_jetpack: block-editor
  _publicize_job_id: "61672088922"
  timeline_notification: "1628526129"
parent_id: "0"
password: ""
status: publish
tags:
  - c++
  - google test
  - visual studio
title: Visual Studio, C++ & Google Test
url: /2021/08/09/visual-studio-c-google-test/
---

Using Visual Studio for C++ and Google Test seems like it should be absolutely
straightforward. I\'m trying to do everything in Visual Studio, so it should
just set it up automatically right at the start when I create a project. Perhaps
I am spoilt from working primarily with Java for many years, but here I was,
trying to set it up and it took an inordinate amount of time.

First things first, the way it works in C++, at least in Visual Studio, unlike
in Java is that the tests are set up in a separate project, but within the same
solution. I will try and remember all the steps I had to undertake to set this
up.

This, however, is not the only requirement. When testing, since we can have only
one main function in the final executable, it also makes sense to put all of
your code, apart from the main function for your application in another project.

<!--more-->

In summary, you end up with three projects

- Your Solution
  - Static Library with all your code \*except\* the main function
  - Executable project with your main function (linked with your static library
    above)
  - Tests project (again, linked with your static library above.

Once you have this set up, you will want to add a reference to the static
library from both the executable project and the tests project

![Add Reference menu](/assets/2021/08/image.png "Add Reference")

Finally, you will also have to update the linker to link with the library
project. I am sure there is a better way of doing this, but I did it by right
clicking the `project -> Properties -> Linker -> General`:

`Additional Library Directories ->` and adding in something like
"`$(SolutionDir)<lib-folder>\$(IntermediateOutputPath)*.obj`"

You may have to add the same into `Linker -> Input -> Additional Dependencies`.

Hope that helps
