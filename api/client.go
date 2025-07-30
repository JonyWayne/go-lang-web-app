package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webApp/internal/types"
)

// Client структура HTTP-клиента
type Client struct {
	BaseURL        string
	HTTPClient  *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}


// GetIP получает IP-информацию
func (c *Client) GetIP(ctx context.Context) (*types.IPResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("%s/ip", c.BaseURL),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result types.IPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("json decode failed: %w", err)
	}

	return &result, nil
}