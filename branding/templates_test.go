package branding_test

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/knadh/listmonk/branding"
	"github.com/knadh/listmonk/internal/i18n"
	"github.com/knadh/listmonk/internal/manager"
)

func templateContext(t *testing.T) (template.FuncMap, map[string]any) {
	t.Helper()
	b, err := os.ReadFile("../i18n/en.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := i18n.New(b)
	if err != nil {
		t.Fatal(err)
	}
	m := manager.New(manager.Config{
		RootURL: "https://mailer.example.com",
		LogoURL: "https://mailer.example.com/public/static/logo.png",
	}, nil, l, log.New(io.Discard, "", 0))
	t.Cleanup(m.Close)
	f := m.GenericTemplateFuncs()
	f["L"] = func() *i18n.I18n { return l }
	f["RootURL"] = func() string { return "https://mailer.example.com" }
	f["UnsubscribeURL"] = func() string { return "https://mailer.example.com/subscription/test" }
	f["MessageURL"] = func() string { return "https://mailer.example.com/campaign/test" }
	f["OptinURL"] = func() string { return "https://mailer.example.com/subscription/optin/test" }
	f["TrackView"] = func() template.HTML { return `<img src="https://mailer.example.com/tracking.png" />` }
	f["Safe"] = func(s string) template.HTML { return template.HTML(s) }
	f["Date"] = func(string) string { return "2026-10-01" }
	data := map[string]any{
		"L":                   l,
		"SiteName":            branding.Current.ProductName,
		"RootURL":             "https://mailer.example.com",
		"LogoURL":             "https://mailer.example.com/public/static/logo.png",
		"EnablePublicSubPage": true,
		"Campaign":            map[string]any{"Subject": "Campaign subject"},
		"Subscriber":          map[string]any{"Name": "Test Subscriber", "LastName": "Subscriber"},
		"Tx":                  map[string]any{"Data": map[string]any{"order_id": "SCR-123", "shipping_date": "2026-10-01"}},
		"Data":                map[string]any{"Title": "Login", "PasswordEnabled": true, "Nonce": "nonce", "NextURI": "/admin"},
	}
	return f, data
}

// Rendering the real templates catches missing helpers and escaping errors that can break delivered mail.
func TestCampaignTemplatesRenderWithBranding(t *testing.T) {
	for _, name := range []string{"default.tpl", "default-archive.tpl", "default-visual.tpl", "sample-tx.tpl"} {
		t.Run(name, func(t *testing.T) {
			f, data := templateContext(t)
			b, err := os.ReadFile(filepath.Join("../static/email-templates", name))
			if err != nil {
				t.Fatal(err)
			}
			tpl, err := template.New("email").Funcs(f).Parse(string(b) + `{{ define "content" }}Subscriber content{{ end }}`)
			if err != nil {
				t.Fatal(err)
			}
			var out bytes.Buffer
			if err := tpl.Execute(&out, data); err != nil {
				t.Fatal(err)
			}
			for _, text := range []string{branding.Current.ProductName, branding.Current.CompanyName} {
				if !strings.Contains(out.String(), text) {
					t.Fatalf("rendered email missing %q", text)
				}
			}
			if strings.Contains(out.String(), "ZgotmplZ") {
				t.Fatal("brand URL was rejected by the HTML renderer")
			}
			if name == "default.tpl" || name == "default-visual.tpl" {
				for _, text := range []string{"/subscription/test", "/campaign/test", "/tracking.png"} {
					if !strings.Contains(out.String(), text) {
						t.Fatalf("rendered email lost %q", text)
					}
				}
			}
		})
	}
}

func TestPublicAndNotificationTemplatesRender(t *testing.T) {
	f, data := templateContext(t)
	for _, test := range []struct{ glob, name string }{
		{"../static/public/templates/*.html", "admin-login"},
		{"../static/email-templates/*.html", "smtp-test"},
	} {
		t.Run(test.name, func(t *testing.T) {
			tpl, err := template.New("root").Funcs(f).ParseGlob(test.glob)
			if err != nil {
				t.Fatal(err)
			}
			var out bytes.Buffer
			if err := tpl.ExecuteTemplate(&out, test.name, data); err != nil {
				t.Fatal(err)
			}
			for _, text := range []string{branding.Current.ProductName, branding.Current.CompanyName, "/public/static/logo.png"} {
				if !strings.Contains(out.String(), text) {
					t.Fatalf("rendered page missing %q", text)
				}
			}
		})
	}
}
