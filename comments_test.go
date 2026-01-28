package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCardComments(t *testing.T) {
	t.Run("returns comments on success", func(t *testing.T) {
		comments := []Comment{
			{ID: "comment-1"},
			{ID: "comment-2"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/cards/42/comments" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comments)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetCardComments(context.Background(), 42)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 comments, got %d", len(result))
		}
	})
}

func TestGetCardComment(t *testing.T) {
	t.Run("returns comment on success", func(t *testing.T) {
		comment := Comment{ID: "comment-1"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/cards/42/comments/comment-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comment)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetCardComment(context.Background(), 42, "comment-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "comment-1" {
			t.Errorf("expected comment ID 'comment-1', got '%s'", result.ID)
		}
	})
}

func TestCreateCardComment(t *testing.T) {
	t.Run("creates comment on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/comments" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]map[string]string
			json.NewDecoder(r.Body).Decode(&body)
			if body["comment"]["body"] != "Test comment" {
				t.Errorf("expected comment body 'Test comment', got '%s'", body["comment"]["body"])
			}

			w.WriteHeader(http.StatusCreated)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		_, err := client.CreateCardComment(context.Background(), 42, "Test comment")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestUpdateCardComment(t *testing.T) {
	t.Run("updates comment on success", func(t *testing.T) {
		comment := Comment{ID: "comment-1"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comment)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.UpdateCardComment(context.Background(), 42, "comment-1", "Updated comment")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "comment-1" {
			t.Errorf("expected comment ID 'comment-1', got '%s'", result.ID)
		}
	})
}

func TestDeleteCardComment(t *testing.T) {
	t.Run("deletes comment on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.DeleteCardComment(context.Background(), 42, "comment-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
