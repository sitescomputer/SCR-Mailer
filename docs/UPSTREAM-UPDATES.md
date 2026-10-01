# Maintaining SCR Mailer with upstream releases

`origin` is `https://github.com/sitescomputer/SCR-Mailer.git`. `upstream` is `https://github.com/knadh/listmonk.git`. The upstream push URL is disabled to prevent accidental pushes to the original project.

## Bring in a selected stable release

Start with committed changes. Replace the example release below with the upstream version you want to integrate:

```powershell
./scripts/sync-upstream.ps1 -Release v6.2.0
```

The script fetches the original project, imports its tags under `upstream/v*`, and creates an `update/upstream-v*` branch from your current `origin/master`. It merges the selected release with a merge commit; it does not overwrite branding, push, publish or deploy anything automatically.

If conflicts occur, combine upstream functional changes with the SCR identity. Do not accept entire conflicting files blindly. Then:

```powershell
git add <resolved-files>
git merge --continue
```

To abandon an unresolved merge, run `git merge --abort`; your original branded branch remains available.

## Verify and submit

Use the Go version in `.go-version`, Node 22, Yarn Classic and GNU Make:

```sh
go test ./...
make dist
docker compose config --quiet
docker build -t scr-mailer:check .
```

Check the SCR logo and company name on login, dashboard, browser tabs, public subscription/opt-in/unsubscribe pages, archives, and test emails. Verify campaign links, unsubscribe and tracking still work. Check new features for newly introduced upstream branding.

```powershell
git push -u origin update/upstream-v6.2.0
```

Open a pull request into `master`. **Merge upstream integration PRs using a merge commit, not squash or rebase**, so Git retains the imported upstream ancestry and can combine future releases correctly. Ordinary SCR feature PRs can use your usual merge policy.

## Publish a branded release

After the integration is tested and merged, use a distinct SCR tag such as `v6.2.0-scr.1`:

```powershell
git switch master
git pull --ff-only origin master
git tag v6.2.0-scr.1
git push origin v6.2.0-scr.1
```

The SCR release workflow builds branded binaries and publishes the Docker image to `ghcr.io/sitescomputer/scr-mailer`. It uses the repository's `GITHUB_TOKEN`, without Docker Hub credentials. GitHub Actions must be enabled, with package publishing allowed. Review package visibility if anonymous image pulls are desired.

Take a database backup before upgrading a live installation. Changes to source defaults do not automatically rewrite existing database templates or production settings. See [the branding guide](../branding/README.md).
