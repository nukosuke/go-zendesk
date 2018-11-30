package zendesk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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

// GetTicketForms fetches ticket forms
func (z Client) GetTicketForms() ([]TicketForm, Page, error) {
	var data struct {
		TicketForms []TicketForm `json:"ticket_forms"`
		Page        Page
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
	type Payload struct {
		TicketForm TicketForm `json:"ticket_form"`
	}

	payload := Payload{TicketForm: ticketForm}
	req, err := z.NewPostRequest("/ticket_forms.json", payload)
	if err != nil {
		return TicketForm{}, err
	}

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return TicketForm{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return TicketForm{}, errors.New(http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return TicketForm{}, err
	}

	var result Payload
	err = json.Unmarshal(body, &result)
	if err != nil {
		return TicketForm{}, err
	}

	return result.TicketForm, nil
}
