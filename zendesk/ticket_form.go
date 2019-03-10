package zendesk

import (
	"encoding/json"
)

// TicketForm is JSON payload struct
type TicketForm struct {
	ID                 int64   `json:"id,omitempty"`
	Name               string  `json:"name"`
	RawName            string  `json:"raw_name,omitempty"`
	DisplayName        string  `json:"display_name,omitempty"`
	RawDisplayName     string  `json:"raw_display_name,omitempty"`
	Position           int64   `json:"position"`
	Active             bool    `json:"active,omitempty"`
	EndUserVisible     bool    `json:"end_user_visible,omitempty"`
	Default            bool    `json:"default,omitempty"`
	TicketFieldIDs     []int64 `json:"ticket_field_ids,omitempty"`
	InAllBrands        bool    `json:"in_all_brands,omitempty"`
	RestrictedBrandIDs []int64 `json:"restricted_brand_ids,omitempty"`
}

// TicketFormAPI an interface containing all ticket form related methods
type TicketFormAPI interface {
	GetTicketForms() ([]TicketForm, Page, error)
	CreateTicketForm(ticketForm TicketForm) (TicketForm, error)
}

// GetTicketForms fetches ticket forms
func (z Client) GetTicketForms() ([]TicketForm, Page, error) {
	var data struct {
		TicketForms []TicketForm `json:"ticket_forms"`
		Page
	}

	body, err := z.Get("/ticket_forms.json")
	if err != nil {
		return []TicketForm{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []TicketForm{}, Page{}, err
	}
	return data.TicketForms, data.Page, nil
}

// CreateTicketForm creates new ticket form
// ref: https://developer.zendesk.com/rest_api/docs/core/ticket_forms#create-ticket-forms
func (z Client) CreateTicketForm(ticketForm TicketForm) (TicketForm, error) {
	var data, result struct {
		TicketForm TicketForm `json:"ticket_form"`
	}

	body, err := z.Post("/ticket_forms.json", data)
	if err != nil {
		return TicketForm{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return TicketForm{}, err
	}
	return result.TicketForm, nil
}
