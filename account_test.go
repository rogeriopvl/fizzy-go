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
			if r.URL.Path != "/test-account/account/settings" {
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

func TestGetAccountJoinCode(t *testing.T) {
	t.Run("returns join code on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/account/join_code" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(JoinCode{
				Code:       "abc123",
				UsageCount: 3,
				UsageLimit: 10,
				URL:        "http://app.fizzy.localhost/897362094/join/abc123",
				Active:     true,
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetAccountJoinCode(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Code != "abc123" {
			t.Errorf("expected code abc123, got %s", result.Code)
		}
		if result.UsageCount != 3 {
			t.Errorf("expected usage count 3, got %d", result.UsageCount)
		}
		if result.UsageLimit != 10 {
			t.Errorf("expected usage limit 10, got %d", result.UsageLimit)
		}
		if !result.Active {
			t.Errorf("expected active true, got false")
		}
	})
}

func TestUpdateAccountJoinCode(t *testing.T) {
	t.Run("updates join code usage limit on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/account/join_code" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]UpdateJoinCodePayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["account_join_code"].UsageLimit != 25 {
				t.Errorf("expected usage limit 25, got %d", body["account_join_code"].UsageLimit)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.UpdateAccountJoinCode(context.Background(), UpdateJoinCodePayload{
			UsageLimit: 25,
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestResetAccountJoinCode(t *testing.T) {
	t.Run("resets join code on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/account/join_code" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.ResetAccountJoinCode(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestUpdateAccountEntropy(t *testing.T) {
	t.Run("updates account entropy on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/account/entropy" {
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
