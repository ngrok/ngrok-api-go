package ngrok

import "context"

type Iter[T any] interface {
	Next(ctx context.Context) bool
	Item() T
	Err() error
}
