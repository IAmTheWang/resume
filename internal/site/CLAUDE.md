# internal/site/ (+ cmd/build/)

`cmd/build/main.go` is the entrypoint. Run it with:

```bash
go run ./cmd/build          # writes to dist/
go run ./cmd/build -out X   # writes to X/ instead — used for side-by-side diffs
```

It cleans the output dir, copies `templates/style.css` with a cache-busting
hash, then renders every entry in `Pages` (`internal/site/pages.go`) via
`Render` (`internal/site/render.go`) and writes each to `dist/`.

## Adding a page or language

A new language needs a `resume.ResumeData` var in `internal/resume/` (see
`internal/resume/CLAUDE.md`) plus an entry in the `Pages` slice in
`internal/site/pages.go`: `Data` (a `*resume.ResumeData` pointer — no
string-based module lookup, just reference the var directly), `Lang` (drives
`<html lang="...">`, `PrintLabel` lookup, and which `:lang()` CSS rules
apply), `Output` (path relative to `dist/`, POSIX slashes), `Title`,
`Detailed` (short Resume vs. full CV rendering — see `internal/resume/CLAUDE.md`
for why EN differs from JA/ZH), `ShowPageSwitch` / `PageSwitchActive` (only
meaningful for EN, which has two pages to switch between).

## Things that will break silently if you're not careful

- **Relative paths only, via `relhref` (`internal/site/relhref.go`).** Every
  internal link and the stylesheet reference is computed relative to the
  current page's own output path. This exists because GitHub Pages serves
  this site from a `/resume/` sub-path, not the domain root — an absolute
  path like `/style.css` resolves to the wrong URL and 404s. `relhref` uses
  Go's `path` package (POSIX, forward-slash) deliberately, never
  `path/filepath` (OS-separator-flavored) — mixing them here is the single
  highest-risk mistake in this package, since a bug only shows up as a 404 on
  the deployed site, not a build error. `relhref_test.go` covers every
  `(target, currentOutput)` pair the real build exercises; add a case there
  before changing the function.
- **`style.css` must be copied, not just linked** (`internal/site/assets.go`,
  `CopyCSS`). `cmd/build`'s `run()` does this before rendering any page — if
  a future refactor drops that call, `actions/upload-pages-artifact` only
  uploads what's physically inside `dist/`, and every page silently loses its
  stylesheet in production even though it works locally (browsers cache the
  old CSS, masking the break until a hard refresh).
- **CSS cache-busting is automatic.** The `?v=<hash>` query string on the
  stylesheet link is an MD5 hash of `style.css`'s own bytes (first 8 hex
  chars, lowercase — matches Python's old `hexdigest()[:8]` byte for byte),
  computed once in `CopyCSS` and threaded through every page via
  `PageView.CSSHref`. You don't need to bump a version number by hand when
  editing the CSS.
- **`NoIndex`** (top of `pages.go`) controls whether every page ships
  `<meta name="robots" content="noindex, nofollow">`. Currently `true` by
  design (see root `CLAUDE.md`) — don't flip it without asking the account
  owner first, since it changes public search-engine visibility of PII-light
  but still personal content.
- **Resume-content fields go through the `esc` template func, not bare
  `{{.}}`.** `esc` wraps Go's stdlib `html.EscapeString` (see `render.go`).
  This is deliberate: `html/template`'s own default escaper additionally
  escapes ASCII `+` to `&#43;` in body text, which the old Python generator's
  `html.escape()`-based `esc()` never did — and this resume's content is full
  of literal `+` (e.g. "Vue 3 + TypeScript", "15+ suites"). Using the plain
  `html.EscapeString`-based func keeps output consistent with what shipped
  before and avoids a wall of cosmetic (but noisy) diff churn. The one
  accepted difference from the old Python output is the apostrophe encoding:
  Python's `html.escape` emits `&#x27;`, Go's `html.EscapeString` emits
  `&#39;` — both render as `'` in every browser; don't chase byte-parity on
  this one.
- **Templates are embedded via `//go:embed templates/*.tmpl.html`**
  (`render.go`), so `go run ./cmd/build` works from any working directory.
  `templates/style.css` at the repo root is deliberately *not* embedded the
  same way — `//go:embed` can't reference a path above its own package
  directory (`..` is rejected at compile time), and moving `style.css` into
  `internal/site/templates/` would blur its actual role (a shared design
  asset, not a per-language template). `CopyCSS` instead resolves the repo
  root at runtime via `runtime.Caller(0)`, so it's CWD-independent without
  needing embed.
