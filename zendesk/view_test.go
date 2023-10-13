package zendesk

import (
	"net/http"
	"testing"
)

func TestGetView(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "view.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	view, err := client.GetView(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get view: %s", err)
	}

	expectedID := int64(360002440594)
	if view.ID != expectedID {
		t.Fatalf("Returned view does not have the expected ID %d. View ID is %d", expectedID, view.ID)
	}
}

func TestGetViews(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "views.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	views, _, err := client.GetViews(ctx)
	if err != nil {
		t.Fatalf("Failed to get views: %s", err)
	}

	if len(views) != 2 {
		t.Fatalf("expected length of views is 2, but got %d", len(views))
	}
}

func TestGetCountTicketsInViewsTestGetViews(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "views_ticket_count.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()
	ids := []string{"25", "78"}
	viewsCount, err := client.GetCountTicketsInViews(ctx, ids)
	if err != nil {
		t.Fatalf("Failed to get views tickets count: %s", err)
	}

	if len(viewsCount) != 2 {
		t.Fatalf("expected length of views ticket counts is 2, but got %d", len(viewsCount))
	}
}
