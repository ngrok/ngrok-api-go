# ngrok API client library for Golang

This library wraps the [ngrok HTTP API](https://ngrok.com/docs/api) to make it
easier to consume in Go.

## Installation

Installation is as simple as using `go get`.

    go get github.com/ngrok/ngrok-api-go/v5

## Documentation

A quickstart guide and a full API reference are included in the [ngrok go API documentation on pkg.go.dev](https://pkg.go.dev/github.com/ngrok/ngrok-api-go/v5)

## Quickstart

Please consult the [documentation](https://pkg.go.dev/github.com/ngrok/ngrok-api-go/v5) for additional examples.

### Create an IP Policy that allows traffic from some subnets

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/ip_policies"
	"github.com/ngrok/ngrok-api-go/v5/ip_policy_rules"
)

func main() {
	fmt.Println(example(context.Background()))
}

func example(ctx context.Context) error {
	// create clients to api resources
	clientConfig := ngrok.NewClientConfig(os.Getenv("NGROK_API_KEY"))
	policies := ip_policies.NewClient(clientConfig)
	policyRules := ip_policy_rules.NewClient(clientConfig)

	// create the ip policy
	policy, err := policies.Create(ctx, &ngrok.IPPolicyCreate{})
	if err != nil {
		return err
	}
	fmt.Println(policy)

	// create rules for each cidr
	for _, cidr := range []string{"24.0.0.0/8", "12.0.0.0/8"} {
		rule, err := policyRules.Create(ctx, &ngrok.IPPolicyRuleCreate{
			CIDR:       cidr,
			IPPolicyID: policy.ID,
			Action:     ngrok.String("allow"),
		})
		if err != nil {
			return err
		}
		fmt.Println(rule)
	}
	return nil
}
```

### List all online tunnels

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/tunnels"
)

func main() {
	fmt.Println(example(context.Background()))
}

func example(ctx context.Context) error {
	// construct the api client
	clientConfig := ngrok.NewClientConfig(os.Getenv("NGROK_API_KEY"))

	// list all online tunnels
	tunnels := tunnels.NewClient(clientConfig)
	iter := tunnels.List(nil)
	for iter.Next(ctx) {
		fmt.Println(iter.Item())
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}
```
