// Code generated for API Clients. DO NOT EDIT.

package ngrok

import (
	"net/http"
	"net/url"
)

var (
	defaultBaseURL, _ = url.Parse("https://api.ngrok.com")
)

type ClientConfig struct {
	APIKey     string
	BaseURL    *url.URL
	HTTPClient *http.Client
	UserAgent  *string
}

func NewClientConfig(apiKey string, opts ...ClientConfigOption) *ClientConfig {
	var configOpts clientConfigOpts
	configOpts.parseOptions(opts)
	return &ClientConfig{
		APIKey:     apiKey,
		BaseURL:    configOpts.baseURL,
		HTTPClient: configOpts.httpClient,
		UserAgent:  configOpts.userAgent,
	}
}

type clientConfigOpts struct {
	baseURL    *url.URL
	httpClient *http.Client
	userAgent  *string
}

type ClientConfigOption func(cc *clientConfigOpts)

func (cc *clientConfigOpts) parseOptions(opts []ClientConfigOption) {
	for _, o := range opts {
		o(cc)
	}
	cc.setDefaults()
}

func (cc *clientConfigOpts) setDefaults() {
	if cc.httpClient == nil {
		cc.httpClient = http.DefaultClient
	}
	if cc.baseURL == nil {
		cc.baseURL = defaultBaseURL
	}
}

func WithBaseURL(baseURL string) ClientConfigOption {
	return func(cc *clientConfigOpts) {
		cc.baseURL, _ = url.Parse(baseURL)
	}
}

func WithHTTPClient(httpClient *http.Client) ClientConfigOption {
	return func(cc *clientConfigOpts) {
		cc.httpClient = httpClient
	}
}

func WithUserAgent(userAgent string) ClientConfigOption {
	return func(cc *clientConfigOpts) {
		cc.userAgent = &userAgent
	}
}
