package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdateRequest   = "getUpdates"
	sendMessageRequest = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

// Method for getUpdates from server telegram
func (c *Client) Update(ctx context.Context, offset int, limit int) ([]Updates, error) {
	q := url.Values{}
	q.Set("offset", strconv.Itoa(offset))
	q.Set("limit", strconv.Itoa(limit))

	data, err := c.DoReq(ctx, getUpdateRequest, q)
	if err != nil {
		return nil, err
	}
	var updates UpdatesResponse
	if err := json.Unmarshal(data, &updates); err != nil {
		return nil, err
	}
	return updates.Result, nil
}

// method for sending message
func (c *Client) SendMessage(ctx context.Context, chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)
	_, err := c.DoReq(ctx, sendMessageRequest, q)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DoReq(ctx context.Context, method string, query url.Values) ([]byte, error) {
	// prepare url for request
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	// Prepare request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	// req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	// create request with foolPath
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return data, nil
}
