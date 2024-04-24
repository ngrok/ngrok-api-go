// Code generated for API Clients. DO NOT EDIT.

package ngrok

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAgent(t *testing.T) {
	url := &url.URL{Path: "/test"}
	apiKey := "testKey"

	testUserAgent := func(cfg *ClientConfig, expected string) {
		client := NewBaseClient(cfg)
		req, err := client.buildRequest(context.TODO(), "GET", url, nil)
		assert.NoError(t, err)
		assert.Equal(t, expected, req.Header.Get("user-agent"))
	}

	testUserAgent(NewClientConfig(apiKey), defaultUserAgent)
	testUserAgent(NewClientConfig(apiKey, WithUserAgent("testAgent")), "testAgent")
}
