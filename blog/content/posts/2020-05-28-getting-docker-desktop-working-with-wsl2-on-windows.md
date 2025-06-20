---
categories:
  - Systems (Administration)
date: "2020-05-28T17:29:52Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:04"
  _publicize_job_id: "44804887047"
  timeline_notification: "1590686992"
parent_id: "0"
password: ""
status: publish
tags:
  - sysadmin
  - docker
  - docker-desktop
  - wsl2
title: Getting Docker Desktop Working with WSL2 on Windows
url: /2020/05/28/getting-docker-desktop-working-with-wsl2-on-windows/
---

I ran into several issues while trying to get this to work. Here are the steps I
had to complete to get it working. Hopefully this will save some hair on your
head ;)

The main step is to go into the settings in Docker Desktop -\> Resources and
make sure that your distribution is enabled for docker.

![enable your distro on docker](/assets/2020/05/image.png)

1.  Make sure that you have no docker packages installed on your WSL
    distribution. Docker Desktop will deploy its own binaries, and any
    pre-existing binaries will confuse it. This issue exhibited itself for me
    with errors related to missing files around credentials.
2.  Remove any `DOCKER_HOST `environment variables. Docker Desktop will sort it
    out. Docker kept hanging for me until I fixed this.
3.  If you want to use docker as non-root user, add yourself to the
    `docker `group.

Errors / Issues I ran into:

`docker.credentials.errors.InitializationError: docker-credential-desktop.exe not installed or not available in PATH` -
Fixed by 1 above.

`docker-compose` from WSL2 errors out - Again, fixed by 1

Unable to run `docker` as non-root user - fixed by 3.

Docker hangs when run as non-root user - fixed by 2.
