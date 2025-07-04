package inscribe_test

import (
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
