---
categories:
  - Software Development
date: "2011-12-28T14:51:31Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:53"
  _publicize_pending: "1"
  _wp_old_slug: "754"
  oc_commit_id: http://drone-ah.com/2011/12/28/hudson-jenkins-and-continuous-integration-1114/1325083894
  original_post_id: "754"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Apache Maven
  - Continuous integration
  - Hudson
  - Maven
  - PMD
  - Subversion
title: Hudson / Jenkins and Continuous Integration [1114]
url: /2011/12/28/hudson-jenkins-and-continuous-integration-1114/
---

Fair Warning: This is more notes for me to remember and document how to do these
things rather than particularly detailed instructions. Therefore, it might be
missing sections and will assume a reasonable knowledge of hudson/jenkins and
not to mention the benefits of continuous integration and builds.

Installing hudson / jenkins is easy enough. I deployed as part of a pre-existing
tomcat6 installation so was as simple as popping the war file into the webapps
folder. Tomcat automatically started it up without issues.

I chose to have hudson use /home/hudson as its home directory. Since I am
running an ubuntu system, I added a line into /etc/defaults/tomcat6. There are
various other ways of doing this but it was a quick fix for me.

You of course need to make sure the directory exists. I also popped in a .m2
folder from my home directory to save it from downloading all the various jar
files and included a settings.xml file with appropriate configurations.

Hudson 2.2 uses maven 3 but I use maven 3 locally as well even though the
projects pom files were built for maven 2. There doesn't seem to be any issues
with this setup.

First step is to create a new job from the home page. This asks for which type a
job you want to create. If you use maven and a standard source control, it is as
simple as choosing the first option: Build a free-style software project.

Give it a name and you are brought to the configuration screen. There are a
number of options here and I started with the basic set:

I chose Subversion for the source control management section and gave it the svn
path. There is a checkout strategy as well and I chose the one to revert and
update which I feel to be a bit cleaner.

I chose to poll the scm every fifteen minutes

```cron
*/15 * * * *
```

and saved.

Running the build pulled the code out of svn and stopped there. This was because
I didn't add a step to build / install it.

Go back into configure the job and add a maven 3 build step. This automatically
selected the clean install goals. Save and build now and the project was checked
out and built without issues.

Success!

There are a number of other options you can play with here but this gives you a
solid starting point.

Later on, I will cover the addition of various other plugins for source analysis
including findbugs and pmd.
