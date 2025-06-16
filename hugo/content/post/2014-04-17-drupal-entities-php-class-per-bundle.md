---
categories:
- drupal
date: "2014-04-17T11:18:08Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:59"
  _publicize_pending: "1"
  _wp_old_slug: "953"
  _wpt_short_url: http://drone-ah.com/2014/04/17/drupal-entities-php-class-per-bundle/
  oc_commit_id: http://drone-ah.com/2014/04/17/drupal-entities-php-class-per-bundle/1397729891
  original_post_id: "953"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
published: true
status: publish
tags:
- drupal-entities
- php
title: Drupal Entities - PHP Class per bundle
type: post
url: /2014/04/17/drupal-entities-php-class-per-bundle/
---

If you would like a bit more polymorphism in your drupal entities, this might
cheer you up :-D

I was looking for a way to have a class hierarchy that matched the bundle
"hierarchy" of entities in drupal. Yes, they are all "subclasses" of ONE parent,
but it is still useful to be able to have a class per bundle.

The
[entity bundle plugin](https://drupal.org/project/entity_bundle_plugin "entity bundle plugin") does
a good job of providing a plugin framework to instantiate classes per bundle
type. There is also
[an example of how to use this](http://bojanz.wordpress.com/2013/07/19/entity-bundle-plugin/ "how to use entity bundle plugin").
However, this was a bit of overkill for me. I did however borrow the idea (and
some code) to implement it in a simpler fashion.

Implement a custom controller and override the create and the query methods

```phg
    class MyEntityAPIController extends EntityAPIController {
      /**
       * Overrides EntityAPIController::query().
       */
      public function query($ids, $conditions, $revision_id = FALSE) {
        $query = $this->buildQuery($ids, $conditions, $revision_id);
        $result = $query->execute();
        $result->setFetchMode(PDO::FETCH_ASSOC);

        // Build the resulting objects ourselves, since the standard PDO ways of
        // doing that are completely useless.
        $objects = array();
        foreach ($result as $row) {
          $row['is_new'] = FALSE;
          $objects[] = $this->create($row);
        }
        return $objects;
      }

      /**
       * Overrides EntityAPIController::create().
       */
      public function create(array $values = array()) {
        if (!isset($values[$this->entityInfo['entity keys']['bundle']])) {
          throw new Exception(t('No bundle provided to MyEntityAPIController::create().'));
        }

        $bundle = $values[$this->entityInfo['entity keys']['bundle']];
          // Add is_new property if it is not set.
        $values += array(
          'is_new' => TRUE,
        );

        $default_class = isset($this->entityInfo['entity class']) ? $this->entityInfo['entity class'] : NULL;
        $class = isset($this->entityInfo['bundles'][$bundle]['entity class']) ? $this->entityInfo['bundles'][$bundle]['entity class'] : $default_class;
        if (!class_exists($class)) {
          $class = "Entity";
        }

        return new $class($values, $this->entityType);
      }

    }
```

I can now define a PHP class for each bundle as follows in hook_entity_info

```php
      $entity['myentity'] = array(
        'label' => t('myentity'),
        'module' => 'mymodule',
        'entity class' => 'MyEntity',  // Default entity class if not defined in bundle
        'controller class' => 'MyEntityAPIController',
        'base table' => 'myentity',
        'entity keys' => array(
          'id' => 'id',
          'label' => 'name',
          'bundle' => 'bundle',
        ),
        'bundles' => array(
          'bundle1' => array(
            'label' => 'Bundle 1',
            'entity class' => 'Bundle1Entity',
          ),
          'bundle2' => array(
            'label' => 'Bundle 2',
            'entity class' => 'Bundle2Entity',
          ),
        ),
```

Don't forget to clear the cache and you should be able to get bundle specific
classes instantiated. It will fall back to the 'entity class' defined for the
entity if the bundle 'entity class' is not defined or cannot be found.
