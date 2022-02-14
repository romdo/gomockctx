// Package gomockctx contains gomock helpers for matching context.Context
// objects.
package gomockctx

import (
	"context"
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

type (
	contextKey   string
	contextValue string
)

var ctxKey contextKey = "gomockctx context ID"

func newCtxID() contextValue {
	id, err := randString(64)
	if err != nil {
		panic(err)
	}

	return contextValue(id)
}

func value(ctx context.Context) contextValue {
	var value contextValue
	v := ctx.Value(ctxKey)
	if s, ok := v.(contextValue); ok {
		value = s
	}

	return value
}

// New returns a context as a child of the given parent, and with a randomized
// gomockctx ID value set, making it a gomockctx context. This can then be used
// with Is to get a gomock Matcher which returns true for the context from New,
// or any child contexts of it.
//
// If crypto/rand returns an error, this will panic trying to generate the
// gomockctx ID. In practise though, crypto/rand should never return a error.
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
	return WithValue(ctxKey, value(ctx))
}

// WithValue creates a generic gomock context matcher which returns true for any
// context which has the specified key/value.
func WithValue(key interface{}, value interface{}) gomock.Matcher {
	return &contextMatcher{
		key:   key,
		value: value,
	}
}

// ID returns the gomockctx ID value in the given context, or a empty string if
// it not a gomockctx context.
func ID(ctx context.Context) string {
	return string(value(ctx))
}

type contextMatcher struct {
	key   interface{}
	value interface{}
}

var _ gomock.Matcher = &contextMatcher{}

func (e *contextMatcher) Matches(x interface{}) bool {
	ctx, ok := x.(context.Context)
	if !ok {
		return false
	}

	return reflect.DeepEqual(e.value, ctx.Value(e.key))
}

func (e *contextMatcher) String() string {
	return fmt.Sprintf(`context with "%+v" = "%+v"`, e.key, e.value)
}
