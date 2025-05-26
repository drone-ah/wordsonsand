---
layout: post
title: Tracking progress of an update statement [1101]
date: 2011-11-02 19:59:02.000000000 +00:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Database Systems
tags:
- computing
- Database
- Databases
- PostgreSQL
- SQL
- SQL keywords
- Update
meta:
  _publicize_pending: '1'
  _edit_last: '48492462'
  oc_metadata: "{\t\tversion:'1.1',\t\ttags: {'computing': {\"text\":\"Computing\",\"slug\":\"computing\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/7644252a-6749-3bf2-8f11-ee98edeb48ad/SocialTag/1\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Computing\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'databases': {\"text\":\"Databases\",\"slug\":\"databases\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/7644252a-6749-3bf2-8f11-ee98edeb48ad/SocialTag/3\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Databases\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'sql-keywords': {\"text\":\"SQL keywords\",\"slug\":\"sql-keywords\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/7644252a-6749-3bf2-8f11-ee98edeb48ad/SocialTag/7\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"SQL
    keywords\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'update': {\"text\":\"Update\",\"slug\":\"update\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/7644252a-6749-3bf2-8f11-ee98edeb48ad/SocialTag/8\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Update\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'postgresql': {\"text\":\"PostgreSQL\",\"slug\":\"postgresql\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/7644252a-6749-3bf2-8f11-ee98edeb48ad/SocialTag/9\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"PostgreSQL\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'sql': {\"text\":\"SQL\",\"slug\":\"sql\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/7644252a-6749-3bf2-8f11-ee98edeb48ad/SocialTag/10\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"SQL\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'database': {\"text\":\"Database\",\"slug\":\"database\",\"source\":{\"url\":\"http://d.opencalais.com/dochash-1/7644252a-6749-3bf2-8f11-ee98edeb48ad/SocialTag/11\",\"subjectURL\":null,\"type\":{\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\",\"_className\":\"ArtifactType\"},\"name\":\"Database\",\"makeMeATag\":true,\"importance\":1,\"_className\":\"SocialTag\",\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  oc_commit_id: http://drone-ah.com/2011/11/02/tracking-progress-of-an-update-statement-1101/1320263945
  restapi_import_id: 591d994f7aad5
  original_post_id: '688'
  _wp_old_slug: '688'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:54:51'
permalink: "/2011/11/02/tracking-progress-of-an-update-statement-1101/"
---

Sometimes there is a need to execute a long running update statement.
This update statement might be modifying millions of rows as was the
case when we went hunting for a way to track the progress of the update.
Hunting around took us
to <http://archives.postgresql.org/pgsql-admin/2002-07/msg00286.php> In
our particular case, we are using postgresql but this should work with
any database server that provides sequences. Our original sql was of the
form:

    update only table1 t1
    set amount = t2.price
    from table2 t2
    where t1.id = t2.id;

There is of course now way of figuring out how many rows had been
updated already. The first step was to create a sequence

    CREATE TEMPORARY SEQUENCE seq_progress START 1;

We can then use this sequence in the update statement to ensure that
each row updated also increments the sequence

    update only table1 t1
    set amount = t2.price
    from table2 t2
    where nextval('seq_progress') != 0
    and t1.id = t2.id;

Once the query is running, you can open another connection to the
database. To get an indication of how far it has got, you can just run
the following

     select nextval('seq_progress');

Bear in mind that this will also increment it by 1 but if you have
millions of rows which is really the only case in which this would be
useful, a few additional increments is hardly going to make a
difference.

Good luck and have fun!
