package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetUsers(ctx context.Context) ([]User, error) {
	endpointURL := c.AccountBaseURL + "/users"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get users request: %w", err)
	}

	var response []User
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	endpointURL := c.AccountBaseURL + "/users/" + userID

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get user request: %w", err)
	}

	var response User
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type UpdateUserPayload struct {
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

func (c *Client) UpdateUser(ctx context.Context, userID string, payload UpdateUserPayload) error {
	endpointURL := c.AccountBaseURL + "/users/" + userID

	body := map[string]UpdateUserPayload{"user": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create update user request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeactivateUser(ctx context.Context, userID string) error {
	endpointURL := c.AccountBaseURL + "/users/" + userID

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create deactivate user request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteUserAvatar(ctx context.Context, userID string) error {
	endpointURL := c.AccountBaseURL + "/users/" + userID + "/avatar"

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete user avatar request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

type RequestEmailChangePayload struct {
	EmailAddress string `json:"email_address"`
}

func (c *Client) RequestUserEmailChange(ctx context.Context, userID string, payload RequestEmailChangePayload) error {
	endpointURL := c.AccountBaseURL + "/users/" + userID + "/email_addresses"

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, payload)
	if err != nil {
		return fmt.Errorf("failed to create request user email change request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusCreated)
	return err
}

func (c *Client) ConfirmUserEmailChange(ctx context.Context, userID, token string) error {
	endpointURL := c.AccountBaseURL + "/users/" + userID + "/email_addresses/" + token + "/confirmation"

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create confirm user email change request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}
