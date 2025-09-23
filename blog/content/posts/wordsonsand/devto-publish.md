---
title: "Publishing from hugo to dev.to"
date: 2025-09-23T09:30:54+01:00
tags:
  - wordsonsand
  - hugo
  - devto
---

I have been pondering federating parts of my blog to [dev.to](https://dev.to/)
for a bit more visibility.

However, I had a couple of issues:

- With html (instead of markdown), it would not pick up the code blocks
  correctly
- With markdown, it would render the relative links incorrectly

What I really wanted was a way to render the hugo markdown into Jekyll style
(which is what forem wants) but with the links rendered.

This was a little more complicated than I would have liked.

## Goals

At a minimum, I wanted two main things:

- Render code blocks correctly
- Relative URLs should be rendered as absolute (because they won't work on
  dev.to)

## Pre-existing Solutions

I found [hugodevto](https://github.com/maelvls/hudevto) which looked promising
except:

- Not a fan of having to manually update hundreds of posts with the devto id
- Had a couple of fiddly bits to get it working (my plain text outputs had some
  troubles for unknown reasons)
- It rendered image urls,
  [but not regular urls.](https://github.com/maelvls/hudevto/issues/2#issuecomment-3302934120)

Ultimately though, it felt a bit bulkier than what I was looking for

## Using Hugo

I had [done enough work with hugo](/tags/hugo) and
[output formats](/tags/inscribe) to have a vague idea of how to make it work.

### Limitations

There are a few limitations to doing it this way though. Hugo
[does not provide render hooks for everything](https://gohugo.io/render-hooks/introduction/).
You will end up with html in the output. However, since Forem (and Jekyll) will
just render them, it fits my use case. It won't work as well if you try and use
this to generate like for like markdown usable in Jekyll.

### Define a new content type

We want a new content type which will output markdown

```toml
[outputFormats.jekyll]
    mediaType = "text/markdown"
    baseName = "index"
    isPlainText = true
    isHTML = false
    notAlternative = true
    path = 'jekyll'        # put the output in `public/jekyll` so it's easier to find

[outputs]
    page = ['html', 'jekyll'] # Output all pages in our jekyll format as well
```

You also need a basic template before hugo will output our markdown files.

[`layouts/_default/single.jekyll.md`](../../../layouts/_default/single.jekyll.md)

```gotemplate
---
title: {{ .Title }}
published: true
date: {{ .Date }}
{{- with .Params.tags }}
tags: [{{ delimit . ", " }}]
{{- end }}
canonical_url: {{ .Permalink }}
---

{{ .Content }}
```

With this, if you `hugo build`, it'll render the `.md` files, but the content
will be html.

### Render code blocks as markdown

We can use the
[code block render hook](https://gohugo.io/render-hooks/code-blocks/) to
"convert them" back to markdown.

[`layouts/_default/_markup/render-codeblock.jekyll.md`](../../../layouts/_default/_markup/render-codeblock.jekyll.md)

````gotemplate
```
{{ .Type }}
{{ .Inner }}
```
````

### Render absolute urls

I already have a pretty extensive
[`render-link`](../../../layouts/_default/_markup/render-link.html) so updating
it was just a case of making a copy of it and replacing relative url references
with absolute ones.

It's fine for it to be in html because Forem will still render it correctly.
They could be rendered as markdown and it should work just as well.

[`layouts/_default/_markup/render-link.jekyll.md`](../../../layouts/_default/_markup/render-link.jekyll.md)

```gotemplate
{{ /* other content */ }}
  <a href="{{ printf "%s#%s" .PageInner.Permalink $u.Fragment | safeURL }}" {{ with .Title }}title="{{ . }}"{{ end }}>{{ $text }}</a>
{{ /* other content */ }}
```

### Output

With these relatively minor changes, I was able to render markdown files I could
then pop into dev.to and it works for the handful I set up.

## Next Steps

### Images

One big glaring omission is images - it's not as relevant for me because I
rarely use images.

I expect it to be easy enough to use the
[image render hook](https://gohugo.io/render-hooks/images/) to achieve this.

### Tags

One of the issues I have is that dev.to has a strict tag limit of four.
Currently I manually edit that when I create the post on dev.to.

It would be better to have a `devto_tags` field because my local content tags
aren't always relevant for dev.to.

I could also write a script to automate the setting of the `devto_tags` field
automatically based on the first four tags, and mapping from my tags to dev.to
tags if necessary by building a small map data set somewhere.

### Automation

Once the above two are done, it would be good to automate it. I could

- [get all the posts](https://developers.forem.com/api/v0#tag/articles/operation/getUserAllArticles)
- get all the local posts
- map them based on the canonical url (which is returned by the endpoint)
- upload updated ones.

Not all my posts are technical, so I'd also want to add a field to the
frontmatter (`devto_published`) and figure out a way to push updates only if
there are changes.

## Conclusion

I have a working solution for now - and if dev.to brings enough traffic / value,
then I'll consider spending a bit more time adding polish.

For the time being though, seems to work!

Feel free to use any of this code (such that it is) as you wish.
