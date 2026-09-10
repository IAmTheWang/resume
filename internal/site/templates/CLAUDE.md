# internal/site/templates/

`*.tmpl.html` files here are parsed as one `html/template` set (embedded via
`//go:embed`, see `internal/site/render.go`) and executed by name — `head`,
`hero`, `summary`, `coreStrengths`, `skills`, `experienceItem`,
`educationLanguages`, `personalStatement` (in `sections.tmpl.html`), plus the
shared `ul` / `langTabsNav` / `pageSwitchNav` helpers in `shared.tmpl.html`.

## Whitespace is load-bearing here

These templates were hand-tuned line-by-line against the old Python
generator's output to keep the rendered HTML close to byte-identical
(verified with `go run ./cmd/build -out dist_go` + `diff -rq dist dist_go`).
Every `{{if}}`/`{{range}}`/`{{end}}` placement and line break was chosen
deliberately so the concatenated output has no stray blank lines. If you add
a new conditional block or loop, check the rendered output's whitespace
against a real page rather than assuming it'll look fine — a misplaced
newline around an action tag is easy to introduce and easy to miss in a
diff-free read-through.

## Escaping: always use `esc`, never bare `{{.}}`

Every resume-content field must go through the `esc` func
(`{{esc .Field}}`), not raw `{{.Field}}`. `esc` wraps stdlib
`html.EscapeString` instead of relying on `html/template`'s own
context-derived escaper, specifically because that escaper additionally
turns ASCII `+` into `&#43;` in body text — see `internal/site/CLAUDE.md` for
why that matters here. `dot` and `br` are the two spots allowed to emit raw,
unescaped HTML (`&nbsp;·&nbsp;` and `<br>`) — both are fixed fragments the
data never controls, not resume content.
