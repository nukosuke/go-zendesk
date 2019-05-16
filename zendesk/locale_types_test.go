package zendesk

import "testing"

func TestLocaleTypeText(t *testing.T) {
	if l := LocaleTypeText(LocaleENUS - 1); l != "" {
		t.Fatalf("expected empty string, but got %v", l)
	}
	if l := LocaleTypeText(LocaleENPH + 1); l != "" {
		t.Fatalf("expected empty string, but got %v", l)
	}
}
