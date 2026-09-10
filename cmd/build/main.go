// Command build renders internal/resume's per-language content into a
// multilingual static site under dist/ (or -out).
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"resume-site/internal/site"
)

func main() {
	outDir := flag.String("out", "dist", "output directory")
	flag.Parse()

	if err := run(*outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(outDir string) error {
	if err := os.RemoveAll(outDir); err != nil {
		return fmt.Errorf("clean output dir: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	cssHash, err := site.CopyCSS(outDir)
	if err != nil {
		return fmt.Errorf("copy css: %w", err)
	}

	for i := range site.Pages {
		page := &site.Pages[i]
		html, err := site.Render(page, cssHash)
		if err != nil {
			return fmt.Errorf("render %s: %w", page.Output, err)
		}

		outPath := filepath.Join(outDir, filepath.FromSlash(page.Output))
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return fmt.Errorf("create dir for %s: %w", page.Output, err)
		}
		if err := os.WriteFile(outPath, []byte(html), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", page.Output, err)
		}
		fmt.Printf("Built %s\n", outPath)
	}
	return nil
}
