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
		result, err := client.GetBoards(context.Background())

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
		_, err := client.GetBoards(context.Background())

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestGetBoard(t *testing.T) {
	t.Run("returns board on success", func(t *testing.T) {
		board := Board{ID: "board-1", Name: "Test Board"}

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

			w.WriteHeader(http.StatusCreated)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.CreateBoard(context.Background(), CreateBoardPayload{Name: "New Board"})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestUpdateBoard(t *testing.T) {
	t.Run("updates board on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.UpdateBoard(context.Background(), "board-1", UpdateBoardPayload{Name: "Updated Board"})

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
