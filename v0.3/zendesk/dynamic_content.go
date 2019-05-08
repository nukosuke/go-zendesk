package zendesk

import (
	"encoding/json"
	"time"
)

// DynamicContentItem is zendesk dynamic content item JSON payload format
// https://developer.zendesk.com/rest_api/docs/support/users
type DynamicContentItem struct {
	ID              int64                   `json:"id,omitempty"`
	URL             string                  `json:"url,omitempty"`
	Name            string                  `json:"name"`
	Placeholder     string                  `json:"placeholder,omitempty"`
	DefaultLocaleID int64                   `json:"default_locale_id"`
	Outdated        bool                    `json:"outdated,omitempty"`
	Variants        []DynamicContentVariant `json:"variants"`
	CreatedAt       time.Time               `json:"created_at,omitempty"`
	UpdatedAt       time.Time               `json:"updated_at,omitempty"`
}

// DynamicContentVariant is zendesk dynamic content variant JSON payload format
// https://developer.zendesk.com/rest_api/docs/support/dynamic_content#json-format-for-variants
type DynamicContentVariant struct {
	ID        int64     `json:"id,omitempty"`
	URL       string    `json:"url,omitempty"`
	Content   string    `json:"content"`
	LocaleID  int64     `json:"locale_id"`
	Outdated  bool      `json:"outdated,omitempty"`
	Active    bool      `json:"active,omitempty"`
	Default   bool      `json:"default,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetDynamicContentItems fetches dynamic content item list
// https://developer.zendesk.com/rest_api/docs/support/dynamic_content#list-items
func (z *Client) GetDynamicContentItems() ([]DynamicContentItem, Page, error) {
	var data struct {
		Items []DynamicContentItem `json:"items"`
		Page
	}

	body, err := z.get("/dynamic_content/items.json")
	if err != nil {
		return []DynamicContentItem{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []DynamicContentItem{}, Page{}, err
	}
	return data.Items, data.Page, nil
}

// CreateDynamicContentItem creates new dynamic content item
// https://developer.zendesk.com/rest_api/docs/support/dynamic_content#create-item
func (z *Client) CreateDynamicContentItem(item DynamicContentItem) (DynamicContentItem, error) {
	var data, result struct {
		Item DynamicContentItem `json:"item"`
	}
	data.Item = item

	body, err := z.post("/groups.json", data)
	if err != nil {
		return DynamicContentItem{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return DynamicContentItem{}, err
	}
	return result.Item, nil
}
