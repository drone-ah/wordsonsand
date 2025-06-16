---
categories:
- Java (EE)
date: "2011-11-05T15:32:32Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:51"
  _publicize_pending: "1"
  _wp_old_slug: "702"
  oc_commit_id: http://drone-ah.com/2011/11/05/jbossws-cxf-pojo-vs-stateless/1320507155
  original_post_id: "702"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- Java (EE)
- Load testing
- Log4j
- Plain Old Java Object
- Software components
- web services
- web.xml
- XML
title: JBossWS CXF - POJO vs Stateless [1104]
url: /2011/11/05/jbossws-cxf-pojo-vs-stateless/
---

Cleaning up a bunch of code to reduce object instantiations got me thinking
about the webservice layer. We are using POJO based webservices but it got to me
wondering whether useless Stateless web service beans would improve memory
usage. More accurately, whether it would improve garbage collection performance.

To test this, the plan was to build two versions of the same web service and
load test it to see the memory and cpu utilisation to compare cost /
performance.

In the process, I also discovered other differences.

<!--more-->

Both the web services were built against this interface

```java
@WebService
public interface WSTest {

    @WebMethod
    public String greetClient(String name);

}
```

That is of course a simple enough interface.

The EJB version of this is as follows:

```java
@WebService(endpointInterface = "uk.co.kraya.wstest.WSTest")
@Stateless
@Remote(WSTest.class)
public class EJBWebTest implements WSTest {

    public String greetClient(String name) {
        return "EJB Says Hello to " + name;
    }

}
```

Simple and straightforward enough and the POJO version is not that different

```java
@WebService
public class PojoWebTest implements WSTest  {

    @WebMethod
    public String greetClient(String name) {
        return "Pojo Says Hello to " + name;
    }

}
```

I also used the following web.xml. This may not be a necessary step any more.

```xml
<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE web-app PUBLIC "-//Sun Microsystems, Inc.//DTD Web Application 2.2//EN" "http://java.sun.com/j2ee/dtds/web-app_2_2.dtd">
<web-app>

        <display-name>Archetype Created Web Application</display-name>
        <servlet>
                <servlet-name>GreetingWebService</servlet-name>
                <servlet-class>uk.co.kraya.wstest.pojo.PojoWebTest</servlet-class>
        </servlet>
       <servlet-mapping>
                <servlet-name>GreetingWebService</servlet-name>
                <url-pattern>/*</url-pattern>
        </servlet-mapping>
</web-app>
```

Now that was complete, there was the deployment to consider. I had assumed that
I could just do one build, package it as an war and just deploy it but as it
turns out, that didn't work. This only deployed the Pojo web service. To deploy
the EJB web service, I had to package it as an EJB (maven) and deploy a jar
file.

This meant that I ended up with two deployments, wstestpojo.war and
wstestejb.jar.

Once deployed, I used SoapUI to load test both the web services and the results
were interesting.

In the grand scheme of things, the difference was pretty minimal between the
two. However, the Stateless EJB web service used a little extra (3% - 5%) CPU
presumably from the pooling of the beans and accessing them.

I used JBoss 5.1 and ran into an issue with log pollution

EJBTHREE-1337: do not get WebServiceContext property from stateless bean
context, it should already have been injected

This was printed every time an EJB based web service was called.

I
[found more information about the EJBTHREE-1337 issue and a workaround for it](http://idevone.wordpress.com/2009/09/14/howto-suppress-ejbthree-1337-warning/ "HOWTO: Suppress EJBTHREE-1337 warning")
The workaround simply involves updating log4j to not log it which is good enough
for me... :-)

As a final note, I am unsure as to how this would scale with complex web
services that have multiple instantiated objects as part of it.  I have a sneaky
suspicion that web services with expensive object instantiations would perform
better if using Stateless EJB Beans as would web services that are memory
hungry. However, this has not been tested.
