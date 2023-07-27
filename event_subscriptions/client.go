// Code generated for API Clients. DO NOT EDIT.

package event_subscriptions

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/internal/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create an Event Subscription.
//
// https://ngrok.com/docs/api#api-event-subscriptions-create
func (c *Client) Create(ctx context.Context, arg *ngrok.EventSubscriptionCreate) (*ngrok.EventSubscription, error) {
	if arg == nil {
		arg = new(ngrok.EventSubscriptionCreate)
	}
	var res ngrok.EventSubscription
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/event_subscriptions")).Execute(&path, arg); err != nil {
		panic(err)
	}
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

// Delete an Event Subscription.
//
// https://ngrok.com/docs/api#api-event-subscriptions-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/event_subscriptions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
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

// Get an Event Subscription by ID.
//
// https://ngrok.com/docs/api#api-event-subscriptions-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.EventSubscription, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.EventSubscription
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/event_subscriptions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
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

// List this Account's Event Subscriptions.
//
// https://ngrok.com/docs/api#api-event-subscriptions-list
func (c *Client) list(ctx context.Context, arg *ngrok.Paging) (*ngrok.EventSubscriptionList, error) {
	if arg == nil {
		arg = new(ngrok.Paging)
	}
	var res ngrok.EventSubscriptionList
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/event_subscriptions")).Execute(&path, arg); err != nil {
		panic(err)
	}
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	queryVals := make(url.Values)
	if arg.BeforeID != nil {
		queryVals.Set("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		queryVals.Set("limit", *arg.Limit)
	}
	apiURL.RawQuery = queryVals.Encode()

	if err := c.apiClient.Do(ctx, "GET", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// List this Account's Event Subscriptions.
//
// https://ngrok.com/docs/api#api-event-subscriptions-list
func (c *Client) List(paging *ngrok.Paging) *Iter {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	if paging.Limit == nil {
		paging.Limit = ngrok.String("100")
	}
	return &Iter{
		client:     c,
		limit:      paging.Limit,
		lastItemID: paging.BeforeID,
		n:          -1,
	}
}

// Iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type Iter struct {
	client     *Client
	n          int
	items      []ngrok.EventSubscription
	err        error
	limit      *string
	lastItemID *string
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *Iter) Next(ctx context.Context) bool {
	// no more if there is an error
	if it.err != nil {
		return false
	}

	// advance the iterator
	it.n += 1

	// is there an available item?
	if it.n < len(it.items) {
		it.lastItemID = ngrok.String(it.Item().ID)
		return true
	}

	// fetch the next page
	resp, err := it.client.list(ctx, &ngrok.Paging{
		BeforeID: it.lastItemID,
		Limit:    it.limit,
	})
	if err != nil {
		it.err = err
		return false
	}

	// page with zero items means there are no more
	if len(resp.EventSubscriptions) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.EventSubscriptions
	return it.Next(ctx)
}

// Item() returns the EventSubscription currently
// pointed to by the iterator.
func (it *Iter) Item() *ngrok.EventSubscription {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter) Err() error {
	return it.err
}

// Update an Event Subscription.
//
// https://ngrok.com/docs/api#api-event-subscriptions-update
func (c *Client) Update(ctx context.Context, arg *ngrok.EventSubscriptionUpdate) (*ngrok.EventSubscription, error) {
	if arg == nil {
		arg = new(ngrok.EventSubscriptionUpdate)
	}
	var res ngrok.EventSubscription
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/event_subscriptions/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
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
