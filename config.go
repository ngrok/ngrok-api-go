package ngrok

import (
	"net/http"
	"net/url"
	"os"
)

var (
	defaultBaseURL, _ = url.Parse("https://api.ngrok.com")
)

type clientConfig struct {
	apiKey     string
	baseURL    *url.URL
	debug      Debug
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
	if c.apiKey == "" {
		c.apiKey = os.Getenv("NGROK_API_KEY")
	}
}

func WithAPIKey(apiKey string) ClientOption {
	return func(cc *clientConfig) {
		cc.apiKey = apiKey
	}
}

func WithBaseURL(baseURL string) ClientOption {
	return func(cc *clientConfig) {
		cc.baseURL, cc.err = url.Parse(baseURL)
	}
}

func WithDebug(debug Debug) ClientOption {
	return func(cc *clientConfig) {
		cc.debug = debug
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(cc *clientConfig) {
		cc.httpClient = httpClient
	}
}
