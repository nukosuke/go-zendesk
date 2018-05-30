package core

import (
	"encoding/json"
	"github.com/nukosuke/go-zendesk/common"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type TicketFieldSystemFieldOption struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"position"`
	RawName  string `json:"raw_name"`
	URL      string `json:"url"`
	Value    string `json:"value"`
}

type TicketFieldCustomFieldOption struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"position"`
	RawName  string `json:"raw_name"`
	URL      string `json:"url"`
	Value    string `json:"value"`
}

type TicketField struct {
	ID                  int64                          `json:"id,omitempty"`
	URL                 string                         `json:"url,omitempty"`
	Type                string                         `json:"type"`
	Title               string                         `json:"title"`
	RawTitle            string                         `json:"raw_title,omitempty"`
	Description         string                         `json:"description,omitempty"`
	RawDescription      string                         `json:"raw_description,omitempty"`
	Position            int64                          `json:"position,omitempty"`
	Active              bool                           `json:"active,omitempty"`
	Required            bool                           `json:"required,omitempty"`
	CollapsedForAgents  bool                           `json:"collapsed_for_agents,omitempty"`
	RegexpForValidation string                         `json:"regexp_for_validation,omitempty"`
	TitleInPortal       string                         `json:"title_in_portal,omitempty"`
	RawTitleInPortal    string                         `json:"raw_title_in_portal,omitempty"`
	VisibleInPortal     bool                           `json:"visible_in_portal,omitempty"`
	EditableInPortal    bool                           `json:"editable_in_portal,omitempty"`
	Tag                 string                         `json:"tag,omitempty"`
	CreatedAt           *time.Time                     `json:"created_at,omitempty"`
	UpdatedAt           *time.Time                     `json:"updated_at,omitempty"`
	SystemFieldOptions  []TicketFieldSystemFieldOption `json:"system_field_options,omitempty"`
	CustomFieldOptions  []TicketFieldCustomFieldOption `json:"custom_field_options,omitempty"`
	SubTypeID           int64                          `json:"sub_type_id,omitempty"`
	Removable           bool                           `json:"removable,omitempty"`
	AgentDescription    string                         `json:"agent_description,omitempty"`
}

func (s Service) GetTicketFields() ([]TicketField, common.Page, error) {
	type Payload struct {
		TicketFields []TicketField `json:"ticket_fields"`
		Page         common.Page
	}

	req, err := http.NewRequest("GET", s.BaseURL.String()+"/ticket_fields.json", nil)
	if err != nil {
		return []TicketField{}, common.Page{}, err
	}

	req.Header.Set("User-Agent", s.UserAgent)
	req.SetBasicAuth(s.Credential.Email(), s.Credential.Secret())

	resp, err := s.HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return []TicketField{}, common.Page{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []TicketField{}, common.Page{}, err
	}

	var payload Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return []TicketField{}, common.Page{}, err
	}

	return payload.TicketFields, payload.Page, nil
}

func (s Service) PostTicketField(ticketField TicketField) (TicketField, error) {
	type Payload struct {
		TicketField TicketField `json:"ticket_field"`
	}

	payload := Payload{TicketField: ticketField}
	jsonStr, err := json.Marshal(payload)
	if err != nil {
		return TicketField{}, err
	}

	req, err := http.NewRequest("POST", s.BaseURL.String()+"/ticket_fields.json", strings.NewReader(string(jsonStr)))
	if err != nil {
		return TicketField{}, err
	}

	req.Header.Set("User-Agent", s.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(s.Credential.Email(), s.Credential.Secret())

	resp, err := s.HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return TicketField{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return TicketField{}, err
	}

	var resultTicketField TicketField
	err = json.Unmarshal(body, &resultTicketField)
	if err != nil {
		return TicketField{}, err
	}

	return resultTicketField, nil
}
