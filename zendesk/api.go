package zendesk

//nolint
//go:generate  mockgen -destination=mock/client.go -package=mock -mock_names=API=Client github.com/nukosuke/go-zendesk/zendesk API

// API an interface containing all of the zendesk client methods
type API interface {
	AttachmentAPI
	BrandAPI
	DynamicContentAPI
	GroupAPI
	LocaleAPI
	TicketFieldAPI
	TicketFormAPI
	TriggerAPI
	UserAPI
}

var _ API = (*Client)(nil)
