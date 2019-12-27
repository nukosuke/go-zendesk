package sideload

import "time"

type TicketDates struct {
	AssigneeUpdatedAt    *time.Time `json:"assignee_updated_at"`
	RequesterUpdatedAt   *time.Time `json:"requester_updated_at"`
	StatusUpdatedAt      *time.Time `json:"status_updated_at"`
	InitiallyAssignedAt  *time.Time `json:"initially_assigned_at"`
	AssignedAt           *time.Time `json:"assigned_at"`
	SolvedAt             *time.Time `json:"solved_at"`
	LatestCommentAddedAt *time.Time `json:"latest_comment_added_at"`
}

func IncludeTicketDates(dates *TicketDates) SideLoader {
	return Include("dates", "ticket.dates", dates)
}
