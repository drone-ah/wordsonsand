#!/bin/env python3

import os
import sys
import datetime
import frontmatter
from pathlib import Path

def find_publishable_files(base_path: Path):
    now = datetime.datetime.now(datetime.timezone.utc)
    year_month = now.strftime("%Y/%m")
    target_dir = base_path / year_month

    if not target_dir.exists():
        print(f"No content in {target_dir}")
        return []

    matched = []
    for path in target_dir.rglob("*.md"):
        try:
            post = frontmatter.load(path)
            publish_date = post.get("publishDate")
            published_at = post.get("publishedAt")

            if publish_date:
                print("isodate:", publish_date)
                pd = publish_date
                if pd < now and not published_at:
                    matched.append(post)
        except Exception as e:
            print(f"Error reading {path}: {e}")

    return matched

def post_bluesky(post):
    from atproto import Client, client_utils
    client = Client()
    username = os.environ.get("APP_BLUESKY_USERNAME")
    password = os.environ.get("APP_BLUESKY_PASSWORD")
    profile = client.login(username, password)
    text = build_text_with_facets(post.content)
    client.send_post(text)

import re
from atproto import client_utils, Client

def build_text_with_facets(text: str) -> client_utils.TextBuilder:
    text_builder = client_utils.TextBuilder()
    cursor = 0

    # Patterns
    url_pattern = re.compile(r'https?://\S+')
    tag_pattern = re.compile(r'#\w+')

    # Combine all matches
    matches = []
    for m in url_pattern.finditer(text):
        matches.append((m.start(), m.end(), 'url', m.group()))
    for m in tag_pattern.finditer(text):
        matches.append((m.start(), m.end(), 'tag', m.group()[1:]))

    matches.sort(key=lambda m: m[0])  # ensure correct order

    for start, end, kind, value in matches:
        if cursor < start:
            text_builder.text(text[cursor:start])

        label = text[start:end]
        if kind == 'url':
            text_builder.link(label, value)
        elif kind == 'tag':
            text_builder.tag(label, value)

        cursor = end

    if cursor < len(text):
        text_builder.text(text[cursor:])

    return text_builder

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python despatch.py <base_path>")
        sys.exit(1)

    base_path = Path(sys.argv[1])
    posts = find_publishable_files(base_path)

    for p in posts:
        if p.get("type") == "bluesky":
            post_bluesky(p)

