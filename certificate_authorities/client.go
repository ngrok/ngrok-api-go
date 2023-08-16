// Code generated for API Clients. DO NOT EDIT.

package certificate_authorities

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/internal/api"
)

// Certificate Authorities are x509 certificates that are used to sign other
//  x509 certificates. Attach a Certificate Authority to the Mutual TLS module
//  to verify that the TLS certificate presented by a client has been signed by
//  this CA. Certificate Authorities  are used only for mTLS validation only and
//  thus a private key is not included in the resource.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Upload a new Certificate Authority
//
// https://ngrok.com/docs/api#api-certificate-authorities-create
func (c *Client) Create(ctx context.Context, arg *ngrok.CertificateAuthorityCreate) (*ngrok.CertificateAuthority, error) {
	var res ngrok.CertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/certificate_authorities")).Execute(&path, arg); err != nil {
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

// Delete a Certificate Authority
//
// https://ngrok.com/docs/api#api-certificate-authorities-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get detailed information about a certficate authority
//
// https://ngrok.com/docs/api#api-certificate-authorities-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.CertificateAuthority, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.CertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all Certificate Authority on this account
//
// https://ngrok.com/docs/api#api-certificate-authorities-list
func (c *Client) list(ctx context.Context, arg *ngrok.Paging) (*ngrok.CertificateAuthorityList, error) {
	if arg == nil {
		arg = new(ngrok.Paging)
	}
	var res ngrok.CertificateAuthorityList
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/certificate_authorities")).Execute(&path, arg); err != nil {
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

// List all Certificate Authority on this account
//
// https://ngrok.com/docs/api#api-certificate-authorities-list
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
	items      []ngrok.CertificateAuthority
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
	if len(resp.CertificateAuthorities) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.CertificateAuthorities
	return it.Next(ctx)
}

// Item() returns the CertificateAuthority currently
// pointed to by the iterator.
func (it *Iter) Item() *ngrok.CertificateAuthority {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter) Err() error {
	return it.err
}

// Update attributes of a Certificate Authority by ID
//
// https://ngrok.com/docs/api#api-certificate-authorities-update
func (c *Client) Update(ctx context.Context, arg *ngrok.CertificateAuthorityUpdate) (*ngrok.CertificateAuthority, error) {
	if arg == nil {
		arg = new(ngrok.CertificateAuthorityUpdate)
	}
	var res ngrok.CertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
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
