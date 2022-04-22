package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// DynamicContentAPI an interface containing all methods associated with zendesk dynamic content
type DynamicContentAPI interface {
	GetDynamicContentItems(ctx context.Context) ([]DynamicContentItem, Page, error)
	CreateDynamicContentItem(ctx context.Context, item DynamicContentItem) (DynamicContentItem, error)
	GetDynamicContentItem(ctx context.Context, id int64) (DynamicContentItem, error)
	UpdateDynamicContentItem(ctx context.Context, id int64, item DynamicContentItem) (DynamicContentItem, error)
	DeleteDynamicContentItem(ctx context.Context, id int64) error
}

// DynamicContentItem is zendesk dynamic content item JSON payload format
//
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
//
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
//
// https://developer.zendesk.com/rest_api/docs/support/dynamic_content#list-items
func (z *Client) GetDynamicContentItems(ctx context.Context) ([]DynamicContentItem, Page, error) {
	var data struct {
		Items []DynamicContentItem `json:"items"`
		Page
	}

	body, err := z.get(ctx, "/dynamic_content/items.json")
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
//
// https://developer.zendesk.com/rest_api/docs/support/dynamic_content#create-item
func (z *Client) CreateDynamicContentItem(ctx context.Context, item DynamicContentItem) (DynamicContentItem, error) {
	var data, result struct {
		Item DynamicContentItem `json:"item"`
	}
	data.Item = item

	body, err := z.post(ctx, "/dynamic_content/items.json", data)
	if err != nil {
		return DynamicContentItem{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return DynamicContentItem{}, err
	}
	return result.Item, nil
}

// GetDynamicContentItem returns the specified dynamic content item.
//
// ref: https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content/#show-item
func (z *Client) GetDynamicContentItem(ctx context.Context, id int64) (DynamicContentItem, error) {
	var result struct {
		Item DynamicContentItem `json:"item"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/dynamic_content/items/%d.json", id))
	if err != nil {
		return DynamicContentItem{}, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return DynamicContentItem{}, err
	}

	return result.Item, nil
}

// UpdateDynamicContentItem updates the specified dynamic content item and returns the updated one
//
// ref: https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content/#update-item
func (z *Client) UpdateDynamicContentItem(ctx context.Context, id int64, item DynamicContentItem) (DynamicContentItem, error) {
	var data, result struct {
		Item DynamicContentItem `json:"item"`
	}
	data.Item = item

	body, err := z.put(ctx, fmt.Sprintf("/dynamic_content/items/%d.json", id), data)
	if err != nil {
		return DynamicContentItem{}, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return DynamicContentItem{}, err
	}

	return result.Item, nil
}

// DeleteDynamicContentItem deletes the specified dynamic content item.
//
// ref: https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content/#delete-item
func (z *Client) DeleteDynamicContentItem(ctx context.Context, id int64) error {
	err := z.delete(ctx, fmt.Sprintf("/dynamic_content/items/%d.json", id))
	if err != nil {
		return err
	}

	return nil
}

// CreateDynamicContentItemVariant creates the specified dynamic content item variant.
//
// https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content_item_variants/#create-variant
func (z *Client) CreateDynamicContentItemVariant(ctx context.Context, itemID int64, variant *DynamicContentVariant) (*DynamicContentVariant, error) {
	var data, result struct {
		Variant *DynamicContentVariant `json:"variant"`
	}
	data.Variant = variant

	body, err := z.post(ctx, "/dynamic_content/items/%d/variants.json", data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Variant, nil
}

// GetDynamicContentItem returns the specified dynamic content item variant.
//
// https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content_item_variants/#show-variant
func (z *Client) GetDynamicContentItemVariant(ctx context.Context, itemID int64, variantID int64) (*DynamicContentVariant, error) {
	var result struct {
		Variant *DynamicContentVariant `json:"variant"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/dynamic_content/items/%d/variants/%d.json", itemID, variantID))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Variant, nil
}

// UpdateDynamicContentItemVariant updates the specified dynamic content item variant and returns the updated one
//
// https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content_item_variants/#update-variant
func (z *Client) UpdateDynamicContentItemVariant(ctx context.Context, itemID int64, variantID int64, variant *DynamicContentVariant) (*DynamicContentVariant, error) {
	var data, result struct {
		Variant *DynamicContentVariant `json:"variant"`
	}
	data.Variant = variant

	body, err := z.put(ctx, fmt.Sprintf("/dynamic_content/items/%d/variants/%d.json", itemID, variantID), data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Variant, nil
}

// DeleteDynamicContentItemVariant deletes the specified dynamic content item variant.
//
// https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content_item_variants/#delete-variant
func (z *Client) DeleteDynamicContentItemVariant(ctx context.Context, itemID int64, variantID int64) error {
	err := z.delete(ctx, fmt.Sprintf("/dynamic_content/items/%d/variants/%d.json", itemID, variantID))
	if err != nil {
		return err
	}

	return nil
}
