package zendesk

import (
	"time"
)

// Brand is struct for brand payload
// https://developer.zendesk.com/rest_api/docs/support/brands
type Brand struct {
	ID                int64      `json:"id,omitempty"`
	URL               string     `json:"url,omitempty"`
	Name              string     `json:"name"`
	BrandURL          string     `json:"brand_url,omitempty"`
	HasHelpCenter     bool       `json:"has_help_center,omitempty"`
	HelpCenterState   string     `json:"help_center_state,omitempty"`
	Active            bool       `json:"active,omitempty"`
	Default           bool       `json:"default,omitempty"`
	Logo              Attachment `json:"logo,omitempty"`
	TicketFieldIDs    []int64    `json:"ticket_field_ids,omitempty"`
	Subdomain         string     `json:"subdomain"`
	HostMapping       string     `json:"host_mapping,omitempty"`
	SignatureTemplate string     `json:"signature_template"`
	CreatedAt         time.Time  `json:"created_at,omitempty"`
	UpdatedAt         time.Time  `json:"updated_at,omitempty"`
}
