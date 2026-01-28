// Package fizzy provides a Go client for the Fizzy API.
package fizzy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultBaseURL = "https://app.fizzy.do"
	DefaultTimeout = 30 * time.Second
)

type Client struct {
	BaseURL        string
	AccountBaseURL string
	BoardBaseURL   string
	AccessToken    string
	HTTPClient     *http.Client
	boardID        string
}

type ClientOption func(*Client)

func WithBoard(boardID string) ClientOption {
	return func(c *Client) {
		c.boardID = boardID
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// WithBaseURL overrides the default API base URL (https://app.fizzy.do).
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

// NewClient creates a new Fizzy API client.
// The accountSlug should include the leading slash (e.g., "/123456").
func NewClient(accountSlug string, accessToken string, opts ...ClientOption) (*Client, error) {
	if accountSlug == "" {
		return nil, fmt.Errorf("accountSlug is required")
	}
	if accessToken == "" {
		return nil, fmt.Errorf("accessToken is required")
	}

	c := &Client{
		BaseURL:     DefaultBaseURL,
		AccessToken: accessToken,
		HTTPClient:  &http.Client{Timeout: DefaultTimeout},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.AccountBaseURL = c.BaseURL + accountSlug
	if c.boardID != "" {
		c.BoardBaseURL = c.AccountBaseURL + "/boards/" + c.boardID
	}

	return c, nil
}

func (c *Client) SetBoard(boardID string) {
	c.boardID = boardID
	if boardID != "" {
		c.BoardBaseURL = c.AccountBaseURL + "/boards/" + boardID
	} else {
		c.BoardBaseURL = ""
	}
}

func (c *Client) newRequest(ctx context.Context, method, url string, body any) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) decodeResponse(req *http.Request, v any, expectedStatus ...int) (int, error) {
	expectedCode := http.StatusOK
	if len(expectedStatus) > 0 {
		expectedCode = expectedStatus[0]
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != expectedCode {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return 0, fmt.Errorf("unexpected status code %d (failed to read error response: %w)", res.StatusCode, err)
		}
		return 0, fmt.Errorf("unexpected status code %d: %s", res.StatusCode, string(body))
	}

	if v != nil {
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return 0, fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return res.StatusCode, nil
}
