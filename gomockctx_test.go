package gomockctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
