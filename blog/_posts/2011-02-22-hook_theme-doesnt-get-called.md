---
layout: post
title: hook_theme doesn't get called
date: 2011-02-22 20:20:11.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - drupal
tags:
  - php
  - hook_theme
  - theming
  - Web Development
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2011/02/22/hook_theme-doesnt-get-called/1298406012
  restapi_import_id: 591d994f7aad5
  original_post_id: "412"
  _wp_old_slug: "412"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:40"
permalink: "/2011/02/22/hook_theme-doesnt-get-called/"
---

I was developing a new module in drupal and it needed a theme function to be
implemented.

As per the instructions, it was implemented as follows (to use a template)

```phg
/**
 * Implementation of hook_theme().
 */
function my_module_results_theme($existing, $type, $theme, $path) {

    return array(
        'my_block' => array(
            'template' => 'my_block',
            'arguments' => array(
                'var1' => NULL
            )
        )
    );
}
```

However, when trying to apply the theme, it didn't work. I tried various things
and identified that the hook above was just not being called. A little bit of
digging helped me discover that themes are cached. This happens even in the dev
mode. To resolve this, go to

`Administer  -> Performance -> Clear Cached Data` (right at the bottom of the
page)

and et voila my theme was now being utilised.
