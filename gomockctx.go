// Package gomockctx contains gomock helpers for matching context.Context
// objects.
package gomockctx

import (
	"context"

	"github.com/golang/mock/gomock"
)

type (
	contextKey   string
	contextValue string
)

var ctxKey contextKey = "gomockctx ID"

func newCtxID() contextValue {
	id, err := randString(32)
	if err != nil {
		panic(err)
	}

	return contextValue(id)
}

func getValue(ctx context.Context) contextValue {
	var value contextValue
	if ctx == nil {
		return value
	}

	v := ctx.Value(ctxKey)
	if s, ok := v.(contextValue); ok {
		value = s
	}

	return value
}

// New returns a context as a child of the given parent, which includes a
// randomized gomockctx ID value set, which makes it a gomockctx context. This
// can then be used with Is to get a gomock Matcher which returns true for the
// context from New, or any child contexts of it.
//
// If crypto/rand returns an error, this will panic trying to generate the
// gomockctx ID. In practice though, crypto/rand should never return a error.
func New(parent context.Context) context.Context {
	return context.WithValue(parent, ctxKey, newCtxID())
}

// Is accepts a context with a gomockctx ID value (as returned from New), and
// returns a gomock.Matcher which returns true for the given context, of any
// child contexts of it.
//
// If ctx was not returned from New, the resulting matcher will ALWAYS return
// false.
func Is(ctx context.Context) gomock.Matcher {
	return WithValue(ctxKey, getValue(ctx))
}

// ID returns the gomockctx ID value in the given context, or a empty string if
// the context does not have a gomockctx ID value.
func ID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	return string(getValue(ctx))
}
