# Installation

SCR Mailer by Sites Computer Resources requires PostgreSQL 12 or newer.

## Docker: build this repository

```sh
git clone https://github.com/sitescomputer/SCR-Mailer.git
cd SCR-Mailer
docker compose up -d --build
```

Open http://localhost:9000 and create the administrator account. Set your actual public URL, SMTP credentials and sender address in Settings before sending campaigns.

The Compose file builds `scr-mailer:local` from `Dockerfile.build`. It includes the branded frontend, subscriber pages and email templates. Keep the existing `listmonk-data` volume and database identifiers when upgrading an existing installation.

### Published release images

After an SCR release workflow succeeds, images are available from `ghcr.io/sitescomputer/scr-mailer`. To use a published version, remove the `build` section from Compose and set `image` to a specific SCR tag, for example `ghcr.io/sitescomputer/scr-mailer:v6.2.0-scr.1` **only after that tag has been published**. Private packages require a registry login.

## Binary

Download a published SCR release from [the repository releases](https://github.com/sitescomputer/SCR-Mailer/releases), or build with `make dist`.

```sh
./scr-mailer --new-config
# Configure PostgreSQL in config.toml.
./scr-mailer --install
./scr-mailer
```

Use `scr-mailer.exe` on Windows. Do not run `--install` against an existing database; use the documented upgrade procedure and take a backup first.

## Configuration and compatibility

The environment variable prefix remains `LISTMONK_` so existing installations and integrations stay compatible. The Docker entrypoint also supports `LISTMONK_*_FILE` secrets.

The container keeps `/listmonk` as its working directory and `/listmonk/uploads` as the upload path for existing mounts. These are technical paths; the application and binary are SCR Mailer.

See [configuration](configuration.md), [branding](https://github.com/sitescomputer/SCR-Mailer/blob/master/branding/README.md), and [upstream updates](https://github.com/sitescomputer/SCR-Mailer/blob/master/docs/UPSTREAM-UPDATES.md).
