package gomockctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAny(t *testing.T) {
	tests := []struct {
		name string
		x    interface{}
		want bool
	}{
		{
			name: "nil",
			x:    nil,
			want: false,
		},
		{
			name: "empty string",
			x:    "",
			want: false,
		},
		{
			name: "string",
			x:    "foo bar",
			want: false,
		},
		{
			name: "int",
			x:    42,
			want: false,
		},
		{
			name: "int8",
			x:    int8(64),
			want: false,
		},
		{
			name: "int16",
			x:    int16(1024),
			want: false,
		},
		{
			name: "int32",
			x:    int32(1123456789),
			want: false,
		},
		{
			name: "int64",
			x:    int64(16123456789),
			want: false,
		},
		{
			name: "uint",
			x:    uint(616),
			want: false,
		},
		{
			name: "uint8",
			x:    uint8(64),
			want: false,
		},
		{
			name: "uint16",
			x:    uint16(1024),
			want: false,
		},
		{
			name: "uint32",
			x:    uint32(1123456789),
			want: false,
		},
		{
			name: "uint64",
			x:    uint64(16123456789),
			want: false,
		},
		{
			name: "byte",
			x:    byte('A'),
			want: false,
		},
		{
			name: "rune",
			x:    rune('A'),
			want: false,
		},
		{
			name: "float32",
			x:    float32(6.16),
			want: false,
		},
		{
			name: "float64",
			x:    float64(6.16),
			want: false,
		},
		{
			name: "bool",
			x:    true,
			want: false,
		},
		{
			name: "slice",
			x:    []string{"foo", "bar"},
			want: false,
		},
		{
			name: "array",
			x:    [2]string{"foo", "bar"},
			want: false,
		},
		{
			name: "channel",
			x:    make(chan bool),
			want: false,
		},
		{
			name: "func",
			x:    func() {},
			want: false,
		},
		{
			name: "context.Background()",
			x:    context.Background(),
			want: true,
		},
		{
			name: "context.TODO()",
			x:    context.TODO(),
			want: true,
		},
		{
			name: "custom context",
			x:    context.WithValue(context.Background(), ctxKey, "foo"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Any()
			require.Implements(t, (*gomock.Matcher)(nil), m)

			got := m.Matches(tt.x)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, "is a context.Context", m.String())
		})
	}
}
