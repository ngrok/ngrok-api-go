package ngrok_test

import (
	"context"
	"testing"

	"github.com/ngrok/ngrok-api-go"
	"github.com/ngrok/ngrok-api-go/ip_policies"
	"github.com/stretchr/testify/require"
)

func TestIPPolicy(t *testing.T) {
	ctx := context.Background()
	c, err := ngrok.NewClient()
	require.NoError(t, err)

	policies := ip_policies.NewClient(c)

	// test that a list call works
	_, err = policies.List(ctx, &ngrok.Page{})
	require.NoError(t, err)

	// test policy creation
	createInstance, err := policies.Create(ctx, &ngrok.IPPolicyCreate{
		Action:      "allow",
		Description: "ngrok-api-go tests",
	})
	require.NoError(t, err)

	// test get
	getInstance, err := policies.Get(ctx, createInstance.ID)
	require.NoError(t, err)
	require.Equal(t, createInstance, getInstance)

	// test update
	metadata := `{"device-id": "malamute-12"`
	updatedInstance, err := policies.Update(ctx, &ngrok.IPPolicyUpdate{
		ID:       createInstance.ID,
		Metadata: ngrok.String(metadata),
	})
	require.NoError(t, err)
	require.Equal(t, updatedInstance.Metadata, metadata)

	// test get after update
	getAfterUpdateInstance, err := policies.Get(ctx, createInstance.ID)
	require.NoError(t, err)
	require.Equal(t, updatedInstance, getAfterUpdateInstance)

	iter := policies.Iter(ctx)
	var iterPolicies []*ngrok.IPPolicy
	for iter.Next() {
		iterPolicies = append(iterPolicies, iter.Item())
	}
	require.NoError(t, iter.Err())
	require.Contains(t, iterPolicies, updatedInstance)

	// test delete
	err = policies.Delete(ctx, createInstance.ID)
	require.NoError(t, err)

	// test 404
	_, err = policies.Get(ctx, createInstance.ID)
	require.True(t, ngrok.IsNotFound(err))
}
