package zendesk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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

	resp, err := z.HTTPClient.Do(req)
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

	var resultTicketForm TicketForm
	err = json.Unmarshal(body, &resultTicketForm)
	if err != nil {
		return TicketForm{}, err
	}

	return resultTicketForm, nil
}
