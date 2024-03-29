
// Code generated by Script. DO NOT EDIT.
// Source: script/codegen/main.go
//
// Generated by this command:
//
//	go run script/codegen/main.go

package zendesk

import "context"

func (z *Client) GetOrganizationMembershipsIterator(ctx context.Context, opts *PaginationOptions) *Iterator[OrganizationMembership] {
	return &Iterator[OrganizationMembership]{
		CommonOptions: opts.CommonOptions,
		pageSize:      opts.PageSize,
		hasMore:       true,
		isCBP:         opts.IsCBP,
		pageAfter:     "",
		pageIndex:     1,
		ctx:           ctx,
		obpFunc:       z.GetOrganizationMembershipsOBP,
		cbpFunc:       z.GetOrganizationMembershipsCBP,
	}
}

func (z *Client) GetOrganizationMembershipsOBP(ctx context.Context, opts *OBPOptions) ([]OrganizationMembership, Page, error) {
	var data struct {
		OrganizationMemberships []OrganizationMembership `json:"organization_memberships"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &OBPOptions{}
	}
	
	u, err := addOptions("/organization_memberships.json", tmp)
	
	if err != nil {
		return nil, Page{}, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.OrganizationMemberships, data.Page, nil
}

func (z *Client) GetOrganizationMembershipsCBP(ctx context.Context, opts *CBPOptions) ([]OrganizationMembership, CursorPaginationMeta, error) {
	var data struct {
		OrganizationMemberships []OrganizationMembership `json:"organization_memberships"`
		Meta    CursorPaginationMeta `json:"meta"`
	}

	tmp := opts
	if tmp == nil {
		tmp = &CBPOptions{}
	}
	
	u, err := addOptions("/organization_memberships.json", tmp)
	
	if err != nil {
		return nil, data.Meta, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, data.Meta, err
	}
	return data.OrganizationMemberships, data.Meta, nil
}

