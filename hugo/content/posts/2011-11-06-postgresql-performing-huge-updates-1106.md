---
categories:
  - Database Systems
date: "2011-11-06T12:45:41Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:52"
  _publicize_pending: "1"
  _wp_old_slug: "718"
  oc_commit_id: http://drone-ah.com/2011/11/06/postgresql-performing-huge-updates-1106/1320583543
  original_post_id: "718"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - Data management
  - Databases
  - memory configuration
  - Performance
  - PostgreSQL
  - SQL
  - Temporary file
title: PostgreSQL performing huge updates [1106]
url: /2011/11/06/postgresql-performing-huge-updates-1106/
---

PostgreSQL is a pretty powerful database server and will work with almost any
settings thrown at it. It is really good at making do with what it has and
performing as it is asked.

We recently found this as we were trying to update every row in a table that had
over eight million entries. We found in the first few tries that the update was
taking over 24 hours to complete which was far too long for an update script.

Our investigation of this led us to the pgsql_tmp folder and the work_mem
configuration parameter.

When the query was being executed, we checked the pgsql_tmp folder to see how
was space being utilised in there. We already knew about the pgsql table from
past experience. We had a server running out of disk space and rapidly. We had
narrowed it down into this folder. In cancelling the query referenced by the tmp
files in here, we were able to free up literally gigabytes of disk space\...

<!--more-->

We had found roughly half a gig of temporary files in here. This led us to
investigate the configuration file.

The one parameter that stuck out was work_mem which was set to a default of 1mb
which I guess might make sense under most circumstances but not in this one.
According to the postgresql documentation

> `work_mem` (`integer`)
>
> Specifies the amount of memory to be used by internal sort operations and hash
> tables before switching to temporary disk files. The value is defaults to one
> megabyte (`1MB`). Note that for a complex query, several sort or hash
> operations might be running in parallel; each one will be allowed to use as
> much memory as this value specifies before it starts to put data into
> temporary files. Also, several running sessions could be doing such operations
> concurrently. So the total memory used could be many times the value
> of `work_mem`; it is necessary to keep this fact in mind when choosing the
> value. Sort operations are used for `ORDER BY`, `DISTINCT`, and merge joins.
> Hash tables are used in hash joins, hash-based aggregation, and hash-based
> processing of `IN` subqueries.

This would tell us that the total memory usage with work_mem could be several
times the value set here and setting it to half a gig would probably be a
terrible idea for a heavily utilised production server. However, for the
migration process when we need to update over 8,000,000 rows, it might be a good
temporary fix.

After updating the work_mem to 512mb, we found that no more tmp files were
created and the whole thing was done in memory.

When updating so many rows, there area a few other things to consider.

Firstly, autovacuum will likely kick in several times to vacuum the table.
You\'ll probably want to disable this for the duration of the update statement
and run a vacuum afterwards.

```sql
    --disable auto vacuum
    ALTER TABLE sometable SET (
      autovacuum_enabled = false, toast.autovacuum_enabled = false
    );
```

You can switch autovacuum back on after the update statement has completed

```sql
    --enable auto vacuum
    ALTER TABLE sometable SET (
      autovacuum_enabled = true, toast.autovacuum_enabled = true
    );
```

A few other things you want to take a look at are the

- fsync parameter (I usually have this set to off anyway since the servers are
  pratically fully redundant)
- checkpoint_segments: I changed this to roughly 5 times the original value
  (check the log to see if it says that its checkpointing too often)
- checkpoint_completion_target: I changed this to 0.9

With all of these updates, we were able to bring the total time of the update
down to a few hours.
