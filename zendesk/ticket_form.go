package zendesk

import (
	"encoding/json"
	"io/ioutil"
)

// TicketForm is JSON payload struct
type TicketForm struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetTicketFormsResponse is response structure
type GetTicketFormsResponse struct {
	TicketForms []TicketForm `json:"ticket_forms"`
	Page        Page
}

// GetTicketForms fetches ticket forms
func (z Client) GetTicketForms() ([]TicketForm, Page, error) {
	req, err := z.NewGetRequest("/ticket_forms.json")
	if err != nil {
		return []TicketForm{}, Page{}, err
	}

	resp, err := z.HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return []TicketForm{}, Page{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []TicketForm{}, Page{}, err
	}

	var payload GetTicketFormsResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return []TicketForm{}, Page{}, err
	}

	return payload.TicketForms, payload.Page, nil
}
