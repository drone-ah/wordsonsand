#!/bin/bash
for file in *.md; do
  if ! grep -q '^url:' "$file"; then
    slug=$(basename "$file" .md | cut -d'-' -f4-)
    awk -v slug="$slug" '
      BEGIN { inserted=0 }
      /^title:/ {
        print
        print "slug: " slug
        inserted=1
        next
      }
      { print }
    ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
  fi
done
