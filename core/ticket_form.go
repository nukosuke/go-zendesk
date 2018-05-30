package core

import (
	"encoding/json"
	"github.com/nukosuke/go-zendesk/common"
	"io/ioutil"
	"net/http"
)

type TicketForm struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetTicketFormsResponse struct {
	TicketForms []TicketForm `json:"ticket_forms"`
	Page        common.Page
}

func (s Service) GetTicketForms() (GetTicketFormsResponse, error) {
	req, err := http.NewRequest("GET", s.BaseURL.String()+"/ticket_forms.json", nil)
	if err != nil {
		return GetTicketFormsResponse{}, err
	}

	req.Header.Set("User-Agent", s.UserAgent)
	req.SetBasicAuth(s.Credential.Email(), s.Credential.Secret())

	resp, err := s.HTTPClient.Do(req)
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
