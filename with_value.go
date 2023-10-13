package gomockctx

import (
	"context"
	"fmt"
	"reflect"

	"go.uber.org/mock/gomock"
)

// WithValue returns a gomock.Matcher which matches any context that has the
// specified key and value.
func WithValue(key interface{}, value interface{}) gomock.Matcher {
	return &valueMatcher{
		key:   key,
		value: value,
	}
}

type valueMatcher struct {
	key   interface{}
	value interface{}
}

var _ gomock.Matcher = &valueMatcher{}

func (cm *valueMatcher) Matches(x interface{}) bool {
	if ctx, ok := x.(context.Context); ok {
		return reflect.DeepEqual(cm.value, ctx.Value(cm.key))
	}

	return false
}

func (cm *valueMatcher) String() string {
	return fmt.Sprintf(`context with "%+v" = "%+v"`, cm.key, cm.value)
}
