package fizzy

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
)

// ListOptions specifies pagination options for list operations.
type ListOptions struct {
	// Limit is the maximum number of items to return. 0 means no limit.
	Limit int
}

// Response wraps an HTTP response with pagination info.
type Response struct {
	StatusCode int
	NextURL    string // URL for next page, empty if no more pages
}

// parseNextLink extracts the "next" URL from a Link header.
func parseNextLink(linkHeader string) string {
	if linkHeader == "" {
		return ""
	}

	// Match pattern: <URL>; rel="next"
	re := regexp.MustCompile(`<([^>]+)>;\s*rel="next"`)
	matches := re.FindStringSubmatch(linkHeader)
	if len(matches) >= 2 {
		return matches[1]
	}

	// Also try without quotes: <URL>; rel=next
	re = regexp.MustCompile(`<([^>]+)>;\s*rel=next`)
	matches = re.FindStringSubmatch(linkHeader)
	if len(matches) >= 2 {
		return matches[1]
	}

	return ""
}

// fetchAllPages handles pagination by iterating through all pages and returning all items.
func fetchAllPages[T any](ctx context.Context, client *Client, req *http.Request, limit int) ([]T, error) {
	var allItems []T

	for {
		var pageItems []T
		resp, err := client.decodeResponseWithPagination(req, &pageItems)
		if err != nil {
			return nil, err
		}

		allItems = append(allItems, pageItems...)

		// Check if we've reached the limit
		if limit > 0 && len(allItems) >= limit {
			allItems = allItems[:limit]
			break
		}

		// Check if there are more pages
		if resp.NextURL == "" {
			break
		}

		// Create request for next page
		req, err = client.newRequest(ctx, http.MethodGet, resp.NextURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create next page request: %w", err)
		}
	}

	return allItems, nil
}
