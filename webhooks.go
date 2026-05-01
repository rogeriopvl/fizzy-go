package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetWebhooks(ctx context.Context, boardID string, opts *ListOptions) ([]Webhook, error) {
	endpointURL := fmt.Sprintf("%s/boards/%s/webhooks", c.AccountBaseURL, boardID)

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get webhooks request: %w", err)
	}

	limit := 0
	if opts != nil {
		limit = opts.Limit
	}

	return fetchAllPages[Webhook](ctx, c, req, limit)
}

func (c *Client) GetWebhook(ctx context.Context, boardID, webhookID string) (*Webhook, error) {
	endpointURL := fmt.Sprintf("%s/boards/%s/webhooks/%s", c.AccountBaseURL, boardID, webhookID)

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get webhook request: %w", err)
	}

	var response Webhook
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateWebhook(ctx context.Context, boardID string, payload CreateWebhookPayload) (*Webhook, error) {
	endpointURL := fmt.Sprintf("%s/boards/%s/webhooks", c.AccountBaseURL, boardID)

	body := map[string]CreateWebhookPayload{"webhook": payload}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook request: %w", err)
	}

	var response Webhook
	_, err = c.decodeResponse(req, &response, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateWebhook(ctx context.Context, boardID, webhookID string, payload UpdateWebhookPayload) (*Webhook, error) {
	endpointURL := fmt.Sprintf("%s/boards/%s/webhooks/%s", c.AccountBaseURL, boardID, webhookID)

	body := map[string]UpdateWebhookPayload{"webhook": payload}

	req, err := c.newRequest(ctx, http.MethodPatch, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create update webhook request: %w", err)
	}

	var response Webhook
	_, err = c.decodeResponse(req, &response, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteWebhook(ctx context.Context, boardID, webhookID string) error {
	endpointURL := fmt.Sprintf("%s/boards/%s/webhooks/%s", c.AccountBaseURL, boardID, webhookID)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete webhook request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) GetWebhookDeliveries(ctx context.Context, boardID, webhookID string, opts *ListOptions) ([]WebhookDelivery, error) {
	endpointURL := fmt.Sprintf("%s/boards/%s/webhooks/%s/deliveries", c.AccountBaseURL, boardID, webhookID)

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get webhook deliveries request: %w", err)
	}

	limit := 0
	if opts != nil {
		limit = opts.Limit
	}

	return fetchAllPages[WebhookDelivery](ctx, c, req, limit)
}

func (c *Client) ActivateWebhook(ctx context.Context, boardID, webhookID string) (*Webhook, error) {
	endpointURL := fmt.Sprintf("%s/boards/%s/webhooks/%s/activation", c.AccountBaseURL, boardID, webhookID)

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create activate webhook request: %w", err)
	}

	var response Webhook
	_, err = c.decodeResponse(req, &response, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
