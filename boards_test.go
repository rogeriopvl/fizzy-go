package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBoards(t *testing.T) {
	t.Run("returns boards on success", func(t *testing.T) {
		boards := []Board{
			{ID: "board-1", Name: "Board 1"},
			{ID: "board-2", Name: "Board 2"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			if r.Header.Get("Authorization") != "Bearer test-token" {
				t.Errorf("unexpected Authorization header: %s", r.Header.Get("Authorization"))
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(boards)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetBoards(context.Background(), nil)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 boards, got %d", len(result))
		}
		if result[0].ID != "board-1" {
			t.Errorf("expected board ID 'board-1', got '%s'", result[0].ID)
		}
	})

	t.Run("returns error on non-200 status", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal error"))
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		_, err := client.GetBoards(context.Background(), nil)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestGetBoard(t *testing.T) {
	t.Run("returns board on success", func(t *testing.T) {
		board := Board{
			ID:                       "board-1",
			Name:                     "Test Board",
			AutoPostponePeriodInDays: 30,
			PublicURL:                "https://example.com/public/board-1",
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/boards/board-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(board)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetBoard(context.Background(), "board-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "board-1" {
			t.Errorf("expected board ID 'board-1', got '%s'", result.ID)
		}
		if result.AutoPostponePeriodInDays != 30 {
			t.Errorf("expected auto postpone period 30, got %d", result.AutoPostponePeriodInDays)
		}
		if result.PublicURL != "https://example.com/public/board-1" {
			t.Errorf("expected public URL to be decoded, got '%s'", result.PublicURL)
		}
	})
}

func TestCreateBoard(t *testing.T) {
	t.Run("creates board on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]CreateBoardPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["board"].Name != "New Board" {
				t.Errorf("expected board name 'New Board', got '%s'", body["board"].Name)
			}
			if body["board"].AutoPostponePeriodInDays != 30 {
				t.Errorf("expected auto postpone period 30, got %d", body["board"].AutoPostponePeriodInDays)
			}

			w.WriteHeader(http.StatusCreated)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.CreateBoard(context.Background(), CreateBoardPayload{
			Name:                     "New Board",
			AutoPostponePeriodInDays: 30,
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestUpdateBoard(t *testing.T) {
	t.Run("updates board on success", func(t *testing.T) {
		autoPostponePeriodInDays := 60

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]UpdateBoardPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["board"].AutoPostponePeriodInDays == nil || *body["board"].AutoPostponePeriodInDays != 60 {
				t.Errorf("expected auto postpone period 60, got %#v", body["board"].AutoPostponePeriodInDays)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.UpdateBoard(context.Background(), "board-1", UpdateBoardPayload{
			Name:                     "Updated Board",
			AutoPostponePeriodInDays: &autoPostponePeriodInDays,
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestDeleteBoard(t *testing.T) {
	t.Run("deletes board on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.DeleteBoard(context.Background(), "board-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestPublishBoard(t *testing.T) {
	t.Run("publishes board on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/publication" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Board{
				ID:        "board-1",
				Name:      "Board 1",
				PublicURL: "https://example.com/public/board-1",
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.PublishBoard(context.Background(), "board-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.PublicURL != "https://example.com/public/board-1" {
			t.Errorf("expected public URL to be returned, got '%s'", result.PublicURL)
		}
	})
}

func TestUnpublishBoard(t *testing.T) {
	t.Run("unpublishes board on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/publication" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.UnpublishBoard(context.Background(), "board-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestUpdateBoardEntropy(t *testing.T) {
	t.Run("updates board entropy on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/entropy" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]EntropyPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["board"].AutoPostponePeriodInDays != 90 {
				t.Errorf("expected auto postpone period 90, got %d", body["board"].AutoPostponePeriodInDays)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Board{
				ID:                       "board-1",
				AutoPostponePeriodInDays: 90,
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.UpdateBoardEntropy(context.Background(), "board-1", EntropyPayload{
			AutoPostponePeriodInDays: 90,
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.AutoPostponePeriodInDays != 90 {
			t.Errorf("expected updated auto postpone period 90, got %d", result.AutoPostponePeriodInDays)
		}
	})
}

func TestGetBoardAccesses(t *testing.T) {
	t.Run("returns board accesses on success", func(t *testing.T) {
		response := BoardAccesses{
			BoardID:   "board-1",
			AllAccess: false,
			Users: []BoardAccess{
				{User: User{ID: "user-1", Name: "Alice"}, HasAccess: true, Involvement: "watching"},
				{User: User{ID: "user-2", Name: "Bob"}, HasAccess: false},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/accesses" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetBoardAccesses(context.Background(), "board-1", nil)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.BoardID != "board-1" {
			t.Errorf("expected board_id 'board-1', got '%s'", result.BoardID)
		}
		if len(result.Users) != 2 {
			t.Fatalf("expected 2 users, got %d", len(result.Users))
		}
		if !result.Users[0].HasAccess {
			t.Errorf("expected first user to have access")
		}
		if result.Users[0].Involvement != "watching" {
			t.Errorf("expected first user involvement 'watching', got '%s'", result.Users[0].Involvement)
		}
	})
}
