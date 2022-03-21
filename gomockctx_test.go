package gomockctx

import (
	"context"
	"regexp"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type key struct{}
	k := key{}
	ctxIDRegexp := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	parent := context.WithValue(context.Background(), k, "the parent")

	ids := map[contextValue]struct{}{}
	limit := 1000

	for i := 0; i < limit; i++ {
		ctx := New(parent)
		require.Equal(t, "the parent", ctx.Value(k))

		v := ctx.Value(ctxKey)
		require.IsType(t, contextValue(""), v)
		require.Len(t, v, 32)
		require.Regexp(t, ctxIDRegexp, v)
		cv, _ := v.(contextValue)
		ids[cv] = struct{}{}
	}

	assert.Len(t, ids, limit)
}

func TestEq(t *testing.T) {
	type strKey string

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *valueMatcher
	}{
		{
			name: "nil",
			args: args{
				ctx: nil,
			},
			want: &valueMatcher{
				key:   ctxKey,
				value: contextValue(""),
			},
		},
		{
			name: "context without gomockctx ID",
			args: args{
				ctx: context.Background(),
			},
			want: &valueMatcher{
				key:   ctxKey,
				value: contextValue(""),
			},
		},
		{
			name: "context with gomockctx ID",
			args: args{
				ctx: context.WithValue(
					context.Background(),
					ctxKey,
					contextValue("z9KZVcfmA4sWJX0yuIIESVcEARlwiAT2"),
				),
			},
			want: &valueMatcher{
				key:   ctxKey,
				value: contextValue("z9KZVcfmA4sWJX0yuIIESVcEARlwiAT2"),
			},
		},
		{
			name: "child context of context with gomockctx ID",
			args: args{
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						ctxKey,
						contextValue("hWEKf4Gtj15iLx4R7IFlHc5ooj5tU4UW"),
					),
					strKey("foo"),
					"bar",
				),
			},
			want: &valueMatcher{
				key:   ctxKey,
				value: contextValue("hWEKf4Gtj15iLx4R7IFlHc5ooj5tU4UW"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Eq(tt.args.ctx)

			assert.Implements(t, (*gomock.Matcher)(nil), got)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIs(t *testing.T) {
	TestEq(t)
}

func TestID(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want string
	}{
		{
			name: "nil",
			ctx:  nil,
			want: "",
		},
		{
			name: "without ID",
			ctx:  context.Background(),
			want: "",
		},
		{
			name: "with ID",
			ctx: context.WithValue(
				context.Background(), ctxKey, contextValue("xI2UWC8MvdYcU22B"),
			),
			want: "xI2UWC8MvdYcU22B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ID(tt.ctx)

			assert.Equal(t, tt.want, got)
		})
	}
}
