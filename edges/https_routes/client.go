// Code generated for API Clients. DO NOT EDIT.

package https_routes

import (
	"bytes"
	"context"
	"fmt"
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

// Create an HTTPS Edge Route
//
// https://ngrok.com/docs/api#api-edges-https-routes-create
func (c *Client) Create(ctx context.Context, arg *ngrok.HTTPSEdgeRouteCreate) (*ngrok.HTTPSEdgeRoute, error) {
	var res ngrok.HTTPSEdgeRoute
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/edges/https/{{ .EdgeID }}/routes")).Execute(&path, arg); err != nil {
		return nil, fmt.Errorf("error building path for create: %w", err)
	}
	arg.EdgeID = ""
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

// Get an HTTPS Edge Route by ID
//
// https://ngrok.com/docs/api#api-edges-https-routes-get
func (c *Client) Get(ctx context.Context, arg *ngrok.EdgeRouteItem) (*ngrok.HTTPSEdgeRoute, error) {
	if arg == nil {
		arg = new(ngrok.EdgeRouteItem)
	}
	var res ngrok.HTTPSEdgeRoute
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}")).Execute(&path, arg); err != nil {
		return nil, fmt.Errorf("error building path for get: %w", err)
	}
	arg.EdgeID = ""
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

// Updates an HTTPS Edge Route by ID. If a module is not specified in the update,
// it will not be modified. However, each module configuration that is specified
// will completely replace the existing value. There is no way to delete an
// existing module via this API, instead use the delete module API.
//
// https://ngrok.com/docs/api#api-edges-https-routes-update
func (c *Client) Update(ctx context.Context, arg *ngrok.HTTPSEdgeRouteUpdate) (*ngrok.HTTPSEdgeRoute, error) {
	if arg == nil {
		arg = new(ngrok.HTTPSEdgeRouteUpdate)
	}
	var res ngrok.HTTPSEdgeRoute
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}")).Execute(&path, arg); err != nil {
		return nil, fmt.Errorf("error building path for update: %w", err)
	}
	arg.EdgeID = ""
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

// Delete an HTTPS Edge Route by ID
//
// https://ngrok.com/docs/api#api-edges-https-routes-delete
func (c *Client) Delete(ctx context.Context, arg *ngrok.EdgeRouteItem) error {
	if arg == nil {
		arg = new(ngrok.EdgeRouteItem)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}")).Execute(&path, arg); err != nil {
		return fmt.Errorf("error building path for delete: %w", err)
	}
	arg.EdgeID = ""
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
