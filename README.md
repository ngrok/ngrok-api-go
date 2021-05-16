# ngrok API client library for Golang

This library wraps the [ngrok HTTP API](https://ngrok.com/docs/api) to make it
easier to consume in Go.

## Installation

Installation is as simple as using `go get`.

    go get github.com/ngrok/ngrok-api-go

## Documentation

A quickstart guide and a full API reference are included in the [ngrok go API documentation on pkg.go.dev](https://pkg.go.dev/github.com/ngrok/ngrok-api-go)

## Quickstart

Please consult the [documentation](https://pkg.go.dev/github.com/ngrok/ngrok-api-go) for additional examples.

### Create an IP Policy that allows traffic from some subnets

```go
import (
        "github.com/ngrok/ngrok-api-go"
        "github.com/ngrok/ngrok-api-go/ip_policy"
        "github.com/ngrok/ngrok-api-go/ip_policy_rules"
)

func example(ctx context.Context) error {
        // create clients to api resources
        apiClient := ngrok.NewClient("<API KEY>")
        policies := ip_policies.NewClient(apiClient)
        policyRules := ip_policy_rules.NewClient(apiClient)

        // create the ip policy
        policy, err := policies.Create(ctx, &ngrok.IPPolicyCreate{
                Action: "allow",
        })
        if err != nil {
                return err
        }

        // create rules for each cidr
        for _, cidr := range []string{"24.0.0.0/8", "12.0.0.0/8"} {
                policyRules.Create(ctx, &ngrok.IPPolicyRuleCreate{
                        CIDR: cidr,
                        IPPolicyID: policy.ID,
                })
        }
        return nil
}
```

### List all online tunnels

```go
import (
        "github.com/ngrok/ngrok-api-go"
        "github.com/ngrok/ngrok-api-go/tunnels"
)

func example(ctx context.Context) error {
        // construct the api client
        apiClient := ngrok.NewClient("<API KEY>")

        // list all online tunnels
        tunnels := tunnels.NewClient(apiClient)
        iter := tunnels.List(ctx, nil)
        for iter.Next() {
                fmt.Println(iter.Item())
        }
        if err := iter.Err(); err != nil {
                return err
        }
        return nil
}
```
