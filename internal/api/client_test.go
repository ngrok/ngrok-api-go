// Code generated for API Clients. DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"testing"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/stretchr/testify/assert"
)

func TestUserAgent(t *testing.T) {
	url := &url.URL{Path: "/test"}
	apiKey := "testKey"

	testUserAgent := func(cfg *ngrok.ClientConfig, expected string) {
		client := NewClient(cfg)
		req, err := client.buildRequest(context.TODO(), "GET", url, nil)
		assert.NoError(t, err)
		assert.Equal(t, expected, req.Header.Get("user-agent"))
	}

	testUserAgent(ngrok.NewClientConfig(apiKey), defaultUserAgent)
	testUserAgent(ngrok.NewClientConfig(apiKey, ngrok.WithUserAgent("testAgent")), "testAgent")
}
