// Code generated for API Clients. DO NOT EDIT.

package event_sources

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v6"
	"github.com/ngrok/ngrok-api-go/v6/internal/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Add an additional type for which this event subscription will trigger
//
// https://ngrok.com/docs/api#api-event-sources-create
func (c *Client) Create(ctx context.Context, arg *ngrok.EventSourceCreate) (*ngrok.EventSource, error) {
	if arg == nil {
		arg = new(ngrok.EventSourceCreate)
	}
	var res ngrok.EventSource
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.SubscriptionID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "POST", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Remove a type for which this event subscription will trigger
//
// https://ngrok.com/docs/api#api-event-sources-delete
func (c *Client) Delete(ctx context.Context, arg *ngrok.EventSourceItem) error {
	if arg == nil {
		arg = new(ngrok.EventSourceItem)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources/{{ .Type }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.SubscriptionID = ""
	arg.Type = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "DELETE", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}

// Get the details for a given type that triggers for the given event subscription
//
// https://ngrok.com/docs/api#api-event-sources-get
func (c *Client) Get(ctx context.Context, arg *ngrok.EventSourceItem) (*ngrok.EventSource, error) {
	if arg == nil {
		arg = new(ngrok.EventSourceItem)
	}
	var res ngrok.EventSource
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources/{{ .Type }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.SubscriptionID = ""
	arg.Type = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "GET", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// List the types for which this event subscription will trigger
//
// https://ngrok.com/docs/api#api-event-sources-list
func (c *Client) List(ctx context.Context, subscriptionId string) (*ngrok.EventSourceList, error) {
	arg := &ngrok.EventSourcePaging{SubscriptionID: subscriptionId}

	var res ngrok.EventSourceList
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.SubscriptionID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "GET", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Update the type for which this event subscription will trigger
//
// https://ngrok.com/docs/api#api-event-sources-update
func (c *Client) Update(ctx context.Context, arg *ngrok.EventSourceUpdate) (*ngrok.EventSource, error) {
	if arg == nil {
		arg = new(ngrok.EventSourceUpdate)
	}
	var res ngrok.EventSource
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/event_subscriptions/{{ .SubscriptionID }}/sources/{{ .Type }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.SubscriptionID = ""
	arg.Type = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "PATCH", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
