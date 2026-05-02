package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetAccountSettings(ctx context.Context) (*Account, error) {
	endpointURL := c.AccountBaseURL + "/account/settings"

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
	endpointURL := c.AccountBaseURL + "/account/entropy"

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

func (c *Client) GetAccountJoinCode(ctx context.Context) (*JoinCode, error) {
	endpointURL := c.AccountBaseURL + "/account/join_code"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get account join code request: %w", err)
	}

	var response JoinCode
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateAccountJoinCode(ctx context.Context, payload UpdateJoinCodePayload) error {
	endpointURL := c.AccountBaseURL + "/account/join_code"

	body := map[string]UpdateJoinCodePayload{"account_join_code": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create update account join code request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) ResetAccountJoinCode(ctx context.Context) error {
	endpointURL := c.AccountBaseURL + "/account/join_code"

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create reset account join code request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}
