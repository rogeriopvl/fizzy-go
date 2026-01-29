package fizzy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// ErrNoBoardSelected is returned when an operation requires a board but none is set.
var ErrNoBoardSelected = errors.New("no board selected: use SetBoard or WithBoard when creating the client")

func (c *Client) GetCards(ctx context.Context, filters CardFilters) ([]Card, error) {
	endpointURL := c.AccountBaseURL + "/cards"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get cards request: %w", err)
	}

	q := req.URL.Query()

	for _, id := range filters.BoardIDs {
		q.Add("board_ids[]", id)
	}
	for _, id := range filters.TagIDs {
		q.Add("tag_ids[]", id)
	}
	for _, id := range filters.AssigneeIDs {
		q.Add("assignee_ids[]", id)
	}
	for _, id := range filters.CreatorIDs {
		q.Add("creator_ids[]", id)
	}
	for _, id := range filters.CloserIDs {
		q.Add("closer_ids[]", id)
	}
	for _, id := range filters.CardIDs {
		q.Add("card_ids[]", id)
	}
	for _, term := range filters.Terms {
		q.Add("terms[]", term)
	}
	if filters.IndexedBy != "" {
		q.Set("indexed_by", filters.IndexedBy)
	}
	if filters.SortedBy != "" {
		q.Set("sorted_by", filters.SortedBy)
	}
	if filters.AssignmentStatus != "" {
		q.Set("assignment_status", filters.AssignmentStatus)
	}
	if filters.CreationStatus != "" {
		q.Set("creation", filters.CreationStatus)
	}
	if filters.ClosureStatus != "" {
		q.Set("closure", filters.ClosureStatus)
	}

	req.URL.RawQuery = q.Encode()

	var response []Card
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetCard(ctx context.Context, cardNumber int) (*Card, error) {
	endpointURL := fmt.Sprintf("%s/cards/%d", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get card by id request: %w", err)
	}

	var response Card
	_, err = c.decodeResponse(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateCard(ctx context.Context, payload CreateCardPayload) error {
	if c.BoardBaseURL == "" {
		return ErrNoBoardSelected
	}

	endpointURL := c.BoardBaseURL + "/cards"

	body := map[string]CreateCardPayload{"card": payload}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create card request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusCreated)
	return err
}

func (c *Client) UpdateCard(ctx context.Context, cardNumber int, payload UpdateCardPayload) (*Card, error) {
	endpointURL := fmt.Sprintf("%s/cards/%d", c.AccountBaseURL, cardNumber)

	body := map[string]UpdateCardPayload{"card": payload}

	req, err := c.newRequest(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create update card request: %w", err)
	}

	var response Card
	_, err = c.decodeResponse(req, &response, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete card request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) DeleteCardImage(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/image", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete card image request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) CloseCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/closure", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create closure card request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) ReopenCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/closure", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete closure card request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// PostponeCard moves a card to the "Not Now" status.
func (c *Client) PostponeCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/not_now", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create post not now request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// TriageCard moves a card from triage into a column.
func (c *Client) TriageCard(ctx context.Context, cardNumber int, columnID string) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/triage", c.AccountBaseURL, cardNumber)

	body := map[string]any{"column_id": columnID}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create post triage request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// UnTriageCard moves a card back to triage (removes it from its column).
func (c *Client) UnTriageCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/triage", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete triage request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) WatchCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/watch", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create post watch request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) UnwatchCard(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/watch", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete watch request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// MarkCardGolden marks a card as golden (starred/highlighted).
func (c *Client) MarkCardGolden(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/goldness", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create post goldness request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

func (c *Client) UnmarkCardGolden(ctx context.Context, cardNumber int) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/goldness", c.AccountBaseURL, cardNumber)

	req, err := c.newRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete goldness request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// AssignCard toggles assignment of a user to/from a card.
func (c *Client) AssignCard(ctx context.Context, cardNumber int, userID string) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/assignments", c.AccountBaseURL, cardNumber)

	body := map[string]string{"assignee_id": userID}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create assignment request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}

// TagCard toggles a tag on/off for a card. Creates the tag if it doesn't exist.
func (c *Client) TagCard(ctx context.Context, cardNumber int, tagTitle string) error {
	endpointURL := fmt.Sprintf("%s/cards/%d/taggings", c.AccountBaseURL, cardNumber)

	body := map[string]string{"tag_title": tagTitle}

	req, err := c.newRequest(ctx, http.MethodPost, endpointURL, body)
	if err != nil {
		return fmt.Errorf("failed to create tagging request: %w", err)
	}

	_, err = c.decodeResponse(req, nil, http.StatusNoContent)
	return err
}
