<h1 align="center">
  gomockctx
</h1>

<p align="center">
  <strong>
    Go package with <a href="https://github.com/golang/mock">gomock</a> helpers
    for matching <a href="https://pkg.go.dev/context">context.Context</a>.
  </strong>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/romdo/gomockctx"><img src="https://img.shields.io/badge/%E2%80%8B-reference-387b97.svg?logo=go&logoColor=white" alt="Go Reference"></a>
  <a href="https://github.com/romdo/gomockctx/actions"><img src="https://img.shields.io/github/workflow/status/romdo/gomockctx/CI.svg?logo=github" alt="Actions Status"></a>
  <a href="https://codeclimate.com/github/romdo/gomockctx"><img src="https://img.shields.io/codeclimate/coverage/romdo/gomockctx.svg?logo=code%20climate" alt="Coverage"></a>
  <a href="https://github.com/romdo/gomockctx/issues"><img src="https://img.shields.io/github/issues-raw/romdo/gomockctx.svg?style=flat&logo=github&logoColor=white" alt="GitHub issues"></a>
  <a href="https://github.com/romdo/gomockctx/pulls"><img src="https://img.shields.io/github/issues-pr-raw/romdo/gomockctx.svg?style=flat&logo=github&logoColor=white" alt="GitHub pull requests"></a>
  <a href="https://github.com/romdo/gomockctx/blob/main/LICENSE"><img src="https://img.shields.io/github/license/romdo/gomockctx.svg?style=flat" alt="License Status"></a>
</p>

## Import

```go
import "github.com/romdo/gomockctx"
```

## Usage

Match against a specific context or any of its child contexts:

```go
// Create a context with a gomockctx ID value.
ctx := gomockctx.New(context.Background())

// Match against a context with a gomockctx ID.
someMock.EXPECT().
	Get(gomockctx.Eq(ctx), "foo").
	Return("bar", nil)

// Use context with gomockctx ID when calling function.
someMock.Get(ctx, "foo")
```

Match against a context containing a specific value:

```go
someMock.EXPECT().
	Get(gomockctx.WithValue(myCtxKey, "hello"), "foo").
	Return("bar", nil)
```

Match against any context:

```go
someMock.EXPECT().
	Get(gomockctx.Any(), "foo").
	Return("bar", nil)
```

## Documentation

Please see the
[Go Reference](https://pkg.go.dev/github.com/romdo/gomockctx#section-documentation).

## License

[MIT](https://github.com/romdo/gomockctx/blob/main/LICENSE)
