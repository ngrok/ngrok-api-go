<!-- Code generated for API Clients. DO NOT EDIT. -->

# ngrok API client library for Golang
[![CI](https://github.com/ngrok/ngrok-api-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ngrok/ngrok-api-go/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/ngrok/ngrok-api-go.svg)](https://pkg.go.dev/github.com/ngrok/ngrok-api-go/v7)

This library wraps the [ngrok HTTP API](https://ngrok.com/docs/api) to make it
easier to consume in Go.

For creating ngrok tunnels directly from your Go application, check out the [ngrok Go Agent SDK](https://github.com/ngrok/ngrok-go) instead.

## Installation

Installation is as simple as using `go get`.

    go get github.com/ngrok/ngrok-api-go/v7

## Support

The best place to get support using this library is through the [ngrok Slack Community](https://ngrok.com/slack). If you find any bugs, please contribute by opening a [new GitHub issue](https://github.com/ngrok/ngrok-api-go/issues/new/choose).

## Documentation

A quickstart guide and a full API reference are included in the [ngrok go API documentation on pkg.go.dev](https://pkg.go.dev/github.com/ngrok/ngrok-api-go/v6)

## Quickstart

Please consult the [documentation](https://pkg.go.dev/github.com/ngrok/ngrok-api-go/v6) for additional examples.

### Create an IP Policy that allows traffic from some subnets

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ngrok/ngrok-api-go/v7"
	"github.com/ngrok/ngrok-api-go/v7/ip_policies"
	"github.com/ngrok/ngrok-api-go/v7/ip_policy_rules"
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

	"github.com/ngrok/ngrok-api-go/v7"
	"github.com/ngrok/ngrok-api-go/v7/tunnels"
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
