package inscribe

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

	return Scribed{
		format: yamlFormat,
	}, nil
}
