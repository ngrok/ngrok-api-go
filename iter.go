// Code generated for API Clients. DO NOT EDIT.

package ngrok

import (
	"context"
)

type Iter[T any] interface {
	Next(context.Context) bool
	Item() T
	Err() error
}
