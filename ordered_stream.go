package streams

import (
	"cmp"
	"sort"
	"sync/atomic"
)

type OrderedStream[T cmp.Ordered] struct {
	Stream[T]
}

func ToOrderedStream[T cmp.Ordered](s *Stream[T]) *OrderedStream[T] {
	return &OrderedStream[T]{
		Stream: Stream[T]{
			data: s.data,
			run:  s.run,
			ran:  atomic.Bool{},
		},
	}
}

func (s *OrderedStream[T]) Sorted(order SortOrder) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			var data []T
			for t := range s.data {
				data = append(data, t)
			}
			if order == DESC {
				sort.Slice(data, func(i, j int) bool {
					return data[i] > data[j]
				})
			} else {
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			}
			for _, t := range data {
				ch <- t
			}
		},
	}
}
