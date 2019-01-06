package content

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/hellozimi/sidago/helpers"
)

type contentType string

var (
	ErrFileExists              = fmt.Errorf("file already exists at destination")
	ErrContentTypeNotSupported = fmt.Errorf("the content type is not supported")
)

const (
	typeUnknown contentType = "unknown"
	typePost    contentType = "post"
	typePage    contentType = "page"
)

type ContentCreator struct {
	path string
	name string
	b    boilerplate
	t    contentType
}

// Execute writes boilerplate to output dir
func (c *ContentCreator) Execute(force bool) error {
	if c.t == typeUnknown {
		return ErrContentTypeNotSupported
	}

	path := c.outputFilePath()
	if _, err := os.Stat(path); !os.IsNotExist(err) && !force {
		return ErrFileExists
	}

	content, err := c.b.content()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(content), 0777)

	if err != nil {
		return err
	}

	return nil
}

func (c *ContentCreator) outputFilePath() string {
	return filepath.Join(c.outputDir(), c.name+".md")
}

func (c *ContentCreator) outputDir() string {
	return filepath.Join(c.path, c.contentDir())
}

func (c *ContentCreator) contentDir() string {
	switch c.t {
	case typePage:
		return "pages"
	case typePost:
		return "posts"
	default:
		return ""
	}
}

// New creates a new content creator object
func New(path string, contentType string, name string) *ContentCreator {
	path = filepath.Clean(path)

	if !helpers.IsSlug(name) {
		name = helpers.Slugify(name)
	}

	t := contentTypeString(contentType)
	slug, date := helpers.SlugAndDateFromFile(name)
	if t == typePost {
		if date.IsZero() {
			date = time.Now()
		}
		dateString := date.Format("2006-01-02")
		name = fmt.Sprintf("%s_%s", dateString, slug)
	}
	b := newBoilerplate(helpers.Unslugify(slug), date, t)

	return &ContentCreator{
		path: path,
		name: name,
		t:    t,
		b:    b,
	}
}

func contentTypeString(in string) contentType {
	switch in {
	case "post":
		return typePost
	case "page":
		return typePage
	default:
		return typeUnknown
	}
}
