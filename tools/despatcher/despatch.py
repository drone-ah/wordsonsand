#!/bin/env python3

import re
import os
import sys
import datetime
import frontmatter
from pathlib import Path

import praw
from blueskysocial import Client, Post

def find_publishable_files(base_path: Path):
    now = datetime.datetime.now(datetime.timezone.utc)
    year_month = now.strftime("%Y/%m")
    target_dir = base_path / year_month

    if not target_dir.exists():
        print(f"No content in {target_dir}")
        return []

    matched = dict()
    for path in target_dir.rglob("*.md"):
        try:
            post = frontmatter.load(path)
            publish_date = post.get("publishDate")
            published_at = post.get("publishedAt")

            if publish_date:
                print("isodate:", publish_date)
                pd = publish_date
                if pd < now and not published_at:
                    matched[path] = post
        except Exception as e:
            print(f"Error reading {path}: {e}")

    return matched

def post_bluesky(post):

    username = os.environ.get("APP_BLUESKY_USERNAME")
    password = os.environ.get("APP_BLUESKY_PASSWORD")
    client = Client()
    client.authenticate(username, password)
    bpost = Post(post.content)
    res = client.post(bpost)
    return res['uri']


def post_reddit(post):
    client_id = os.environ.get("APP_REDDIT_CLIENT_ID")
    client_secret = os.environ.get("APP_REDDIT_CLIENT_SECRET")
    refresh_token = os.environ.get("APP_REDDIT_REFRESH_TOKEN")

    reddit = praw.Reddit(
        client_id=client_id,
        client_secret=client_secret,
        refresh_token=refresh_token,
        user_agent="despatcher",
    )

    print(reddit.user.me())

    sub_name = post.get("subreddit")
    title = post.get("title")
    if sub_name == "" or sub_name is None or title == "" or title is None:
        return

    sub = reddit.subreddit(sub_name)

    posted = sub.submit(title=title, url=post.get("url"), selftext=post.content)
    return posted.url

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python despatch.py <base_path>")
        sys.exit(1)

    base_path = Path(sys.argv[1])
    posts = find_publishable_files(base_path)

    if len(posts) == 0:
       sys.exit(0)

    for path, p in posts.items():
        url = ""

        ptype = p.get("type")
        if ptype == "bluesky":
            url = post_bluesky(p)

        if ptype == "reddit":
            url = post_reddit(p)

            if url != '' and url is not None:
                now = datetime.datetime.now(datetime.timezone.utc).isoformat()
                p["publishedAt"] = now
                p["publishedTo"] = url
                frontmatter.dump(p, path)

