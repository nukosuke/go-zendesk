package zendesk

// FieldTypes is an array for field type string
type FieldTypes []string

// SystemConditionFieldTypes is an array of condition field type string which defined by system
// https://developer.zendesk.com/rest_api/docs/core/triggers#conditions-reference
var SystemConditionFieldTypes = FieldTypes{
	"group_id",
	"assignee_id",
	"requester_id",
	"organization_id",
	"current_tags",
	"via_id",
	"recipient",
	"type",
	"status",
	"priority",
	"description_includes_word",
	"locale_id",
	"satisfaction_score",
	"subject_includes_word",
	"comment_includes_word",
	"current_via_id",
	"update_type",
	"comment_is_public",
	"ticket_is_public",
	"reopens",
	"replies",
	"agent_stations",
	"group_stations",
	"in_business_hours",
	"requester_twitter_followers_count",
	"requester_twitter_statuses_count",
	"requester_twitter_verified",
	"ticket_type_id",
	"current_via_id",
	"exact_created_at",
	"NEW",
	"OPEN",
	"PENDING",
	"SOLVED",
	"CLOSED",
	"assigned_at",
	"updated_at",
	"requester_updated_at",
	"assignee_updated_at",
	"due_date",
	"until_due_date",
}

// SystemActionFieldTypes is an array of action field type string which defined by system
// https://developer.zendesk.com/rest_api/docs/core/triggers#actions-reference
var SystemActionFieldTypes = FieldTypes{
	"status",
	"type",
	"priority",
	"group_id",
	"assignee_id",
	"set_tags",
	"current_tags",
	"remove_tags",
	"satisfaction_score",
	"notification_user",
	"notification_group",
	"notification_target",
	"tweet_requester",
	"cc",
	"locale_id",
	"subject",
	"comment_value",
	"comment_value_html",
	"comment_mode_is_public",
}

// Include checks if FieldTypes include target type string
func (types FieldTypes) Include(target string) bool {
	for _, t := range types {
		if t == target {
			return true
		}
	}
	return false
}
