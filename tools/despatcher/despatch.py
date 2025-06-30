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
                    matched.append(path)
        except Exception as e:
            print(f"Error reading {path}: {e}")

    return matched

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python despatch.py <base_path>")
        sys.exit(1)

    base_path = Path(sys.argv[1])
    files = find_publishable_files(base_path)

    for f in files:
        print(f)

