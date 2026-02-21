package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

// PinCard pins a card for the current user.
func (c *Client) PinCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/pin", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create pin card request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// UnpinCard unpins a card for the current user.
func (c *Client) UnpinCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/pin", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create unpin card request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// GetMyPins returns the current user's pinned cards.
func (c *Client) GetMyPins(ctx context.Context) ([]Card, error) {
	endpointURL := c.BaseURL + "/my/pins"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get my pins request: %w", err)
	}

	var response []Card
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
