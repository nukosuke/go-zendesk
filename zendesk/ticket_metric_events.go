package zendesk

import (
	"context"
	"encoding/json"
	"time"
)

// SLA represents the SLA payload as part of a ticket metric event
type SLA struct {
	Target        int  `json:"target"`
	BusinessHours bool `json:"business_hours"`
	Policy        struct {
		ID          json.RawMessage `json:"id"`
		Title       string          `json:"title"`
		Description string          `json:"description"`
	} `json:"policy"`
}

// TicketMetricEvent represents a ticket metrc event
type TicketMetricEvent struct {
	ID         int64     `json:"id"`
	InstanceID int       `json:"instance_id"`
	Metric     string    `json:"metric"`
	TicketID   int       `json:"ticket_id"`
	Time       time.Time `json:"time"`
	Type       string    `json:"type"`

	// optional fields depending on type
	Status   *TimeDuration `json:"status,omitempty"`
	SLA      *SLA          `json:"sla,omitempty"`
	GroupSLA *SLA          `json:"group_sla,omitempty"`
	Deleted  bool          `json:"deleted,omitempty"`
}

// Timestamp is used to unmarshal a UNIX timestamp into time.Time
type Timestamp struct {
	time.Time
}

func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	var raw int64
	err := json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	p.Time = time.Unix(raw, 0)
	return nil
}

// TicketMetricEventsPage represents the page information of the TicketMetricEvents API response
type TicketMetricEventsPage struct {
	Count       int       `json:"count"`
	EndTime     Timestamp `json:"end_time"`
	EndOfStream bool      `json:"end_of_stream"`
	NextPage    string    `json:"next_page"`
}

// TicketMetricEventsAPI is the interface of the TicketMetricEvents API
type TicketMetricEventsAPI interface {
	GetTicketMetricEvents(ctx context.Context, start time.Time) ([]TicketMetricEvent, TicketMetricEventsPage, error)
}

// TicketMetricEventsOptions represents the options for the  GetTicketMetricEvents method
type TicketMetricEventsOptions struct {
	StartTime int64 `url:"start_time"`
}

// GetTicketMetricEvents
func (z *Client) GetTicketMetricEvents(ctx context.Context, start time.Time) ([]TicketMetricEvent, TicketMetricEventsPage, error) {
	var data struct {
		TicketMetricEventsPage
		TicketMetricEvents []TicketMetricEvent `json:"ticket_metric_events"`
	}

	opts := TicketMetricEventsOptions{
		StartTime: start.Unix(),
	}
	u, err := addOptions("/incremental/ticket_metric_events.json", opts)
	if err != nil {
		return nil, TicketMetricEventsPage{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return nil, TicketMetricEventsPage{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, TicketMetricEventsPage{}, err
	}
	return data.TicketMetricEvents, data.TicketMetricEventsPage, nil
}
