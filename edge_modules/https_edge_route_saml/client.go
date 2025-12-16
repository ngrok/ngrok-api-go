// Code generated for API Clients. DO NOT EDIT.

package https_edge_route_saml

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

func (c *Client) Replace(ctx context.Context, arg *ngrok.EdgeRouteSAMLReplace) (*ngrok.EndpointSAML, error) {
	if arg == nil {
		arg = new(ngrok.EdgeRouteSAMLReplace)
	}
	var res ngrok.EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("replace_path").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		return nil, fmt.Errorf("error building path for replace: %w", err)
	}
	arg.EdgeID = ""
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg.Module

	if err := c.apiClient.Do(ctx, "PUT", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) Get(ctx context.Context, arg *ngrok.EdgeRouteItem) (*ngrok.EndpointSAML, error) {
	if arg == nil {
		arg = new(ngrok.EdgeRouteItem)
	}
	var res ngrok.EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
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

func (c *Client) Delete(ctx context.Context, arg *ngrok.EdgeRouteItem) error {
	if arg == nil {
		arg = new(ngrok.EdgeRouteItem)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/edges/https/{{ .EdgeID }}/routes/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
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
