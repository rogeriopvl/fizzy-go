package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetAccountSettings(ctx context.Context) (*Account, error) {
	endpointURL := c.BaseURL + "/account/settings"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get account settings request: %w", err)
	}

	var response Account
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateAccountEntropy(ctx context.Context, payload EntropyPayload) (*Account, error) {
	endpointURL := c.BaseURL + "/account/entropy"

	body := map[string]EntropyPayload{"entropy": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create account entropy request: %w", err)
	}

	var response Account
	_, err = c.decodeResponse(req, &response, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
