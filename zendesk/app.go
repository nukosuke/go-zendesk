package zendesk

import (
	"context"
	"encoding/json"
	"time"
)

// AppInstallation is a struct representing an app that has been installed from the Zendesk Marketplace.
// https://developer.zendesk.com/api-reference/ticketing/apps/apps/#example-responses-11
type AppInstallation struct {
	Id       int64  `json:"id"`
	AppId    int64  `json:"app_id"`
	Product  string `json:"product"`
	Settings struct {
		Name  string `json:"name"`
		Title string `json:"title"`
	} `json:"settings"`
	SettingsObjects []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"settings_objects"`
	Enabled                   bool      `json:"enabled"`
	Updated                   string    `json:"updated"`
	UpdatedAt                 time.Time `json:"updated_at"`
	CreatedAt                 time.Time `json:"created_at"`
	RecurringPayment          bool      `json:"recurring_payment"`
	Collapsible               bool      `json:"collapsible"`
	Paid                      bool      `json:"paid"`
	HasUnpaidSubscription     bool      `json:"has_unpaid_subscription"`
	HasIncompleteSubscription bool      `json:"has_incomplete_subscription"`
}

// AppAPI is an interface containing all methods associated with zendesk apps
type AppAPI interface {
	ListInstallations(ctx context.Context) ([]AppInstallation, error)
}

// ListInstallations shows all apps installed in the current account.
// ref: https://developer.zendesk.com/api-reference/ticketing/apps/apps/#list-app-installations
func (z *Client) ListInstallations(ctx context.Context) ([]AppInstallation, error) {
	var out struct {
		Installations []AppInstallation `json:"installations"`
	}

	body, err := z.get(ctx, "/apps/installations")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	return out.Installations, err
}
