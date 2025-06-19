---
categories:
  - Database Systems
  - Kraya
  - Software
date: "2009-03-10T16:00:21Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:38"
  _publicize_job_id: "15683357895"
  _wp_old_slug: "204"
  geo_public: "0"
  oc_commit_id: http://drone-ah.com/2009/03/10/database-systems-compared/1236790940
  original_post_id: "204"
  restapi_import_id: 591d994f7aad5
  timeline_notification: "1520950046"
parent_id: "0"
password: ""
status: publish
tags:
  - Database Systems
  - DBase
  - edFringe
  - Filemaker Pro
  - Foxpro
  - GUI
  - HSQLDB
  - IBM
  - IBM DB2
  - insurance
  - International Business Machines Corporation
  - java
  - lightweight applications
  - megabus
  - Microsoft
  - Microsoft Access
  - Microsoft Corporation
  - Microsoft SQL Server
  - MP3
  - MP3 library
  - MusicBrainz
  - MySQL
  - online course
  - open source products
  - Oracle
  - Oracle Corporation
  - PostgreSQL
  - SQLite
  - technology sector
  - Visual Foxpro
  - web frontend
  - Windows
title: Database Systems Compared
url: /2009/03/10/database-systems-compared/
---

My first experiences of a computer started with
[DBase III+](http://en.wikipedia.org/wiki/DBase "Dbase on Wikipedia")which is
now [dBASE](http://www.dbase.com/ "dBASE"), then went on to
[Foxpro](http://en.wikipedia.org/wiki/FoxPro_2 "Foxpro 2 on Wikipedia"), now
[Microsoft Visual Foxpro](http://msdn.microsoft.com/en-us/vfoxpro/bb190288.aspx "Microsoft Visual Foxpro").

I have since used:

- [Filemaker Pro](http://www.filemaker.co.uk/ "Filemaker Pro"),
- [Microsoft Access](http://office.microsoft.com/en-us/access/default.aspx "Microsoft Access"),
- [Microsoft SQL Server](http://www.microsoft.com/sqlserver/2008/en/us/default.aspx "Microsoft SQL Server"),
- [MySQL](http://www.mysql.com/ "MySQL"),
- [PostgreSQL](http://www.postgresql.org/ "PostgreSQL"),
- [SQLite](http://www.sqlite.org/ "SQLite") and
- [HSQLDB](http://hsqldb.org/ "HSQLDB").

I have not yet used:

- [IBM DB2](http://www.ibm.com/software/data/db2/ "IBM DB2"),
- [Oracle](http://www.oracle.com/index.html "Oracle").

[Wikipedia has a list of database systems](http://en.wikipedia.org/wiki/Comparison_of_relational_database_management_systems "Compare DB Systems").

<!--more-->

Having worked with this range of database systems and having done copious
amounts of research into DB2, Oracle and other DB systems I have not mentioned,
I like answering the age old questions. Which is the best database system?

Ah! if only it was that simple. There is no database system that is appropriate
for any given requirement. But then, if you have been in the technology sector
long enough, you would already know that. It\'s all about using the right tool
for the job.

I separate these systems into two broad categories and Oracle. There are the
Desktop based database systems:

- DBase
- Foxpro
- SQLite
- HSQLDB
- Filemaker Pro
- Microsoft Access
- MySQL

DBase, FoxPro, Filemaker Pro and Microsoft Access are essentially a GUI frontend
that has a database backing.

Access is the best choice for this purpose under the majority of circumstances.
Filemaker Pro is relevant in some. The usual reason to use DBase or FoxPro is
simply that the developer is used to it. This is not a good enough reason.

I have used DBase III+ for developing an office management suite back in 1994. I
have since used Filemaker Pro to develop a simple contact management database in
1998, Microsoft Access to develop a patient management system for a clinic.

SQLite, HSQLDB and MySQL are database engines that are to be utilised by popping
a frontend on top; sometimes the frontend is Microsoft Access. Microsoft Access
can also be used for its database engine.

Access is usually the worst choice for this except as a stopgap. There are
exceptions to this. One is for a web frontend if the site is not too busy and
its running on a microsoft platform. You don\'t have to go to the hassle of
installing anything on the server. The drivers will take care of it all.

HSQLDB becomes an obvious choice for a light java based application and SQLite
for any other lightweight applications.

MySQL is substantially more powerful and scales a lot better. I include it in
this section because it is a server grade database system that can also work
well in a desktop environment.

I have used Access for several web based systems and I have used HSQLDB for unit
testing hibernate and for a quick and dirty MP3 library that linked into
[musicBrainz](http://musicbrainz.org/ "Musicbrainz"). I have used SQLite in
passing to be utilised by open source products.

I have used MySQL with an Access frontend as a management suite for a website as
well.

And we have the server based database systems:

- MySQL
- Microsoft SQL Server
- IBM DB2
- PostgreSQL

MySQL was used as the backed database system for the edFringe.com website. This
was the perfect choice since the most important requirement was speed.
Particuarly with the Query Cache and Master Slave replication, MySQL was the
best choice.

SQL Server was used as the backend system for an online course for the Scottish
Enterprise around 1999/2000. While MySQL would have been a good choice this, it
was not of production quality at the time.

We have also used Ms SQL Server for an insurance company since all the
infrastructure was based on Windows and PostgreSQL did not have a viable Windows
version at the time.

We use PostgreSQL for megabus. While speed is absolutely critical, it is a
ticketing system which means that transactionality is absolutely critical.

While MySQL now has transactionality with innodb, it is still nowhere near as
good as the transactionality provided by PostgreSQL through MVCC (Multi-version
Concurrency Control). We could have used Ms SQL Server but the cost savings are
dramatic.

To summarise, each system has a specific use, specific strengths and weaknesses
and which should be used is highly dependent on what it is to be used for. I am
hopeful that the summary of what we have used each of these systems for us
useful in determining which one is best placed to solve any specific problem :-D

We have not yet used Oracle and it was a strong contender for megabus but the
serious heavyweight functionality provided by Oracle comes at a price and it is
not yet a cost effective option.
