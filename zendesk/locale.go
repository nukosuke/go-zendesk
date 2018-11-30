package zendesk

import (
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
