package zendesk

import (
	"context"
	"encoding/json"
	"time"
)

type Macro struct {
	Actions     []Actions   `json:"actions"`
	Active      bool        `json:"active"`
	CreatedAt   time.Time   `json:"created_at"`
	Description interface{} `json:"description"`
	ID          int64       `json:"id"`
	Position    int         `json:"position"`
	Restriction interface{} `json:"restriction"`
	Title       string      `json:"title"`
	UpdatedAt   time.Time   `json:"updated_at"`
	URL         string      `json:"url"`
}

type Actions struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type MacroListOptions struct {
	Access       string `json:"access"`
	Active       string `json:"active"`
	Category     int    `json:"category"`
	GroupId      int    `json:"group_id"`
	Include      string `json:"include"`
	OnlyViewable bool   `json:"only_viewable"`

	PageOptions

	// SortBy can take "created_at", "updated_at", "usage_1h", "usage_24h",
	// "usage_7d", "usage_30d", "alphabetical"
	SortBy string `url:"sort_by,omitempty"`

	// SortOrder can take "asc" or "desc"
	SortOrder string `url:"sort_order,omitempty"`
}

type MacroAPI interface {
	GetMacros(ctx context.Context, opts *MacroListOptions) ([]Macro, Page, error)
}

// GetMacros get macro list
//
// ref: https://developer.zendesk.com/rest_api/docs/support/macros#list-macros
func (z *Client) GetMacros(ctx context.Context, opts *MacroListOptions) ([]Macro, Page, error) {
	var data struct {
		Macros []Macro `json:"macros"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &MacroListOptions{}
	}

	u, err := addOptions("/macros.json", tmp)
	if err != nil {
		return nil, Page{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return nil, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.Macros, data.Page, nil
}
