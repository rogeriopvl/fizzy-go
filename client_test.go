package fizzy

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Run("creates client with required parameters", func(t *testing.T) {
		client, err := NewClient("/test-account", "test-token")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if client.AccessToken != "test-token" {
			t.Errorf("expected AccessToken 'test-token', got '%s'", client.AccessToken)
		}
		if client.BaseURL != DefaultBaseURL {
			t.Errorf("expected BaseURL '%s', got '%s'", DefaultBaseURL, client.BaseURL)
		}
		if client.AccountBaseURL != DefaultBaseURL+"/test-account" {
			t.Errorf("expected AccountBaseURL '%s', got '%s'", DefaultBaseURL+"/test-account", client.AccountBaseURL)
		}
		if client.BoardBaseURL != "" {
			t.Errorf("expected empty BoardBaseURL, got '%s'", client.BoardBaseURL)
		}
	})

	t.Run("returns error when accountSlug is empty", func(t *testing.T) {
		_, err := NewClient("", "test-token")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("returns error when accessToken is empty", func(t *testing.T) {
		_, err := NewClient("/test-account", "")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("applies WithBoard option", func(t *testing.T) {
		client, err := NewClient("/test-account", "test-token", WithBoard("board-123"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBoardURL := DefaultBaseURL + "/test-account/boards/board-123"
		if client.BoardBaseURL != expectedBoardURL {
			t.Errorf("expected BoardBaseURL '%s', got '%s'", expectedBoardURL, client.BoardBaseURL)
		}
	})

	t.Run("applies WithBaseURL option", func(t *testing.T) {
		customURL := "https://custom.fizzy.do"
		client, err := NewClient("/test-account", "test-token", WithBaseURL(customURL))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if client.BaseURL != customURL {
			t.Errorf("expected BaseURL '%s', got '%s'", customURL, client.BaseURL)
		}
		if client.AccountBaseURL != customURL+"/test-account" {
			t.Errorf("expected AccountBaseURL '%s', got '%s'", customURL+"/test-account", client.AccountBaseURL)
		}
	})

	t.Run("applies WithHTTPClient option", func(t *testing.T) {
		customHTTPClient := &http.Client{Timeout: 60 * time.Second}
		client, err := NewClient("/test-account", "test-token", WithHTTPClient(customHTTPClient))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if client.HTTPClient != customHTTPClient {
			t.Error("expected custom HTTP client to be set")
		}
	})

	t.Run("WithBaseURL and WithBoard work together", func(t *testing.T) {
		customURL := "https://custom.fizzy.do"
		client, err := NewClient("/test-account", "test-token",
			WithBaseURL(customURL),
			WithBoard("board-123"),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBoardURL := customURL + "/test-account/boards/board-123"
		if client.BoardBaseURL != expectedBoardURL {
			t.Errorf("expected BoardBaseURL '%s', got '%s'", expectedBoardURL, client.BoardBaseURL)
		}
	})
}

func TestSetBoard(t *testing.T) {
	client, _ := NewClient("/test-account", "test-token")

	t.Run("sets board URL", func(t *testing.T) {
		client.SetBoard("board-456")
		expectedBoardURL := DefaultBaseURL + "/test-account/boards/board-456"
		if client.BoardBaseURL != expectedBoardURL {
			t.Errorf("expected BoardBaseURL '%s', got '%s'", expectedBoardURL, client.BoardBaseURL)
		}
	})

	t.Run("clears board URL when empty", func(t *testing.T) {
		client.SetBoard("board-456")
		client.SetBoard("")
		if client.BoardBaseURL != "" {
			t.Errorf("expected empty BoardBaseURL, got '%s'", client.BoardBaseURL)
		}
	})
}
