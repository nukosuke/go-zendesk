package zendesk

import "testing"

func TestConditionFieldText(t *testing.T) {
	if cond := ConditionFieldText(ConditionFieldGroupID); cond != "group_id" {
		t.Fatal(`expected "group_id", but got ` + cond)
	}
}
