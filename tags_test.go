package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTags(t *testing.T) {
	t.Run("returns tags on success", func(t *testing.T) {
		tags := []Tag{
			{ID: "tag-1", Title: "bug"},
			{ID: "tag-2", Title: "feature"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/tags" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tags)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetTags(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 tags, got %d", len(result))
		}
		if result[0].Title != "bug" {
			t.Errorf("expected tag title 'bug', got '%s'", result[0].Title)
		}
	})
}
