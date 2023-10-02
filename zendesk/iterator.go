package zendesk

import (
	"context"
)

// PaginationOptions struct represents general pagination options.
// PageSize specifies the number of items per page, IsCBP indicates if it's cursor-based pagination,
// SortBy and SortOrder describe how to sort the items in Offset Based Pagination, and Sort describes how to sort items in Cursor Based Pagination.
type PaginationOptions struct {
	CommonOptions
	PageSize int  //default is 100
	IsCBP    bool //default is true
}

// NewPaginationOptions() returns a pointer to a new PaginationOptions struct with default values (PageSize is 100, IsCBP is true).
func NewPaginationOptions() *PaginationOptions {
	return &PaginationOptions{
		PageSize: 100,
		IsCBP:    true,
	}
}

type CommonOptions struct {
	Active        bool     `url:"active,omitempty"`
	Role          string   `url:"role,omitempty"`
	Roles         []string `url:"role[],omitempty"`
	PermissionSet int64    `url:"permission_set,omitempty"`

	// SortBy can take "assignee", "assignee.name", "created_at", "group", "id",
	// "locale", "requester", "requester.name", "status", "subject", "updated_at"
	SortBy string `url:"sort_by,omitempty"`

	// SortOrder can take "asc" or "desc"
	SortOrder string `url:"sort_order,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Id        int64
}

// CBPOptions struct is used to specify options for listing objects in CBP (Cursor Based Pagination).
// It embeds the CursorPagination struct for pagination and provides an option Sort for sorting the result.
type CBPOptions struct {
	CursorPagination
	CommonOptions
}

// OBPOptions struct is used to specify options for listing objects in OBP (Offset Based Pagination).
// It embeds the PageOptions struct for pagination and provides options for sorting the result;
// SortBy specifies the field to sort by, and SortOrder specifies the order (either 'asc' or 'desc').
type OBPOptions struct {
	PageOptions
	CommonOptions
}

// ObpFunc defines the signature of the function used to list objects in OBP.
type ObpFunc[T any] func(ctx context.Context, opts *OBPOptions) ([]T, Page, error)

// CbpFunc defines the signature of the function used to list objects in CBP.
type CbpFunc[T any] func(ctx context.Context, opts *CBPOptions) ([]T, CursorPaginationMeta, error)

// terator struct provides a convenient and genric way to iterate over pages of objects in either OBP or CBP.
// It holds state for iteration, including the current page size, a flag indicating more pages, pagination type (OBP or CBP), and sorting options.
type Iterator[T any] struct {
	CommonOptions
	// generic fields
	pageSize int
	hasMore  bool
	isCBP    bool

	// OBP fields
	pageIndex int

	// CBP fields
	pageAfter string

	// common fields
	ctx     context.Context
	obpFunc ObpFunc[T]
	cbpFunc CbpFunc[T]
}

// HasMore() returns a boolean indicating whether more pages are available for iteration.
func (i *Iterator[T]) HasMore() bool {
	return i.hasMore
}

// GetNext() retrieves the next batch of objects according to the current pagination and sorting options.
// It updates the state of the iterator for subsequent calls.
// In case of an error, it sets hasMore to false and returns an error.
func (i *Iterator[T]) GetNext() ([]T, error) {
	if !i.isCBP {
		obpOps := &OBPOptions{
			PageOptions: PageOptions{
				PerPage: i.pageSize,
				Page:    i.pageIndex,
			},
			CommonOptions: i.CommonOptions,
		}
		results, page, err := i.obpFunc(i.ctx, obpOps)
		if err != nil {
			i.hasMore = false
			return nil, err
		}
		i.hasMore = page.HasNext()
		i.pageIndex++
		return results, nil
	}

	cbpOps := &CBPOptions{
		CursorPagination: CursorPagination{
			PageSize:  i.pageSize,
			PageAfter: i.pageAfter,
		},
		CommonOptions: i.CommonOptions,
	}
	results, meta, err := i.cbpFunc(i.ctx, cbpOps)
	if err != nil {
		i.hasMore = false
		return nil, err
	}
	i.hasMore = meta.HasMore
	i.pageAfter = meta.AfterCursor
	return results, nil
}
