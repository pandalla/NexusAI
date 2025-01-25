package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"nexus-ai/kling/auth"
	"nexus-ai/kling/config"
)

type Client struct {
	httpClient *http.Client
	cfg        *config.KlingAIConfig
}

func NewClient(cfg *config.KlingAIConfig) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		cfg: cfg,
	}
}

func (c *Client) CreateRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	token, err := auth.GenerateToken(c.cfg.AccessKey, c.cfg.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT: %w", err)
	}

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := c.cfg.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	return req, nil
}

func (c *Client) DoRequest(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed: %s, body: %s", resp.Status, body)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
