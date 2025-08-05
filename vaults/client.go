// Code generated for API Clients. DO NOT EDIT.

package vaults

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v7"
	"github.com/ngrok/ngrok-api-go/v7/internal/api"
)

// Vaults is an api service for securely storing and managing sensitive data such
// as secrets, credentials, and tokens.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new Vault
//
// https://ngrok.com/docs/api#api-vaults-create
func (c *Client) Create(ctx context.Context, arg *ngrok.VaultCreate) (*ngrok.Vault, error) {
	if arg == nil {
		arg = new(ngrok.VaultCreate)
	}
	var res ngrok.Vault
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/vaults")).Execute(&path, arg); err != nil {
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

// Update an existing Vault by ID
//
// https://ngrok.com/docs/api#api-vaults-update
func (c *Client) Update(ctx context.Context, arg *ngrok.VaultUpdate) (*ngrok.Vault, error) {
	if arg == nil {
		arg = new(ngrok.VaultUpdate)
	}
	var res ngrok.Vault
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/vaults/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Delete a Vault
//
// https://ngrok.com/docs/api#api-vaults-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/vaults/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get a Vault by ID
//
// https://ngrok.com/docs/api#api-vaults-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.Vault, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.Vault
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/vaults/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all Vaults owned by account
//
// https://ngrok.com/docs/api#api-vaults-list
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.Vault] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/vaults")).Execute(&path, paging); err != nil {
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
	return &iterVault{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type iterVault struct {
	client *Client
	n      int
	items  []ngrok.Vault
	err    error

	nextPage *url.URL
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *iterVault) Next(ctx context.Context) bool {
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
	var resp ngrok.VaultList
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
	if len(resp.Vaults) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.Vaults
	return it.Next(ctx)
}

// Item() returns the Vault currently
// pointed to by the iterator.
func (it *iterVault) Item() *ngrok.Vault {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iterVault) Err() error {
	return it.err
}
