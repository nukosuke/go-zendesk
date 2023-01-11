package zendesk

import (
	"context"
	"encoding/json"
	"time"
)

// App is a struct representing an app listed on the Zendesk Marketplace.
// https://developer.zendesk.com/api-reference/ticketing/apps/apps/#json-format
type App struct {
	ID                          int       `json:"id"`
	OwnerID                     int       `json:"owner_id"`
	Name                        string    `json:"name"`
	SingleInstall               bool      `json:"single_install"`
	DefaultLocale               string    `json:"default_locale"`
	AuthorName                  string    `json:"author_name"`
	AuthorEmail                 string    `json:"author_email"`
	AuthorURL                   string    `json:"author_url"`
	ShortDescription            string    `json:"short_description"`
	LongDescription             string    `json:"long_description"`
	RawLongDescription          string    `json:"raw_long_description"`
	InstallationInstructions    string    `json:"installation_instructions"`
	RawInstallationInstructions string    `json:"raw_installation_instructions"`
	Visibility                  string    `json:"visibility"`
	Enabled                     bool      `json:"enabled"`
	Installable                 bool      `json:"installable"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	FrameworkVersion            string    `json:"framework_version"`
	Featured                    bool      `json:"featured"`
	Promoted                    bool      `json:"promoted"`
	Products                    []string  `json:"products"`
	Version                     string    `json:"version"`
	MarketingOnly               bool      `json:"marketing_only"`
	Deprecated                  bool      `json:"deprecated"`
	Obsolete                    bool      `json:"obsolete"`
	Paid                        bool      `json:"paid"`
	State                       string    `json:"state"`
	ClosedPreview               bool      `json:"closed_preview"`
	TermsConditionsURL          string    `json:"terms_conditions_url"`
}

// AppAPI is an interface containing all methods associated with zendesk apps
type AppAPI interface {
	ListInstallations(ctx context.Context) ([]App, error)
}

// ListInstallations shows all apps installed in the current account.
// ref: https://developer.zendesk.com/api-reference/ticketing/apps/apps/#list-app-installations
func (z *Client) ListInstallations(ctx context.Context) ([]App, error) {
	var out struct {
		Installations []App `json:"installations"`
	}

	body, err := z.get(ctx, "/apps/installations")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	return out.Installations, err
}
