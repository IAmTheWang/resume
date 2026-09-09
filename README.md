# Resume site

Static resume + detailed CV, published via GitHub Pages. Content lives in
`data/resume_data.py`; `scripts/build_site.py` renders it into `dist/index.html`
and `dist/cv.html`. GitHub Actions (`.github/workflows/deploy.yml`) rebuilds
and deploys on every push to `main`.

## Local preview

```bash
python3 scripts/build_site.py
open dist/index.html
```

## Updating content

Edit `data/resume_data.py`, then push to `main` — the site rebuilds
automatically. No need to run the build script yourself unless you want to
preview locally first.
