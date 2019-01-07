package content

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestContentTypeString(t *testing.T) {
	cases := []struct {
		a string
		e contentType
	}{
		{a: "post", e: typePost},
		{a: "page", e: typePage},
		{a: "something", e: typeUnknown},
	}

	for _, x := range cases {
		res := contentTypeString(x.a)
		if res != x.e {
			t.Errorf("expected contentType to equal '%s' but was '%s'", x.e, res)
		}
	}
}

func TestContentCreator(t *testing.T) {
	tmpDir, err := ioutil.TempDir("./", "sida")
	if err != nil {
		t.Fatalf("Coudln't create TempDir")
	}
	defer os.RemoveAll(tmpDir)

	tmpPostsDir := filepath.Join(tmpDir, "posts")
	err = os.Mkdir(tmpPostsDir, 0777)
	if err != nil {
		t.Fatalf("Coudln't create tmpPostsDir")
	}
	tmpPagesDir := filepath.Join(tmpDir, "pages")
	err = os.Mkdir(tmpPagesDir, 0777)
	if err != nil {
		t.Fatalf("Coudln't create tmpPagesDir")
	}

	ymdToday := time.Now().Format("2006-01-02")

	cases := []struct {
		a  *ContentCreator
		en string
		et contentType
		// output dir
		eod string
		// output file path
		eofp string
		// execute error
		err error
	}{
		{
			a:    New(tmpDir, "page", "My Page"),
			en:   "my-page",
			et:   typePage,
			eod:  filepath.Join(tmpDir, "pages"),
			eofp: filepath.Join(tmpDir, "pages", "my-page.md"),
		},
		{
			a:    New(tmpDir, "post", "My New Post"),
			en:   fmt.Sprintf("%s_my-new-post", ymdToday),
			et:   typePost,
			eod:  filepath.Join(tmpDir, "posts"),
			eofp: filepath.Join(tmpDir, "posts", fmt.Sprintf("%s_my-new-post.md", ymdToday)),
		},
		{
			a:    New(tmpDir, "post", "2018-03-24_birthday-post"),
			en:   "2018-03-24_birthday-post",
			et:   typePost,
			eod:  filepath.Join(tmpDir, "posts"),
			eofp: filepath.Join(tmpDir, "posts", "2018-03-24_birthday-post.md"),
		},
		{
			a:    New(tmpDir, "something", "Something"),
			en:   "something",
			et:   typeUnknown,
			eod:  tmpDir,
			eofp: filepath.Join(tmpDir, "something.md"),
			err:  ErrContentTypeNotSupported,
		},
	}

	for _, x := range cases {
		if x.a.name != x.en {
			t.Errorf("expected name to equal '%s' but was '%s'", x.en, x.a.name)
		}

		if x.a.t != x.et {
			t.Errorf("expected type to equal '%s' but was '%s'", x.et, x.a.t)
		}

		if x.a.outputDir() != x.eod {
			t.Errorf("expected outputDir to equal '%s' but was '%s'", x.eod, x.a.outputDir())
		}

		if x.a.outputFilePath() != x.eofp {
			t.Errorf("expected outputFilePath to equal '%s' but was '%s'", x.eofp, x.a.outputFilePath())
		}

		err := x.a.Execute(false)
		if err != x.err {
			t.Errorf("expected execute err to equal '%v' but was '%v'", x.err, err)
		}
	}
}
