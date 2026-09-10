# internal/resume/

One Go file per language — the only place resume content should be edited.
`internal/site/pages.go` references these directly (`&resume.EN`, `&resume.JA`,
`&resume.ZH`); nothing else in this repo hardcodes resume content.

| File | Language | Pages it feeds |
|---|---|---|
| `data_en.go` (`resume.EN`) | English | `index.html` (short Resume) **and** `cv.html` (Detailed CV) |
| `data_ja.go` (`resume.JA`) | 日本語 | `ja/index.html` (one page only) |
| `data_zh.go` (`resume.ZH`) | 中文 | `zh/index.html` (one page only) |

## Why English gets two pages and Japanese/Chinese get one

The private source repo (`~/personal/resumes`) only ever produces **one**
combined resume document per language for Japanese and Chinese — there's no
short-form "Resume" variant for those, only the full detailed one. English is
the only language with both a short Resume and a separate Detailed CV. This
site mirrors that reality on purpose: `resume.JA` / `resume.ZH` leave
`HasSummary` false and `Summary` empty (only `SummaryLong` is set), and
`internal/site/pages.go` renders them with `Detailed: true` — the same
rendering path as the English Detailed CV. Don't invent a short-form summary
for JA/ZH; it doesn't exist upstream.

## Origin and syncing

These are trimmed, privacy-redacted, Go-struct ports of the private
source-of-truth files:

| This file | Private source (`~/personal/resumes/`) |
|---|---|
| `data_en.go` | `resume_data.py` |
| `data_ja.go` | `resume_data_jp.py` |
| `data_zh.go` | `resume_data_cn.py` |

Redaction applied on every copy: phone number and detailed street address are
dropped; `Location` (a city-level string), `LinkedInLabel`, and `LinkedInURL`
are added instead. Field names are normalized across all three files so
`internal/site` can render any of them generically via the shared
`resume.ResumeData` struct.

**These files have already diverged from the private source on some content**
(e.g. the `Education` field) after the account owner manually edited them
directly in this repo. Don't blindly re-sync from the private repo and
overwrite local edits — diff first and check with the user if a private-repo
value disagrees with what's here.

## Required fields (all three files)

`Name`, `Tagline`, `Location`, `Email`, `LinkedInLabel`, `LinkedInURL`,
`SummaryLong` (paragraphs), `CoreStrengths`, `Skills` (`[]SkillEntry`),
`Experience` (`[]Job` — see any existing entry for the shape:
`Company`/`Role`/`Period`/`Location`/`Meta`/`Blurb`/`Responsibilities`/
`Achievements`), `Education` (`Education{Degree, School}`), `Languages`
(`[]LangEntry`), `PersonalStatement`.

`data_en.go` additionally sets `HasSummary: true` and `Summary` (a single
short paragraph) — required only because EN is the sole dataset rendered with
`Detailed: false` (on `index.html`).

## Content asymmetries are real, not translation gaps

Don't assume the three language files are structurally parallel beyond field
*names* and *types*:

- `Languages` length varies: EN/JA have 4 entries (including a work-authorization
  / visa entry), ZH has only 2 (no visa entry, no "Chinese: native" entry).
- The same job's `Responsibilities` bullet count differs per language — e.g.
  EN has several for the GTMap role, JA/ZH have only one. These are
  independently authored per language, not literal translations of each
  other.
- Company names are localized differently — ZH appends parenthetical
  qualifiers (e.g. `"Welby, Inc.（日本）"`) that EN/JA don't.

When editing, treat each language's data as its own source of truth; don't
"fix" one language to match another's structure.
