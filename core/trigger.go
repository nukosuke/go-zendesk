package core

import (
	"encoding/json"
	"github.com/zenform/go-zendesk/common"
	"io/ioutil"
	"net/http"
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
	Field string `json:"field"`
	Value string `json:"value"`
}

// Trigger is zendesk trigger JSON payload format
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

type GetTriggersResponse struct {
	Triggers []Trigger `json:"triggers"`
	*common.Paginatable
}

// GetTriggers fetch trigger list
func (s *Service) GetTriggers() (GetTriggersResponse, error) {
	req, err := http.NewRequest("GET", s.BaseURL.String()+"/triggers.json", nil)
	if err != nil {
		return GetTriggersResponse{}, err
	}

	req.Header.Set("User-Agent", s.UserAgent)

	//TODO: read email & token from s
	req.SetBasicAuth("user@example.com/token", "apitoken")

	resp, err := s.HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return GetTriggersResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetTriggersResponse{}, err
	}

	var triggers GetTriggersResponse
	err = json.Unmarshal(body, &triggers)
	if err != nil {
		return GetTriggersResponse{}, err
	}

	return triggers, nil
}
