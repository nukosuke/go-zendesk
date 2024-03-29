
// Code generated by Script. DO NOT EDIT.
// Source: script/codegen/main.go
//
// Generated by this command:
//
//	go run script/codegen/main.go

package zendesk

import "context"

func (z *Client) GetViewsIterator(ctx context.Context, opts *PaginationOptions) *Iterator[View] {
	return &Iterator[View]{
		CommonOptions: opts.CommonOptions,
		pageSize:      opts.PageSize,
		hasMore:       true,
		isCBP:         opts.IsCBP,
		pageAfter:     "",
		pageIndex:     1,
		ctx:           ctx,
		obpFunc:       z.GetViewsOBP,
		cbpFunc:       z.GetViewsCBP,
	}
}

func (z *Client) GetViewsOBP(ctx context.Context, opts *OBPOptions) ([]View, Page, error) {
	var data struct {
		Views []View `json:"views"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &OBPOptions{}
	}
	
	u, err := addOptions("/views.json", tmp)
	
	if err != nil {
		return nil, Page{}, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.Views, data.Page, nil
}

func (z *Client) GetViewsCBP(ctx context.Context, opts *CBPOptions) ([]View, CursorPaginationMeta, error) {
	var data struct {
		Views []View `json:"views"`
		Meta    CursorPaginationMeta `json:"meta"`
	}

	tmp := opts
	if tmp == nil {
		tmp = &CBPOptions{}
	}
	
	u, err := addOptions("/views.json", tmp)
	
	if err != nil {
		return nil, data.Meta, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, data.Meta, err
	}
	return data.Views, data.Meta, nil
}

