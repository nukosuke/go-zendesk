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
