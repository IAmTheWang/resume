#!/usr/bin/env python3
"""Renders data/resume_data*.py into a multilingual static site under dist/.

No templating library — plain f-strings. Run:

    python3 scripts/build_site.py
"""

from __future__ import annotations

import hashlib
import html
import importlib
import posixpath
import shutil
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
sys.path.insert(0, str(ROOT / "data"))

DIST = ROOT / "dist"

# Set to False once you want the site indexed by search engines.
NOINDEX = True

FAVICON = (
    '<link rel="icon" href="data:image/svg+xml,'
    '<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22>'
    '<text y=%22.9em%22 font-size=%2290%22>%F0%9F%93%84</text></svg>">'
)

LANG_TABS = [
    ("en", "EN", "index.html"),
    ("ja", "日本語", "ja/index.html"),
    ("zh", "中文", "zh/index.html"),
]

PRINT_LABEL = {
    "en": "Print / Save PDF",
    "ja": "印刷 / PDF保存",
    "zh": "打印 / 导出PDF",
}


def esc(text: str) -> str:
    return html.escape(text)


def truncate(text: str, limit: int = 200) -> str:
    text = text.strip()
    return text if len(text) <= limit else text[: limit - 1].rsplit(" ", 1)[0] + "…"


def relhref(target: str, current_output: str) -> str:
    current_dir = posixpath.dirname(current_output) or "."
    return posixpath.relpath(target, start=current_dir)


def contact_line(d) -> str:
    return (
        f'{esc(d.LOCATION)} &nbsp;·&nbsp; '
        f'<a href="mailto:{esc(d.EMAIL)}">{esc(d.EMAIL)}</a> &nbsp;·&nbsp; '
        f'<a href="{esc(d.LINKEDIN_URL)}">{esc(d.LINKEDIN_LABEL)}</a>'
    )


def lang_tabs_nav(current_output: str, current_lang: str) -> str:
    links = []
    for code, label, target in LANG_TABS:
        cls = ' class="active"' if code == current_lang else ""
        links.append(f'<a href="{relhref(target, current_output)}"{cls}>{esc(label)}</a>')
    return '<nav class="lang-tabs">' + "".join(links) + "</nav>"


def page_switch_nav(current_output: str, active: str) -> str:
    def link(target: str, label: str, key: str) -> str:
        cls = ' class="active"' if key == active else ""
        return f'<a href="{relhref(target, current_output)}"{cls}>{label}</a>'

    return (
        '<nav class="page-switch">'
        + link("index.html", "Resume", "resume")
        + link("cv.html", "Detailed CV", "cv")
        + "</nav>"
    )


def hero(d, page: dict) -> str:
    switch = page_switch_nav(page["output"], page["page_switch_active"]) if page["show_page_switch"] else ""
    return f"""
<header class="hero">
  <h1>{esc(d.NAME)}</h1>
  <p class="tagline">{esc(d.TAGLINE)}</p>
  <p class="contact-line">{contact_line(d)}</p>
  {lang_tabs_nav(page["output"], page["lang"])}
  {switch}
  <button onclick="window.print()" class="no-print print-btn">{esc(PRINT_LABEL[page["lang"]])}</button>
</header>
""".strip()


def ul(items: list[str]) -> str:
    lis = "\n".join(f"    <li>{esc(item)}</li>" for item in items)
    return f"  <ul>\n{lis}\n  </ul>"


def skills_section(d) -> str:
    rows = "\n".join(
        f'    <div class="skill-row"><dt>{esc(label)}</dt><dd>{esc(value)}</dd></div>'
        for label, value in d.SKILLS
    )
    return f"""
<section>
  <h2>Skills</h2>
  <dl class="skills-grid">
{rows}
  </dl>
</section>
""".strip()


def education_languages_section(d) -> str:
    degree, school = d.EDUCATION
    lang_items = "\n".join(
        f"    <li><strong>{esc(label)}:</strong> {esc(value)}</li>" for label, value in d.LANGUAGES
    )
    return f"""
<section>
  <h2>Education</h2>
  <p>{esc(degree)}<br>{esc(school)}</p>
</section>
<section>
  <h2>Languages &amp; Work Authorization</h2>
  <ul class="plain-list">
{lang_items}
  </ul>
</section>
""".strip()


def experience_item(job: dict, detailed: bool) -> str:
    meta_line = f" &nbsp;·&nbsp; {esc(job['meta'])}" if job.get("meta") else ""
    parts = [
        '<div class="experience-item">',
        '  <div class="role-line">',
        f"    <h3>{esc(job['role'])}</h3>",
        f'    <span class="period">{esc(job["period"])}</span>',
        "  </div>",
        f'  <p class="company-line">{esc(job["company"])} &nbsp;·&nbsp; {esc(job["location"])}{meta_line}</p>',
        f'  <p class="blurb">{esc(job["blurb"])}</p>',
    ]
    if detailed and job.get("responsibilities"):
        parts.append('  <p class="subhead">Key Contributions</p>')
        parts.append(ul(job["responsibilities"]))
    parts.append('  <p class="subhead">Achievements</p>' if detailed else "")
    parts.append(ul(job["achievements"]))
    parts.append("</div>")
    return "\n".join(p for p in parts if p)


def summary_section(d, detailed: bool) -> str:
    if detailed:
        return f"""
<section>
  <h2>Summary</h2>
{ul(d.SUMMARY_LONG)}
</section>
""".strip()
    return f"""
<section>
  <h2>Summary</h2>
  <p>{esc(d.SUMMARY)}</p>
</section>
""".strip()


def core_strengths_section(d) -> str:
    return f"""
<section>
  <h2>Core Strengths</h2>
{ul(d.CORE_STRENGTHS)}
</section>
""".strip()


def personal_statement_section(d) -> str:
    return f"""
<section>
  <h2>Personal Statement</h2>
  <p>{esc(d.PERSONAL_STATEMENT)}</p>
</section>
""".strip()


def og_description(d, detailed: bool) -> str:
    text = d.SUMMARY_LONG[0] if detailed or not hasattr(d, "SUMMARY") else d.SUMMARY
    return truncate(text)


def head(d, page: dict, css_hash: str) -> str:
    css_href = relhref("style.css", page["output"]) + f"?v={css_hash}"
    robots = '<meta name="robots" content="noindex, nofollow">\n' if NOINDEX else ""
    title = esc(page["title"])
    description = esc(og_description(d, page["detailed"]))
    return f"""<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{title}</title>
{robots}{FAVICON}
<meta property="og:title" content="{title}">
<meta property="og:description" content="{description}">
<meta property="og:type" content="website">
<link rel="stylesheet" href="{css_href}">
</head>"""


def build_page_body(d, page: dict) -> str:
    detailed = page["detailed"]
    exp_items = "\n".join(experience_item(job, detailed=detailed) for job in d.EXPERIENCE)
    sections = [hero(d, page), summary_section(d, detailed)]
    if detailed:
        sections.append(core_strengths_section(d))
    sections.append(skills_section(d))
    sections.append(f"<section>\n  <h2>Experience</h2>\n{exp_items}\n</section>")
    sections.append(education_languages_section(d))
    if detailed:
        sections.append(personal_statement_section(d))
    return "\n".join(sections)


def build_page(page: dict, css_hash: str) -> str:
    d = importlib.import_module(page["module"])
    body = build_page_body(d, page)
    return f"""<!doctype html>
<html lang="{page["lang"]}">
{head(d, page, css_hash)}
<body>
<div class="page">
{body}
</div>
</body>
</html>
"""


PAGES = [
    {
        "module": "resume_data",
        "lang": "en",
        "output": "index.html",
        "title": "Hongcheng Wang — Resume",
        "detailed": False,
        "show_page_switch": True,
        "page_switch_active": "resume",
    },
    {
        "module": "resume_data",
        "lang": "en",
        "output": "cv.html",
        "title": "Hongcheng Wang — Detailed CV",
        "detailed": True,
        "show_page_switch": True,
        "page_switch_active": "cv",
    },
    {
        "module": "resume_data_ja",
        "lang": "ja",
        "output": "ja/index.html",
        "title": "職務経歴書 — 王鴻乗",
        "detailed": True,
        "show_page_switch": False,
        "page_switch_active": None,
    },
    {
        "module": "resume_data_zh",
        "lang": "zh",
        "output": "zh/index.html",
        "title": "个人简历 — 王鸿乘",
        "detailed": True,
        "show_page_switch": False,
        "page_switch_active": None,
    },
]


def main() -> None:
    DIST.mkdir(exist_ok=True)
    css_src = ROOT / "templates" / "style.css"
    css_bytes = css_src.read_bytes()
    css_hash = hashlib.md5(css_bytes).hexdigest()[:8]
    shutil.copy(css_src, DIST / "style.css")

    for page in PAGES:
        out_path = DIST / page["output"]
        out_path.parent.mkdir(parents=True, exist_ok=True)
        out_path.write_text(build_page(page, css_hash), encoding="utf-8")
        print(f"Built {out_path}")


if __name__ == "__main__":
    main()
