# Go streams APIs
Yet another go streams api using generics


## Usage

```go
package main

import (
	"fmt"
	
	"github.com/vkumbhar94/go-streams"
)

func main() {
	stream := streams.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	filtered := streams.Filter(stream, func(i int) bool {
		return i%2 == 0
	})
	mapped := streams.Map(filtered, func(i int) int {
		return i * 2
	})
	limited := streams.Limit(mapped, 3)
	
	collected := streams.Collect(limited)
	fmt.Println(collected)
}

```