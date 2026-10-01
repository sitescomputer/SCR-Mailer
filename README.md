# SCR Mailer

<img src="branding/scr-logo.png" alt="Sites Computer Resources" width="180" />

**SCR Mailer by Sites Computer Resources** is a self-hosted newsletter, mailing list, and transactional email application. It uses PostgreSQL and includes an admin dashboard, visual email editor, subscriber management, campaign tracking, and delivery integrations.

![SCR Mailer dashboard preview with sample data](docs/docs/content/images/scr-dashboard.png)

## Run with Docker

```sh
git clone https://github.com/sitescomputer/SCR-Mailer.git
cd SCR-Mailer
docker compose up -d --build
```

Open http://localhost:9000 and complete administrator setup. Configure your real public URL, SMTP service and sender address in Settings before sending mail.

The Compose file builds **this fork**, including SCR branding, rather than downloading an upstream image. Release images are published to `ghcr.io/sitescomputer/scr-mailer` when an SCR release tag is pushed and the release workflow succeeds. See [installation](docs/docs/content/installation.md).

## Build from source

Use the Go version in `.go-version`, Node.js 22, Yarn Classic and GNU Make (on Windows, use WSL or the development container):

```sh
make dist
./scr-mailer --new-config
# Edit config.toml with your PostgreSQL connection settings.
./scr-mailer --install
./scr-mailer
```

## Branding and upstream updates

- [Branding configuration](branding/README.md)
- [Maintaining upstream updates](docs/UPSTREAM-UPDATES.md)
- [Documentation](docs/docs/content/index.md)
- [Support and issues](https://github.com/sitescomputer/SCR-Mailer/issues)

The product and company identity are shared by the Go backend and Vue frontend through `branding/brand.json`. The supplied SCR logo is preserved unchanged in `branding/scr-logo.png`.

For an existing installation, review `scripts/apply-branding.sql` and the branding guide. A database backup is required before database upgrades. Existing custom sender addresses, root URLs and templates need deployment-specific review.

## License and attribution

SCR Mailer is based on [listmonk](https://github.com/knadh/listmonk), created by Kailash Nadh and contributors. The original [AGPLv3 license](LICENSE) is preserved. See [upstream attribution](UPSTREAM.md).
