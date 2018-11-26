package zendesk

import "testing"

func TestViaTypeText(t *testing.T) {
	if viaType := ViaTypeText(ViaWebForm); viaType != "web_form" {
		t.Fatal(`expect "web_form", but got "` + viaType + `"`)
	}
}
