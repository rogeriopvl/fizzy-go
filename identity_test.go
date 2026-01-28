package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMyIdentity(t *testing.T) {
	t.Run("returns identity on success", func(t *testing.T) {
		identity := GetMyIdentityResponse{
			Accounts: []Account{
				{ID: "acc-1", Name: "Account 1", Slug: "/123"},
				{ID: "acc-2", Name: "Account 2", Slug: "/456"},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/my/identity" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			if r.Header.Get("Authorization") != "Bearer test-token" {
				t.Errorf("unexpected Authorization header: %s", r.Header.Get("Authorization"))
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(identity)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetMyIdentity(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result.Accounts) != 2 {
			t.Fatalf("expected 2 accounts, got %d", len(result.Accounts))
		}
		if result.Accounts[0].Slug != "/123" {
			t.Errorf("expected slug '/123', got '%s'", result.Accounts[0].Slug)
		}
	})

	t.Run("returns error on failure", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		_, err := client.GetMyIdentity(context.Background())

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
