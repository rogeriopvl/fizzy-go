package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetActivities(t *testing.T) {
	t.Run("returns activities on success", func(t *testing.T) {
		activities := []Activity{
			{
				ID:            "act-1",
				Action:        "card_closed",
				EventableType: "Card",
				Description:   "Closed card",
			},
			{
				ID:            "act-2",
				Action:        "comment_created",
				EventableType: "Comment",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/activities" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(activities)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetActivities(context.Background(), nil)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 activities, got %d", len(result))
		}
		if result[0].Action != "card_closed" {
			t.Errorf("expected first action 'card_closed', got '%s'", result[0].Action)
		}
	})

	t.Run("applies filters as query parameters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			creators := q["creator_ids[]"]
			boards := q["board_ids[]"]

			if len(creators) != 2 || creators[0] != "user-1" || creators[1] != "user-2" {
				t.Errorf("expected creator_ids[]=[user-1, user-2], got %v", creators)
			}
			if len(boards) != 1 || boards[0] != "board-1" {
				t.Errorf("expected board_ids[]=[board-1], got %v", boards)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Activity{})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		_, err := client.GetActivities(context.Background(), &ActivityFilters{
			CreatorIDs: []string{"user-1", "user-2"},
			BoardIDs:   []string{"board-1"},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("preserves particulars as raw json for action-specific decoding", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{
				"id": "act-1",
				"action": "card_triaged",
				"particulars": {"column": "In Progress"},
				"eventable_type": "Card"
			}]`))
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetActivities(context.Background(), nil)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Fatalf("expected 1 activity, got %d", len(result))
		}

		var particulars struct {
			Column string `json:"column"`
		}
		if err := json.Unmarshal(result[0].Particulars, &particulars); err != nil {
			t.Fatalf("failed to decode particulars: %v", err)
		}
		if particulars.Column != "In Progress" {
			t.Errorf("expected column 'In Progress', got '%s'", particulars.Column)
		}
	})
}
