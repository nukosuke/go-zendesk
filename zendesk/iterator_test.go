package zendesk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for NewPaginationOptions function
func TestNewPaginationOptions(t *testing.T) {
	opts := NewPaginationOptions()

	assert.Equal(t, 100, opts.PageSize)
	assert.Equal(t, true, opts.IsCBP)
}

// Test for HasMore function
func TestHasMore(t *testing.T) {
	iter := &Iterator[int]{
		hasMore: true,
	}

	result := iter.HasMore()

	assert.Equal(t, true, result)
}

// Mock functions for GetNext testing
func mockObpFunc(ctx context.Context, opts *OBPOptions) ([]int, Page, error) {
	nextPage := "2"
	return []int{1, 2, 3}, Page{NextPage: &nextPage, Count: 3}, nil
}

func mockCbpFunc(ctx context.Context, opts *CBPOptions) ([]int, CursorPaginationMeta, error) {
	return []int{1, 2, 3}, CursorPaginationMeta{HasMore: true, AfterCursor: "3"}, nil
}

// Test for GetNext function
func TestGetNext(t *testing.T) {
	ctx := context.Background()

	iter := &Iterator[int]{
		pageSize:  2,
		hasMore:   true,
		isCBP:     false,
		pageIndex: 1,
		ctx:       ctx,
		obpFunc:   mockObpFunc,
		cbpFunc:   mockCbpFunc,
	}

	results, err := iter.GetNext()

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, results)
	assert.Equal(t, true, iter.HasMore())
}
