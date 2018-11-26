package zendesk

import (
	"encoding/json"
	"errors"
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

// GetTriggersResponse is response structure of triggers list
type GetTriggersResponse struct {
	Triggers []Trigger `json:"triggers"`
	Page     Page
}

// GetTriggers fetch trigger list
func (z *Client) GetTriggers() ([]Trigger, Page, error) {
	req, err := z.NewGetRequest("/triggers.json")
	if err != nil {
		return []Trigger{}, Page{}, err
	}

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return []Trigger{}, Page{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Trigger{}, Page{}, err
	}

	var payload GetTriggersResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return []Trigger{}, Page{}, err
	}

	return payload.Triggers, payload.Page, nil
}

// CreateTrigger creates new trigger
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#create-trigger
func (z Client) CreateTrigger(trigger Trigger) (Trigger, error) {
	type Payload struct {
		Trigger Trigger `json:"trigger"`
	}

	payload := Payload{Trigger: trigger}
	req, err := z.NewPostRequest("/triggers.json", payload)
	if err != nil {
		return Trigger{}, err
	}

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return Trigger{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return Trigger{}, errors.New(http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Trigger{}, err
	}

	var result Payload
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Trigger{}, err
	}

	return result.Trigger, nil
}
