package zendesk

import (
	"time"
)

// DynamicContentItem is zendesk dynamic content item JSON payload format
// https://developer.zendesk.com/rest_api/docs/support/users
type DynamicContentItem struct {
	ID              int64                   `json:"id,omitempty"`
	URL             string                  `json:"url,omitempty"`
	Name            string                  `json:"name"`
	Placeholder     string                  `json:"placeholder,omitempty"`
	DefaultLocaleID int64                   `json:"default_locale_id"`
	Outdated        bool                    `json:"outdated,omitempty"`
	Variants        []DynamicContentVariant `json:"variants"`
	CreatedAt       time.Time               `json:"created_at,omitempty"`
	UpdatedAt       time.Time               `json:"updated_at,omitempty"`
}

// DynamicContentVariant is zendesk dynamic content variant JSON payload format
// https://developer.zendesk.com/rest_api/docs/support/dynamic_content#json-format-for-variants
type DynamicContentVariant struct {
	ID        int64     `json:"id,omitempty"`
	URL       string    `json:"url,omitempty"`
	Content   string    `json:"content"`
	LocaleID  int64     `json:"locale_id"`
	Outdated  bool      `json:"outdated,omitempty"`
	Active    bool      `json:"active,omitempty"`
	Default   bool      `json:"default,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
