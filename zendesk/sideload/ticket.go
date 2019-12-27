package sideload

import "time"

//"assignee_updated_at": null,
//"requester_updated_at": "2019-06-06T10:02:04Z",
//"status_updated_at": "2019-06-06T10:02:04Z",
//"initially_assigned_at": null,
//"assigned_at": null,
//"solved_at": null,
//"latest_comment_added_at": "2019-06-06T10:02:04Z"
type TicketDates struct {
	AssigneeUpdatedAt *time.Time `json:"assignee_updated_at"`
	RequesterUpdatedAt *time.Time `json:"requester_updated_at"`
	StatusUpdatedAt *time.Time `json:"status_updated_at"`
	InitiallyAssignedAt *time.Time `json:"initially_assigned_at"`
	AssignedAt *time.Time `json:"assigned_at"`
	SolvedAt *time.Time `json:"solved_at"`
	LatestCommentAddedAt *time.Time `json:"latest_comment_added_at"`
}
