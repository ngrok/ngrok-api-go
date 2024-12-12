// Code generated for API Clients. DO NOT EDIT.

package reserved_domains

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v6"
	"github.com/ngrok/ngrok-api-go/v6/internal/api"
)

// Reserved Domains are hostnames that you can listen for traffic on. Domains
//  can be used to listen for http, https or tls traffic. You may use a domain
//  that you own by creating a CNAME record specified in the returned resource.
//  This CNAME record points traffic for that domain to ngrok's edge servers.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-create
func (c *Client) Create(ctx context.Context, arg *ngrok.ReservedDomainCreate) (*ngrok.ReservedDomain, error) {
	if arg == nil {
		arg = new(ngrok.ReservedDomainCreate)
	}
	var res ngrok.ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/reserved_domains")).Execute(&path, arg); err != nil {
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

// Delete a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get the details of a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.ReservedDomain, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all reserved domains on this account.
//
// https://ngrok.com/docs/api#api-reserved-domains-list
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.ReservedDomain] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/reserved_domains")).Execute(&path, paging); err != nil {
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
	return &iterReservedDomain{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type iterReservedDomain struct {
	client *Client
	n      int
	items  []ngrok.ReservedDomain
	err    error

	nextPage *url.URL
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *iterReservedDomain) Next(ctx context.Context) bool {
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
	var resp ngrok.ReservedDomainList
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
	if len(resp.ReservedDomains) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.ReservedDomains
	return it.Next(ctx)
}

// Item() returns the ReservedDomain currently
// pointed to by the iterator.
func (it *iterReservedDomain) Item() *ngrok.ReservedDomain {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iterReservedDomain) Err() error {
	return it.err
}

// Update the attributes of a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-update
func (c *Client) Update(ctx context.Context, arg *ngrok.ReservedDomainUpdate) (*ngrok.ReservedDomain, error) {
	if arg == nil {
		arg = new(ngrok.ReservedDomainUpdate)
	}
	var res ngrok.ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Detach the certificate management policy attached to a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-delete-certificate-management-policy
func (c *Client) DeleteCertificateManagementPolicy(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_certificate_management_policy_path").Parse("/reserved_domains/{{ .ID }}/certificate_management_policy")).Execute(&path, arg); err != nil {
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

// Detach the certificate attached to a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-delete-certificate
func (c *Client) DeleteCertificate(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_certificate_path").Parse("/reserved_domains/{{ .ID }}/certificate")).Execute(&path, arg); err != nil {
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
