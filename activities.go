package fizzy

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetActivities(ctx context.Context, filters *ActivityFilters) ([]Activity, error) {
	endpointURL := c.AccountBaseURL + "/activities"

	req, err := c.newRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get activities request: %w", err)
	}

	q := req.URL.Query()
	if filters != nil {
		for _, id := range filters.CreatorIDs {
			q.Add("creator_ids[]", id)
		}
		for _, id := range filters.BoardIDs {
			q.Add("board_ids[]", id)
		}
	}
	req.URL.RawQuery = q.Encode()

	limit := 0
	if filters != nil {
		limit = filters.Limit
	}

	return fetchAllPages[Activity](ctx, c, req, limit)
}
