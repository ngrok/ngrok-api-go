<!-- Code generated for API Clients. DO NOT EDIT. -->

## v5.1.0

ENHANCEMENTS:

* Added `owner_id` field to the `api_key`, `credential`, and `ssh_credential` resources. If supplied at credential creation, ownership will be assigned to the specified User or Bot. Only admins may specify an owner other than themselves. Defaults to the authenticated User or Bot.
* Added `failover_backend`, `http_response_backend`, and `tunnel_group_backend` resources. A Failover backend defines failover behavior within a list of referenced backends. Traffic is sent to the first backend in the list. If that backend is offline or no connection can be established, ngrok attempts to connect to the next backend in the list until one is successful.

## v5.0.0

### Breaking Changes

`v5.0.0` has breaking changes. Notably, properties that
previously had pointers to slices or maps no longer have pointers. For instance,

```go
Hostports: *[]string `json:"hostports,omitempty"`
```

are now

```go
Hostports: []string `json:"hostports,omitempty"`
```

### Additions

New clients have been generated for `ApplicationSessions` and `ApplicationUsers`. 

