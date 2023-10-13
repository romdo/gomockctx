package gomockctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWithValue(t *testing.T) {
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *valueMatcher
	}{
		{
			name: "nil",
			args: args{
				key:   nil,
				value: nil,
			},
			want: &valueMatcher{
				key:   nil,
				value: nil,
			},
		},
		{
			name: "string",
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: &valueMatcher{
				key:   "foo",
				value: "bar",
			},
		},
		{
			name: "gomockctx ctxKey",
			args: args{
				key:   ctxKey,
				value: "FrAcGnKKpVk1rB3AWC9S8Dnff04svNtN",
			},
			want: &valueMatcher{
				key:   ctxKey,
				value: "FrAcGnKKpVk1rB3AWC9S8Dnff04svNtN",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithValue(tt.args.key, tt.args.value)

			assert.Implements(t, (*gomock.Matcher)(nil), got)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_valueMatcher_Matches(t *testing.T) {
	ctx := context.Background()

	type stringKey string
	strKey := stringKey("strKey")
	strKeyVal1 := "foo"
	strKeyVal2 := "bar"
	ctxWithStrKey := context.WithValue(ctx, strKey, strKeyVal1)

	structKey := struct{}{}
	structKeyVal1 := struct{ name string }{name: "foo"}
	structKeyVal2 := struct{ name string }{name: "bar"}
	ctxWithStructKey := context.WithValue(ctx, structKey, structKeyVal1)

	ctxKeyVal1 := "XAzb0Cr7yLuLzO369vTodjxKL3GUpspE"
	ctxKeyVal2 := "r11X0FOejbPamvLiWhAuGiSqXzdmGnIm"
	ctxWithCtxKey := context.WithValue(ctx, ctxKey, ctxKeyVal1)

	type fields struct {
		key   interface{}
		value interface{}
	}
	type args struct {
		x interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: nil},
			want: false,
		},
		{
			name: "empty string",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: ""},
			want: false,
		},
		{
			name: "string",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: "hello world"},
			want: false,
		},
		{
			name: "int",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: int(616)},
			want: false,
		},
		{
			name: "int8",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: int8(64)},
			want: false,
		},
		{
			name: "int16",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: int16(1024)},
			want: false,
		},
		{
			name: "int32",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: int32(1123456789)},
			want: false,
		},
		{
			name: "int64",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: int64(16123456789)},
			want: false,
		},
		{
			name: "uint",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: uint(616)},
			want: false,
		},
		{
			name: "uint8",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: uint8(64)},
			want: false,
		},
		{
			name: "uint16",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: uint16(1024)},
			want: false,
		},
		{
			name: "uint32",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: uint32(1123456789)},
			want: false,
		},
		{
			name: "uint64",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: uint64(16123456789)},
			want: false,
		},
		{
			name: "byte",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: byte('A')},
			want: false,
		},
		{
			name: "rune",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: rune('A')},
			want: false,
		},
		{
			name: "float32",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: float32(6.16)},
			want: false,
		},
		{
			name: "float64",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: float64(6.16)},
			want: false,
		},
		{
			name: "bool",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: 616},
			want: false,
		},
		{
			name: "slice",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: []string{"foo", "bar"}},
			want: false,
		},
		{
			name: "array",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: [2]string{"foo", "bar"}},
			want: false,
		},
		{
			name: "channel",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: make(chan bool)},
			want: false,
		},
		{
			name: "func",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: func() {}},
			want: false,
		},
		{
			name: "context with strKey",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: ctxWithStrKey},
			want: true,
		},
		{
			name: "context with different strKey",
			fields: fields{
				key:   strKey,
				value: strKeyVal2,
			},
			args: args{x: ctxWithStrKey},
			want: false,
		},
		{
			name: "context without strKey",
			fields: fields{
				key:   strKey,
				value: strKeyVal1,
			},
			args: args{x: ctx},
			want: false,
		},
		{
			name: "context with structKey",
			fields: fields{
				key:   structKey,
				value: structKeyVal1,
			},
			args: args{x: ctxWithStructKey},
			want: true,
		},
		{
			name: "context with different structKey",
			fields: fields{
				key:   structKey,
				value: structKeyVal2,
			},
			args: args{x: ctxWithStructKey},
			want: false,
		},
		{
			name: "context without structKey",
			fields: fields{
				key:   structKey,
				value: structKeyVal1,
			},
			args: args{x: ctx},
			want: false,
		},
		{
			name: "context with ctxKey",
			fields: fields{
				key:   ctxKey,
				value: ctxKeyVal1,
			},
			args: args{x: ctxWithCtxKey},
			want: true,
		},
		{
			name: "context with different ctxKey",
			fields: fields{
				key:   ctxKey,
				value: ctxKeyVal2,
			},
			args: args{x: ctxWithCtxKey},
			want: false,
		},
		{
			name: "context without ctxKey",
			fields: fields{
				key:   ctxKey,
				value: ctxKeyVal1,
			},
			args: args{x: ctx},
			want: false,
		},
		{
			name: "context with ctxKey and empty value",
			fields: fields{
				key:   ctxKey,
				value: contextValue(""),
			},
			args: args{x: ctxWithCtxKey},
			want: false,
		},
		{
			name: "context with different ctxKey and empty value",
			fields: fields{
				key:   ctxKey,
				value: contextValue(""),
			},
			args: args{x: ctxWithCtxKey},
			want: false,
		},
		{
			name: "context without ctxKey and empty value",
			fields: fields{
				key:   ctxKey,
				value: contextValue(""),
			},
			args: args{x: ctx},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &valueMatcher{
				key:   tt.fields.key,
				value: tt.fields.value,
			}

			got := vm.Matches(tt.args.x)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_valueMatcher_String(t *testing.T) {
	type stringKey string
	type structKey struct{ name string }

	type fields struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "stringKey",
			fields: fields{
				key:   stringKey("foo"),
				value: "hello world",
			},
			want: `context with "foo" = "hello world"`,
		},
		{
			name: "structKey",
			fields: fields{
				key:   structKey{name: "bar"},
				value: "okay then",
			},
			want: `context with "{name:bar}" = "okay then"`,
		},
		{
			name: "gomockctx ctxKey",
			fields: fields{
				key:   ctxKey,
				value: "foobar",
			},
			want: `context with "gomockctx ID" = "foobar"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &valueMatcher{
				key:   tt.fields.key,
				value: tt.fields.value,
			}

			got := vm.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
