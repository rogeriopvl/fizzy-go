package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetBoards(ctx context.Context, opts *ListOptions) ([]Board, error) {
	endpointURL := c.AccountBaseURL + "/boards"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	limit := 0
	if opts != nil {
		limit = opts.Limit
	}

	return fetchAllPages[Board](ctx, c, req, limit)
}

func (c *Client) GetBoard(ctx context.Context, boardID string) (*Board, error) {
	endpointURL := c.AccountBaseURL + "/boards/" + boardID

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var response Board
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateBoard(ctx context.Context, payload CreateBoardPayload) error {
	endpointURL := c.AccountBaseURL + "/boards"

	body := map[string]CreateBoardPayload{"board": payload}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create board request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusCreated)
	return err
}

func (c *Client) UpdateBoard(ctx context.Context, boardID string, payload UpdateBoardPayload) error {
	endpointURL := c.AccountBaseURL + "/boards/" + boardID

	body := map[string]UpdateBoardPayload{"board": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create update board request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteBoard(ctx context.Context, boardID string) error {
	endpointURL := c.AccountBaseURL + "/boards/" + boardID

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete board request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PublishBoard(ctx context.Context, boardID string) (*Board, error) {
	endpointURL := c.AccountBaseURL + "/boards/" + boardID + "/publication"

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create publish board request: %w", err)
	}

	var response Board
	_, err = c.decodeResponse(req, &response, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UnpublishBoard(ctx context.Context, boardID string) error {
	endpointURL := c.AccountBaseURL + "/boards/" + boardID + "/publication"

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create unpublish board request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) UpdateBoardEntropy(ctx context.Context, boardID string, payload EntropyPayload) (*Board, error) {
	endpointURL := c.AccountBaseURL + "/boards/" + boardID + "/entropy"

	body := map[string]EntropyPayload{"board": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create board entropy request: %w", err)
	}

	var response Board
	_, err = c.decodeResponse(req, &response, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
