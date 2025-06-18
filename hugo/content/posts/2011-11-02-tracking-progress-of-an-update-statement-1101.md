---
categories:
  - Database Systems
date: "2011-11-02T19:59:02Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:51"
  _publicize_pending: "1"
  _wp_old_slug: "688"
  original_post_id: "688"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - computing
  - Database
  - Databases
  - PostgreSQL
  - SQL
  - SQL keywords
  - Update
title: Tracking progress of an update statement [1101]
url: /2011/11/02/tracking-progress-of-an-update-statement-1101/
---

Sometimes there is a need to execute a long running update statement. This
update statement might be modifying millions of rows as was the case when we
went hunting for a way to track the progress of the update. Hunting around took
us to <http://archives.postgresql.org/pgsql-admin/2002-07/msg00286.php> In our
particular case, we are using postgresql but this should work with any database
server that provides sequences. Our original sql was of the form:

```sql
update only table1 t1
set amount = t2.price
from table2 t2
where t1.id = t2.id;
```

There is of course now way of figuring out how many rows had been updated
already. The first step was to create a sequence

```sql
CREATE TEMPORARY SEQUENCE seq_progress START 1;
```

<!--more-->

We can then use this sequence in the update statement to ensure that each row
updated also increments the sequence

```sql
update only table1 t1
set amount = t2.price
from table2 t2
where nextval('seq_progress') != 0
and t1.id = t2.id;
```

Once the query is running, you can open another connection to the database. To
get an indication of how far it has got, you can just run the following

```sql
select nextval('seq_progress');
```

Bear in mind that this will also increment it by 1 but if you have millions of
rows which is really the only case in which this would be useful, a few
additional increments is hardly going to make a difference.

Good luck and have fun!
