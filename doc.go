// Code generated for API Clients. DO NOT EDIT.

// Package ngrok makes it easy to work with the ngrok API from Go. The package
// is fully code generated and should always be up to date with the latest
// ngrok API.
//
// Full documentation of the ngrok API can be found at: https://ngrok.com/docs/api
//
// # Versioning
//
// This package follows the best practices outlined for Go modules. All releases
// are tagged and any breaking changes will be reflected as a new major
// version. You should only import this package for production applications by
// pointing at a stable tagged version.
//
// # Quickstart
//
// The following example code demonstrates typical initialization and usage of
// the package to make an API call:
//
//	import (
//	    "context"
//	    "fmt"
//	    "math/rand"
//
//	    "github.com/ngrok/ngrok-api-go/v7"
//	    "github.com/ngrok/ngrok-api-go/v7/reserved_domains"
//	)
//
//	func example(ctx context.Context) error {
//	        domains := reserved_domains.NewClient(ngrok.NewClientConfig("<API KEY>"))
//	        d, err := domains.Create(ctx, &ngrok.ReservedDomainCreate{
//	                Name: fmt.Sprintf("hello-gopher-%x", rand.Int()),
//	        })
//	        if err != nil {
//	                return err
//	        }
//	        fmt.Println("reserved domain", d.Domain)
//	        return nil
//	}
//
// # Package Layout and API Clients
//
// API client configuration and all of the datatypes exchanged by the API are
// defined in this base package. There are subpackages for every API service
// and a Client type defined in those packages with methods to interact with
// that API service. It's usually easiest to find the subpackage of the service
// you want to work with and begin consulting the documentation there.
//
// It is recommended to construct the service-specific clients once at
// initialization time.
//
//	import (
//	    "github.com/ngrok/ngrok-api-go/v7"
//	    "github.com/ngrok/ngrok-api-go/v7/ip_policies"
//	    "github.com/ngrok/ngrok-api-go/v7/ip_policy_rules"
//	)
//
//	// Construct the root api Client object
//	apiClientConfig := ngrok.NewClientConfig("<API KEY>")
//
//	// then construct service-specific clients
//	policies := ip_policies.NewClient(apiClientConfig)
//	rules := ip_policy_rules.NewClient(apiClientConfig)
//
// # Functional Option Configuration
//
// The ClientConfig object in the root package supports functional options for
// configuration. The most common option to use is `WithHTTPClient()` which
// allows the caller to specify a different net/http.Client object. This allows
// the caller full customization over the transport if needed for use with
// proxies, custom TLS setups, observability and tracing, etc.
//
//	ngrokAPI := ngrok.NewClientConfig("<API KEY>", ngrok.WithHTTPClient(yourHTTPClient))
//
// # Nullable arguments
//
// Some arguments to methods in the ngrok API are optional and must be
// meaningfully distinguished from zero values, especially in Update() methods.
// This allows the API to distinguish between choosing not to update a value
// vs. setting it to zero or the empty string. For these arguments, ngrok
// follows the industry standard practice of using pointers to the primitive
// types and providing convenince functions like ngrok.String() and
// ngrok.Bool() for the caller to wrap literals as pointer values. For example:
//
//	creds := credentials.NewClient(ngrok.NewClientConfig("<API KEY>"))
//	c, err := creds.Update(ctx, &ngrok.CredentialUpdate{
//	        ID:          "cr_1kYzunEyn6XHHlqyMBLrj5nxkoz",
//	        Description: ngrok.String("this optional description is a pointer to a string"),
//	})
//
// # Automatic Paging
//
// All List methods in the ngrok API are paged. This package abstracts that
// problem away from you by returning an iterator from any List API call. As
// you advance the iterator it will transparently fetch new pages of values for
// you behind the scenes. Note that the context supplied to the initial List()
// call will be used for all subsequent page fetches so it must be long enough
// to work through the entire list. Here's an example of paging through all of
// the TLS certificates on your account. Note that you must check for an error
// after Next() returns false to determine if the iterator failed to fetch the
// next page of results.
//
//	certs := tls_certificates.NewClient(ngrok.NewClientConfig("<API KEY>"))
//	certsIter := certs.List(nil)
//	for certsIter.Next(ctx) {
//	        fmt.Println(certsIter.Item())
//	}
//	if err := certsIter.Err(); err != nil {
//	        return err
//	}
//
// # Error Handling
//
// All errors returned by the ngrok API are returned as structured payloads for
// easy error handling. Most non-networking errors returned by API calls in
// this package will be an ngrok.Error type. The ngrok.Error type exposes
// important metadata that will help you handle errors. Specifically it
// includes the HTTP status code of any failed operation as well as an error
// code value that uniquely identifies the failure condition.
//
// There are two helper functions that will make error handling easy:
// IsNotFound and IsErrorCode. IsNotFound helps identify the common case of
// accessing an API resource that no longer exists:
//
//	domains := reserved_domains.NewClient(ngrok.NewClientConfig("<API KEY>"))
//	d, err := domains.Get(ctx, "rd_1bXG9oRzwO4wECTdws3hlVw6jCg")
//	switch {
//	case ngrok.IsNotFound(err):
//	         // maybe this is an expected condition
//	case err != nil:
//	        return err
//	}
//
// IsErrorCode helps you identify specific ngrok errors by their unique ngrok
// error code. All ngrok error codes are documented at
// https://ngrok.com/docs/errors To check for a specific error condition, you
// would structure your code like the following example:
//
//	domains := reserved_domains.NewClient(ngrok.NewClientConfig("<API KEY>"))
//	d, err := domains.Create(ctx, &ngrok.ReservedDomainCreate{
//	        Region: "invalid",
//	        Name:   "gopher",
//	})
//	switch {
//	case ngrok.IsErrorCode(err, 400):
//	         // handle this error
//	case err != nil:
//	        return err
//	}
//
// # Pretty Printing
//
// All ngrok datatypes in this package define String() and GoString() methods
// so that they can be formatted into strings in helpful representations. The
// GoString() method is defined to pretty-print an object for debugging
// purposes with the "%#v" formatting verb.
package ngrok
