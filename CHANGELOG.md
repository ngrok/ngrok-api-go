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
