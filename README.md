[![GoDoc](https://godoc.org/github.com/arnumina/logfmt?status.svg)](https://godoc.org/github.com/arnumina/logfmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/arnumina/logfmt)](https://goreportcard.com/report/github.com/arnumina/logfmt)
![CI](https://github.com/arnumina/logfmt/workflows/CI/badge.svg)

# logfmt

This module allows to encode a list of key/value pairs in **logfmt** format.

## Example

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/arnumina/logfmt"
)

func main() {
	buf := bytes.Buffer{}
	logfmt.Encode(&buf, "year", 2020, "month", "July", "Tuesday", true)

	fmt.Println(buf.String())
}
```
```
year=2020 month="July" Tuesday=true
```

## Link

- [logfmt](https://brandur.org/logfmt)

---
Copyright (c) 2020 Institut National de l'Audiovisuel
