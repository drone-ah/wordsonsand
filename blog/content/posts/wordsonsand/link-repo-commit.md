---
title: "Automatically link to repo at current commit"
date: 2025-07-08T15:39:14+01:00
categories:
  - wordsonsand
tags:
  - wordsonsand
  - hugo
  - git
  - automation
  - markdown
  - github
---

I like writing blog posts, particularly about code, and I like to link to code
on my repo from my blog post. I do this a lot. Until now, I've just been copying
and pasting the full link to github. However, I ran into a problem today.

I moved a file that was referenced in a blog post. I then had to go to that blog
post and update the links - this is fine if I remember the linked blog posts -
but that's not scalable.

Also:

- Finding the link on GitHub, then copying and pasting is annoying
- It worries me a little that the version they are linked to could be vastly
  different from what I mention on the post.

I was already trying to create permalinks using tags, but that is laborious and
error prone.

What if I could get Hugo to:

- Automatically link to GitHub if the relative link is not within `blog/content`
- What if I could get it to link it to the file at the last commit of the post.

The last one is something to bear in mind. If I update a blog post, I'll have to
ensure that the links are still relevant.

<!-- more -->

Alternatively, let's allow an override at the page level where you can provide a
specific commit to link to:

## Step 1: Enable git info

To be able to get the commit of the post, we need to enable
[git info](https://gohugo.io/methods/page/gitinfo/)

[hugo.toml](../../../hugo.toml)

```toml
enableGitInfo = true
```

## Step 2: Update rendering of link

[layout/\_default/\_markup/render-link.html](../../../layouts/_default/_markup/render-link.html)

```gotmpl
{{- $linkPath := .Destination -}}                           {{/* e.g. "../scripts/tool.sh" */}}
{{- $currentPath := .Page.File.Path -}}                     {{/* e.g. "posts/foo.md" */}}
{{- $currentDir := path.Dir $currentPath -}}                {{/* e.g. "posts" */}}

{{- $combined := path.Join $currentDir $linkPath -}}        {{/* e.g. "posts/../scripts/tool.sh" */}}
{{- $resolved := path.Clean $combined -}}                   {{/* e.g. "scripts/tool.sh" */}}

{{- $fullRepoPath := path.Join "blog/content" $resolved -}} {{/* e.g. "blog/content/scripts/tool.sh" */}}

{{- $isInContent := strings.HasPrefix $fullRepoPath "blog/content/" -}}

{{- if $isInContent -}}
    <span class="unpublished">{{ $text }}</span>
{{- else -}}
    {{- $commit := or .Page.Params.link_commit .Page.GitInfo.Hash -}}
    <a href="https://github.com/drone-ah/wordsonsand/blob/{{ $commit }}/{{ $fullRepoPath }}" {{ with .Title }}title="{{ . }}"{{ end }}>{{ $text }}</a>
```

## Bonus: Allow per-post commit override

In the above code:

```gotmpl
{{- $commit := or .Page.Params.link_commit .Page.GitInfo.Hash -}}
```

The commit id is picked up from the page parameter `link_commit`, or if that
doesn't exist, from the last commit of the page.

You can therefore, set the commit to use for a post with:

```yaml
link_commit: <custom-commit-to-link-to>
```

## Conclusion

Necessity might be the mother of invention, but sometimes it takes a fine-tuned
sense of frustration to detect minor needs.
