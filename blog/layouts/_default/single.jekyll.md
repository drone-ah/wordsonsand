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
