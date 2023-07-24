// Code generated for API Clients. DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"

	"github.com/ngrok/ngrok-api-go/v5"
)

const (
	apiVersion = "2"
)

var (
	defaultUserAgent = "ngrok-api-go/" + Version + "/" + runtime.Version()
)

type Client struct {
	cfg *ngrok.ClientConfig
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{cfg: cfg}
}

func (c *Client) Do(ctx context.Context, method string, reqURL *url.URL, reqBody interface{}, respBody interface{}) error {
	req, err := c.buildRequest(ctx, method, reqURL, reqBody)
	if err != nil {
		return err
	}
	resp, err := c.cfg.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if err = c.readResponse(resp, respBody); err != nil {
		return err
	}
	return nil
}

func (c *Client) buildRequest(ctx context.Context, method string, reqURL *url.URL, reqBody interface{}) (*http.Request, error) {
	body, err := c.buildRequestBody(reqBody)
	if err != nil {
		return nil, err
	}
	reqURLString := c.cfg.BaseURL.ResolveReference(reqURL).String()
	r, err := http.NewRequestWithContext(ctx, method, reqURLString, body)
	if err != nil {
		return nil, err
	}

	r.Header.Set("authorization", fmt.Sprintf("Bearer %s", c.cfg.APIKey))
	r.Header.Set("user-agent", c.userAgent())
	r.Header.Set("ngrok-version", apiVersion)
	if body != nil {
		r.Header.Set("content-type", "application/json")
	}
	return r, nil
}

func (c *Client) buildRequestBody(reqBody interface{}) (io.Reader, error) {
	if reqBody == nil {
		return nil, nil
	}
	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonBytes), nil
}

func (c *Client) readResponse(resp *http.Response, out interface{}) error {
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode >= 400 {
		return c.readErrorResponse(resp)
	}
	return c.readResponseBody(resp, out)
}

// read an error response body
func (c *Client) readErrorResponse(resp *http.Response) error {
	var out ngrok.Error
	err := c.readResponseBody(resp, &out)
	if err != nil {
		return err
	}
	return &out
}

// unmarshal a response body
func (c *Client) readResponseBody(resp *http.Response, out interface{}) error {
	if out == nil {
		return nil
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.buildUnmarshalError(resp, bodyBytes, err)
	}
	if err := json.Unmarshal(bodyBytes, out); err != nil {
		return c.buildUnmarshalError(resp, bodyBytes, err)
	}
	return nil
}

// if an error occurs while trying to read a response body, construct a new
// error explaining the unmarshalling failure
func (c *Client) buildUnmarshalError(resp *http.Response, bodyBytes []byte, err error) error {
	return &ngrok.Error{
		Msg:        fmt.Sprintf("failed to unmarshal response body: %s. body: %s", err, bodyBytes),
		StatusCode: int32(resp.StatusCode),
		Details: map[string]string{
			"unmarshal_error": err.Error(),
			"invalid_body":    string(bodyBytes),
			"operation_id":    resp.Header.Get("ngrok-operation-id"),
		},
	}
}

// Returns a user agent override if one was set on the client config. Otherwise,
// returns the default user agent.
func (c *Client) userAgent() string {
	if c.cfg.UserAgent != nil {
		return *c.cfg.UserAgent
	}
	return defaultUserAgent
}
