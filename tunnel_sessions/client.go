// Code generated for API Clients. DO NOT EDIT.

package tunnel_sessions

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v7"
	"github.com/ngrok/ngrok-api-go/v7/internal/api"
)

// Tunnel Sessions represent instances of ngrok agents or SSH reverse tunnel
//  sessions that are running and connected to the ngrok service. Each tunnel
//  session can include one or more Tunnels.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// List all online tunnel sessions running on this account.
//
// https://ngrok.com/docs/api#api-tunnel-sessions-list
func (c *Client) List(paging *ngrok.Paging) ngrok.Iter[*ngrok.TunnelSession] {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/tunnel_sessions")).Execute(&path, paging); err != nil {
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
	return &iterTunnelSession{
		client:   c,
		n:        -1,
		nextPage: apiURL,
	}
}

// iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type iterTunnelSession struct {
	client *Client
	n      int
	items  []ngrok.TunnelSession
	err    error

	nextPage *url.URL
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *iterTunnelSession) Next(ctx context.Context) bool {
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
	var resp ngrok.TunnelSessionList
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
	if len(resp.TunnelSessions) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.TunnelSessions
	return it.Next(ctx)
}

// Item() returns the TunnelSession currently
// pointed to by the iterator.
func (it *iterTunnelSession) Item() *ngrok.TunnelSession {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *iterTunnelSession) Err() error {
	return it.err
}

// Get the detailed status of a tunnel session by ID
//
// https://ngrok.com/docs/api#api-tunnel-sessions-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.TunnelSession, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.TunnelSession
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/tunnel_sessions/{{ .ID }}")).Execute(&path, arg); err != nil {
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

// Issues a command instructing the ngrok agent to restart. The agent restarts
// itself by calling exec() on platforms that support it. This operation is notably
// not supported on Windows. When an agent restarts, it reconnects with a new
// tunnel session ID.
//
// https://ngrok.com/docs/api#api-tunnel-sessions-restart
func (c *Client) Restart(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("restart_path").Parse("/tunnel_sessions/{{ .ID }}/restart")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "POST", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}

// Issues a command instructing the ngrok agent that started this tunnel session to
// exit.
//
// https://ngrok.com/docs/api#api-tunnel-sessions-stop
func (c *Client) Stop(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("stop_path").Parse("/tunnel_sessions/{{ .ID }}/stop")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "POST", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}

// Issues a command instructing the ngrok agent to update itself to the latest
// version. After this call completes successfully, the ngrok agent will be in the
// update process. A caller should wait some amount of time to allow the update to
// complete (at least 10 seconds) before making a call to the Restart endpoint to
// request that the agent restart itself to start using the new code. This call
// will never update an ngrok agent to a new major version which could cause
// breaking compatibility issues. If you wish to update to a new major version,
// that must be done manually. Still, please be aware that updating your ngrok
// agent could break your integration. This call will fail in any of the following
// circumstances: there is no update available the ngrok agent's configuration
// disabled update checks the agent is currently in process of updating the agent
// has already successfully updated but has not yet been restarted
//
// https://ngrok.com/docs/api#api-tunnel-sessions-update
func (c *Client) Update(ctx context.Context, id string) error {
	arg := &ngrok.TunnelSessionsUpdate{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/tunnel_sessions/{{ .ID }}/update")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "POST", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}
