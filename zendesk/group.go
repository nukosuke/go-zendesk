package zendesk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Group is struct for group payload
// https://developer.zendesk.com/rest_api/docs/support/groups
type Group struct {
	ID        int64     `json:"id,omitempty"`
	URL       string    `json:"url,omitempty"`
	Name      string    `json:"name"`
	Deleted   bool      `json:"deleted,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetGroups fetches group list
// https://developer.zendesk.com/rest_api/docs/support/groups#list-groups
func (z *Client) GetGroups() ([]Group, Page, error) {
	type Payload struct {
		Groups []Group `json:"groups"`
		Page   Page
	}

	req, err := z.NewGetRequest("/groups.json")
	if err != nil {
		return []Group{}, Page{}, err
	}

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return []Group{}, Page{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []Group{}, Page{}, errors.New(http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Group{}, Page{}, err
	}

	var payload Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return []Group{}, Page{}, err
	}

	return payload.Groups, payload.Page, nil
}

// CreateGroup creates new group
// https://developer.zendesk.com/rest_api/docs/support/groups#create-group
func (z *Client) CreateGroup(group Group) (Group, error) {
	type Payload struct {
		Group Group `json:"group"`
	}

	payload := Payload{Group: group}
	req, err := z.NewPostRequest("/groups.json", payload)
	if err != nil {
		return Group{}, err
	}

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return Group{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return Group{}, errors.New(http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Group{}, err
	}

	var result Payload
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Group{}, err
	}

	return result.Group, nil
}
