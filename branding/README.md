# SCR branding

`brand.json` is the shared product/company identity for the Go backend and Vue admin interface. `scr-logo.png` is the exact supplied SCR logo. Copies under the frontend, public pages and documentation are delivery assets; SVG files wrap the same PNG without redrawing it.

The company URL currently points to the SCR Mailer repository because a company website was not supplied. Change `companyURL` to the real website when ready. Set the actual deployment address and sender email through **Settings → General**. Example addresses are placeholders and must be configured before sending.

The admin color overrides are in `frontend/src/assets/_scr-brand.scss`. Public page colors are in `static/public/static/style.css`; visual editor colors are in `frontend/email-builder/src/theme.ts`.

Template helpers `ProductName`, `CompanyName`, `CompanyURL`, `DocsURL` and `LogoURL` are available in subscriber pages and emails. Empty logo/favicon settings fall back to the included SCR assets. Existing custom URLs remain respected.

## Existing installations

1. Back up the database and deploy the branded application built from this repository.
2. Configure the site name, public root URL, SMTP service and real sender address in Settings.
3. Review `scripts/apply-branding.sql` before applying it. It only changes upstream default names, disables the upstream release feed, and replaces exact legacy footer strings in non-visual templates.
4. Edit existing visual templates in the editor so the saved JSON source and rendered HTML both receive the branding. Review custom templates and draft campaigns individually; repository defaults apply to new installations.
5. Check login, dashboard, subscription, unsubscribe, opt-in, archive and delivered test emails.

Database names, volume keys, Go imports and `LISTMONK_*`/`X-Listmonk-*` integration identifiers remain compatible with upstream. License and copyright notices are retained in `LICENSE` and `UPSTREAM.md`.
