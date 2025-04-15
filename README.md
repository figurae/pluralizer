# pluralizer
Simple pluralizer lib written in Go

# Usage
```go
package main

import (
	"fmt"
	"github.com/figurae/pluralizer"
)

func main() {
	singular := "cat"
	plural := pluralizer.ToPlural(singular)
	fmt.Println(plural)
}
```
