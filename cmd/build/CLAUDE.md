# cmd/build/

`main.go` is the entrypoint for the whole site generator:

```bash
go run ./cmd/build          # writes to dist/
go run ./cmd/build -out X   # writes to X/ instead
```

It's intentionally thin — clean the output dir, copy the CSS, loop over
`site.Pages` and render+write each one. All the actual logic (page config,
templates, `relhref`, escaping) lives in `internal/site`; see
[`internal/site/CLAUDE.md`](../../internal/site/CLAUDE.md) for the things
that break silently if changed carelessly (relative paths, CSS caching,
`NoIndex`, the `esc` template func).

Don't add rendering logic here — this file should stay a thin CLI wrapper
around `internal/site`.
