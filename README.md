# Java equivalent stream APIs in go

- This library provides a set of APIs to perform operations on a collection of elements similar to Java Stream APIs.

- Similar to Java, intermediate operations are lazy and terminal operations are eager i.e. intermediate operations are not performed until a terminal operation is called.

- when intermediate operations are performed, a new stream is returned, and the original stream is not modified.

- Intermediate operations that can short circuit will halt once the condition is met, and the rest of the elements will not be processed. To avoid go routine leaks, the stream is cleared after obtaining the result.

- The library make use of go routines to perform operations in parallel.
- The library is designed to be used with a collection of elements, and not with a channel.
- The library is not thread safe.
- The library is not designed to be used with infinite streams.
- Stream can only be used once, and it is not reusable.


## Usage

```go
package main

import (
	"fmt"

	"github.com/vkumbhar94/go-streams/pkg/streams"
)

func main() {
	// var args example of using streams
	fmt.Println(
		streams.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
			Filter(
				func(i int) bool {
					return i%2 == 0
				}).
			Map(func(i int) int {
				return i * 2
			}).
			Limit(3).
			Collect(),
	)

	// from slice example of using streams
	fmt.Println(
		streams.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
			Filter(
				func(i int) bool {
					return i%2 == 0
				},
			).
			Count(),
	)

	// from slice example of using streams
	fmt.Println(
		streams.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
			Filter(
				func(i int) bool {
					return i%2 == 0
				},
			).
			Reduce(0,
				func(sum, i int) int {
					return sum + i
				},
			),
	)

	fmt.Println("stream peek start")
	// from slice example of using streams
	fmt.Println(
		streams.ToNumberStream(
			streams.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
				Filter(
					func(i int) bool {
						return i%2 == 0
					},
				),
		).
			Reverse().
			Map(
				func(i int) int {
					return i * 2
				},
			).
			Skip(1).
			Limit(3).
			Filter(func(i int) bool {
				return i%2 == 0
			}).
			Peek(
				func(i int) {
					fmt.Println("before", i)
				},
			).
			DropWhile(
				func(i int) bool {
					return i < 5
				},
			).
			Peek(
				func(i int) {
					fmt.Println("after", i)
				},
			).
			Sum(),
	)
	fmt.Println("stream peek end")

	fmt.Println("for each start")
	// from slice example of using streams
	streams.ToNumberStream(
		streams.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
			Filter(
				func(i int) bool {
					return i%2 == 0
				},
			),
	).ForEach(func(i int) {
		fmt.Println(i)
	})
	fmt.Println("for each end")

	// from map example of using streams
	fmt.Println(
		streams.FromMap(map[int]string{1: "one", 2: "two", 3: "three"}).
			Filter(
				func(e streams.MapEntry[int, string]) bool {
					return e.V == "two"
				},
			).
			Collect(),
	)
}

```