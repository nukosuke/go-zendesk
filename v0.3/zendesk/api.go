package zendesk

//go:generate  mockgen -destination=mock/client.go -package=mock -mock_names=API=Client github.com/nukosuke/go-zendesk/v0.3/zendesk API

// API an interface containing all of the zendesk client methods
type API interface {
	AttachmentAPI
	BrandAPI
	GroupAPI
	LocaleAPI
	TicketFieldAPI
	TicketFormAPI
	TriggerAPI
	UserAPI
}

var _ API = (*Client)(nil)
