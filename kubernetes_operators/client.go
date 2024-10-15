// Code generated for API Clients. DO NOT EDIT.

package kubernetes_operators

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v6"
	"github.com/ngrok/ngrok-api-go/v6/internal/api"
)

// KubernetesOperators is used by the Kubernetes Operator to register and
//  manage its own resource, as well as for users to see active kubernetes
//  clusters.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new Kubernetes Operator
//
// https://ngrok.com/docs/api#api-kubernetes-operators-create
func (c *Client) Create(ctx context.Context, arg *ngrok.KubernetesOperatorCreate) (*ngrok.KubernetesOperator, error) {
	if arg == nil {
		arg = new(ngrok.KubernetesOperatorCreate)
	}
	var res ngrok.KubernetesOperator
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/kubernetes_operators")).Execute(&path, arg); err != nil {
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

// Update an existing Kubernetes operator by ID.
//
// https://ngrok.com/docs/api#api-kubernetes-operators-update
func (c *Client) Update(ctx context.Context, arg *ngrok.KubernetesOperatorUpdate) (*ngrok.KubernetesOperator, error) {
	if arg == nil {
		arg = new(ngrok.KubernetesOperatorUpdate)
	}
	var res ngrok.KubernetesOperator
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/kubernetes_operators/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Delete a Kubernetes Operator
//
// https://ngrok.com/docs/api#api-kubernetes-operators-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/kubernetes_operators/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get of a Kubernetes Operator
//
// https://ngrok.com/docs/api#api-kubernetes-operators-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.KubernetesOperator, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.KubernetesOperator
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/kubernetes_operators/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all Kubernetes Operators owned by this account
//
// https://ngrok.com/docs/api#api-kubernetes-operators-list
func (c *Client) List(paging *ngrok.Paging) *Iter {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/kubernetes_operators")).Execute(&path, paging); err != nil {
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
	return &Iter{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// Iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type Iter struct {
	client *Client
	n      int
	items  []ngrok.KubernetesOperator
	err    error

	nextPage *url.URL
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
		return true
	}

	// no more items, do we have a next page?
	if it.nextPage == nil {
		return false
	}

	// fetch the next page
	var resp ngrok.KubernetesOperatorList
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
	if len(resp.Operators) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.Operators
	return it.Next(ctx)
}

// Item() returns the KubernetesOperator currently
// pointed to by the iterator.
func (it *Iter) Item() *ngrok.KubernetesOperator {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter) Err() error {
	return it.err
}
