package fizzy

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetColumns(t *testing.T) {
	t.Run("returns columns on success", func(t *testing.T) {
		columns := []Column{
			{ID: "col-1", Name: "To Do"},
			{ID: "col-2", Name: "In Progress"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/boards/board-1/columns" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(columns)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL), WithBoard("board-1"))
		result, err := client.GetColumns(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 columns, got %d", len(result))
		}
	})

	t.Run("returns error when no board selected", func(t *testing.T) {
		client, _ := NewClient("/test-account", "test-token")
		_, err := client.GetColumns(context.Background())

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNoBoardSelected) {
			t.Errorf("expected ErrNoBoardSelected, got %v", err)
		}
	})
}

func TestGetColumn(t *testing.T) {
	t.Run("returns column on success", func(t *testing.T) {
		column := Column{ID: "col-1", Name: "To Do"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/boards/board-1/columns/col-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(column)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL), WithBoard("board-1"))
		result, err := client.GetColumn(context.Background(), "col-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "col-1" {
			t.Errorf("expected column ID 'col-1', got '%s'", result.ID)
		}
	})
}

func TestCreateColumn(t *testing.T) {
	t.Run("creates column on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/columns" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]CreateColumnPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["column"].Name != "New Column" {
				t.Errorf("expected column name 'New Column', got '%s'", body["column"].Name)
			}

			w.WriteHeader(http.StatusCreated)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL), WithBoard("board-1"))
		err := client.CreateColumn(context.Background(), CreateColumnPayload{Name: "New Column"})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("returns error when no board selected", func(t *testing.T) {
		client, _ := NewClient("/test-account", "test-token")
		err := client.CreateColumn(context.Background(), CreateColumnPayload{Name: "New Column"})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNoBoardSelected) {
			t.Errorf("expected ErrNoBoardSelected, got %v", err)
		}
	})
}

func TestUpdateColumn(t *testing.T) {
	t.Run("updates column on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL), WithBoard("board-1"))
		err := client.UpdateColumn(context.Background(), "col-1", UpdateColumnPayload{Name: "Updated Column"})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestDeleteColumn(t *testing.T) {
	t.Run("deletes column on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL), WithBoard("board-1"))
		err := client.DeleteColumn(context.Background(), "col-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
