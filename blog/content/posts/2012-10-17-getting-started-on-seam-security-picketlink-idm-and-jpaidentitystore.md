---
categories:
  - Java (EE)
date: "2012-10-17T10:40:28Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:56"
  _last_editor_used_jetpack: block-editor
  _publicize_pending: "1"
  _wp_old_slug: "886"
  oc_commit_id: http://drone-ah.com/2012/10/17/getting-started-on-seam-security-picketlink-idm-and-jpaidentitystore/1350466831
  original_post_id: "886"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Java (EE)
  - Identity management
  - idm
  - picketlink
  - seam
  - seam-security
  - seam3
title: Getting started on seam-security, picketlink IDM and JPAIdentityStore
url: /2012/10/17/getting-started-on-seam-security-picketlink-idm-and-jpaidentitystore/
---

I love how JBoss 7(.1) has everything working out of the box - not much fiddling
with jars or suchlike and with Arquillian, everything really was a treat to get
started on a new project. This was until I had to sort out security with
seam-security.

To be fair, the main issue was just poor documentation. It took me a day to sort
out what should essentially have taken an hour(or two)

The documentation you get to from
http://www.seamframework.org/Seam3/SecurityModule seems to be out of date. The
fact that the page referes to version 3.0.0.Alpha1 and Alpha2 should have tipped
me off but the url for the doc suggested it was the latest.

The more up to date documentation I found was
at http://docs.jboss.org/seam/3/3.1.0.Final/reference/en-US/html/pt04.html

I followed
[chapter 33](http://docs.jboss.org/seam/3/3.1.0.Final/reference/en-US/html/security-identitymanagement.html "Identity Management")
on there and I won\'t repeat it here for the sake of brevity.

What follows are the additional steps I had to take to get it to work.

<!--more-->

I ran into a javax.enterprise.inject.CreationException, the relevant part of the
stack trace being:

```
    Caused by: java.lang.IllegalArgumentException: targetClass parameter may not be null
        at org.jboss.solder.properties.query.PropertyQuery.(PropertyQuery.java:54) [solder-impl-3.1.0.Final.jar:3.1.0.Final]
        at org.jboss.solder.properties.query.PropertyQueries.createQuery(PropertyQueries.java:39) [solder-impl-3.1.0.Final.jar:3.1.0.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.findNamedProperty(JpaIdentityStore.java:441) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.configureRoleTypeName(JpaIdentityStore.java:877) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.bootstrap(JpaIdentityStore.java:328) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        at org.picketlink.idm.impl.configuration.IdentityConfigurationImpl.createRealmMap(IdentityConfigurationImpl.java:192) [picketlink-idm-core-1.5.0.Alpha02.jar:1.5.0.Alpha02]
        at org.picketlink.idm.impl.configuration.IdentityConfigurationImpl.buildIdentitySessionFactory(IdentityConfigurationImpl.java:147) [picketlink-idm-core-1.5.0.Alpha02.jar:1.5.0.Alpha02]
        ... 109 more
```

To resolve this,  I had to add in the \@IdentityEntity Annotation to the
IdentityObjectType class

```java
    @Entity
    @IdentityEntity(EntityType.IDENTITY_ROLE_NAME)
    public class IdentityObjectType {
    ...
```

The next exception was org.picketlink.idm.common.exception.IdentityException:
Error creating identity object. The relevant part of the strack trace being:

```
    Caused by: java.lang.NullPointerException
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.lookupIdentityType(JpaIdentityStore.java:966) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.createIdentityObject(JpaIdentityStore.java:999) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        ... 87 more
```

It turned out that the entitymanager was not being picked up and it was null.
This part was probably in the documentation earlier with regards to configuring
seam but I had skipped directly to the security section so missed it. We need to
define the persistence unit with the beans.xml. I have included my full file
below.

```xml
    <?xml version="1.0"?>
    <beans xmlns="http://java.sun.com/xml/ns/javaee"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xmlns:em="urn:java:javax.persistence"
        xmlns:s="urn:java:ee"
        xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://jboss.org/schema/cdi/beans_1_0.xsd">

        <interceptors>
            <class>org.jboss.seam.security.SecurityInterceptor</class>
        </interceptors>

        <em:EntityManager>
            <s:Produces />
            <em:PersistenceContext unitName="invision-users" />
        </em:EntityManager>
    </beans>
```

This brought us further forward still. The next exception was:

```
    javax.persistence.NoResultException: No entity found for query
        at org.hibernate.ejb.QueryImpl.getSingleResult(QueryImpl.java:286) [hibernate-entitymanager-4.0.1.Final.jar:4.0.1.Final]
        at org.hibernate.ejb.criteria.CriteriaQueryCompiler$3.getSingleResult(CriteriaQueryCompiler.java:264) [hibernate-entitymanager-4.0.1.Final.jar:4.0.1.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.lookupCredentialTypeEntity(JpaIdentityStore.java:1112) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.updateCredential(JpaIdentityStore.java:1633) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        at org.picketlink.idm.impl.repository.WrapperIdentityStoreRepository.updateCredential(WrapperIdentityStoreRepository.java:310) [picketlink-idm-core-1.5.0.Alpha02.jar:1.5.0.Alpha02]
        at org.picketlink.idm.impl.api.session.managers.AttributesManagerImpl.updatePassword(AttributesManagerImpl.java:563) [picketlink-idm-core-1.5.0.Alpha02.jar:1.5.0.Alpha02]
```

This was related to missing data in the database. It needed a credential type. I
created one for password.

```sql
    INSERT INTO CredentialType(id, name) VALUES (1, 'password');
```

This brought us forward on to the next exception:
org.picketlink.idm.common.exception.IdentityException: Exception creating
relationship

with the relevant part of

```
    Caused by: javax.persistence.NoResultException: No entity found for query
        at org.hibernate.ejb.QueryImpl.getSingleResult(QueryImpl.java:286) [hibernate-entitymanager-4.0.1.Final.jar:4.0.1.Final]
        at org.hibernate.ejb.criteria.CriteriaQueryCompiler$3.getSingleResult(CriteriaQueryCompiler.java:264) [hibernate-entitymanager-4.0.1.Final.jar:4.0.1.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.lookupRelationshipType(JpaIdentityStore.java:1127) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        at org.jboss.seam.security.management.picketlink.JpaIdentityStore.createRelationship(JpaIdentityStore.java:1066) [seam-security-3.1.0.Final.jar:3.1.0.Final]
        ... 86 more
```

This was solved by adding in a relationship type

```sql
    INSERT INTO RelationshipType(id, name) VALUES (1, 'JBOSS_IDENTITY_MEMBERSHIP');
```

Both the sql statements were put into import.sql and hibernate is configured to
create tables. My test case is as follows. It was taken
from <https://github.com/seam/seam-example-confbuzz/blob/develop/src/test/java/seam/example/confbuzz/test/integration/LoginIntegrationTest.java> and
modified.

```java
    @RunWith(Arquillian.class)
    public class LoginIntegrationTest {

        @Inject
        private IdentitySession identitySession;

        @Inject
        private Identity identity;

        @Inject
        @DefaultTransaction
        SeamTransaction tx;

        @Deployment(name = "authentication")
        public static Archive createLoginDeployment() {
            // This is the simplest way to test the full archive as you will be
            // deploying it
            final MavenDependencyResolver resolver =
                    DependencyResolvers.use(MavenDependencyResolver.class)
                        .loadMetadataFromPom("pom.xml")
                        .goOffline();

            Archive archive = ShrinkWrap
                    .create(WebArchive.class)
                    .addPackages(true, "uk.co.kraya.test-seam.auth")
                    .addAsResource("META-INF/test-persistence.xml", "META-INF/persistence.xml")
                    .addAsResource("META-INF/beans.xml", "META-INF/beans.xml")
                    .addAsResource("test-import.sql", "import.sql")
                    .addAsLibraries(resolver.artifact("org.jboss.seam.security:seam-security").resolveAsFiles());

            System.out.println(archive.toString(true));

            return archive;

        }

        @Before
        public void setupTestUser() throws IdentityException, SystemException,
                NotSupportedException, RollbackException,
                HeuristicRollbackException, HeuristicMixedException {

            if (!tx.isActive())
                tx.begin();

            final PersistenceManager pm = identitySession.getPersistenceManager();
            final AttributesManager am = identitySession.getAttributesManager();
            final RelationshipManager rm = identitySession.getRelationshipManager();

            // Setup the group we want our user to belong to
            final Group memberGroup = pm.createGroup("member2", "USER2");
            final User user = pm.createUser("test");

            am.updatePassword(user, "password");

            rm.associateUser(memberGroup, user);

            tx.commit();
        }

        @Test
        public void assertUserCanAuthenticate(Credentials credentials) {
            credentials.setUsername("test");
            credentials.setCredential(new PasswordCredential("password"));
            assertEquals(identity.login(), Identity.RESPONSE_LOGIN_SUCCESS);
        }
```

Do comment and let me know if it helped :-D
