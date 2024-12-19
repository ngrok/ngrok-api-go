// Code generated for API Clients. DO NOT EDIT.

package ssh_certificate_authorities

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v6"
	"github.com/ngrok/ngrok-api-go/v6/internal/api"
)

// An SSH Certificate Authority is a pair of an SSH Certificate and its private
//  key that can be used to sign other SSH host and user certificates.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new SSH Certificate Authority
//
// https://ngrok.com/docs/api#api-ssh-certificate-authorities-create
func (c *Client) Create(ctx context.Context, arg *ngrok.SSHCertificateAuthorityCreate) (*ngrok.SSHCertificateAuthority, error) {
	if arg == nil {
		arg = new(ngrok.SSHCertificateAuthorityCreate)
	}
	var res ngrok.SSHCertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/ssh_certificate_authorities")).Execute(&path, arg); err != nil {
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

// Delete an SSH Certificate Authority
//
// https://ngrok.com/docs/api#api-ssh-certificate-authorities-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/ssh_certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get detailed information about an SSH Certficate Authority
//
// https://ngrok.com/docs/api#api-ssh-certificate-authorities-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.SSHCertificateAuthority, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.SSHCertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/ssh_certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all SSH Certificate Authorities on this account
//
// https://ngrok.com/docs/api#api-ssh-certificate-authorities-list
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.SSHCertificateAuthority] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/ssh_certificate_authorities")).Execute(&path, paging); err != nil {
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
	return &iterSSHCertificateAuthority{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type iterSSHCertificateAuthority struct {
	client *Client
	n      int
	items  []ngrok.SSHCertificateAuthority
	err    error

	nextPage *url.URL
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *iterSSHCertificateAuthority) Next(ctx context.Context) bool {
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
	var resp ngrok.SSHCertificateAuthorityList
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
	if len(resp.SSHCertificateAuthorities) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.SSHCertificateAuthorities
	return it.Next(ctx)
}

// Item() returns the SSHCertificateAuthority currently
// pointed to by the iterator.
func (it *iterSSHCertificateAuthority) Item() *ngrok.SSHCertificateAuthority {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iterSSHCertificateAuthority) Err() error {
	return it.err
}

// Update an SSH Certificate Authority
//
// https://ngrok.com/docs/api#api-ssh-certificate-authorities-update
func (c *Client) Update(ctx context.Context, arg *ngrok.SSHCertificateAuthorityUpdate) (*ngrok.SSHCertificateAuthority, error) {
	if arg == nil {
		arg = new(ngrok.SSHCertificateAuthorityUpdate)
	}
	var res ngrok.SSHCertificateAuthority
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/ssh_certificate_authorities/{{ .ID }}")).Execute(&path, arg); err != nil {
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
