package zendesk

import (
	"encoding/json"
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

// GetTicketForms 
func (z Client) GetTicketForms() (GetTicketFormsResponse, error) {
	req, err := http.NewRequest("GET", z.BaseURL.String()+"/ticket_forms.json", nil)
	if err != nil {
		return GetTicketFormsResponse{}, err
	}

	req.Header.Set("User-Agent", z.UserAgent)
	req.SetBasicAuth(z.Credential.Email(), z.Credential.Secret())

	resp, err := z.HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return GetTicketFormsResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetTicketFormsResponse{}, err
	}

	var ticketForms GetTicketFormsResponse
	err = json.Unmarshal(body, &ticketForms)
	if err != nil {
		return GetTicketFormsResponse{}, err
	}

	return ticketForms, nil
}
