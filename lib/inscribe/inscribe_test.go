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

	expected := `---
title: New Title
keyId: abc123
---

This is the description body.-
`

	if o.String() != expected {
		t.Errorf("output (%s) doesn't match\n (%s)", o.String(), expected)
	}

}
