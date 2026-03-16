package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWebhooks(t *testing.T) {
	t.Run("returns webhooks on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/webhooks" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Webhook{
				{ID: "hook-1", Name: "Production API"},
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetWebhooks(context.Background(), "board-1", nil)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Fatalf("expected 1 webhook, got %d", len(result))
		}
		if result[0].ID != "hook-1" {
			t.Errorf("expected webhook ID 'hook-1', got '%s'", result[0].ID)
		}
	})
}

func TestGetWebhook(t *testing.T) {
	t.Run("returns a webhook on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/webhooks/hook-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Webhook{
				ID:                "hook-1",
				Name:              "Production API",
				PayloadURL:        "https://api.example.com/webhooks",
				SubscribedActions: []string{"card_closed"},
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetWebhook(context.Background(), "board-1", "hook-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.PayloadURL != "https://api.example.com/webhooks" {
			t.Errorf("expected payload URL to be decoded, got '%s'", result.PayloadURL)
		}
	})
}

func TestCreateWebhook(t *testing.T) {
	t.Run("creates a webhook on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/webhooks" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]CreateWebhookPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["webhook"].URL != "https://api.example.com/webhooks" {
				t.Errorf("expected webhook URL to be encoded, got '%s'", body["webhook"].URL)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Webhook{
				ID:         "hook-1",
				PayloadURL: "https://api.example.com/webhooks",
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.CreateWebhook(context.Background(), "board-1", CreateWebhookPayload{
			Name:              "Production API",
			URL:               "https://api.example.com/webhooks",
			SubscribedActions: []string{"card_closed"},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "hook-1" {
			t.Errorf("expected created webhook ID 'hook-1', got '%s'", result.ID)
		}
	})
}

func TestUpdateWebhook(t *testing.T) {
	t.Run("updates a webhook on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/webhooks/hook-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]UpdateWebhookPayload
			json.NewDecoder(r.Body).Decode(&body)
			if len(body["webhook"].SubscribedActions) != 1 || body["webhook"].SubscribedActions[0] != "card_closed" {
				t.Errorf("expected subscribed actions to be encoded, got %#v", body["webhook"].SubscribedActions)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Webhook{
				ID:                "hook-1",
				SubscribedActions: []string{"card_closed"},
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.UpdateWebhook(context.Background(), "board-1", "hook-1", UpdateWebhookPayload{
			Name:              "Production API",
			SubscribedActions: []string{"card_closed"},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result.SubscribedActions) != 1 || result.SubscribedActions[0] != "card_closed" {
			t.Errorf("expected updated subscribed actions, got %#v", result.SubscribedActions)
		}
	})
}

func TestDeleteWebhook(t *testing.T) {
	t.Run("deletes a webhook on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/webhooks/hook-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.DeleteWebhook(context.Background(), "board-1", "hook-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestActivateWebhook(t *testing.T) {
	t.Run("reactivates a webhook on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/boards/board-1/webhooks/hook-1/activation" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Webhook{
				ID:     "hook-1",
				Active: true,
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.ActivateWebhook(context.Background(), "board-1", "hook-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !result.Active {
			t.Error("expected webhook to be active")
		}
	})
}
