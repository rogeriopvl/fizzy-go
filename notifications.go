package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetNotifications(ctx context.Context, opts *ListOptions) ([]Notification, error) {
	endpointURL := c.AccountBaseURL + "/notifications"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get notifications request: %w", err)
	}

	limit := 0
	if opts != nil {
		limit = opts.Limit
	}

	return fetchAllPages[Notification](ctx, c, req, limit)
}

func (c *Client) GetNotification(ctx context.Context, notificationID string) (*Notification, error) {
	endpointURL := fmt.Sprintf("%s/notifications/%s", c.AccountBaseURL, notificationID)

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get notification request: %w", err)
	}

	var response Notification
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) MarkNotificationRead(ctx context.Context, notificationID string) error {
	endpointURL := fmt.Sprintf("%s/notifications/%s/reading", c.AccountBaseURL, notificationID)

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create mark notification as read request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) MarkNotificationUnread(ctx context.Context, notificationID string) error {
	endpointURL := fmt.Sprintf("%s/notifications/%s/reading", c.AccountBaseURL, notificationID)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete notification request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) MarkAllNotificationsRead(ctx context.Context) error {
	endpointURL := c.AccountBaseURL + "/notifications/bulk_reading"

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create bulk notifications reading request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) GetNotificationSettings(ctx context.Context) (*NotificationSettings, error) {
	endpointURL := c.AccountBaseURL + "/notifications/settings"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get notification settings request: %w", err)
	}

	var response NotificationSettings
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateNotificationSettings(ctx context.Context, payload UpdateNotificationSettingsPayload) error {
	endpointURL := c.AccountBaseURL + "/notifications/settings"

	body := map[string]UpdateNotificationSettingsPayload{"user_settings": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create update notification settings request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}
