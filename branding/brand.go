// Package branding contains the SCR Mailer identity shared with the frontend.
package branding

import (
	_ "embed"
	"encoding/json"
	"html/template"
)

//go:embed brand.json
var identityJSON []byte

// Identity is the public product identity. Technical integration names remain compatible with upstream.
type Identity struct {
	ProductName   string `json:"productName"`
	CompanyName   string `json:"companyName"`
	CompanyURL    string `json:"companyURL"`
	RepositoryURL string `json:"repositoryURL"`
	DocsURL       string `json:"docsURL"`
	SupportURL    string `json:"supportURL"`
}

// Current is embedded into every build, including unpacked development builds.
var Current = func() Identity {
	var identity Identity
	if err := json.Unmarshal(identityJSON, &identity); err != nil {
		panic(err)
	}
	return identity
}()

// TemplateFuncs provides the identity in public pages, notifications and campaign templates.
func TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"ProductName": func() string { return Current.ProductName },
		"CompanyName": func() string { return Current.CompanyName },
		"CompanyURL":  func() string { return Current.CompanyURL },
		"DocsURL":     func() string { return Current.DocsURL + "/index.md" },
	}
}
