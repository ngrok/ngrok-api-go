// Code generated for API Clients. DO NOT EDIT.

package failover

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v7"
	"github.com/ngrok/ngrok-api-go/v7/internal/api"
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
		return nil, fmt.Errorf("error building path for create: %w", err)
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
		return fmt.Errorf("error building path for delete: %w", err)
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
		return nil, fmt.Errorf("error building path for get: %w", err)
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
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.FailoverBackend] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/backends/failover")).Execute(&path, paging); err != nil {
		return &iterList{err: fmt.Errorf("error building path for list: %w", err)}
	}
	var apiURL = &url.URL{Path: path.String()}
	queryVals := make(url.Values)
	if paging.BeforeID != nil {
		queryVals.Set("before_id", *paging.BeforeID)
	}
	if paging.Limit != nil {
		queryVals.Set("limit", *paging.Limit)
	}
	apiURL.RawQuery = queryVals.Encode()
	return &iterList{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type iterList struct {
	client *Client
	n      int
	items  []ngrok.FailoverBackend
	err    error

	nextPage *url.URL
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *iterList) Next(ctx context.Context) bool {
	// no more if there is an error
	if it.err != nil {
		return false
	}

	// advance the iterator
	it.n += 1

	// is there an available item?
	if it.n < len(it.items) {
		return true
	}

	// no more items, do we have a next page?
	if it.nextPage == nil {
		return false
	}

	// fetch the next page
	var resp ngrok.FailoverBackendList
	err := it.client.apiClient.Do(ctx, "GET", it.nextPage, nil, &resp)
	if err != nil {
		it.err = err
		return false
	}

	// parse the next page URI as soon as we get it and store it
	// so we can use it on the next fetch
	if resp.NextPageURI != nil {
		it.nextPage, it.err = url.Parse(*resp.NextPageURI)
		if it.err != nil {
			return false
		}
	} else {
		it.nextPage = nil
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
func (it *iterList) Item() *ngrok.FailoverBackend {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iterList) Err() error {
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
		return nil, fmt.Errorf("error building path for update: %w", err)
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
