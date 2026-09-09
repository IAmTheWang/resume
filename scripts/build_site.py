#!/usr/bin/env python3
"""Renders data/resume_data.py into dist/index.html and dist/cv.html.

No templating library — plain f-strings, matching the style of the docx
build scripts this data was trimmed from. Run:

    python3 scripts/build_site.py
"""

from __future__ import annotations

import html
import shutil
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
sys.path.insert(0, str(ROOT / "data"))

import resume_data as d  # noqa: E402

DIST = ROOT / "dist"


def esc(text: str) -> str:
    return html.escape(text)


def contact_line() -> str:
    return (
        f'{esc(d.LOCATION)} &nbsp;·&nbsp; '
        f'<a href="mailto:{esc(d.EMAIL)}">{esc(d.EMAIL)}</a> &nbsp;·&nbsp; '
        f'<a href="{esc(d.LINKEDIN_URL)}">{esc(d.LINKEDIN_LABEL)}</a>'
    )


def nav(active: str) -> str:
    def link(href: str, label: str, key: str) -> str:
        cls = ' class="active"' if key == active else ""
        return f'<a href="{href}"{cls}>{label}</a>'

    return (
        '<nav class="page-switch">'
        + link("index.html", "Resume", "resume")
        + link("cv.html", "Detailed CV", "cv")
        + "</nav>"
    )


def hero(active: str) -> str:
    return f"""
<header class="hero">
  <h1>{esc(d.NAME)}</h1>
  <p class="tagline">{esc(d.TAGLINE)}</p>
  <p class="contact-line">{contact_line()}</p>
  {nav(active)}
</header>
""".strip()


def ul(items: list[str]) -> str:
    lis = "\n".join(f"    <li>{esc(item)}</li>" for item in items)
    return f"  <ul>\n{lis}\n  </ul>"


def skills_section() -> str:
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


def education_languages_section() -> str:
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


def build_resume_page() -> str:
    exp_items = "\n".join(experience_item(job, detailed=False) for job in d.EXPERIENCE)
    return f"""<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{esc(d.NAME)} — Resume</title>
<link rel="stylesheet" href="style.css">
</head>
<body>
<div class="page">
{hero("resume")}
<section>
  <h2>Summary</h2>
  <p>{esc(d.SUMMARY)}</p>
</section>
{skills_section()}
<section>
  <h2>Experience</h2>
{exp_items}
</section>
{education_languages_section()}
</div>
</body>
</html>
"""


def build_cv_page() -> str:
    exp_items = "\n".join(experience_item(job, detailed=True) for job in d.EXPERIENCE)
    summary_items = ul(d.SUMMARY_LONG)
    strengths_items = ul(d.CORE_STRENGTHS)
    return f"""<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{esc(d.NAME)} — Detailed CV</title>
<link rel="stylesheet" href="style.css">
</head>
<body>
<div class="page">
{hero("cv")}
<section>
  <h2>Summary</h2>
{summary_items}
</section>
<section>
  <h2>Core Strengths</h2>
{strengths_items}
</section>
{skills_section()}
<section>
  <h2>Experience</h2>
{exp_items}
</section>
{education_languages_section()}
<section>
  <h2>Personal Statement</h2>
  <p>{esc(d.PERSONAL_STATEMENT)}</p>
</section>
</div>
</body>
</html>
"""


def main() -> None:
    DIST.mkdir(exist_ok=True)
    shutil.copy(ROOT / "templates" / "style.css", DIST / "style.css")
    (DIST / "index.html").write_text(build_resume_page(), encoding="utf-8")
    (DIST / "cv.html").write_text(build_cv_page(), encoding="utf-8")
    print(f"Built {DIST / 'index.html'} and {DIST / 'cv.html'}")


if __name__ == "__main__":
    main()
