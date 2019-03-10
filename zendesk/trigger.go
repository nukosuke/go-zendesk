package zendesk

import (
	"encoding/json"
	"time"
)

// TriggerCondition zendesk trigger condition
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#conditions-reference
type TriggerCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// TriggerAction is zendesk trigger action
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#actions
type TriggerAction struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Trigger is zendesk trigger JSON payload format
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#json-format
type Trigger struct {
	ID         int64  `json:"id,omitempty"`
	Title      string `json:"title"`
	Active     bool   `json:"active,omitempty"`
	Position   int64  `json:"position,omitempty"`
	Conditions struct {
		All []TriggerCondition `json:"all"`
		Any []TriggerCondition `json:"any"`
	} `json:"conditions"`
	Actions     []TriggerAction `json:"actions"`
	Description string          `json:"description,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
}

// TriggerAPI an interface containing all trigger related methods
type TriggerAPI interface {
	GetTriggers() ([]Trigger, Page, error)
	CreateTrigger(trigger Trigger) (Trigger, error)
}

// GetTriggers fetch trigger list
func (z *Client) GetTriggers() ([]Trigger, Page, error) {
	var data struct {
		Triggers []Trigger `json:"triggers"`
		Page
	}

	body, err := z.Get("/triggers.json")
	if err != nil {
		return []Trigger{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []Trigger{}, Page{}, err
	}
	return data.Triggers, data.Page, nil
}

// CreateTrigger creates new trigger
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#create-trigger
func (z Client) CreateTrigger(trigger Trigger) (Trigger, error) {
	var data, result struct {
		Trigger Trigger `json:"trigger"`
	}

	body, err := z.Post("/triggers.json", data)
	if err != nil {
		return Trigger{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Trigger{}, err
	}
	return result.Trigger, nil
}
