package sida

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestIsSida(t *testing.T) {
	a := makeFolders(t, "sida-1", "")
	defer os.RemoveAll(a)
	b := makeFolders(t, "sida-2", "layout")
	defer os.RemoveAll(b)
	c := makeFolders(t, "sida-3", "layout", "posts")
	defer os.RemoveAll(c)
	d := makeFolders(t, "sida-4", "layout", "pages", "posts")
	defer os.RemoveAll(d)
	e := makeFolders(t, "sida-4", "layout", "pages", "posts")
	defer os.RemoveAll(e)

	makeConfig(t, filepath.Join(d, "config.toml"), `
title = "My config"`)
	makeConfig(t, filepath.Join(e, "config.toml"), `{"title": "My config"}`)

	cases := []struct {
		a string
		e bool
	}{
		{a: a, e: false},
		{a: b, e: false},
		{a: c, e: false},
		{a: d, e: true},
		{a: e, e: false},
		{a: "/usr/local/var/sida_test", e: false},
	}

	for _, x := range cases {
		res := IsSida(x.a)
		if res != x.e {
			t.Errorf("expected result to equal '%v' but was '%v'", x.e, res)
		}
	}
}

func makeFolders(t *testing.T, main string, in ...string) string {
	tmpDir, err := ioutil.TempDir("./", main)
	if err != nil {
		t.Fatalf("Coudln't create TempDir")
	}

	if len(in) > 0 {
		for _, f := range in {
			if f == "" {
				continue
			}
			err := os.Mkdir(filepath.Join(tmpDir, f), 0777)
			if err != nil {
				t.Fatalf("couldn't create folder inside tmp folder: %v", err)
			}
		}
	}

	return tmpDir
}

func makeConfig(t *testing.T, path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0777)
	if err != nil {
		t.Fatalf("couldn't write config file: %v", err)
	}
}
