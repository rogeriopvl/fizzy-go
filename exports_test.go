package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccountExport(t *testing.T) {
	t.Run("starts account export on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/account/exports" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Export{
				ID:        "exp-1",
				Status:    "pending",
				CreatedAt: "2026-04-02T12:34:56Z",
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.CreateAccountExport(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "exp-1" {
			t.Errorf("expected id 'exp-1', got '%s'", result.ID)
		}
		if result.Status != "pending" {
			t.Errorf("expected status 'pending', got '%s'", result.Status)
		}
	})
}

func TestGetAccountExport(t *testing.T) {
	t.Run("returns account export status on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/account/exports/exp-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Export{
				ID:          "exp-1",
				Status:      "completed",
				CreatedAt:   "2026-04-02T12:34:56Z",
				DownloadURL: "https://example.com/export.zip",
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetAccountExport(context.Background(), "exp-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Status != "completed" {
			t.Errorf("expected status 'completed', got '%s'", result.Status)
		}
		if result.DownloadURL != "https://example.com/export.zip" {
			t.Errorf("expected download_url, got '%s'", result.DownloadURL)
		}
	})
}

func TestCreateUserDataExport(t *testing.T) {
	t.Run("starts user data export on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/users/user-1/data_exports" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Export{
				ID:        "exp-2",
				Status:    "pending",
				CreatedAt: "2026-04-02T12:34:56Z",
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.CreateUserDataExport(context.Background(), "user-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "exp-2" {
			t.Errorf("expected id 'exp-2', got '%s'", result.ID)
		}
	})
}

func TestGetUserDataExport(t *testing.T) {
	t.Run("returns user data export status on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/users/user-1/data_exports/exp-2" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Export{
				ID:          "exp-2",
				Status:      "completed",
				DownloadURL: "https://example.com/user-export.zip",
			})
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetUserDataExport(context.Background(), "user-1", "exp-2")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Status != "completed" {
			t.Errorf("expected status 'completed', got '%s'", result.Status)
		}
		if result.DownloadURL == "" {
			t.Errorf("expected non-empty download_url")
		}
	})
}
