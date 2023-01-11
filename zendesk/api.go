package zendesk

//nolint
//go:generate  mockgen -destination=mock/client.go -package=mock -mock_names=API=Client github.com/nukosuke/go-zendesk/zendesk API

// API an interface containing all of the zendesk client methods
type API interface {
	AppAPI
	AttachmentAPI
	AutomationAPI
	BaseAPI
	BrandAPI
	CustomRoleAPI
	DynamicContentAPI
	GroupAPI
	GroupMembershipAPI
	LocaleAPI
	MacroAPI
	OrganizationAPI
	OrganizationMembershipAPI
	SearchAPI
	SLAPolicyAPI
	TagAPI
	TargetAPI
	TicketAuditAPI
	TicketAPI
	TicketCommentAPI
	TicketFieldAPI
	TicketFormAPI
	TriggerAPI
	UserAPI
	UserFieldAPI
	ViewAPI
	WebhookAPI
}

var _ API = (*Client)(nil)
