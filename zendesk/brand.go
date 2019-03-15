package zendesk

import (
	"encoding/json"
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

// BrandAPI an interface containing all methods associated with zendesk brands
type BrandAPI interface {
	CreateBrand(brand Brand) (Brand, error)
}

// GetBrands fetches brand list
// https://developer.zendesk.com/rest_api/docs/support/brands#list-brands
func (z *Client) GetBrands() ([]Brand, Page, error) {
	var data struct {
		Brands []Brand `json:"brands"`
		Page
	}

	body, err := z.Get("/brands.json")
	if err != nil {
		return []Brand{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []Brand{}, Page{}, err
	}
	return data.Brands, data.Page, nil
}

// CreateBrand creates new brand
// https://developer.zendesk.com/rest_api/docs/support/brands#create-brand
func (z *Client) CreateBrand(brand Brand) (Brand, error) {
	var data, result struct {
		Brand Brand `json:"brand"`
	}
	data.Brand = brand

	body, err := z.Post("/brands.json", data)
	if err != nil {
		return Brand{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Brand{}, err
	}
	return result.Brand, nil
}
