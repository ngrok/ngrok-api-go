// Code generated for API Clients. DO NOT EDIT.

package ip_policies

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/internal/api"
)

// IP Policies are reusable groups of CIDR ranges with an allow or deny
//  action. They can be attached to endpoints via the Endpoint Configuration IP
//  Policy module. They can also be used with IP Restrictions to control source
//  IP ranges that can start tunnel sessions and connect to the API and dashboard.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new IP policy. It will not apply to any traffic until you associate to
// a traffic source via an endpoint configuration or IP restriction.
//
// https://ngrok.com/docs/api#api-ip-policies-create
func (c *Client) Create(ctx context.Context, arg *ngrok.IPPolicyCreate) (*ngrok.IPPolicy, error) {
	if arg == nil {
		arg = new(ngrok.IPPolicyCreate)
	}
	var res ngrok.IPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/ip_policies")).Execute(&path, arg); err != nil {
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

// Delete an IP policy. If the IP policy is referenced by another object for the
// purposes of traffic restriction it will be treated as if the IP policy remains
// but has zero rules.
//
// https://ngrok.com/docs/api#api-ip-policies-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/ip_policies/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get detailed information about an IP policy by ID.
//
// https://ngrok.com/docs/api#api-ip-policies-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.IPPolicy, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.IPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/ip_policies/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all IP policies on this account
//
// https://ngrok.com/docs/api#api-ip-policies-list
func (c *Client) List(paging *ngrok.Paging) *Iter {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/ip_policies")).Execute(&path, paging); err != nil {
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
	items  []ngrok.IPPolicy
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
	var resp ngrok.IPPolicyList
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
	if len(resp.IPPolicies) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.IPPolicies
	return it.Next(ctx)
}

// Item() returns the IPPolicy currently
// pointed to by the iterator.
func (it *Iter) Item() *ngrok.IPPolicy {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter) Err() error {
	return it.err
}

// Update attributes of an IP policy by ID
//
// https://ngrok.com/docs/api#api-ip-policies-update
func (c *Client) Update(ctx context.Context, arg *ngrok.IPPolicyUpdate) (*ngrok.IPPolicy, error) {
	if arg == nil {
		arg = new(ngrok.IPPolicyUpdate)
	}
	var res ngrok.IPPolicy
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/ip_policies/{{ .ID }}")).Execute(&path, arg); err != nil {
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
