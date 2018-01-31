package core

import (
	"time"
)

// zendesk trigger condition
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#conditions-reference
type TriggerCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// zendesk trigger action
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#actions
type TriggerAction struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// zendesk trigger JSON payload format
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#json-format
type Trigger struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Active     bool   `json:"active"`
	Position   int64  `json:"position"`
	Conditions struct {
		All []TriggerCondition `json:"all"`
		Any []TriggerCondition `json:"any"`
	} `json:"conditions"`
	Actions     []TriggerAction `json:"actions"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
