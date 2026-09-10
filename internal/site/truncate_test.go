package site

import (
	"strings"
	"testing"
)

func TestTruncate(t *testing.T) {
	short := "a short sentence"
	if got := truncate(short, 200); got != short {
		t.Errorf("truncate(short) = %q, want unchanged %q", got, short)
	}

	if got := truncate("  padded  ", 200); got != "padded" {
		t.Errorf("truncate should trim whitespace first, got %q", got)
	}

	long := strings.Repeat("word ", 60) // 300 chars, well over the 200 limit
	got := truncate(long, 200)
	if !strings.HasSuffix(got, "…") {
		t.Errorf("truncate(long) = %q, want it to end with an ellipsis", got)
	}
	if strings.HasSuffix(got, " …") {
		t.Errorf("truncate(long) = %q, should have backed off the trailing space before appending …", got)
	}
	if r := []rune(got); len(r) > 200 {
		t.Errorf("truncate(long) produced %d runes, want <= 200", len(r))
	}

	// CJK text: must slice by rune, not byte, or this corrupts multi-byte characters.
	cjk := strings.Repeat("日本語のテスト文章です。", 30) // well over 200 runes, no spaces to back off to
	gotCJK := truncate(cjk, 200)
	runesCJK := []rune(gotCJK)
	if runesCJK[len(runesCJK)-1] != '…' {
		t.Errorf("truncate(cjk) = %q, want it to end with an ellipsis", gotCJK)
	}
	if len(runesCJK) != 200 {
		t.Errorf("truncate(cjk) produced %d runes, want exactly 200 (199 content + ellipsis)", len(runesCJK))
	}
}
