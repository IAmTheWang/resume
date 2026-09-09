# data/

One Python module per language — the only place resume content should be
edited. `scripts/build_site.py` imports these by module name (see the `PAGES`
list) and renders them; nothing else in this repo hardcodes resume content.

| Module | Language | Pages it feeds |
|---|---|---|
| `resume_data.py` | English | `index.html` (short Resume) **and** `cv.html` (Detailed CV) |
| `resume_data_ja.py` | 日本語 | `ja/index.html` (one page only) |
| `resume_data_zh.py` | 中文 | `zh/index.html` (one page only) |

## Why English gets two pages and Japanese/Chinese get one

The private source repo (`~/personal/resumes`) only ever produces **one**
combined resume document per language for Japanese and Chinese — there's no
short-form "Resume" variant for those, only the full detailed one. English is
the only language with both a short Resume and a separate Detailed CV. This
site mirrors that reality on purpose: `resume_data_ja.py` / `resume_data_zh.py`
have no `SUMMARY` field (only `SUMMARY_LONG`), and `build_site.py` renders
them with `detailed=True` — the same rendering path as the English Detailed
CV. Don't invent a short-form summary for JA/ZH; it doesn't exist upstream.

## Origin and syncing

These are trimmed, privacy-redacted copies of the private source-of-truth
files:

| This file | Private source (`~/personal/resumes/`) |
|---|---|
| `resume_data.py` | `resume_data.py` |
| `resume_data_ja.py` | `resume_data_jp.py` |
| `resume_data_zh.py` | `resume_data_cn.py` |

Redaction applied on every copy: `PHONE` and detailed `ADDRESS` are dropped;
`LOCATION` (a city-level string), `LINKEDIN_LABEL`, and `LINKEDIN_URL` are
added instead. Field names were normalized across all three modules so
`scripts/build_site.py` can render any of them generically.

**These files have already diverged from the private source on some content**
(e.g. the `EDUCATION` field) after the account owner manually edited them
directly in this repo. Don't blindly re-sync from the private repo and
overwrite local edits — diff first and check with the user if a private-repo
value disagrees with what's here.

## Required fields (all three modules)

`NAME`, `TAGLINE`, `LOCATION`, `EMAIL`, `LINKEDIN_LABEL`, `LINKEDIN_URL`,
`SUMMARY_LONG` (list of paragraphs), `CORE_STRENGTHS` (list), `SKILLS` (list
of `(category, value)` tuples), `EXPERIENCE` (list of dicts — see any existing
entry for the shape: `company`/`role`/`period`/`location`/`meta`/`blurb`/
`responsibilities`/`achievements`), `EDUCATION` (a `(degree, school)` tuple),
`LANGUAGES` (list of `(label, value)` tuples), `PERSONAL_STATEMENT`.

`resume_data.py` additionally has `SUMMARY` (a single short paragraph) —
required only because it's the sole module rendered with `detailed=False`.
