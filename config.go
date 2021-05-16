package ngrok

import (
	"net/http"
	"net/url"
)

var (
	defaultBaseURL, _ = url.Parse("https://api.ngrok.com")
)

type clientConfig struct {
	baseURL    *url.URL
	httpClient *http.Client
	err        error
}

type ClientOption func(cc *clientConfig)

func (cc *clientConfig) parseOptions(opts []ClientOption) error {
	for _, o := range opts {
		o(cc)
	}
	cc.setDefaults()
	return cc.err
}

func (c *clientConfig) setDefaults() {
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	if c.baseURL == nil {
		c.baseURL = defaultBaseURL
	}
}

func WithBaseURL(baseURL string) ClientOption {
	return func(cc *clientConfig) {
		cc.baseURL, cc.err = url.Parse(baseURL)
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(cc *clientConfig) {
		cc.httpClient = httpClient
	}
}
