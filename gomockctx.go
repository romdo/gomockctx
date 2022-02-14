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

// ID returns the gomockctx ID value in the given context, or a empty string if
// the context does not have a gomockctx ID value.
func ID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	return string(value(ctx))
}

// Any returns a gomock.Matcher which matches any context.Context object.
func Any() gomock.Matcher {
	return gomock.AssignableToTypeOf(
		reflect.TypeOf((*context.Context)(nil)).Elem(),
	)
}

// WithValue returns a gomock.Matcher which matches any context that has the
// specified key and value.
func WithValue(key interface{}, value interface{}) gomock.Matcher {
	return &contextMatcher{
		key:   key,
		value: value,
	}
}

type contextMatcher struct {
	key   interface{}
	value interface{}
}

var _ gomock.Matcher = &contextMatcher{}

func (cm *contextMatcher) Matches(x interface{}) bool {
	if ctx, ok := x.(context.Context); ok {
		return reflect.DeepEqual(cm.value, ctx.Value(cm.key))
	}

	return false
}

func (cm *contextMatcher) String() string {
	return fmt.Sprintf(`context with "%+v" = "%+v"`, cm.key, cm.value)
}

func newCtxID() contextValue {
	id, err := randString(64)
	if err != nil {
		panic(err)
	}

	return contextValue(id)
}

func value(ctx context.Context) contextValue {
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
