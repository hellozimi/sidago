package fs

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIsDirEmpty(t *testing.T) {

	tmpDir, err := ioutil.TempDir("./", "sida")
	if err != nil {
		t.Fatalf("Coudln't create TempDir")
	}
	defer os.Remove(tmpDir)
	tmpFile, err := ioutil.TempFile("./", "sida")
	if err != nil {
		t.Fatalf("Coudln't create TempFile")
	}
	defer os.Remove(tmpFile.Name())

	cases := []struct {
		path   string
		expect bool
	}{
		{"./", false},
		{tmpDir, true},
		{tmpFile.Name(), false},
		{"/usr/local/var/sida", false},
	}

	for i, test := range cases {
		result, err := IsDirEmpty(test.path)
		if result != test.expect {
			t.Errorf("%d: Path %s got result %v but expected %v with error %v", i, test.path, result, test.expect, err)
		}
	}
}
