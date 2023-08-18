// Code generated for API Clients. DO NOT EDIT.

package ssh_credentials

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/internal/api"
)

// SSH Credentials are SSH public keys that can be used to start SSH tunnels
//  via the ngrok SSH tunnel gateway.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new ssh_credential from an uploaded public SSH key. This ssh credential
// can be used to start new tunnels via ngrok's SSH gateway.
//
// https://ngrok.com/docs/api#api-ssh-credentials-create
func (c *Client) Create(ctx context.Context, arg *ngrok.SSHCredentialCreate) (*ngrok.SSHCredential, error) {
	var res ngrok.SSHCredential
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/ssh_credentials")).Execute(&path, arg); err != nil {
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

// Delete an ssh_credential by ID
//
// https://ngrok.com/docs/api#api-ssh-credentials-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/ssh_credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Get detailed information about an ssh_credential
//
// https://ngrok.com/docs/api#api-ssh-credentials-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.SSHCredential, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.SSHCredential
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/ssh_credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// List all ssh credentials on this account
//
// https://ngrok.com/docs/api#api-ssh-credentials-list
func (c *Client) list(ctx context.Context, arg *ngrok.Paging) (*ngrok.SSHCredentialList, error) {
	if arg == nil {
		arg = new(ngrok.Paging)
	}
	var res ngrok.SSHCredentialList
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/ssh_credentials")).Execute(&path, arg); err != nil {
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

// List all ssh credentials on this account
//
// https://ngrok.com/docs/api#api-ssh-credentials-list
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
	items      []ngrok.SSHCredential
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
	if len(resp.SSHCredentials) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.SSHCredentials
	return it.Next(ctx)
}

// Item() returns the SSHCredential currently
// pointed to by the iterator.
func (it *Iter) Item() *ngrok.SSHCredential {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter) Err() error {
	return it.err
}

// Update attributes of an ssh_credential by ID
//
// https://ngrok.com/docs/api#api-ssh-credentials-update
func (c *Client) Update(ctx context.Context, arg *ngrok.SSHCredentialUpdate) (*ngrok.SSHCredential, error) {
	if arg == nil {
		arg = new(ngrok.SSHCredentialUpdate)
	}
	var res ngrok.SSHCredential
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/ssh_credentials/{{ .ID }}")).Execute(&path, arg); err != nil {
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
