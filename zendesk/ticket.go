package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Ticket struct {
	ID              int64     `json:"id,omitempty"`
	URL             string    `json:"url,omitempty"`
	ExternalID      string    `json:"external_id,omitempty"`
	Type            string    `json:"type,omitempty"`
	Subject         string    `json:"subject,omitempty"`
	RawSubject      string    `json:"raw_subject,omitempty"`
	Description     string    `json:"description,omitempty"`
	Priority        string    `json:"priority,omitempty"`
	Status          string    `json:"status,omitempty"`
	Recipient       string    `json:"recipient,omitempty"`
	RequesterID     int64     `json:"requester_id,omitempty"`
	SubmitterID     int64     `json:"submitter_id,omitempty"`
	AssigneeID      int64     `json:"assignee_id,omitempty"`
	OrganizationID  int64     `json:"organization_id,omitempty"`
	GroupID         int64     `json:"group_id,omitempty"`
	CollaboratorIDs []int64   `json:"collaborator_ids,omitempty"`
	FollowerIDs     []int64   `json:"follower_ids,omitempty"`
	EmailCCIDs      []int64   `json:"email_cc_ids,omitempty"`
	ForumTopicID    int64     `json:"forum_topic_id,omitempty"`
	ProblemID       int64     `json:"problem_id,omitempty"`
	HasIncidents    bool      `json:"has_incidents,omitempty"`
	DueAt           time.Time `json:"due_at,omitempty"`
	Tags            []string  `json:"tags,omitempty"`

	// TODO: Via          #123
	// TODO: CustomFields #122

	SatisfactionRating struct {
		ID      int64  `json:"id"`
		Score   string `json:"score"`
		Comment string `json:"comment"`
	} `json:"satisfaction_rating,omitempty"`

	SharingAgreementIDs []int64   `json:"sharing_agreement_ids,omitempty"`
	FollowupIDs         []int64   `json:"followup_ids,omitempty"`
	ViaFollowupSourceID int64     `json:"via_followup_source_id,omitempty"`
	MacroIDs            []int64   `json:"macro_ids,omitempty"`
	TicketFormID        int64     `json:"ticket_form_id,omitempty"`
	BrandID             int64     `json:"brand_id,omitempty"`
	AllowChannelback    bool      `json:"allow_channelback,omitempty"`
	AllowAttachments    bool      `json:"allow_attachments,omitempty"`
	IsPublic            bool      `json:"is_public,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`

	// Collaborators is POST only
	Collaborators Collaborators `json:"collaborators,omitempty"`

	// Comment is POST only and required
	Comment TicketComment `json:"comment,omitempty"`

	// TODO: TicketAudit (POST only) #126
}

// TODO: This is temporary struct for ticket support. #125
//       Need to make it into correct TicketComment structure.
//       https://developer.zendesk.com/rest_api/docs/support/ticket_comments
type TicketComment struct {
	Body string `json:"body"`
}

// TicketAPI an interface containing all ticket related methods
type TicketAPI interface {
	GetTicket(ctx context.Context, id int64) (Ticket, error)
}

// GetTicket gets a specified ticket
//
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#show-ticket
func (z *Client) GetTicket(ctx context.Context, ticketID int64) (Ticket, error) {
	var result struct {
		Ticket Ticket `json:"ticket"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/tickets/%d.json", ticketID))
	if err != nil {
		return Ticket{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Ticket{}, err
	}

	return result.Ticket, err
}

// CreateTicket create a new ticket
//
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#create-ticket
func (z *Client) CreateTicket(ctx context.Context, ticket Ticket) (Ticket, error) {
	var data, result struct {
		Ticket Ticket `json:"ticket"`
	}
	data.Ticket = ticket

	body, err := z.post(ctx, "/tickets.json", data)
	if err != nil {
		return Ticket{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Ticket{}, err
	}
	return result.Ticket, nil
}
