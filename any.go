package gomockctx

import (
	"context"

	"go.uber.org/mock/gomock"
)

// Any returns a gomock.Matcher which matches any context.Context object.
func Any() gomock.Matcher {
	return &anyMatcher{}
}

type anyMatcher struct{}

var _ gomock.Matcher = &anyMatcher{}

func (cm *anyMatcher) Matches(x interface{}) bool {
	_, ok := x.(context.Context)

	return ok
}

func (cm *anyMatcher) String() string {
	return "is a context.Context"
}
