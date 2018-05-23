package core

import (
	"encoding/json"
	"github.com/nukosuke/go-zendesk/common"
	"io/ioutil"
	"net/http"
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
	ID                  int64                          `json:"id"`
	URL                 string                         `json:"url"`
	Type                string                         `json:"type"`
	Title               string                         `json:"title"`
	RawTitle            string                         `json:"raw_title"`
	Description         string                         `json:"description"`
	RawDescription      string                         `json:"raw_description"`
	Position            int64                          `json:"position"`
	Active              bool                           `json:"active"`
	Required            bool                           `json:"required"`
	CollapsedForAgents  bool                           `json:"collapsed_for_agents"`
	RegexpForValidation string                         `json:"regexp_for_validation"`
	TitleInPortal       string                         `json:"title_in_portal"`
	RawTitleInPortal    string                         `json:"raw_title_in_portal"`
	VisibleInPortal     bool                           `json:"visible_in_portal"`
	EditableInPortal    bool                           `json:"editable_in_portal"`
	Tag                 string                         `json:"tag"`
	CreatedAt           time.Time                      `json:"created_at"`
	UpdatedAt           time.Time                      `json:"updated_at"`
	SystemFieldOptions  []TicketFieldSystemFieldOption `json:"system_field_options"`
	CustomFieldOptions  []TicketFieldCustomFieldOption `json:"custom_field_options"`
	SubTypeID           int64                          `json:"sub_type_id"`
	Removable           bool                           `json:"removable"`
	AgentDescription    string                         `json:"agent_description"`
}

type GetTicketFieldsResponse struct {
	TicketFields []TicketField `json:"ticket_fields"`
	*common.Paginatable
}

func (s Service) GetTicketFields() (GetTicketFieldsResponse, error) {
	req, err := http.NewRequest("GET", s.BaseURL.String()+"/ticket_fields.json", nil)
	if err != nil {
		return GetTicketFieldsResponse{}, err
	}

	req.Header.Set("User-Agent", s.UserAgent)
	req.SetBasicAuth(s.Credential.Email(), s.Credential.Secret())

	resp, err := s.HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return GetTicketFieldsResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetTicketFieldsResponse{}, err
	}

	var ticketFields GetTicketFieldsResponse
	err = json.Unmarshal(body, &ticketFields)
	if err != nil {
		return GetTicketFieldsResponse{}, err
	}

	return ticketFields, nil
}
