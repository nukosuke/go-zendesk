package zendesk

import (
	"encoding/json"
	"time"
)

// TicketFieldSystemFieldOption is struct for value of `system_field_options`
type TicketFieldSystemFieldOption struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"position"`
	RawName  string `json:"raw_name"`
	URL      string `json:"url"`
	Value    string `json:"value"`
}

// TicketField is struct for ticket_field payload
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
	CustomFieldOptions  []CustomFieldOption            `json:"custom_field_options,omitempty"`
	SubTypeID           int64                          `json:"sub_type_id,omitempty"`
	Removable           bool                           `json:"removable,omitempty"`
	AgentDescription    string                         `json:"agent_description,omitempty"`
}

// GetTicketFields fetches ticket field list
// ref: https://developer.zendesk.com/rest_api/docs/core/ticket_fields#list-ticket-fields
func (z Client) GetTicketFields() ([]TicketField, Page, error) {
	var data struct {
		TicketFields []TicketField `json:"ticket_fields"`
		Page         Page
	}

	body, err := z.Get("/ticket_fields.json")
	if err != nil {
		return []TicketField{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []TicketField{}, Page{}, err
	}
	return data.TicketFields, data.Page, nil
}

// CreateTicketField creates new ticket field
// ref: https://developer.zendesk.com/rest_api/docs/core/ticket_fields#create-ticket-field
func (z Client) CreateTicketField(ticketField TicketField) (TicketField, error) {
	var data, result struct {
		TicketField TicketField `json:"ticket_field"`
	}

	body, err := z.Post("/ticket_fields.json", data)
	if err != nil {
		return TicketField{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return TicketField{}, err
	}
	return result.TicketField, nil
}
