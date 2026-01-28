package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	t.Run("returns users on success", func(t *testing.T) {
		users := []User{
			{ID: "user-1", Name: "Alice"},
			{ID: "user-2", Name: "Bob"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/users" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetUsers(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 users, got %d", len(result))
		}
	})
}

func TestGetUser(t *testing.T) {
	t.Run("returns user on success", func(t *testing.T) {
		user := User{ID: "user-1", Name: "Alice"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/users/user-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetUser(context.Background(), "user-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "user-1" {
			t.Errorf("expected user ID 'user-1', got '%s'", result.ID)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("updates user on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/users/user-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			var body map[string]UpdateUserPayload
			json.NewDecoder(r.Body).Decode(&body)
			if body["user"].Name != "Alice Updated" {
				t.Errorf("expected user name 'Alice Updated', got '%s'", body["user"].Name)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.UpdateUser(context.Background(), "user-1", UpdateUserPayload{Name: "Alice Updated"})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestDeactivateUser(t *testing.T) {
	t.Run("deactivates user on success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/users/user-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.DeactivateUser(context.Background(), "user-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
