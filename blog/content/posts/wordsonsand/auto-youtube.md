---
title: "Auto Youtube"
date: 2025-07-03T13:18:15+01:00
draft: true
---

Uploading and setting up YouTube Videos is a nightmare. There area lot of things
to get just right. I also want to link to it to/from blog posts and social media
posts. Making sure that I'm using the correct link everywhere is by itself
annoying.

It gets more complicated when scheduling multiple videos.

I wanted to make it easier

<!-- more -->

I (at the time of writing) use hugo for this blog site, and I regularly link
from the YouTube description to a page on here. Since I want to save having to
copy and paste that link into YouTube, leveraging the CMS felt sensible.

## YouTube Content Type

[archetypes/youtube.md](https://github.com/drone-ah/wordsonsand/tree/main/blog/archetypes/youtube.md)

```yaml
---
title: "{{ replace .Name "-" " " | title }}"
publishDate: {{ .Date }}
youtubeId: ""
playlist: ""
tags: []
chapters:
  - "0:00 Intro"
links:
  - title: <title>
  - url: <url>
outputs: ["plain"]
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
- `tags`: Used to add hashtags at the end of the video
- `chapters`: Added to description to demarcate chapters
- `links`: adds each link to the description. Can be other youtube videos
- `outputs`: needed to output in plaintext format
- `_build` and `sitemap`: prevent this file getting linked/crawled

### Layout (plain text)

We'll create a plaintext template to render it as text

[layouts/youtube/single.plain.txt](https://github.com/drone-ah/wordsonsand/tree/main/blog/layout/youtube/single.plain.txt)

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

As fas as I could see, there is no way (currently) in hugo to specify a default
output type for a `type` (i.e. youtube) of content, only a `kind` (e.g. page) of
content.

We, therefore need `outputs: ["plain"]` to the frontmatter

[hugo.toml](https://github.com/drone-ah/wordsonsand/tree/main/blog/hugo.toml)

```toml
[outputFormats.plain]
	mediaType = "text/plain"
	baseName = "index"
	isPlainText = true
	isHTML = false
	noUgly = true
```
