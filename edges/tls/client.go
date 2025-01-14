// Code generated for API Clients. DO NOT EDIT.

package tls

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v7"
	"github.com/ngrok/ngrok-api-go/v7/internal/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a TLS Edge
//
// https://ngrok.com/docs/api#api-edges-tls-create
func (c *Client) Create(ctx context.Context, arg *ngrok.TLSEdgeCreate) (*ngrok.TLSEdge, error) {
	if arg == nil {
		arg = new(ngrok.TLSEdgeCreate)
	}
	var res ngrok.TLSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/edges/tls")).Execute(&path, arg); err != nil {
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

// Get a TLS Edge by ID
//
// https://ngrok.com/docs/api#api-edges-tls-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.TLSEdge, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.TLSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/edges/tls/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Returns a list of all TLS Edges on this account
//
// https://ngrok.com/docs/api#api-edges-tls-list
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.TLSEdge] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/edges/tls")).Execute(&path, paging); err != nil {
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
	return &iterTLSEdge{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type iterTLSEdge struct {
	client *Client
	n      int
	items  []ngrok.TLSEdge
	err    error

	nextPage *url.URL
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *iterTLSEdge) Next(ctx context.Context) bool {
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
	var resp ngrok.TLSEdgeList
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
	if len(resp.TLSEdges) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.TLSEdges
	return it.Next(ctx)
}

// Item() returns the TLSEdge currently
// pointed to by the iterator.
func (it *iterTLSEdge) Item() *ngrok.TLSEdge {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iterTLSEdge) Err() error {
	return it.err
}

// Updates a TLS Edge by ID. If a module is not specified in the update, it will
// not be modified. However, each module configuration that is specified will
// completely replace the existing value. There is no way to delete an existing
// module via this API, instead use the delete module API.
//
// https://ngrok.com/docs/api#api-edges-tls-update
func (c *Client) Update(ctx context.Context, arg *ngrok.TLSEdgeUpdate) (*ngrok.TLSEdge, error) {
	if arg == nil {
		arg = new(ngrok.TLSEdgeUpdate)
	}
	var res ngrok.TLSEdge
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/edges/tls/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Delete a TLS Edge by ID
//
// https://ngrok.com/docs/api#api-edges-tls-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/edges/tls/{{ .ID }}")).Execute(&path, arg); err != nil {
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
