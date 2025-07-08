---
title: "Generate YouTube Descriptions from Hugo"
date: 2025-07-03T13:18:15+01:00
categories:
  - wordsonsand
tags:
  - wordsonsand
  - projector
  - hugo
  - youtube
  - automation
---

Uploading and setting up YouTube videos is fiddly. There are a lot of things to
get right - title, description, chapters, links, tags - the list goes on.

I also want to link to and from blog posts and social posts - and making sure
those links stay in sync is a hassle.

It gets more complicated when scheduling multiple videos.

I wanted to make it easier

<!-- more -->

I (at the time of writing) use hugo for this blog site, and I regularly link
from the YouTube description to a page on here. Since I want to save having to
copy and paste that link into YouTube, leveraging the CMS felt sensible.

## A YouTube Content Type

[archetypes/youtube.md](../../../archetypes/youtube.md)

```yaml
---
title: "{{ replace .Name "-" " " | title }}"
publishDate: {{ .Date }}
youtubeId: ""
playlist: ""
categoryId: 20
tags: []
chapters:
  - "0:00 Intro"
links:
  - title: <title>
    url: <url>
_build:
  list: never
  render: always
  publishResources: false
sitemap: false
---
```

- `title`: Title for the youtube video
- `publishDate`: When should the video go live
- `youtubeId`: The video id from YouTube, used to build links
- `playlist`: Which playlist is this a part of? Used to build links
- `categoryId`:
  [YouTube category Id](https://mixedanalytics.com/blog/list-of-youtube-video-category-ids/) -
  e.g. gaming
- `tags`: Used to add hashtags at the end of the video
- `chapters`: Added to description to demarcate chapters
- `links`: adds each link to the description. Can be other youtube videos
- `outputs`: needed to output in plaintext format
- `_build` and `sitemap`: prevent this file getting linked/crawled

### Cascaded Properties

We also want to prevent these pages from showing up on:

- List Pages
- Sitemap

We also want to prevent them from being published. We could add `_build` to each
of the `md` files, or we can cascade it (thanks to
[jmooring](https://discourse.gohugo.io/u/jmooring) from the gohugo discourse).

[content/youtube/\_index.md](https://github.com/drone-ah/wordsonsand/tree/main/blog/content/youtube/_index.md)

```yaml
title: "YouTube"
cascade:
  _build:
    list: never
    render: always
    publishResources: false
  sitemap: false
```

### Layout (plain text)

We need a plaintext template to render it as text

[layouts/\_default/single.plain.txt](https://github.com/drone-ah/wordsonsand/tree/main/blog/layout/_default/single.plain.txt)

```gotmpl
{{ .Content | plainify | htmlUnescape }}

{{- if .Params.links }}
Links:
{{- range .Params.links }}
{{ .title }}: {{ .url | absURL }}
{{- end }}
{{ end }}

{{- if .Params.chapters }}
{{ range .Params.chapters }}
{{- . }}
{{ end -}}
{{ end }}

{{- if .Params.tags }}
{{ range .Params.tags }}#{{ . }} {{ end }}
{{ end }}
```

### Enable plaintext output

We also need to define plain as an output format.

As far as I could see, there is no way (currently) in hugo to specify a default
output type for a `type` (i.e. youtube) of content, only a `kind` (e.g. page) of
content.

However, we can add this to the cascade as well:

```yaml
cascade:
  outputs: ["plain"]
```

I also created a `layouts/_default/list.plain.txt` file to avoid the error:

`WARN  found no layout file for "plain" for kind "section": You should create a template file which matches Hugo Layouts Lookup Rules for this combination.`

The contents of this file doesn't really matter as we shouldn't be rendering or
using it.

[hugo.toml](https://github.com/drone-ah/wordsonsand/tree/main/blog/hugo.toml)

```toml
[outputFormats.plain]
	mediaType = "text/plain"
	baseName = "index"
	isPlainText = true
	isHTML = false
	noUgly = true
```

## Auto link to YouTube

I'd like to be able to link to a local markdown file, and have that resolve to
the correct YouTube URL.

### From Posts

[layouts/\_default/\_markup/render-link.html](https://github.com/drone-ah/wordsonsand/blob/main/blog/layouts/_default/_markup/render-link.html)

```gotmpl
{{- if eq $page.Type "youtube" -}}
  {{- $href = printf "https://www.youtube.com/watch?v=%s" $page.Params.youtubeId -}}
{{- else -}}
  <a href="{{ $page.RelPermalink | safeURL }}" {{ with .Title }}title="{{ . }}"{{ end }}>{{ $text }}</a>
{{- end -}}

<a href="{{ $href | safeURL }}">{{ $text }}</a>
```

You know what would be nicer? If it took the user to the video in the playlist -
if playlist is defined

```gotmpl
{{- if eq $page.Type "youtube" -}}
  {{- $href = printf "https://www.youtube.com/watch?v=%s" $page.Params.youtubeId -}}
  {{- with $page.Params.playlist }}
    {{- $href = printf "%s&list=%s" $href . -}}
  {{- end }}
{{- else -}}
  <a href="{{ $page.RelPermalink | safeURL }}" {{ with .Title }}title="{{ . }}"{{ end }}>{{ $text }}</a>
{{- end -}}

<a href="{{ $href | safeURL }}">{{ $text }}</a>
```

### From YouTube Description

Let's render links to YouTube from the `links` property:

[layouts/\_default/single.plain.txt](https://github.com/drone-ah/wordsonsand/tree/main/blog/layout/default/single.plain.txt)

```gotmpl
{{ with .Params.links }}
Links:
{{- $this := $.Page }}
{{ range . -}}
  {{- $target := $this.GetPage .url -}}
  {{- if and $target (eq $target.Type "youtube") -}}
    {{- $href := printf "https://www.youtube.com/watch?v=%s" $target.Params.youtubeId -}}
    {{- with $target.Params.playlist -}}
      {{- $href = printf "%s&list=%s" $href . -}}
    {{- end -}}
    {{ .title }}: {{ $href }}
  {{- else if $target -}}
    {{ .title }}: {{ $target.Permalink }}
  {{- else -}}
    {{ .title }}: {{ .url | absURL }}
  {{- end }}
{{ end }}
{{ end }}
```

### Future Links

While we're at it, let's skip rendering any links that go live in the future:

```gotmpl
{{- $target := $this.GetPage .url -}}
{{- if and $target (eq $target.Type "youtube") (not ($target.PublishDate.After now)) -}}
  {{- $href := printf "https://www.youtube.com/watch?v=%s" $target.Params.youtubeId -}}
  {{- with $target.Params.playlist -}}
    {{- $href = printf "%s&list=%s" $href . -}}
  {{- end -}}
  {{ .title }}: {{ $href }}
{{- else if and $target (ne $target.Type "youtube") -}}
  {{ .title }}: {{ $target.Permalink }}
{{- else if not $target -}}
  {{ .title }}: {{ .url | absURL }}
{{- end }}
```

## Next

This covers the Hugo-side of things.

There are two more parts, that I'd like to happen automatically:

- [Uploading the video](./projector-upload.md)
- [Syncing metadata](./projector-sync.md)

## Links / References

- [Suggestions from `jmooring` on hugo discourse](https://discourse.gohugo.io/t/generating-youtube-descriptions-using-hugo/55233/2?u=drone.ah)
- [cascade](https://gohugo.io/configuration/cascade)
