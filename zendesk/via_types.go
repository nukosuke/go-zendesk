package zendesk

// https://developer.zendesk.com/rest_api/docs/support/triggers#via-types
const (
	// ViaWebForm : Web form
	ViaWebForm = 0
	// ViaMail : Email
	ViaMail = 4
	// ViaChat : Chat
	ViaChat = 29
	// ViaTwitter : Twitter
	ViaTwitter = 30
	// ViaTwitterDM : Twitter DM
	ViaTwitterDM = 26
	// TwitterFavorite : Twitter like
	ViaTwitterFavorite = 23
	// ViaVoicemail : Voicemail
	ViaVoicemail = 33
	// ViaPhoneCallInbound : Phone call (incoming)
	ViaPhoneCallInbound = 34
	// ViaPhoneCallOutbound : Phone call (outbound)
	ViaPhoneCallOutbound = 35
	// ViaAPIVoicemail : CTI voicemail
	ViaAPIVoicemail = 44
	// ViaAPIPhoneCallInbound : CTI phone call (inbound)
	ViaAPIPhoneCallInbound = 45
	// ViaAPIPhoneCallOutbound : CTI phone call (outbound)
	ViaAPIPhoneCallOutbound = 46
	// ViaSMS : SMS
	ViaSMS = 57
	// ViaGetSatisfaction : Get Satisfaction
	ViaGetSatisfaction = 16
	// ViaWebWidget : Web Widget
	ViaWebWidget = 48
	// ViaMobileSDK : Mobile SDK
	ViaMobileSDK = 49
	// ViaMobile : Mobile
	ViaMobile = 56
	// ViaHelpCenter : Help Center post
	ViaHelpCenter = 50
	// ViaWebService : Web service (API)
	ViaWebService = 5
	// ViaRule : Trigger, automation
	ViaRule = 8
	// ViaClosedTicket : Closed ticket
	ViaClosedTicket = 27
	// ViaTicketSharing : Ticket Sharing
	ViaTicketSharing = 31
	// ViaFacebookPost : Facebook post
	ViaFacebookPost = 38
	// ViaFacebookMessage : Facebook private message
	ViaFacebookMessage = 41
	// ViaSatisfactionPrediction : Satisfaction prediction
	ViaSatisfactionPrediction = 54
	// ViaAnyChannel : Channel framework
	ViaAnyChannel = 55
)
