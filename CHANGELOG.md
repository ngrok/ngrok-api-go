<!-- Code generated for API Clients. DO NOT EDIT. -->

## v6.0.0
### Breaking Changes
* Renamed the Policy Module to the Traffic Policy Module on HTTP Edge Routes, TCP Edges, and TLS Edges, which allows you to configure rules that can be used to influence and control traffic to and from your upstream service. The Traffic Policy itself is now specified as either a JSON or YAML string.

## v5.4.1

FIXES:

* Fixed an issue with pagination for some resources, like `reserved_addrs`, by using the `next_page_uri` returned in the response instead of the id of the last resource.

## v5.4.0

ENHANCEMENTS:

* Added support for the Bot User API. The Bot User API allows you to manage the bots that are registered to your ngrok account. You can automate the creation, management, and deletion of bot users in your account.

* Added support for the Policy Module on HTTP Edge Routes, TCP Edges, and TLS Edges, which allows you to configure rules that can be used to influence and control traffic to and from your upstream service.

## v5.3.1

INTERNAL:

* Moved location of BaseClient

## v5.3.0

ENHANCEMENTS:

* Added support for the Traffic Policy module on TCP edges, TLS edges, and HTTPS edge routes. A Traffic Policy allows traffic to and from your edges to be shaped and influenced by defining rules that contain expression to filter the traffic and actions to take effect.
* Added `static_backend` resource, which defines a static address at which traffic should be forwarded.

## v5.2.0

ENHANCEMENTS:

* Added support for the User Agent Filter module on HTTPS edge routes. The User Agent Filter module enables you to block bots, crawlers, or certain browsers from accessing your web application. It allows or denies traffic based on the User-Agent header of incoming HTTP requests.
* Added support for the JWT Validation module on HTTPS edge routes. The JWT Validation module allows or denies traffic based on validation of the JSON Web Token (JWT) provided within the request that was initiated to your ngrok endpoints.

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

