# option

[![Go Reference](https://pkg.go.dev/badge/github.com/wlrgo/option.svg)](https://pkg.go.dev/github.com/wlrgo/option)

`Option[T]` for values that may be absent.

This package is experimental. It follows the Rust `Option` API as closely as Go allows, and adds a few Go-oriented helpers such as `Get`, `FromPtr`, and `OkOr`.

```bash
go get github.com/wlrgo/option
```

```go
package main

import (
	"fmt"

	"github.com/wlrgo/option"
)

func main() {
	opt := option.Some(2)
	if v, ok := opt.Get(); ok {
		fmt.Println(v)
	}
}
```

See [pkg.go.dev/github.com/wlrgo/option](https://pkg.go.dev/github.com/wlrgo/option) for the full API.
