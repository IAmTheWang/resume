# internal/

Go convention: packages here can't be imported by any module outside this
repo (not that anything would — this is a `main`-only site generator, not a
library). Two packages:

| Package | Contents | See |
|---|---|---|
| [`resume/`](resume/CLAUDE.md) | Per-language resume content (Go structs) — the single source of truth for what's on the site | CLAUDE.md |
| [`site/`](site/CLAUDE.md) | Rendering pipeline: page config, `html/template` templates, `relhref`, CSS/cache-busting | CLAUDE.md |
