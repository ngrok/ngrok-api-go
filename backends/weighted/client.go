// Code generated for API Clients. DO NOT EDIT.

package weighted

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v6"
	"github.com/ngrok/ngrok-api-go/v6/internal/api"
)

// A Weighted Backend balances traffic among the referenced backends. Traffic
//  is assigned proportionally to each based on its weight. The percentage of
//  traffic is calculated by dividing a backend's weight by the sum of all
//  weights.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new Weighted backend
//
// https://ngrok.com/docs/api#api-weighted-backends-create
func (c *Client) Create(ctx context.Context, arg *ngrok.WeightedBackendCreate) (*ngrok.WeightedBackend, error) {
	if arg == nil {
		arg = new(ngrok.WeightedBackendCreate)
	}
	var res ngrok.WeightedBackend
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/backends/weighted")).Execute(&path, arg); err != nil {
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

// Delete a Weighted backend by ID.
//
// https://ngrok.com/docs/api#api-weighted-backends-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/backends/weighted/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get detailed information about a Weighted backend by ID
//
// https://ngrok.com/docs/api#api-weighted-backends-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.WeightedBackend, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.WeightedBackend
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/backends/weighted/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all Weighted backends on this account
//
// https://ngrok.com/docs/api#api-weighted-backends-list
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.WeightedBackend] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/backends/weighted")).Execute(&path, paging); err != nil {
		panic(err)
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
	return &iter{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type iter struct {
	client *Client
	n      int
	items  []ngrok.WeightedBackend
	err    error

	nextPage *url.URL
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *iter) Next(ctx context.Context) bool {
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
	var resp ngrok.WeightedBackendList
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

// Item() returns the WeightedBackend currently
// pointed to by the iterator.
func (it *iter) Item() *ngrok.WeightedBackend {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iter) Err() error {
	return it.err
}

// Update Weighted backend by ID
//
// https://ngrok.com/docs/api#api-weighted-backends-update
func (c *Client) Update(ctx context.Context, arg *ngrok.WeightedBackendUpdate) (*ngrok.WeightedBackend, error) {
	if arg == nil {
		arg = new(ngrok.WeightedBackendUpdate)
	}
	var res ngrok.WeightedBackend
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/backends/weighted/{{ .ID }}")).Execute(&path, arg); err != nil {
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
