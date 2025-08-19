// Code generated for API Clients. DO NOT EDIT.

package ip_restrictions

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v7"
	"github.com/ngrok/ngrok-api-go/v7/internal/api"
)

// An IP restriction is a restriction placed on the CIDRs that are allowed to
//  initiate traffic to a specific aspect of your ngrok account. An IP
//  restriction has a type which defines the ingress it applies to. IP
//  restrictions can be used to enforce the source IPs that can make API
//  requests, log in to the dashboard, start ngrok agents, and connect to your
//  public-facing endpoints.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new IP restriction
//
// https://ngrok.com/docs/api#api-ip-restrictions-create
func (c *Client) Create(ctx context.Context, arg *ngrok.IPRestrictionCreate) (*ngrok.IPRestriction, error) {
	var res ngrok.IPRestriction
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/ip_restrictions")).Execute(&path, arg); err != nil {
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

// Delete an IP restriction
//
// https://ngrok.com/docs/api#api-ip-restrictions-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/ip_restrictions/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get detailed information about an IP restriction
//
// https://ngrok.com/docs/api#api-ip-restrictions-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.IPRestriction, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.IPRestriction
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/ip_restrictions/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all IP restrictions on this account
//
// https://ngrok.com/docs/api#api-ip-restrictions-list
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.IPRestriction] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/ip_restrictions")).Execute(&path, paging); err != nil {
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
	items  []ngrok.IPRestriction
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
	var resp ngrok.IPRestrictionList
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
	if len(resp.IPRestrictions) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.IPRestrictions
	return it.Next(ctx)
}

// Item() returns the IPRestriction currently
// pointed to by the iterator.
func (it *iterList) Item() *ngrok.IPRestriction {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iterList) Err() error {
	return it.err
}

// Update attributes of an IP restriction by ID
//
// https://ngrok.com/docs/api#api-ip-restrictions-update
func (c *Client) Update(ctx context.Context, arg *ngrok.IPRestrictionUpdate) (*ngrok.IPRestriction, error) {
	if arg == nil {
		arg = new(ngrok.IPRestrictionUpdate)
	}
	var res ngrok.IPRestriction
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/ip_restrictions/{{ .ID }}")).Execute(&path, arg); err != nil {
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
