# Upgrade

Back up PostgreSQL before upgrading an existing SCR Mailer installation. Keep the backup and previous application version together for rollback.

## Binary

Stop the running application or service. Replace it with a published SCR Mailer binary, then run:

```sh
./scr-mailer --upgrade
./scr-mailer
```

Use the `scr-mailer` service descriptors supplied in this repository when installing a new service. Existing `/etc/listmonk` configuration directories are retained for compatibility.

## Docker: builds from this repository

Update your checkout to a reviewed SCR release or merged commit. Keep the existing database volume and credentials, then rebuild the application:

```sh
docker compose build app
docker compose up -d app
```

The provided Compose startup command runs idempotent installation checks and pending upgrades before starting SCR Mailer. Do not delete database volumes during an upgrade.

For deployments using a published image, set `image` to a published SCR tag, remove the local `build` section, and use `docker compose pull app` followed by `docker compose up -d app`.

## Preview builds

The manually triggered SCR preview-image workflow publishes the `nightly` image tag. Use a tested SCR release for production.

## Rollback

Stop the application, restore the pre-upgrade database backup, and restore the matching application version. Database schema changes may prevent an older binary from running against a newer database.

## Original developer updates

See [the upstream update workflow](https://github.com/sitescomputer/SCR-Mailer/blob/master/docs/UPSTREAM-UPDATES.md). Import upstream releases into an integration branch, preserve SCR branding in conflicts, test, and merge using a merge commit.

## Upgrading to v4.x.x

v4 is a major upgrade from prior versions with significant changes to certain important features and behaviour. It is the first version to have multi-user support and full fledged user management. Prior versions only had a simple BasicAuth for both admin login (browser prompt) and API calls, with the username and password defined in the TOML configuration file.

It is safe to upgrade an older installation with `--upgrade`, but there are a few important things to keep in mind. The upgrade automatically imports the `admin_username` and `admin_password` defined in the TOML configuration into the new user management system.

1. **New login UI**: Once you upgrade an older installation, the admin dashboard will no longer show the native browser prompt for login. Instead, a new login UI rendered by SCR Mailer is displayed at the URI `/admin/login`.

1. **API credentials**: If you are using APIs to interact with SCR Mailer, after logging in, go to Settings -> Users and create a new API user with the necessary permissions. Change existing API integrations to use these credentials instead of the old username and password defined in the legacy TOML configuration file or environment variables.

1. **Credentials in TOML file or old environment variables**: The admin dashboard shows a warning until the `admin_username` and `admin_password` fields are removed from the configuration file or old environment variables. In v4.x.x, these are irrelevant as user credentials are stored in the database and managed from the admin UI. IMPORTANT: if you are using APIs to interact with SCR Mailer, follow the previous step before removing the legacy credentials.


## Railway
- Head to your dashboard, and select your SCR Mailer project.
- Select the GitHub deployment service.
- In the Deployment tab, head to the latest deployment, click on the three vertical dots to the right, and select "Redeploy".

![Railway Redeploy option](https://user-images.githubusercontent.com/55474996/226517149-6dc512d5-f862-46f7-a57d-5e55b781ff53.png)
