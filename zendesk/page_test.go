package zendesk

import "testing"

func TestHasNext(t *testing.T) {
	pageURL := "https://example.com/pages/2"

	if (Page{NextPage: &pageURL}).HasNext() != true {
		t.Fatalf("expext true, but got false")
	}
	if (Page{NextPage: nil}).HasNext() != false {
		t.Fatalf("expect false, but got true")
	}
}

func TestHasPrev(t *testing.T) {
	pageURL := "https://example.com/pages/1"

	if (Page{PreviousPage: &pageURL}).HasPrev() != true {
		t.Fatalf("expect true, but got false")
	}
	if (Page{PreviousPage: nil}).HasPrev() != false {
		t.Fatalf("expect false, but got true")
	}
}
