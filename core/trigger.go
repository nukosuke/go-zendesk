package core

import (
	"time"
)

type TriggerCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type TriggerAction struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

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
