-- Optional, idempotent update for an EXISTING database. Back up first and review before running.
-- This only replaces upstream defaults and exact legacy footer strings; custom identities stay intact.
BEGIN;
UPDATE settings SET value = '"SCR Mailer"'::jsonb
WHERE key = 'app.site_name' AND value IN ('"Mailing list"'::jsonb, '"listmonk"'::jsonb, '"Listmonk"'::jsonb);
UPDATE settings SET value = 'false'::jsonb WHERE key = 'app.check_updates';
UPDATE templates SET body = replace(
    replace(replace(body, '{{ L.T "public.poweredBy" }} ', ''), 'href="https://listmonk.app"', 'href="{{ CompanyURL }}"'),
    '>listmonk</a>', '>{{ ProductName }} by {{ CompanyName }}</a>'
)
WHERE type <> 'campaign_visual' AND body LIKE '%href="https://listmonk.app"%>listmonk</a>%';
COMMIT;
-- Keep the real root URL, sender addresses, custom logo URLs and existing campaigns unchanged.
-- Review stored visual templates in the editor so their JSON source and HTML stay consistent.
