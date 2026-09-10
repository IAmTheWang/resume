package site

import "strings"

// truncate mirrors Python's truncate(text, limit=200): strip, and if the
// trimmed text is longer than limit, cut at limit-1 runes then back off to
// the last space (so no word is cut mid-token) and append a single "…".
//
// Operates on []rune throughout, not bytes — JA/ZH content is full of
// multi-byte UTF-8 characters, and Python's text[:limit] slices by Unicode
// character, not byte, so a byte-based port would corrupt CJK text.
func truncate(text string, limit int) string {
	text = strings.TrimSpace(text)
	runes := []rune(text)
	if len(runes) <= limit {
		return text
	}

	cut := runes[:limit-1]
	if idx := lastIndexRune(cut, ' '); idx >= 0 {
		cut = cut[:idx]
	}
	return string(cut) + "…"
}

func lastIndexRune(rs []rune, r rune) int {
	for i := len(rs) - 1; i >= 0; i-- {
		if rs[i] == r {
			return i
		}
	}
	return -1
}
