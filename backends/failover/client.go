// Code generated for API Clients. DO NOT EDIT.

package failover

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/internal/api"
)

// A Failover backend defines failover behavior within a list of referenced
//  backends. Traffic is sent to the first backend in the list. If that backend
//  is offline or no connection can be established, ngrok attempts to connect to
//  the next backend in the list until one is successful.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new Failover backend
//
// https://ngrok.com/docs/api#api-failover-backends-create
func (c *Client) Create(ctx context.Context, arg *ngrok.FailoverBackendCreate) (*ngrok.FailoverBackend, error) {
	if arg == nil {
		arg = new(ngrok.FailoverBackendCreate)
	}
	var res ngrok.FailoverBackend
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/backends/failover")).Execute(&path, arg); err != nil {
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

// Delete a Failover backend by ID.
//
// https://ngrok.com/docs/api#api-failover-backends-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/backends/failover/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get detailed information about a Failover backend by ID
//
// https://ngrok.com/docs/api#api-failover-backends-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.FailoverBackend, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.FailoverBackend
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/backends/failover/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all Failover backends on this account
//
// https://ngrok.com/docs/api#api-failover-backends-list
func (c *Client) list(ctx context.Context, arg *ngrok.Paging) (*ngrok.FailoverBackendList, error) {
	if arg == nil {
		arg = new(ngrok.Paging)
	}
	var res ngrok.FailoverBackendList
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/backends/failover")).Execute(&path, arg); err != nil {
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

// List all Failover backends on this account
//
// https://ngrok.com/docs/api#api-failover-backends-list
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
	items      []ngrok.FailoverBackend
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
	if len(resp.Backends) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.Backends
	return it.Next(ctx)
}

// Item() returns the FailoverBackend currently
// pointed to by the iterator.
func (it *Iter) Item() *ngrok.FailoverBackend {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter) Err() error {
	return it.err
}

// Update Failover backend by ID
//
// https://ngrok.com/docs/api#api-failover-backends-update
func (c *Client) Update(ctx context.Context, arg *ngrok.FailoverBackendUpdate) (*ngrok.FailoverBackend, error) {
	if arg == nil {
		arg = new(ngrok.FailoverBackendUpdate)
	}
	var res ngrok.FailoverBackend
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/backends/failover/{{ .ID }}")).Execute(&path, arg); err != nil {
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
