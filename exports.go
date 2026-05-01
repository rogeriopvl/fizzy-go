package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) CreateAccountExport(ctx context.Context) (*Export, error) {
	endpointURL := c.BaseURL + "/account/exports"

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create account export request: %w", err)
	}

	var response Export
	_, err = c.decodeResponse(req, &response, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAccountExport(ctx context.Context, exportID string) (*Export, error) {
	endpointURL := c.BaseURL + "/account/exports/" + exportID

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get account export request: %w", err)
	}

	var response Export
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateUserDataExport(ctx context.Context, userID string) (*Export, error) {
	endpointURL := c.AccountBaseURL + "/users/" + userID + "/data_exports"

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create user data export request: %w", err)
	}

	var response Export
	_, err = c.decodeResponse(req, &response, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetUserDataExport(ctx context.Context, userID, exportID string) (*Export, error) {
	endpointURL := c.AccountBaseURL + "/users/" + userID + "/data_exports/" + exportID

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get user data export request: %w", err)
	}

	var response Export
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
