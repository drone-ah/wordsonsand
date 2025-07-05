package inscribe_test

import (
	"strings"
	"testing"

	"github.com/drone-ah/wordsonsand/lib/inscribe"
)

type fMatter struct {
	Title string `yaml:"title"`
	KeyId string `yaml:"keyId"`
}

func TestReadContent(t *testing.T) {
	s, err := inscribe.NewScribedFromFile("testdata/md-with-frontmatter.md")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if s.Content != "This is the description body.-\n" {
		t.Errorf("content (%s) does not match", s.Content)
	}

}

func TestReadFrontMatter(t *testing.T) {
	s, err := inscribe.NewScribedFromFile("testdata/md-with-frontmatter.md")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var fm fMatter
	s.FrontMatter(&fm)

	if fm.Title != "Test Title" {
		t.Errorf("title (%s) doesn't match", fm.Title)
	}

	if fm.KeyId != "abc123" {
		t.Errorf("Key (%s) doesn't match", fm.KeyId)
	}
}

type withMap struct {
	Hashes map[string]string `yaml:"hashes"`
}

func TestReadFrontMatterMap(t *testing.T) {
	s, err := inscribe.NewScribedFromFile("testdata/md-with-frontmatter.md")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var wm withMap
	s.FrontMatter(&wm)

	if wm.Hashes["description"] != "hashhash" {
		t.Errorf("Key (%s) doesn't match", wm.Hashes["description"])
	}
}

const expectedWritten = `---
title: "New Title"
keyId: "abc123"
hashes:
  description: hashhash
---

This is the description body.-
`

func TestWriteFrontMatter(t *testing.T) {
	s, err := inscribe.NewScribedFromFile("testdata/md-with-frontmatter.md")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var fm fMatter
	s.FrontMatter(&fm)

	fm.Title = "New Title"

	var o strings.Builder
	s.Write(fm, &o)
	if err != nil {
		t.Errorf("unexpected error in Write: %v", err)
	}

	if o.String() != expectedWritten {
		t.Errorf("output (%s) doesn't match\n (%s)", o.String(), expectedWritten)
	}

}

type titleOnly struct {
	Title string `yaml:"title"`
}

func TestWritePartialFrontMatter(t *testing.T) {
	s, err := inscribe.NewScribedFromFile("testdata/md-with-frontmatter.md")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var to titleOnly
	s.FrontMatter(&to)

	to.Title = "New Title"

	var o strings.Builder
	s.Write(to, &o)
	if err != nil {
		t.Errorf("unexpected error in Write: %v", err)
	}

	if o.String() != expectedWritten {
		t.Errorf("output (%s) doesn't match\n (%s)", o.String(), expectedWritten)
	}
}

const expectedNewKey = `---
title: "New Title"
keyId: "abc123"
hashes:
  description: hashhash
newKey: something
---

This is the description body.-
`

type newKey struct {
	Title  string `yaml:"title"`
	NewKey string `yaml:"newKey"`
}

func TestWritingNewKeyToFrontMatter(t *testing.T) {
	s, err := inscribe.NewScribedFromFile("testdata/md-with-frontmatter.md")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var nk newKey
	s.FrontMatter(&nk)

	nk.Title = "New Title"
	nk.NewKey = "something"

	var o strings.Builder
	s.Write(nk, &o)
	if err != nil {
		t.Errorf("unexpected error in Write: %v", err)
	}

	if o.String() != expectedNewKey {
		t.Errorf("output (%s) doesn't match\n (%s)", o.String(), expectedNewKey)
	}
}

const expectedMapKey = `---
title: "New Title"
keyId: "abc123"
hashes:
  description: hashhash
  new: something
newKey: something
---

This is the description body.-
`

func TestWritingMaps(t *testing.T) {
	s, err := inscribe.NewScribedFromFile("testdata/md-with-frontmatter.md")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var wm withMap
	s.FrontMatter(&wm)

	wm.Hashes["new"] = "something"

	var o strings.Builder
	s.Write(wm, &o)
	if err != nil {
		t.Errorf("unexpected error in Write: %v", err)
	}

	if o.String() != expectedMapKey {
		t.Errorf("output (%s) doesn't match\n (%s)", o.String(), expectedMapKey)
	}
}
