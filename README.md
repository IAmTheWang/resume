# Resume site

Live at https://iamthewang.github.io/resume/

Static resume site, published via GitHub Pages, with EN / 日本語 / 中文 tabs.
`scripts/build_site.py` renders `data/resume_data*.py` into `dist/`. GitHub
Actions (`.github/workflows/deploy.yml`) rebuilds and deploys on every push
to `main`.

Pages:
- `dist/index.html` — English Resume (short form)
- `dist/cv.html` — English Detailed CV
- `dist/ja/index.html` — 日本語（単一ページ、Detailed CV 相当の詳細度）
- `dist/zh/index.html` — 中文（单页，详细程度对应 Detailed CV）

日本語・中文はそれぞれ 1 ページのみ。私有リポジトリ側でも日本語・中国語版の
応募書類は元々 1 種類しか作っていない（英語だけ Resume/Detailed CV の 2 種）
ため、存在しない「精簡版」を捏造しない方針。

## Local preview

```bash
python3 scripts/build_site.py
open dist/index.html
```

## Updating content

- English: edit `data/resume_data.py`
- 日本語: edit `data/resume_data_ja.py`
- 中文: edit `data/resume_data_zh.py`

Push to `main` — the site rebuilds automatically. No need to run the build
script yourself unless you want to preview locally first.

## Notes

- Contact info on every page is redacted for public visibility: no phone
  number, no street address — just email, LinkedIn, and city.
- Pages are `noindex, nofollow` by default (see `NOINDEX` in
  `scripts/build_site.py`) — reachable by direct link, not meant to show up
  in search results. Flip that flag if you want search engines to index it.
