---
categories:
- Java (EE)
date: "2008-12-23T22:14:42Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:35"
  _publicize_job_id: "5181454935"
  _wp_old_slug: "103"
  original_post_id: "103"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
- Code Coverage
- Data Testing
- DbUnit
- DOM
- Domain Model
- Domain Object Model
- Ejb3Unit
- Entity Testing
- Hibernate
- java
- SuperCsv
- testing
- Testing Framework
- TestNG
title: Hibernate Domain Model Testing
type: post
url: /2008/12/23/hibernate-domain-model-testing/
---

One of my pet peeves with Hibernate has always been how difficult it was to test
it. I want to test the persistence of data, loading the data back and any
specific funtionality with the domain model.

Simple? NO! The main problem was the management of the data set. I had set up,
in the past fairly interesting classes to test the functionality using
reflection, and injecting the data from the classes themselves through the data
provider mechanism of [TestNG](http://testng.org/d "TestNG"). However, this was
error prone and clunky at best. It also made dependency management of data quite
cumbersome.

With a view to resolving this, I also looked at
[DbUnit](http://dbunit.sourceforge.net/ "DbUnit"),
[unitils](http://unitils.org/ "Unitils") and
[Ejb3Unit](http://ejb3unit.sourceforge.net/ "Ejb3Unit"). They all did some
things that I liked but lacked some functionality that was important.

<!--more-->

This led me to write a simple testing infrastructure. The goal was
straightforward.

- I need to be able to define data in a CSV (actually it was seperated by the
  pipe character |, so PSV) based on entities.
- The framework should automatically persist the data (and fail on errors)
- It should test that it can load all that data back
- It should run as many automated tests on the DOM as possible.

The framework uses the CSV files to read the data for each of the classes (using
the excellent [SuperCsv](http://supercsv.sourceforge.net/ "SuperCsv") library).
It needs an Id field for internal reference. As long as the id's match within
the CSV files for the relationships, it will be persisted correctly into the
database even when the persisted id's are different.

For example, I could have a Contact.csv with 5 records (ids 1 through 5) and a
Company.csv with 3 records (ids 1 through 3).

The Contact.csv records can map to the id specified in the Company.csv file and
when the records get persisted, they will be associated correctly, even if the
id's in the database end up being different.

The framework also looks for the CSV file which has the same name as the class
within the location defined within the configuration file. This means that as
long as the filename matches the class name, the data loading is automatic.

For simple classes, the Test case is as simple as:

> public class CompanyTest extends DOMTest<Company> {
>
> public CompanyTest() { super(Company.class); } }

The system (with the help of testNG) is also easily flexible to define object
model dependencies. Just override the persist method (which just calls the
super.persist) and define the groups to be persist and <object>.persist

in this particular case, it would be

> @override
>
> @Test(groups={"persist", "Company.persist"}
>
> public void persist() {
>
> super.persist();
>
> }

For all dependent classes, I then depend on the Company.persist group (For the
ContactTest class for example, since it needs to link to the Company object)

You can specify OneToOne and ManyToOne relationships with just the CSV files -
just defining the field name and the id of the object to pull in.

ManyToMany is more complex and requires an interim object to be created within
the test section. If the Contact to Company relationship above was ManyToMany,
we would create a ContactCompany class with just the two fields - Contact &
Company, then create a csv file with three fields, id, Contact, & Company. The
framework currently always needs an id field.

You would then need to write a method within the ContactTest or CompanyTest(I
use the owning side) to read the CSV file in and pump the data. This process is
a little bit complex just now.

With an appropriate amount of test data, you are able to write a test suite that
can consistently test your domain model. More importantly, you can configure it
to drop the database at the start of each run so that once the tests are
complete, you have a database structure and data than can be used for testing of
higher level components (EJB/Spring/UI/WebApp)

We currently use this framework to test the domain model as well as distribute a
data set for development and testing of the higher tier functionalities.

For the future, there are several additional features this framework needs:

- It currently needs the setters/getters & constructors to be public. This needs
  to be FIXED
- Refactor the ManyToMany Relationship code to make it easier and simpler to
  test and pump data
- See if we can ensure that additional tests which data is done within a
  transaction and rolled back so that the database is left in the "CSV Imported"
  state on completion of tests
- Easier Dependency management if possible

This framework is still inside the walls of Kraya, but once the above issues are
resolved and it is in a releasable state, it will be published into the open
source community. If you are interested in getting a hold of it, email me and
I'll provide you with the latest version.

The easier and quicker it is to test, the more time we can spend on writing
code... :-) The higher the coverage of the tests, the more confident you can be
of your final product.

To more testing...
