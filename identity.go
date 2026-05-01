package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetMyIdentity(ctx context.Context) (*GetMyIdentityResponse, error) {
	endpointURL := c.BaseURL + "/my/identity"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var response GetMyIdentityResponse
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateAccessToken(ctx context.Context, payload CreateAccessTokenPayload) (*PersonalAccessToken, error) {
	endpointURL := c.AccountBaseURL + "/my/access_tokens"

	body := map[string]CreateAccessTokenPayload{"access_token": payload}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token request: %w", err)
	}

	var response PersonalAccessToken
	_, err = c.decodeResponse(req, &response, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateMyTimezone(ctx context.Context, timezoneName string) error {
	endpointURL := c.BaseURL + "/my/timezone"

	body := map[string]string{"timezone_name": timezoneName}

	req, err := c.newRequest(ctx, http.MethodPatch, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create update timezone request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// DeleteSession logs out the current user by deleting the session.
func (c *Client) DeleteSession(ctx context.Context) error {
	endpointURL := c.BaseURL + "/session"

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}
