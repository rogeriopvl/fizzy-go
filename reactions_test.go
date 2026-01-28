package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCommentReactions(t *testing.T) {
	t.Run("returns reactions on success", func(t *testing.T) {
		reactions := []Reaction{
			{ID: "reaction-1", Content: "👍"},
			{ID: "reaction-2", Content: "🎉"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/cards/42/comments/comment-1/reactions" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reactions)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetCommentReactions(context.Background(), 42, "comment-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 reactions, got %d", len(result))
		}
	})
}

func TestCreateCommentReaction(t *testing.T) {
	t.Run("creates reaction on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/comments/comment-1/reactions" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]map[string]string
			json.NewDecoder(r.Body).Decode(&body)
			if body["reaction"]["content"] != "👍" {
				t.Errorf("expected reaction content '👍', got '%s'", body["reaction"]["content"])
			}

			w.WriteHeader(http.StatusCreated)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.CreateCommentReaction(context.Background(), 42, "comment-1", "👍")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Content != "👍" {
			t.Errorf("expected reaction content '👍', got '%s'", result.Content)
		}
	})
}

func TestDeleteCommentReaction(t *testing.T) {
	t.Run("deletes reaction on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/comments/comment-1/reactions/reaction-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.DeleteCommentReaction(context.Background(), 42, "comment-1", "reaction-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
