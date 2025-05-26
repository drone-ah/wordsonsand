---
layout: post
title: Visual Studio, C++ &amp; Google Test
date: 2021-08-09 16:22:08.000000000 +01:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories: []
tags:
- c++
- google-test
- visual studio
meta:
  _last_editor_used_jetpack: block-editor
  _publicize_job_id: '61672088922'
  timeline_notification: '1628526129'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:55:04'
permalink: "/2021/08/09/visual-studio-c-google-test/"
---

Using Visual Studio for C++ and Google Test seems like it should be
absolutely straightforward. I\'m trying to do everything in Visual
Studio, so it should just set it up automatically right at the start
when I create a project. Perhaps I am spoilt from working primarily with
Java for many years, but here I was, trying to set it up and it took an
inordinate amount of time.

First things first, the way it works in C++, at least in Visual Studio,
unlike in Java is that the tests are set up in a separate project, but
within the same solution. I will try and remember all the steps I had to
undertake to set this up.

This, however, is not the only requirement. When testing, since we can
have only one main function in the final executable, it also makes sense
to put all of your code, apart from the main function for your
application in another project.

In summary, you end up with three projects

-   Your Solution
    -   Static Library with all your code \*except\* the main function
    -   Executable project with your main function (linked with your
        static library above)
    -   Tests project (again, linked with your static library above.

Once you have this set up, you will want to add a reference to the
static library from both the executable project and the tests project

<figure class="wp-block-image size-large">
<a href="https://drone-ah.com/wp-content/uploads/2021/08/image.png"><img
src="%7B%7Bsite.baseurl%7D%7D/assets/2021/08/image.png?w=622"
class="wp-image-1255" /></a>
</figure>

Finally, you will also have to update the linker to link with the
library project. I am sure there is a better way of doing this, but I
did it by right clicking the project -\> Properties -\> Linker -\>
General:

Additional Library Directories -\> and adding in something like
\"`$(SolutionDir)<lib-folder>\$(IntermediateOutputPath)*.obj`\"

You may have to add the same into Linker -\> Input -\> Additional
Dependencies.

Hope that helps

\
