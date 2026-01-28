package fizzy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNotifications(t *testing.T) {
	t.Run("returns notifications on success", func(t *testing.T) {
		notifications := []Notification{
			{ID: "notif-1", Title: "New comment"},
			{ID: "notif-2", Title: "Card assigned"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/notifications" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(notifications)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetNotifications(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 notifications, got %d", len(result))
		}
	})
}

func TestGetNotification(t *testing.T) {
	t.Run("returns notification on success", func(t *testing.T) {
		notification := Notification{ID: "notif-1", Title: "New comment"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/test-account/notifications/notif-1" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(notification)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		result, err := client.GetNotification(context.Background(), "notif-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "notif-1" {
			t.Errorf("expected notification ID 'notif-1', got '%s'", result.ID)
		}
	})
}

func TestMarkNotificationRead(t *testing.T) {
	t.Run("marks notification as read", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/notifications/notif-1/reading" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.MarkNotificationRead(context.Background(), "notif-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestMarkNotificationUnread(t *testing.T) {
	t.Run("marks notification as unread", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/notifications/notif-1/reading" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.MarkNotificationUnread(context.Background(), "notif-1")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestMarkAllNotificationsRead(t *testing.T) {
	t.Run("marks all notifications as read", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/test-account/notifications/bulk_reading" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}

			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client, _ := NewClient("/test-account", "test-token", WithBaseURL(server.URL))
		err := client.MarkAllNotificationsRead(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
