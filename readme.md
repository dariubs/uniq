go uniq
=======

![test workflow](https://github.com/dariubs/uniq/actions/workflows/test.yml/badge.svg) [![Go Reference](https://pkg.go.dev/badge/github.com/dariubs/uniq.svg)](https://pkg.go.dev/github.com/dariubs/uniq) [![Go Report Card](https://goreportcard.com/badge/github.com/dariubs/uniq)](https://goreportcard.com/report/github.com/dariubs/uniq)

make slice items unique in go

install
-------

```
go get github.com/dariubs/uniq
```

example
-------

```go
package main

import (
	"fmt"

	"github.com/dariubs/uniq"
)

func main() {
	wstring := []string{"hello", "hello", "holla", "hello", "hi", "holla", "hi"}
	wuniq := uniq.UniqString(wstring)
	fmt.Printf("%v\n", wuniq) // [hello holla hi]

	xint := []int{1, 2, 2, 3, 4, 5, 5, 5, 6, 6, 7}
	xuniq := uniq.UniqInt(xint)
	fmt.Printf("%v\n", xuniq) // [1 2 3 4 5 6 7]

	yuint := []uint{1, 2, 2, 3, 4, 5, 5, 5, 6, 6, 7}
	yuniq := uniq.UniqUint(yuint)
	fmt.Printf("%v\n", yuniq) // [1 2 3 4 5 6 7]
}

```

by
--

- [Dariush Abbsi](https://github.com/dariubs)

License
-------

[MIT](license)
