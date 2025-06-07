---
layout: post
title: Maven2, EJB3 and JBoss
date: 2008-12-28 18:53:50.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Java (EE)
tags:
  - Eclipse
  - EJB3
  - ERP
  - InVision
  - Maven 2
  - Out of Container Testing
  - Process Dashboard
  - Request Tracker
meta:
  _edit_last: "48492462"
  restapi_import_id: 591d994f7aad5
  original_post_id: "110"
  _wp_old_slug: "110"
  _publicize_job_id: "5181443435"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:35"
permalink: "/2008/12/28/maven2-ejb3-and-jboss/"
---

I started work on a project called InVision about a year ago but have probably
spent about a week or two worth of effort on it in total... :-(

The Project aim was to bring together the easy time logging capabilities of
[Process Dashboard](http://processdash.sourceforge.net/ "The Software Process Dashboard Initiative")
along with the project management capabilities of Microsoft Project (including
the Server Component). It is also to be integrated into our request tracking
System - [Request Tracker](http://bestpractical.com/rt/ "Request Tracker").
Eventually, it is also to integrate with our accounting system and turn into an
ERP (Enterprise Resource Planning) system and MIS (Management Information
System). There are plans to integrate with our Wiki and our Document Management
System too.

But these are all lofty goals.  One of our recent projects introduced me to the
[Spring Framework](http://www.springframework.net/ "Spring.NET Application Framework").
While I am still not a fan of Spring, the scale of the project and the way of
approaching it gave me some ideas and additional tools to work with. I wanted to
bring these into the InVision Project.

The key one here was Maven 2. InVision already used EJB3 and JBoss (4.2 as it
happened). There was one additional issue for me to resolve and that was out of
container testing. Something that is very easy to do with Spring but a little
more troublesome with EJB3 since it doesn't have an out of container
framework...

<!-- more -->

I have grown to be a big fan of Maven 2 and using Maven 2 to configure an EJB
project is not as easy or straightforward as I would have liked: I wanted to
separate the whole project into four parts

- Domain Model (or just the entity beans); Also referred to as a Hibernate
  Archive (HAR)
- Stateful/Stateless Beans (Just the Beans, since I don't consider entities
  beans in EJB3)
- Application Client (J2SE Application)
- Web App (Using SEAM)
- I would also need an EAR project to deploy the DomainModel, Beans & WebApp as
  one pacakge into JBoss.

I have not got as far as the SEAM project yet but the other ones were
straightforward enough to set up with Maven 2.

Both the Domain Model and the Beans project had to be set up as ejb projects and
use the maven-ejb-plugin

```xml
 <build>
     <plugins>
         <plugin>
             <groupId>org.apache.maven.plugins</groupId>
             <artifactId>maven-ejb-plugin</artifactId>
             <configuration>
                 <ejbVersion>3.0</ejbVersion>
             </configuration>
         </plugin>
     </plugins>
 </build>
```

I set up the persistence context within the Domain Model

```xml
<persistence-unit name="em">
    <provider>org.hibernate.ejb.HibernatePersistence</provider>
    <jta-data-source>java:/datasource</jta-data-source>
</persistence-unit>
```

I could then reference the context from the Beans project by injecting it with

```java
@PersistenceContext(unitName="em")
```

Easy enough!

Now configuring the EAR project: This was configured as an ear package which
depended on the other two projects with the following configuration

```
<build>
<plugins>
<plugin>
<groupId>org.apache.maven.plugins</groupId>
<artifactId>maven-ear-plugin</artifactId>
<configuration>
<version>5</version>
 <modules>
 <ejbModule>
 <groupId>uk.co.kraya.invision</groupId>
 <artifactId>beans</artifactId>
 </ejbModule>
 <ejbModule>
 <groupId>uk.co.kraya.invision</groupId>
 <artifactId>DomainModel</artifactId>
 </ejbModule>
 </modules>
 <jboss>
 <version>4.2</version>
 <data-sources>
 <data-source>invision-ds.xml</data-source>
 </data-sources>
 </jboss>
 </configuration>
 </plugin>
 <plugin>
 <groupId>org.codehaus.mojo</groupId>
 <artifactId>jboss-maven-plugin</artifactId>
 <configuration>
 <jbossHome><jboss-home-path></jbossHome>
 <hostName><hostname></hostName>
 <port>8080</port>
 </configuration>
 </plugin>
 </plugins>
 </build>
```

With this configured, from the EAR project, I could do mvn ear:deploy to deploy
to JBoss.

Additionally, within eclipse, I created a new run-type that ran ear:undeploy
package ear:deploy to re-deploy the package to JBoss. Works a treat

There are still a few kinks to be ironed out.

I still need to install (mvn install) the two projects before the EAR will pick
it up to deploy. I need to get the ear re-deploy to re-build the other projects.
Something to look at another day.

I had manually deployed the DataSource file to JBoss. It might be possible to do
this via Maven.

I also very much liked the Eclipse automatic deploy feature. It is possible to
use the eclipse plugin on maven to get Eclipse to identify this as a JBoss
deployable project but I ran into some problems and gave up. Ideally, Eclipse
would auto-deploy the project.

However, the above is less relevant once Out-Of-Container testing is in place.
Now, this does work, but I will leave that to another day...
