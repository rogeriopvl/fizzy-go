package fizzy

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCards(t *testing.T) {
	t.Run("returns cards on success", func(t *testing.T) {
		cards := []Card{
			{ID: "card-1", Number: 1, Title: "Card 1"},
			{ID: "card-2", Number: 2, Title: "Card 2"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/cards" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cards)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetCards(context.Background(), CardFilters{})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 cards, got %d", len(result))
		}
	})

	t.Run("applies filters to query string", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()

			boardIDs := query["board_ids[]"]
			if len(boardIDs) != 2 || boardIDs[0] != "b1" || boardIDs[1] != "b2" {
				t.Errorf("unexpected board_ids: %v", boardIDs)
			}

			tagIDs := query["tag_ids[]"]
			if len(tagIDs) != 1 || tagIDs[0] != "tag1" {
				t.Errorf("unexpected tag_ids: %v", tagIDs)
			}

			if query.Get("indexed_by") != "golden" {
				t.Errorf("unexpected indexed_by: %s", query.Get("indexed_by"))
			}

			if query.Get("sorted_by") != "newest" {
				t.Errorf("unexpected sorted_by: %s", query.Get("sorted_by"))
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Card{})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		_, err := client.GetCards(context.Background(), CardFilters{
			BoardIDs:  []string{"b1", "b2"},
			TagIDs:    []string{"tag1"},
			IndexedBy: "golden",
			SortedBy:  "newest",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestGetCard(t *testing.T) {
	t.Run("returns card on success", func(t *testing.T) {
		card := Card{ID: "card-1", Number: 42, Title: "Test Card"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/cards/42" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(card)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetCard(context.Background(), 42)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Number != 42 {
			t.Errorf("expected card number 42, got %d", result.Number)
		}
	})
}

func TestCreateCard(t *testing.T) {
	t.Run("creates card on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/cards" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]CreateCardPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["card"].Title != "New Card" {
				t.Errorf("expected card title 'New Card', got '%s'", body["card"].Title)
			}

			w.WriteHeader(http.StatusCreated)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL), WithBoard("board-1"))
		err := client.CreateCard(context.Background(), CreateCardPayload{Title: "New Card"})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("returns error when no board selected", func(t *testing.T) {
		client, _ := NewClient("/test-account", "test-token")
		err := client.CreateCard(context.Background(), CreateCardPayload{Title: "New Card"})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNoBoardSelected) {
			t.Errorf("expected ErrNoBoardSelected, got %v", err)
		}
	})
}

func TestUpdateCard(t *testing.T) {
	t.Run("updates card on success", func(t *testing.T) {
		card := Card{ID: "card-1", Number: 42, Title: "Updated Card"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(card)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.UpdateCard(context.Background(), 42, UpdateCardPayload{Title: "Updated Card"})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Title != "Updated Card" {
			t.Errorf("expected title 'Updated Card', got '%s'", result.Title)
		}
	})
}

func TestDeleteCard(t *testing.T) {
	t.Run("deletes card on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.DeleteCard(context.Background(), 42)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestCloseCard(t *testing.T) {
	t.Run("closes card on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/closure" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.CloseCard(context.Background(), 42)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestReopenCard(t *testing.T) {
	t.Run("reopens card on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/closure" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.ReopenCard(context.Background(), 42)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestTriageCard(t *testing.T) {
	t.Run("triages card to column", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/triage" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			if body["column_id"] != "col-1" {
				t.Errorf("expected column_id 'col-1', got '%v'", body["column_id"])
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.TriageCard(context.Background(), 42, "col-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestAssignCard(t *testing.T) {
	t.Run("assigns user to card", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/assignments" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body)
			if body["assignee_id"] != "user-1" {
				t.Errorf("expected assignee_id 'user-1', got '%s'", body["assignee_id"])
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.AssignCard(context.Background(), 42, "user-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestTagCard(t *testing.T) {
	t.Run("tags card", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/taggings" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body)
			if body["tag_title"] != "bug" {
				t.Errorf("expected tag_title 'bug', got '%s'", body["tag_title"])
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.TagCard(context.Background(), 42, "bug")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
