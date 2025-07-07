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

## TODO

- Check Validity of
  - [ ] Video Id
  - [ ] Playlist ID
