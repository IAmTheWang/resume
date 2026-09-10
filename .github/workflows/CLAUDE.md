# .github/workflows/

`deploy.yml` — the only workflow. Triggers on push to `main` (plus manual
`workflow_dispatch`). Two jobs:

1. **build**: checks out the repo, sets up Go, runs `go run ./cmd/build`,
   uploads `dist/` as a Pages artifact.
2. **deploy**: takes that artifact and publishes it via
   `actions/deploy-pages`.

## Things that will break if changed carelessly

- **`permissions: contents: read, pages: write, id-token: write`** at the top
  is required — `actions/deploy-pages` authenticates via OIDC and the deploy
  step 403s without `id-token: write` specifically. Don't trim this block.
- **This workflow deploying successfully is not enough on its own.** The
  repo's Settings → Pages → Source must be set to **"GitHub Actions"** (a
  one-time manual step in the GitHub web UI — no workflow file can set it).
  If `deploy` fails with a Pages-environment error even though `build`
  succeeded, check this setting before debugging the YAML.
- **Pinned action versions need periodic bumping.** GitHub Actions began
  forcing JS-based actions onto Node 24 in mid-2026 and is removing the Node
  20 runtime from runners entirely; actions pinned to old majors that still
  ship a Node 20 runtime will eventually stop working. Check
  `actions/checkout`, `actions/setup-go`, `actions/upload-pages-artifact`,
  and `actions/deploy-pages` for newer majors periodically (e.g.
  `gh api repos/actions/<name>/releases/latest`) rather than assuming the
  versions here stay current indefinitely.
- **No secrets are configured or needed** — `deploy-pages` uses the ambient
  `GITHUB_TOKEN` via OIDC, not a PAT. Don't add one unless a future step
  genuinely requires it.
- **No `go.sum`/dependency-download step exists** because `cmd/build` only
  uses the Go standard library — same "zero third-party deps" property the
  old Python build had. If a future change adds a real dependency, `go run`
  will still work in CI (it resolves modules on demand), but consider whether
  the zero-dependency property is worth preserving before adding one.
