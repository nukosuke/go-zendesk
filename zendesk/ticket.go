package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CustomField struct {
	ID int64 `json:"id"`
	// Valid types are string or []string.
	Value interface{} `json:"value"`
}

// UnmarshalJSON Custom Unmarshal function required because a custom field's value can be
// a string or array of strings.
func (cf *CustomField) UnmarshalJSON(data []byte) error {
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	cf.ID = int64(temp["id"].(float64))

	switch v := temp["value"].(type) {
	case string, nil, bool:
		cf.Value = v
	case []interface{}:
		var list []string

		for _, v := range temp["value"].([]interface{}) {
			if s, ok := v.(string); ok {
				list = append(list, s)
			} else {
				return fmt.Errorf("%T is an invalid type for custom field value", v)
			}
		}

		cf.Value = list
	default:
		return fmt.Errorf("%T is an invalid type for custom field value", v)
	}

	return nil
}

type Ticket struct {
	ID              int64         `json:"id,omitempty"`
	URL             string        `json:"url,omitempty"`
	ExternalID      string        `json:"external_id,omitempty"`
	Type            string        `json:"type,omitempty"`
	Subject         string        `json:"subject,omitempty"`
	RawSubject      string        `json:"raw_subject,omitempty"`
	Description     string        `json:"description,omitempty"`
	Priority        string        `json:"priority,omitempty"`
	Status          string        `json:"status,omitempty"`
	CustomStatusID  int64         `json:"custom_status_id,omitempty"`
	Recipient       string        `json:"recipient,omitempty"`
	RequesterID     int64         `json:"requester_id,omitempty"`
	SubmitterID     int64         `json:"submitter_id,omitempty"`
	AssigneeID      int64         `json:"assignee_id,omitempty"`
	OrganizationID  int64         `json:"organization_id,omitempty"`
	GroupID         int64         `json:"group_id,omitempty"`
	CollaboratorIDs []int64       `json:"collaborator_ids,omitempty"`
	FollowerIDs     []int64       `json:"follower_ids,omitempty"`
	EmailCCIDs      []int64       `json:"email_cc_ids,omitempty"`
	ForumTopicID    int64         `json:"forum_topic_id,omitempty"`
	ProblemID       int64         `json:"problem_id,omitempty"`
	HasIncidents    bool          `json:"has_incidents,omitempty"`
	DueAt           *time.Time    `json:"due_at,omitempty"`
	Tags            []string      `json:"tags,omitempty"`
	CustomFields    []CustomField `json:"custom_fields,omitempty"`

	Via *Via `json:"via,omitempty"`

	SatisfactionRating *struct {
		ID      int64  `json:"id"`
		Score   string `json:"score"`
		Comment string `json:"comment"`
	} `json:"satisfaction_rating,omitempty"`

	SharingAgreementIDs []int64    `json:"sharing_agreement_ids,omitempty"`
	FollowupIDs         []int64    `json:"followup_ids,omitempty"`
	ViaFollowupSourceID int64      `json:"via_followup_source_id,omitempty"`
	MacroIDs            []int64    `json:"macro_ids,omitempty"`
	TicketFormID        int64      `json:"ticket_form_id,omitempty"`
	BrandID             int64      `json:"brand_id,omitempty"`
	AllowChannelback    bool       `json:"allow_channelback,omitempty"`
	AllowAttachments    bool       `json:"allow_attachments,omitempty"`
	IsPublic            bool       `json:"is_public,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`

	// Collaborators is POST only
	Collaborators *Collaborators `json:"collaborators,omitempty"`

	// Comment is POST only and required
	Comment *TicketComment `json:"comment,omitempty"`

	// Requester is POST only and can be used to create a ticket for a nonexistent requester
	Requester *Requester `json:"requester,omitempty"`

	// safe update fields
	// https://developer.zendesk.com/documentation/ticketing/managing-tickets/creating-and-updating-tickets/#protecting-against-ticket-update-collisions
	UpdatedStamp *time.Time `json:"updated_stamp,omitempty"`
	SafeUpdate   bool       `json:"safe_update,omitempty"`

	// TODO: TicketAudit (POST only) #126
}

// Requester is the struct that can be passed to create a new requester on ticket creation
// https://develop.zendesk.com/hc/en-us/articles/360059146153#creating-a-ticket-with-a-new-requester
type Requester struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	LocaleID string `json:"locale_id,omitempty"`
}

// Via is information about source of Ticket or TicketComment
type Via struct {
	Channel string `json:"channel"`
	Source  struct {
		From map[string]interface{} `json:"from"`
		To   map[string]interface{} `json:"to"`
		Rel  string                 `json:"rel"`
	} `json:"source"`
}

// TicketListOptions struct is used to specify options for listing tickets in OBP (Offset Based Pagination).
// It embeds the PageOptions struct for pagination and provides options for sorting the result;
// SortBy specifies the field to sort by, and SortOrder specifies the order (either 'asc' or 'desc').
type TicketListOptions struct {
	PageOptions

	// SortBy can take "assignee", "assignee.name", "created_at", "group", "id",
	// "locale", "requester", "requester.name", "status", "subject", "updated_at"
	SortBy string `url:"sort_by,omitempty"`

	// SortOrder can take "asc" or "desc"
	SortOrder string `url:"sort_order,omitempty"`
}

// TicketListCBPOptions struct is used to specify options for listing tickets in CBP (Cursor Based Pagination).
// It embeds the CursorPagination struct for pagination and provides an option Sort for sorting the result.
type TicketListCBPOptions struct {
	CursorPagination
	Sort string `url:"sort,omitempty"`
}

// TicketListCBPResult struct represents the result of a ticket list operation in CBP. It includes an array of Ticket objects, and Meta that holds pagination metadata.
type TicketListCBPResult struct {
	Tickets []Ticket             `json:"tickets"`
	Meta    CursorPaginationMeta `json:"meta"`
}

// PaginationOptions struct represents general pagination options.
// PageSize specifies the number of items per page, IsCBP indicates if it's cursor-based pagination,
// SortBy and SortOrder describe how to sort the items in Offset Based Pagination, and Sort describes how to sort items in Cursor Based Pagination.
type PaginationOptions struct {
	PageSize int  //default is 100
	IsCBP    bool //default is true

	SortBy string
	// SortOrder can take "asc" or "desc"
	SortOrder string
	Sort      string
}

// NewPaginationOptions() returns a pointer to a new PaginationOptions struct with default values (PageSize is 100, IsCBP is true).
func NewPaginationOptions() *PaginationOptions {
	return &PaginationOptions{
		PageSize: 100,
		IsCBP:    true,
	}
}

// TicketIterator struct provides a convenient way to iterate over pages of tickets in either OBP or CBP.
// It holds state for iteration, including the current page size, a flag indicating more pages, pagination type (OBP or CBP), and sorting options.
type TicketIterator struct {
	// generic fields
	pageSize int
	hasMore  bool
	isCBP    bool

	// OBP fields
	sortBy string
	// SortOrder can take "asc" or "desc"
	sortOrder string
	pageIndex int

	// CBP fields
	sort      string
	pageAfter string

	// common fields
	client *Client
	ctx    context.Context
}

// HasMore() returns a boolean indicating whether more pages are available for iteration.
func (i *TicketIterator) HasMore() bool {
	return i.hasMore
}

// GetNext() retrieves the next batch of tickets according to the current pagination and sorting options.
// It updates the state of the iterator for subsequent calls.
// In case of an error, it sets hasMore to false and returns an error.
func (i *TicketIterator) GetNext() ([]Ticket, error) {
	if i.isCBP {
		cbpOps := &TicketListCBPOptions{
			CursorPagination: CursorPagination{
				PageSize:  i.pageSize,
				PageAfter: i.pageAfter,
			},
		}
		if i.sort != "" {
			cbpOps.Sort = i.sort
		}
		ticketListCBPResult, err := i.client.GetTicketsCBP(i.ctx, cbpOps)
		if err != nil {
			i.hasMore = false
			return nil, err
		}
		i.hasMore = ticketListCBPResult.Meta.HasMore
		i.pageAfter = ticketListCBPResult.Meta.AfterCursor
		return ticketListCBPResult.Tickets, nil
	} else {
		obpOps := &TicketListOptions{
			PageOptions: PageOptions{
				PerPage: i.pageSize,
				Page:    i.pageIndex,
			},
		}
		if i.sortBy != "" {
			obpOps.SortBy = i.sortBy
		}
		if i.sortOrder != "" {
			obpOps.SortOrder = i.sortOrder
		}
		tickets, page, err := i.client.GetTickets(i.ctx, obpOps)
		if err != nil {
			i.hasMore = false
			return nil, err
		}
		i.hasMore = page.HasNext()
		i.pageIndex++
		return tickets, nil
	}
}

// TicketAPI an interface containing all ticket related methods
type TicketAPI interface {
	GetTicketsEx(ctx context.Context, opts *PaginationOptions) *TicketIterator
	GetTickets(ctx context.Context, opts *TicketListOptions) ([]Ticket, Page, error)
	GetTicketsCBP(ctx context.Context, opts *TicketListCBPOptions) (*TicketListCBPResult, error)
	GetOrganizationTickets(ctx context.Context, organizationID int64, ops *TicketListOptions) ([]Ticket, Page, error)
	GetTicket(ctx context.Context, id int64) (Ticket, error)
	GetMultipleTickets(ctx context.Context, ticketIDs []int64) ([]Ticket, error)
	CreateTicket(ctx context.Context, ticket Ticket) (Ticket, error)
	UpdateTicket(ctx context.Context, ticketID int64, ticket Ticket) (Ticket, error)
	DeleteTicket(ctx context.Context, ticketID int64) error
}

// GetTicketsEx returns a TicketIterator to iterate over tickets
//
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#list-tickets
func (z *Client) GetTicketsEx(ctx context.Context, opts *PaginationOptions) *TicketIterator {
	return &TicketIterator{
		pageSize:  opts.PageSize,
		hasMore:   true,
		isCBP:     opts.IsCBP,
		sort:      opts.Sort,
		pageAfter: "",
		sortOrder: opts.SortOrder,
		sortBy:    opts.SortBy,
		pageIndex: 1,
		client:    z,
		ctx:       ctx,
	}
}

// GetTickets get ticket list with offset based pagination
//
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#list-tickets
func (z *Client) GetTickets(ctx context.Context, opts *TicketListOptions) ([]Ticket, Page, error) {
	var data struct {
		Tickets []Ticket `json:"tickets"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &TicketListOptions{}
	}

	u, err := addOptions("/tickets.json", tmp)
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
	return data.Tickets, data.Page, nil
}

// GetTicketsCBP get ticket list with cursor based pagination
//
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#list-tickets
func (z *Client) GetTicketsCBP(ctx context.Context, opts *TicketListCBPOptions) (*TicketListCBPResult, error) {
	var data TicketListCBPResult

	tmp := opts
	if tmp == nil {
		tmp = &TicketListCBPOptions{}
	}

	u, err := addOptions("/tickets.json", tmp)
	if err != nil {
		return nil, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrganizationTickets get organization ticket list
//
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#list-tickets
func (z *Client) GetOrganizationTickets(
	ctx context.Context, organizationID int64, opts *TicketListOptions,
) ([]Ticket, Page, error) {
	var data struct {
		Tickets []Ticket `json:"tickets"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &TicketListOptions{}
	}

	path := fmt.Sprintf("/organizations/%d/tickets.json", organizationID)
	u, err := addOptions(path, tmp)
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
	return data.Tickets, data.Page, nil
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

// GetMultipleTickets gets multiple specified tickets
//
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#show-multiple-tickets
func (z *Client) GetMultipleTickets(ctx context.Context, ticketIDs []int64) ([]Ticket, error) {
	var result struct {
		Tickets []Ticket `json:"tickets"`
	}

	var req struct {
		IDs string `url:"ids,omitempty"`
	}
	idStrs := make([]string, len(ticketIDs))
	for i := 0; i < len(ticketIDs); i++ {
		idStrs[i] = strconv.FormatInt(ticketIDs[i], 10)
	}
	req.IDs = strings.Join(idStrs, ",")

	u, err := addOptions("/tickets/show_many.json", req)
	if err != nil {
		return nil, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Tickets, nil
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

// UpdateTicket update an existing ticket
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#update-ticket
func (z *Client) UpdateTicket(ctx context.Context, ticketID int64, ticket Ticket) (Ticket, error) {
	var data, result struct {
		Ticket Ticket `json:"ticket"`
	}
	data.Ticket = ticket

	path := fmt.Sprintf("/tickets/%d.json", ticketID)
	body, err := z.put(ctx, path, data)
	if err != nil {
		return Ticket{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Ticket{}, err
	}

	return result.Ticket, nil
}

// DeleteTicket deletes the specified ticket
// ref: https://developer.zendesk.com/rest_api/docs/support/tickets#delete-ticket
func (z *Client) DeleteTicket(ctx context.Context, ticketID int64) error {
	err := z.delete(ctx, fmt.Sprintf("/tickets/%d.json", ticketID))

	if err != nil {
		return err
	}

	return nil
}
