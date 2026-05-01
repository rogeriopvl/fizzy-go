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

func TestCreateAccessToken(t *testing.T) {
	t.Run("creates an access token on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/my/access_tokens" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]CreateAccessTokenPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["access_token"].Description != "Fizzy CLI" {
				t.Errorf("expected description 'Fizzy CLI', got '%s'", body["access_token"].Description)
			}
			if body["access_token"].Permission != "write" {
				t.Errorf("expected permission 'write', got '%s'", body["access_token"].Permission)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(PersonalAccessToken{
				Token:       "secret-token",
				Description: "Fizzy CLI",
				Permission:  "write",
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.CreateAccessToken(context.Background(), CreateAccessTokenPayload{
			Description: "Fizzy CLI",
			Permission:  "write",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Token != "secret-token" {
			t.Errorf("expected created token to be returned, got '%s'", result.Token)
		}
	})
}

func TestUpdateMyTimezone(t *testing.T) {
	t.Run("updates timezone on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/my/timezone" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body)
			if body["timezone_name"] != "America/New_York" {
				t.Errorf("expected timezone_name 'America/New_York', got '%s'", body["timezone_name"])
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.UpdateMyTimezone(context.Background(), "America/New_York")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
