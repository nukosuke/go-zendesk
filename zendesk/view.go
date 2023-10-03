package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type (
	// View is struct for group membership payload
	// https://developer.zendesk.com/api-reference/ticketing/business-rules/views/
	View struct {
		ID          int64     `json:"id,omitempty"`
		Active      bool      `json:"active"`
		Description string    `json:"description"`
		Position    int64     `json:"position"`
		Title       string    `json:"title"`
		CreatedAt   time.Time `json:"created_at,omitempty"`
		UpdatedAt   time.Time `json:"updated_at,omitempty"`

		// Conditions Conditions
		// Execution Execution
		// Restriction Restriction
	}

	ViewCount struct {
		ViewID int64  `json:"view_id"`
		URL    string `json:"url"`
		Value  int64  `json:"value"`
		Pretty string `json:"pretty"`
		Fresh  bool   `json:"fresh"`
	}

	// ViewAPI encapsulates methods on view
	ViewAPI interface {
		GetView(context.Context, int64) (View, error)
		GetViews(context.Context) ([]View, Page, error)
		GetTicketsFromView(context.Context, int64, *TicketListOptions) ([]Ticket, Page, error)
		GetCountTicketsInViews(ctx context.Context, ids []string) ([]ViewCount, error)
	}
)

// GetViews gets all views
// ref: https://developer.zendesk.com/api-reference/ticketing/business-rules/views/#list-views
func (z *Client) GetViews(ctx context.Context) ([]View, Page, error) {
	var result struct {
		Views []View `json:"views"`
		Page
	}

	body, err := z.get(ctx, "/views.json")

	if err != nil {
		return []View{}, Page{}, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return []View{}, Page{}, err
	}

	return result.Views, result.Page, nil
}

// GetView gets a given view
// ref: https://developer.zendesk.com/api-reference/ticketing/business-rules/views/#show-view
func (z *Client) GetView(ctx context.Context, viewID int64) (View, error) {
	var result struct {
		View View `json:"view"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/views/%d.json", viewID))

	if err != nil {
		return View{}, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return View{}, err
	}

	return result.View, nil
}

// GetTicketsFromView gets the tickets of the specified view
// ref: https://developer.zendesk.com/api-reference/ticketing/business-rules/views/#list-tickets-from-a-view
func (z *Client) GetTicketsFromView(ctx context.Context, viewID int64, opts *TicketListOptions,
) ([]Ticket, Page, error) {
	var result struct {
		Tickets []Ticket `json:"tickets"`
		Page
	}
	tmp := opts
	if tmp == nil {
		tmp = &TicketListOptions{}
	}

	path := fmt.Sprintf("/views/%d/tickets.json", viewID)
	url, err := addOptions(path, tmp)
	if err != nil {
		return nil, Page{}, err
	}

	body, err := z.get(ctx, url)

	if err != nil {
		return []Ticket{}, Page{}, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return []Ticket{}, Page{}, err
	}

	return result.Tickets, result.Page, nil
}

// GetCountTicketsInViews count tickets in views using views ids
// ref https://developer.zendesk.com/api-reference/ticketing/business-rules/views/#count-tickets-in-views
func (z *Client) GetCountTicketsInViews(ctx context.Context, ids []string) ([]ViewCount, error) {
	var result struct {
		ViewCounts []ViewCount `json:"view_counts"`
	}
	idsURLParameter := strings.Join(ids, ",")
	body, err := z.get(ctx, fmt.Sprintf("/views/count_many?ids=%s", idsURLParameter))

	if err != nil {
		return []ViewCount{}, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return []ViewCount{}, err
	}
	return result.ViewCounts, nil
}
