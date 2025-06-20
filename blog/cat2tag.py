import re
import os
import sys
import frontmatter

def add_cat_to_tag(filepath):
    post = frontmatter.load(filepath)

    categories = post.get('categories', [])
    tags = post.get('tags', [])
    if tags is None:
        tags = []

    # Normalize to lists
    if isinstance(categories, str):
        categories = [categories]
    if isinstance(tags, str):
        tags = [tags]

    if categories:
        first_category = categories[0]
        if first_category not in tags:
            tags.insert(0, first_category)
            post['tags'] = tags
            print(f"Added category '{first_category}' to tags in {filepath}")
        else:
            print(f"Category '{first_category}' already in tags in {filepath}")
    else:
        print(f"No categories in {filepath}")

    frontmatter.dump(post, filepath)

def lowercase_tags(filepath):
    post = frontmatter.load(filepath)

    tags = post.get('tags', [])
    if tags is None:
        tags = []

    for i, t in enumerate(tags):
        tags[i] = normalize_tag(t)

    print(tags)

    frontmatter.dump(post, filepath)
                             
def normalize_tag(tag):
    # Remove unwanted characters, keep letters, numbers, spaces, and hyphens
    tag = re.sub(r'[^\w\s-]', '', tag)
    # Replace spaces/underscores with hyphens, make lowercase
    return re.sub(r'[\s_]+', '-', tag.strip().lower())

def remap_tags(filepath):
    post = frontmatter.load(filepath)
    tags = post.get('tags', [])
    if tags is None:
        tags = []

    mappings = {
        'systems-administration': 'sysadmin',
        'game-development': 'gamedev',
        'android-development': 'android-dev',
        'short-stories': 'short-story',
        'the-elder-scrolls-v-skyrim': 'skyrim',
        'windows': 'microsoft-windows',
        'cache': 'caching',
        'compiling-tools': 'compilers',
        'dev-env': 'devex',
        'dom': 'domain-model',
        'domain-object-model': 'domain-model',
        'international-business-machines-corporation': 'ibm',
        'computing': '',
        'control-key': '',
        'dag': 'directed-acyclic-graph',
        'global-file-system': '',
        'guest-os': '',
        'hd': '',
        'intent': '',
        'jboss-application-server': 'jboss',
        'plain-old-java-object': 'pojo',
        'software-component': '',
        'software-components': '',
        'system-software': '',
        'the-elder-scrolls-ii-daggerfall': 'daggerfall',
        'the-elder-scrolls-iv-oblivion': 'oblivion',
        'united-states-of-america': 'usa',
        'software-development': 'sfeng',
        'software-engineering': 'sfeng',



    }

    for i, t in enumerate(tags):
        if t in mappings:
            print("found in mapping")
            tags[i] = mappings[t]

    post["tags"] = dedupe(tags)
    frontmatter.dump(post, filepath)

def dedupe(seq):
        seen = set('')
        return [x for x in seq if not (x in seen or seen.add(x)) and len(x) > 0]

def main(directory):
    for root, _, files in os.walk(directory):
        for file in files:
            if file.endswith('.md') or file.endswith('.markdown'):
                print("processing: ", file)
                remap_tags(os.path.join(root, file))

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python add_category_to_tags.py <path_to_markdown_dir>")
        sys.exit(1)
    main(sys.argv[1])
