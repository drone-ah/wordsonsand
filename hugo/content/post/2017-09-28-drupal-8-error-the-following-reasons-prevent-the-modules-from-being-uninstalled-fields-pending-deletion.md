---
categories:
- drupal
date: "2017-09-28T15:53:15Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:01"
  _publicize_job_id: "9761742115"
  geo_public: "0"
parent_id: "0"
password: ""
published: true
status: publish
tags: []
title: 'Drupal 8 error: "The following reasons prevent the modules from being uninstalled:
  Fields pending deletion"'
type: post
url: /2017/09/28/drupal-8-error-the-following-reasons-prevent-the-modules-from-being-uninstalled-fields-pending-deletion/
---

When you try and uninstall a module which has a field that you have used, it can
throw the following error:

```
The following reasons prevent the modules from being uninstalled: Fields pending
deletion\
```

This is an issue in both Drupal 7 and Drupal 8. This is due to the fact that
drupal doesn\'t actually delete the data for the field when you delete the
field. It deletes the data during cron runs. If cron hasn\'t been run enough
times since you deleted the field, drupal won\'t let you uninstall the module.

To force drupal to purge the data, you can run the following command

```bash
drush php-eval \'field_purge_batch(500);\'\
```

Increase 500 to a high enough number to wipe out the data. Afte this has
completed, you should be able to uninstall the module

References:\
[Module uninstall dependencies (drupal stackexchange)](https://drupal.stackexchange.com/questions/184690/module-uninstall-dependencies)\
[Message \"Required by Drupal (Fields Pending Deletion)\" baffles users](https://www.drupal.org/node/1331922)\
[Can\'t uninstall YAML because of following reason: Fields pending deletion](https://www.drupal.org/node/2835035)
