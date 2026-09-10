package site

import (
	"path"
	"strings"
)

// relhref computes a POSIX-style relative path from the directory of
// currentOutput to target, mirroring Python's
// posixpath.relpath(target, start=posixpath.dirname(currentOutput) or ".").
//
// This must stay POSIX/forward-slash (path, not path/filepath) regardless of
// the build host's OS: GitHub Pages serves this site from a /resume/
// sub-path, so every internal link and asset href is computed relative to
// the current page rather than absolute-rooted, and a wrong result here
// silently produces a 404 in production rather than a build error.
func relhref(target, currentOutput string) string {
	currentDir := path.Dir(currentOutput)

	currentSegs := splitClean(currentDir)
	targetSegs := splitClean(target)

	n := 0
	for n < len(currentSegs) && n < len(targetSegs) && currentSegs[n] == targetSegs[n] {
		n++
	}

	up := strings.Repeat("../", len(currentSegs)-n)
	down := strings.Join(targetSegs[n:], "/")

	result := up + down
	if result == "" {
		return "."
	}
	return result
}

// splitClean cleans a POSIX path and splits it into non-empty segments.
// "." cleans away to no segments; "a/b/../c" cleans to "a/c" first.
func splitClean(p string) []string {
	cleaned := path.Clean(p)
	if cleaned == "." || cleaned == "" {
		return nil
	}
	return strings.Split(cleaned, "/")
}
