package site

import "testing"

func TestRelhref(t *testing.T) {
	cases := []struct {
		target, currentOutput, want string
	}{
		// The 5 real (target, currentOutput) pairs exercised by this site's build.
		{"style.css", "index.html", "style.css"},
		{"style.css", "ja/index.html", "../style.css"},
		{"index.html", "ja/index.html", "../index.html"},
		{"cv.html", "index.html", "cv.html"},
		{"ja/index.html", "index.html", "ja/index.html"},

		// Synthetic deeper-nesting cases, cross-checked against posixpath.relpath.
		{"a/b/c.html", "x/y/z.html", "../../a/b/c.html"},
		{"index.html", "index.html", "index.html"},
		{"zh/index.html", "ja/index.html", "../zh/index.html"},
	}

	for _, c := range cases {
		got := relhref(c.target, c.currentOutput)
		if got != c.want {
			t.Errorf("relhref(%q, %q) = %q, want %q", c.target, c.currentOutput, got, c.want)
		}
	}
}
