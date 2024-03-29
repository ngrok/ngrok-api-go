// Code generated for API Clients. DO NOT EDIT.

package ngrok_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/ip_policies"
	"github.com/stretchr/testify/require"
)

func TestIPPolicy(t *testing.T) {
	var opts []ngrok.ClientConfigOption

	var mock mockTransport
	if os.Getenv("TEST_NO_MOCK") != "true" {
		opts = append(opts, ngrok.WithHTTPClient(&http.Client{
			Transport: &mock,
		}))
	} else if os.Getenv("TEST_DEBUG") == "true" {
		// dump requests and responses to stdout
		opts = append(opts, ngrok.WithHTTPClient(&http.Client{
			Transport: &debugTransport{
				rt:  http.DefaultTransport,
				out: os.Stderr,
			},
		}))
	}

	ctx := context.Background()
	clientConfig := ngrok.NewClientConfig(os.Getenv("NGROK_API_KEY"), opts...)
	policies := ip_policies.NewClient(clientConfig)

	// test policy creation
	mock.SetResponse(201, `{"id":"ipp_1sbMfZquosZtu5mZPgA91UDFaDC","uri":"https://api.ngrok.com/ip_policies/ipp_1sbMfZquosZtu5mZPgA91UDFaDC","created_at":"2021-05-16T03:48:59Z","description":"ngrok-api-go tests","metadata":""}`)
	createInstance, err := policies.Create(ctx, &ngrok.IPPolicyCreate{
		Description: "ngrok-api-go tests",
	})
	require.NoError(t, err)

	// test get
	mock.SetResponse(200, `{"id":"ipp_1sbMfZquosZtu5mZPgA91UDFaDC","uri":"https://api.ngrok.com/ip_policies/ipp_1sbMfZquosZtu5mZPgA91UDFaDC","created_at":"2021-05-16T03:48:59Z","description":"ngrok-api-go tests","metadata":""}`)
	getInstance, err := policies.Get(ctx, createInstance.ID)
	require.NoError(t, err)
	require.Equal(t, createInstance, getInstance)

	// test update
	mock.SetResponse(200, `{"id":"ipp_1sbMfZquosZtu5mZPgA91UDFaDC","uri":"https://api.ngrok.com/ip_policies/ipp_1sbMfZquosZtu5mZPgA91UDFaDC","created_at":"2021-05-16T03:48:59Z","description":"ngrok-api-go tests","metadata":"{\"device-id\": \"malamute-12\"}"}`)
	metadata := `{"device-id": "malamute-12"}`
	updatedInstance, err := policies.Update(ctx, &ngrok.IPPolicyUpdate{
		ID:       createInstance.ID,
		Metadata: ngrok.String(metadata),
	})
	require.NoError(t, err)
	require.Equal(t, updatedInstance.Metadata, metadata)

	// test get after update
	mock.SetResponse(200, `{"id":"ipp_1sbMfZquosZtu5mZPgA91UDFaDC","uri":"https://api.ngrok.com/ip_policies/ipp_1sbMfZquosZtu5mZPgA91UDFaDC","created_at":"2021-05-16T03:48:59Z","description":"ngrok-api-go tests","metadata":"{\"device-id\": \"malamute-12\"}"}`)
	getAfterUpdateInstance, err := policies.Get(ctx, createInstance.ID)
	require.NoError(t, err)
	require.Equal(t, updatedInstance, getAfterUpdateInstance)

	mock.SetResponse(200, `{"ip_policies":[{"id":"ipp_1sbMfZquosZtu5mZPgA91UDFaDC","uri":"https://api.ngrok.com/ip_policies/ipp_1sbMfZquosZtu5mZPgA91UDFaDC","created_at":"2021-05-16T03:48:59Z","description":"ngrok-api-go tests","metadata":"{\"device-id\": \"malamute-12\"}"},{"id":"ipp_1qXI4T0q6cgkoOVvqSEjU7LiWIr","uri":"https://api.ngrok.com/ip_policies/ipp_1qXI4T0q6cgkoOVvqSEjU7LiWIr","created_at":"2021-03-31T19:35:16Z","description":"martin demo","metadata":""}],"uri":"https://api.ngrok.com/ip_policies","next_page_uri":null}`)
	mock.SetResponse(200, `{"ip_policies":[],"uri":"https://api.ngrok.com/ip_policies","next_page_uri":null}`)
	iter := policies.List(nil)
	var iterPolicies []*ngrok.IPPolicy
	for iter.Next(ctx) {
		iterPolicies = append(iterPolicies, iter.Item())
	}
	require.NoError(t, iter.Err())
	require.Contains(t, iterPolicies, updatedInstance)

	// test delete
	mock.SetResponse(204, "")
	err = policies.Delete(ctx, createInstance.ID)
	require.NoError(t, err)

	// test 404
	mock.SetResponse(404, `{"status_code":404,"msg":"Resource not found","details":{"operation_id":"op_1sbMfWvXaRA26gTJoBPIgyPD8MF"}}`)
	_, err = policies.Get(ctx, createInstance.ID)
	require.True(t, ngrok.IsNotFound(err))
}
