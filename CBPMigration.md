# OBP to CBP migration

Zendesk is [obsoleting](https://support.zendesk.com/hc/en-us/articles/4408846180634-Introducing-Pagination-Changes-Zendesk-API#h_01F7Y57A0G5M3R8JXGCQTBKVWA) Offset Based Pagination (OBP), it's recomended to start adopting Cursor Based Pagination (CBP). This SDK have created pagination Iterators to help facilite the change.

To use the pagination iterator, start with `NewPaginationOptions()` function, it will return a `PaginationOptions` object, you can specify the default page size in `PageSize` variable. By default, `PageSize` is 100. Then you can call the `client.GetXXXXXIterator(ctx, ops)` to return an object pagination iterator, with the iterator, you can iterator through the objects with `HasMore()` and `GetNext()` until `HasMore` return `false`. 

```go
ops := NewPaginationOptions()
// ops.PageSize = 50 // PageSize can be set to 50
it := client.GetTicketsIterator(ctx, ops)
for it.HasMore() {
    tickets, err := it.GetNext()
    if err == nil {
        for _, ticket := range tickets {
            println(ticket.Subject)
        }
    }
}
```

If the API endpoint requires more options like organization ID, it can be set into the `Id` attribute like below example:

```go
ops := NewPaginationOptions()
ops.Sort = "updated_at"
ops.PageSize = 10
ops.Id = 360363695492
it := client.GetOrganizationTicketsIterator(ctx, ops)

for it.HasMore() {
    tickets, err := it.GetNext()
    if err == nil {
        for _, ticket := range tickets {
            println(ticket.Subject)
        }
    }
}
```

For any API specific parameters, they are predefined in the `CommonOptions` struct, we can set these attributes in the `PaginationOptions` object. 
If new attributes are introduced to any existing or new API endpoints, it can be added into this struct.

```go
type CommonOptions struct {
	Active        bool     `url:"active,omitempty"`
	Role          string   `url:"role,omitempty"`
	Roles         []string `url:"role[],omitempty"`
	PermissionSet int64    `url:"permission_set,omitempty"`

	// SortBy can take "assignee", "assignee.name", "created_at", "group", "id",
	// "locale", "requester", "requester.name", "status", "subject", "updated_at"
	SortBy string `url:"sort_by,omitempty"`

	// SortOrder can take "asc" or "desc"
	SortOrder      string `url:"sort_order,omitempty"`
	Sort           string `url:"sort,omitempty"`
	Id             int64
	GroupID        int64 `json:"group_id,omitempty" url:"group_id,omitempty"`
	UserID         int64 `json:"user_id,omitempty" url:"user_id,omitempty"`
	OrganizationID int64 `json:"organization_id,omitempty" url:"organization_id,omitempty"`

	Access            string `json:"access"`
	Category          int    `json:"category"`
	Include           string `json:"include" url:"include,omitempty"`
	OnlyViewable      bool   `json:"only_viewable"`
	Query             string `url:"query"`
	EndUserVisible    bool   `url:"end_user_visible,omitempty"`
	FallbackToDefault bool   `url:"fallback_to_default,omitempty"`
	AssociatedToBrand bool   `url:"associated_to_brand,omitempty"`
	CategoryID        string `url:"category_id,omitempty"`

	IncludeInlineImages string `url:"include_inline_images,omitempty"`
}
```

## To regenerate CBP(Cursor Based Pagination), OBP(Offset Based Pagination) helper function and Iterators

If a new API endpoint supports CBP, add a new element to the funcData in script/codegen/main.go file like this:

```go
{
    FuncName:    "Automations",
    ObjectName:  "Automation",
    ApiEndpoint: "/automation.json",
    JsonName:    "automations",
    FileName:    "automation",
},
```

should use the script to generate the helper functions and the iterator
`go run script/codegen/main.go`

