// Code generated for API Clients. DO NOT EDIT.

package tls_edge_backend

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v6"
	"github.com/ngrok/ngrok-api-go/v6/internal/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

func (c *Client) Replace(ctx context.Context, arg *ngrok.EdgeBackendReplace) (*ngrok.EndpointBackend, error) {
	if arg == nil {
		arg = new(ngrok.EdgeBackendReplace)
	}
	var res ngrok.EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("replace_path").Parse("/edges/tls/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
		panic(err)
	}
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

func (c *Client) Get(ctx context.Context, id string) (*ngrok.EndpointBackend, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.EndpointBackend
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/edges/tls/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
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

func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/edges/tls/{{ .ID }}/backend")).Execute(&path, arg); err != nil {
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
