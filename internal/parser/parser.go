package parser

import (
	"bytes"
	"errors"

	gotoml "github.com/pelletier/go-toml"
)

var tomlDelim = []byte("---")

var (
	NoFrontmatter        = errors.New("no frontmatter found")
	MalformedFrontmatter = errors.New("malformed frontmatter")
)

// ParseFrontmatterBody parses toml frontmatter from byte input
// using the '---' delimiter
func ParseFrontmatterBody(input []byte) (map[string]interface{}, []byte, error) {
	front := make(map[string]interface{})
	body := make([]byte, 0)

	input = bytes.TrimSpace(input)

	if !HasFrontmatter(input) {
		return front, body, NoFrontmatter
	}

	input = bytes.TrimPrefix(input, tomlDelim)

	parts := bytes.SplitN(input, tomlDelim, 2)
	if len(parts) != 2 {
		return front, body, MalformedFrontmatter
	}

	t, err := gotoml.LoadBytes(parts[0])
	if err != nil {
		return front, body, err
	}

	front = t.ToMap()
	body = bytes.TrimSpace(parts[1])

	return front, body, nil
}

// HasFrontmatter checks if input has frontmatter
func HasFrontmatter(input []byte) bool {
	return bytes.HasPrefix(bytes.TrimSpace(input), tomlDelim)
}
