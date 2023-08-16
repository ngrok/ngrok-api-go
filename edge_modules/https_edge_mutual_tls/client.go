// Code generated for API Clients. DO NOT EDIT.

package https_edge_mutual_tls

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/internal/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

func (c *Client) Replace(ctx context.Context, arg *ngrok.EdgeMutualTLSReplace) (*ngrok.EndpointMutualTLS, error) {
	if arg == nil {
		arg = new(ngrok.EdgeMutualTLSReplace)
	}
	var res ngrok.EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("replace_path").Parse("/edges/https/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
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

func (c *Client) Get(ctx context.Context, id string) (*ngrok.EndpointMutualTLS, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.EndpointMutualTLS
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/edges/https/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
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
	if err := template.Must(template.New("delete_path").Parse("/edges/https/{{ .ID }}/mutual_tls")).Execute(&path, arg); err != nil {
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
