package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCardStep(t *testing.T) {
	t.Run("returns step on success", func(t *testing.T) {
		step := Step{ID: "step-1", Content: "Write tests", Completed: false}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/cards/42/steps/step-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(step)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetCardStep(context.Background(), 42, "step-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "step-1" {
			t.Errorf("expected step ID 'step-1', got '%s'", result.ID)
		}
	})
}

func TestCreateCardStep(t *testing.T) {
	t.Run("creates step on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/steps" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			if body["step"]["content"] != "Write tests" {
				t.Errorf("expected step content 'Write tests', got '%v'", body["step"]["content"])
			}
			if body["step"]["completed"] != false {
				t.Errorf("expected step completed false, got '%v'", body["step"]["completed"])
			}

			w.WriteHeader(http.StatusCreated)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.CreateCardStep(context.Background(), 42, "Write tests", false)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Content != "Write tests" {
			t.Errorf("expected step content 'Write tests', got '%s'", result.Content)
		}
	})
}

func TestUpdateCardStep(t *testing.T) {
	t.Run("updates step on success", func(t *testing.T) {
		step := Step{ID: "step-1", Content: "Write tests", Completed: true}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/steps/step-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(step)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		completed := true
		result, err := client.UpdateCardStep(context.Background(), 42, "step-1", nil, &completed)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !result.Completed {
			t.Error("expected step to be completed")
		}
	})
}

func TestDeleteCardStep(t *testing.T) {
	t.Run("deletes step on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/cards/42/steps/step-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.DeleteCardStep(context.Background(), 42, "step-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
