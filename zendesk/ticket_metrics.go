package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// TimeDuration represents a time in business or calendar days
type TimeDuration struct {
	Business int `json:"business"`
	Calendar int `json:"calendar"`
}

type TicketMetric struct {
	AgentWaitTimeInMinutes       TimeDuration `json:"agent_wait_time_in_minutes"`
	AssignedAt                   time.Time    `json:"assigned_at"`
	AssigneeStations             int          `json:"assignee_stations"`
	AssigneeUpdatedAt            time.Time    `json:"assignee_updated_at"`
	CreatedAt                    time.Time    `json:"created_at"`
	CustomStatusUpdatedAt        time.Time    `json:"custom_status_updated_at"`
	FirstResolutionTimeInMinutes TimeDuration `json:"first_resolution_time_in_minutes"`
	FullResolutionTimeInMinutes  TimeDuration `json:"full_resolution_time_in_minutes"`
	GroupStations                int          `json:"group_stations"`
	ID                           int          `json:"id"`
	InitiallyAssignedAt          time.Time    `json:"initially_assigned_at"`
	LatestCommentAddedAt         time.Time    `json:"latest_comment_added_at"`
	OnHoldTimeInMinutes          TimeDuration `json:"on_hold_time_in_minutes"`
	Reopens                      int          `json:"reopens"`
	Replies                      int          `json:"replies"`
	ReplyTimeInMinutes           TimeDuration `json:"reply_time_in_minutes"`
	ReplyTimeInSeconds           struct {
		Calendar int `json:"calendar"`
	} `json:"reply_time_in_seconds"`
	RequesterUpdatedAt         time.Time    `json:"requester_updated_at"`
	RequesterWaitTimeInMinutes TimeDuration `json:"requester_wait_time_in_minutes"`
	SolvedAt                   time.Time    `json:"solved_at"`
	StatusUpdatedAt            time.Time    `json:"status_updated_at"`
	TicketID                   int          `json:"ticket_id"`
	UpdatedAt                  time.Time    `json:"updated_at"`
}

type TicketMetricListOptions struct {
	PageOptions

	SortBy string `url:"sort_by,omitempty"`

	SortOrder string `url:"sort_order,omitempty"`
}

// TicketMetricsAPI is an interface containing all methods for the ticket
// metrics API
type TicketMetricsAPI interface {
	GetTicketMetrics(ctx context.Context, opts ...TicketMetricListOptions) ([]TicketMetric, Page, error)
	GetTicketMetric(ctx context.Context, ticketMetricsID int64) (TicketMetric, error)
	GetTicketMetricByTicket(ctx context.Context, ticketID int64) (TicketMetric, error)
}

// GetTicketMetrics get ticket metrics list with offset based pagination
//
// ref: https://developer.zendesk.com/api-reference/ticketing/tickets/ticket_metrics/#list-ticket-metrics
func (z *Client) GetTicketMetrics(ctx context.Context, opts *TicketMetricListOptions) ([]TicketMetric, Page, error) {
	var data struct {
		TicketMetrics []TicketMetric `json:"ticket_metrics"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &TicketMetricListOptions{}
	}

	u, err := addOptions("/ticket_metrics.json", tmp)
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
	return data.TicketMetrics, data.Page, nil

}

// ref: https://developer.zendesk.com/api-reference/ticketing/tickets/ticket_metrics/#show-ticket-metrics
func (z *Client) GetTicketMetric(ctx context.Context, ticketMetricsID int64) (TicketMetric, error) {
	var result struct {
		TicketMetric TicketMetric `json:"ticket_metric"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/ticket_metrics/%d.json", ticketMetricsID))
	if err != nil {
		return TicketMetric{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return TicketMetric{}, err
	}

	return result.TicketMetric, err
}

// ref: https://developer.zendesk.com/api-reference/ticketing/tickets/ticket_metrics/#show-ticket-metrics
func (z *Client) GetTicketMetricByTicket(ctx context.Context, ticketID int64) (TicketMetric, error) {
	var result struct {
		TicketMetric TicketMetric `json:"ticket_metric"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/tickets/%d/metrics.json", ticketID))
	if err != nil {
		return TicketMetric{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return TicketMetric{}, err
	}

	return result.TicketMetric, err
}
