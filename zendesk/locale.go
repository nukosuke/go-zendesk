package zendesk

import (
	"encoding/json"
	"time"
)

// Locale is zendesk locale JSON payload format
// https://developer.zendesk.com/rest_api/docs/support/locales
type Locale struct {
	ID        int64     `json:"id"`
	URL       string    `json:"url"`
	Locale    string    `json:"locale"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LocaleAPI an interface containing all of the local related zendesk methods
type LocaleAPI interface {
	GetLocales() ([]Locale, error)
}

// GetLocales lists the translation locales available for the account.
// https://developer.zendesk.com/rest_api/docs/support/locales#list-locales
func (z *Client) GetLocales() ([]Locale, error) {
	var data struct {
		Locales []Locale `json:"locales"`
	}

	body, err := z.Get("/locales.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data.Locales, nil
}
