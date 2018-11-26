package zendesk

import "testing"

func TestActionFieldText(t *testing.T) {
	if action := ActionFieldText(ActionFieldStatus); action != "status" {
		t.Fatal(`expected "status", but got ` + action)
	}
}
