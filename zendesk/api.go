package zendesk

// API an interface containing all of the zendesk client methods
type API interface {
	GroupAPI
	LocaleAPI
	TicketFieldAPI
	TicketFormAPI
}
