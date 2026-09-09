# templates/

`style.css` is the only file here — one stylesheet shared by every page in
every language, copied verbatim into `dist/style.css` by
`scripts/build_site.py` (see that folder's `CLAUDE.md` for why the copy step
matters).

## Design system

CSS custom properties on `:root`: `--accent`/`--accent-dark`/`--accent-soft`
(the blue accent family), `--ink`/`--muted` (text colors), `--line` (hairline
borders), `--page-bg`/`--card-bg` (the site renders as a white "card" floating
on a light gray canvas — see `.page`). Reuse these tokens rather than
introducing new hardcoded colors.

`color-scheme: light` plus explicit `background`/`color` on `body` force light
mode even if the visitor's OS/browser is set to force-dark — deliberate, so a
resume never renders as unreadable inverted colors.

## CJK font isolation — don't collapse this into one font-family list

`:lang(ja)` and `:lang(zh)` each set their own font stack (Hiragino/Noto Sans
JP for Japanese, PingFang/Microsoft YaHei/Noto Sans SC for Chinese), separate
from the base Latin stack (Inter, loaded from Google Fonts, plus system
fallbacks). This exists because of Han Unification: the same Unicode code
point can have different standard glyph shapes in Chinese vs. Japanese, and a
single fallback list risks rendering, say, Chinese text in a Japanese font —
subtly wrong stroke shapes to a native reader. Keep the `:lang()` split when
touching typography.

## Print stylesheet (`@media print`)

This is the section most likely to silently regress — it undoes several
screen-only decorations that either waste ink or render broken once the
browser drops background colors on print:

- `.avatar`, `nav`, `.no-print` are hidden entirely.
- `.page`'s card treatment (shadow, radius, margin, padding, canvas
  background) is stripped and `max-width` forced to `100%` — otherwise the
  card's own padding stacks with the printer's physical page margins and
  wastes paper.
- The experience timeline (`.experience-item`'s left border + `::before`
  dot) is a screen-only flourish — printed copies fall back to a plain
  top-to-bottom layout, because dots/lines drawn purely with
  `background-color` can render as broken outlines or vanish once print mode
  drops background colors.
- `.period` pill backgrounds are stripped back to plain text for the same
  reason.
- `a[href]::after { content: "" !important; }` suppresses the URL some
  browsers auto-append after printed links (e.g. `email@x.com
  (mailto:email@x.com)`).

If you add a new decorative element (another badge, another background-color
trick), add its print-mode teardown here in the same pass — don't leave it to
be discovered later via a bad-looking PDF.
