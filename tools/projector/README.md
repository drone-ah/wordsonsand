# projector

A helper tool to manage YouTube video descriptions for Content Creators.

This tool is of medium complexity and requires a third party tool that will help
generate the actual descriptions.

I
[use hugo to generate the descriptions](https://drone-ah.com/2025/07/03/generate-youtube-descriptions-from-hugo/)

It currently has two commands:

- `validate`: Checks that the source files and the generated descriptions exist
  and can be parsed. This part could do with some improvements
- `sync`: Checks the descriptions of videos in the last 30 days to check if the
  description hash has changed. If so, it will update it on YouTube.

More
[details can be found on the blog post](https://drone-ah.com/2025/07/07/projector-keep-youtube-descriptions-synced/)

## How it works

### Yaml File

This tool expects the descriptions for the youtube files to be defined in a yaml
file. The supported fields and default content can be found in
[the archetype](../../blog/archetypes/youtube.md)

You can find some examples of active descriptions in
[the content section](../../blog/content/youtube/)

### Description File

Hugo then generates the description files.

### projector sync

This tool (projector) then uses information from both those directories to
detect if the description has changed, and if so, to update them.

## TODO

- Check Validity of
  - [ ] Video Id
  - [ ] Playlist ID
