package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountSettings(t *testing.T) {
	t.Run("returns account settings on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/account/settings" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Account{
				ID:                       "acc-1",
				Name:                     "37signals",
				CardsCount:               5,
				AutoPostponePeriodInDays: 30,
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetAccountSettings(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.CardsCount != 5 {
			t.Errorf("expected cards count 5, got %d", result.CardsCount)
		}
		if result.AutoPostponePeriodInDays != 30 {
			t.Errorf("expected auto postpone period 30, got %d", result.AutoPostponePeriodInDays)
		}
	})
}

func TestUpdateAccountEntropy(t *testing.T) {
	t.Run("updates account entropy on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/account/entropy" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]EntropyPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["entropy"].AutoPostponePeriodInDays != 30 {
				t.Errorf("expected auto postpone period 30, got %d", body["entropy"].AutoPostponePeriodInDays)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Account{
				ID:                       "acc-1",
				AutoPostponePeriodInDays: 30,
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.UpdateAccountEntropy(context.Background(), EntropyPayload{
			AutoPostponePeriodInDays: 30,
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.AutoPostponePeriodInDays != 30 {
			t.Errorf("expected updated auto postpone period 30, got %d", result.AutoPostponePeriodInDays)
		}
	})
}
