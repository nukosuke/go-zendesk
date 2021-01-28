package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMacros(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "macros.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	macros, _, err := client.GetMacros(ctx, &MacroListOptions{
		PageOptions: PageOptions{
			Page:    1,
			PerPage: 10,
		},
		SortBy:    "id",
		SortOrder: "asc",
	})
	if err != nil {
		t.Fatalf("Failed to get macros: %s", err)
	}

	expectedLength := 2
	if len(macros) != expectedLength {
		t.Fatalf("Returned macros does not have the expected length %d. Macros length is %d", expectedLength, len(macros))
	}
}

func TestGetMacro(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "macro.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	macro, err := client.GetMacro(ctx, 2)
	if err != nil {
		t.Fatalf("Failed to get macro: %s", err)
	}

	expectedID := int64(360111062754)
	if macro.ID != expectedID {
		t.Fatalf("Returned macro does not have the expected ID %d. Macro id is %d", expectedID, macro.ID)
	}

}

func TestCreateMacro(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "macro.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	macro, err := client.CreateMacro(ctx, Macro{
		Title: "nyanyanyanya",
		// Comment: MacroComment{
		// 	Body: "(●ↀ ω ↀ )",
		// },
	})
	if err != nil {
		t.Fatalf("Failed to create macro: %s", err)
	}

	expectedID := int64(4)
	if macro.ID != expectedID {
		t.Fatalf("Returned macro does not have the expected ID %d. Macro id is %d", expectedID, macro.ID)
	}
}

func TestUpdateMacro(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "macro.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	macro, err := client.UpdateMacro(ctx, 2, Macro{})
	if err != nil {
		t.Fatalf("Failed to update macro: %s", err)
	}

	expectedID := int64(2)
	if macro.ID != expectedID {
		t.Fatalf("Returned macro does not have the expected ID %d. Macro id is %d", expectedID, macro.ID)
	}
}

func TestUpdateMacroFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "macro.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.UpdateMacro(ctx, 2, Macro{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestDeleteMacro(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteMacro(ctx, 437)
	if err != nil {
		t.Fatalf("Failed to delete macro field: %s", err)
	}
}
