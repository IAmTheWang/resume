package site

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"runtime"
)

// repoRoot returns the repository root, resolved from this source file's own
// location so the build works regardless of the caller's current working
// directory — the Go equivalent of Python's Path(__file__).resolve().parent.parent.
func repoRoot() string {
	_, thisFile, _, _ := runtime.Caller(0)
	// this file is internal/site/assets.go — up two levels to the repo root.
	return filepath.Dir(filepath.Dir(filepath.Dir(thisFile)))
}

// CopyCSS reads templates/style.css from the repo root, writes a verbatim
// byte copy into distDir/style.css, and returns the first 8 hex chars of its
// MD5 digest for use as a cache-busting query string. Uses filepath (OS path
// semantics), not path — this is a real filesystem operation, unlike
// relhref, which must stay POSIX for URL correctness.
func CopyCSS(distDir string) (hash string, err error) {
	srcPath := filepath.Join(repoRoot(), "templates", "style.css")
	b, err := os.ReadFile(srcPath)
	if err != nil {
		return "", err
	}

	sum := md5.Sum(b)
	hash = hex.EncodeToString(sum[:])[:8]

	if err := os.WriteFile(filepath.Join(distDir, "style.css"), b, 0o644); err != nil {
		return "", err
	}
	return hash, nil
}
