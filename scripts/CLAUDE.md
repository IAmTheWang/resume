# scripts/

`build_site.py` is the entire static site generator — no templating library
(Jinja2, etc.), just Python f-strings. Run it with:

```bash
python3 scripts/build_site.py
```

It reads every module listed in `PAGES` (each entry names a `data/` module,
an output path, and per-page flags), renders each to HTML, copies
`templates/style.css` into `dist/`, and writes everything under `dist/`
(gitignored — CI rebuilds it fresh on every deploy, nothing here is
hand-edited).

## Adding a page or language

Add an entry to the `PAGES` list: `module` (importable name from `data/`,
must be on `sys.path` — it already is, via the `sys.path.insert` at the top),
`lang` (drives `<html lang="...">` and which `:lang()` CSS rules apply),
`output` (path relative to `dist/`), `title`, `detailed` (short Resume vs.
full CV rendering — see `data/CLAUDE.md` for why EN differs from JA/ZH),
`show_page_switch` / `page_switch_active` (only meaningful for EN, which has
two pages to switch between).

## Things that will break silently if you're not careful

- **Relative paths only.** Every internal link and the stylesheet reference
  is computed with `relhref()` (built on `posixpath.relpath`), keyed off each
  page's own `output` path. This exists because GitHub Pages serves this site
  from a `/resume/` sub-path, not the domain root — an absolute path like
  `/style.css` resolves to the wrong URL and 404s. If you add a new asset or
  cross-page link, route it through `relhref()`; don't hardcode a path.
- **`style.css` must be copied, not just linked.** `main()` does
  `shutil.copy` into `dist/` before rendering pages — if a future refactor
  drops that call, `actions/upload-pages-artifact` will only upload what's
  physically inside `dist/`, and every page will silently lose its stylesheet
  in production even though it works locally (browsers cache the old CSS,
  masking the break until a hard refresh).
- **CSS cache-busting is automatic.** The `?v=<hash>` query string on the
  stylesheet link is an md5 of `style.css`'s own bytes, computed once in
  `main()` and threaded through every page — you don't need to bump a version
  number by hand when you edit the CSS.
- **`NOINDEX`** (top of the file) controls whether every page ships
  `<meta name="robots" content="noindex, nofollow">`. Currently `True` by
  design (see root `CLAUDE.md`) — don't flip it without asking the account
  owner first, since it changes public search-engine visibility of PII-light
  but still personal content.
