package streams

import (
	"cmp"
	"sort"
	"sync/atomic"
)

type MapFun[T, R any] func(T) R

type FilterFun[T any] func(T) bool

type Stream[T any] struct {
	data chan T
	run  func()
	ran  atomic.Bool
}

func (s *Stream[T]) Run() {
	if s.ran.CompareAndSwap(false, true) {
		go s.run()
	}
}

func New[T any](data ...T) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			func(tasks []T) {
				defer close(ch)
				for _, task := range tasks {
					ch <- task
				}
			}(data)
		},
	}
}

func Map[T, R any](s *Stream[T], mapper MapFun[T, R]) *Stream[R] {
	ch := make(chan R)
	return &Stream[R]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				r := mapper(t)
				ch <- r
			}
		},
	}
}

func Filter[T any](s *Stream[T], filter FilterFun[T]) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				if filter(t) {
					ch <- t
				}
			}
		},
	}
}

func Limit[T any](s *Stream[T], i int) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			func() {
				defer close(ch)
				for t := range s.data {
					if i > 0 {
						ch <- t
						i--
					} else {
						break
					}
				}
			}()

			for range s.data {
			}

		},
	}
}

type SortOrder int

const (
	ASC SortOrder = iota
	DESC
)

func Sorted[T cmp.Ordered](s *Stream[T], order SortOrder) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			result := Collect(s)
			sort.Slice(result, func(i, j int) bool {
				if order == DESC {
					return result[i] > result[j]
				}
				return result[i] < result[j]
			})
			for _, r := range result {
				ch <- r
			}
		},
	}
}

func Reduce[T any, R any](s *Stream[T], result R, f func(ans R, i T) R) R {
	s.Run()
	for t := range s.data {
		result = f(result, t)
	}
	return result
}

func ForEach[T any](stream *Stream[T], f func(i T)) {
	stream.Run()
	for t := range stream.data {
		f(t)
	}
}
