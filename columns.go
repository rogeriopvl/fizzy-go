package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetColumns(ctx context.Context) ([]Column, error) {
	if c.BoardBaseURL == "" {
		return nil, ErrNoBoardSelected
	}

	endpointURL := c.BoardBaseURL + "/columns"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get columns request: %w", err)
	}

	var response []Column
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetColumn(ctx context.Context, columnID string) (*Column, error) {
	if c.BoardBaseURL == "" {
		return nil, ErrNoBoardSelected
	}

	endpointURL := c.BoardBaseURL + "/columns/" + columnID

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var response Column
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateColumn(ctx context.Context, payload CreateColumnPayload) error {
	if c.BoardBaseURL == "" {
		return ErrNoBoardSelected
	}

	endpointURL := c.BoardBaseURL + "/columns"

	body := map[string]CreateColumnPayload{"column": payload}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create column request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusCreated)
	return err
}

func (c *Client) UpdateColumn(ctx context.Context, columnID string, payload UpdateColumnPayload) error {
	if c.BoardBaseURL == "" {
		return ErrNoBoardSelected
	}

	endpointURL := c.BoardBaseURL + "/columns/" + columnID

	body := map[string]UpdateColumnPayload{"column": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create update column request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteColumn(ctx context.Context, columnID string) error {
	if c.BoardBaseURL == "" {
		return ErrNoBoardSelected
	}

	endpointURL := c.BoardBaseURL + "/columns/" + columnID

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete column request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}
