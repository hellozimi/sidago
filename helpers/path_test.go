package helpers

import "testing"

func TestFileAndExt(t *testing.T) {
	cases := []struct {
		a string
		e string
		f string
	}{
		{a: "/users/root/file", e: "", f: "file"},
		{a: "/users/root/my_mov.mkv", e: ".mkv", f: "my_mov"},
	}

	for _, x := range cases {
		f, e := FileAndExt(x.a)
		if f != x.f {
			t.Errorf("expected %s file to equal %s but was %s", x.a, x.f, f)
		}
		if e != x.e {
			t.Errorf("expected %s ext to equal %s but was %s", x.a, x.e, e)
		}
	}
}
