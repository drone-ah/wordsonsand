package inscribe

import (
	"bytes"
	"errors"
	"io"
	"os"
)

// A Scribed is a representation of a file that contains frontmatter and markdown content
type Scribed struct {
	// HACK: We could store a Reader/Writer if we wanted more flexibility
	format      Format
	frontmatter []byte
	Content     string
}

// NewScribed will return a new Scribed object representing a file with frontmatter
func NewScribedFromFile(path string) (Scribed, error) {
	//TODO: We could auto detect format based on delimiters being present
	// For the time being though, let's just use YAML

	s := Scribed{
		format: yamlFormat,
	}

	f, err := os.Open(path)
	if err != nil {
		return Scribed{}, err
	}
	defer f.Close()

	err = s.splitFrontmatter(f)
	if err != nil {
		return Scribed{}, err
	}

	return s, nil
}

func (s *Scribed) FrontMatter(out any) error {
	return s.format.Unmarshal(s.frontmatter, out)
}

func (s *Scribed) Write(fm any, out io.Writer) error {
	// Merge frontmatter
	data, err := s.format.Merge(s.frontmatter, fm)
	if err != nil {
		return err
	}

	io.WriteString(out, s.format.Delimiter+"\n")
	out.Write(data)
	io.WriteString(out, s.format.Delimiter+"\n\n")
	io.WriteString(out, s.Content)

	return nil
}

// splitFrontmatter will split frontmatter from Content and store them
func (s *Scribed) splitFrontmatter(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	parts := bytes.SplitN(data, []byte("\n"+s.format.Delimiter+"\n"), 2)
	if len(parts) != 2 {
		return errors.New("invalid frontmatter format")
	}

	// Remove the opening '---\n' from the first part
	s.frontmatter = bytes.TrimPrefix(parts[0], []byte("---\n"))

	s.Content = string(bytes.TrimLeft(parts[1], "\r\n"))

	return nil
}
