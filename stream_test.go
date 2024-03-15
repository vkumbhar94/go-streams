package streams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamCollect(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	collected := Collect(stream)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, collected)
}

func TestStreamMap(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	mapped := Map(stream, func(i int) int {
		return i * 2
	})
	collected := Collect(mapped)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, collected)
}

func TestStreamFilter(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	filtered := Filter(stream, func(i int) bool {
		return i%2 == 0
	})
	collected := Collect(filtered)
	assert.Equal(t, []int{2, 4}, collected)
}

func TestStreamLimit(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	limited := Limit(stream, 3)
	collected := Collect(limited)
	assert.Equal(t, []int{1, 2, 3}, collected)
}

func TestStreamSorted(t *testing.T) {
	stream := New(5, 3, 1, 4, 2)
	sorted := Sorted(stream, ASC)
	collected := Collect(sorted)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, collected)
}
func TestStreamReverseSorted(t *testing.T) {
	stream := New(5, 3, 1, 4, 2)
	sorted := Sorted(stream, DESC)
	collected := Collect(sorted)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, collected)
}

func TestStreamReduce(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	reduced := Reduce(stream, 0, func(ans, i int) int {
		return ans + i
	})
	assert.Equal(t, 15, reduced)
}

func TestStreamSum(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	sum := Sum(stream)
	assert.Equal(t, 15, sum)
}
func TestFloatStreamSum(t *testing.T) {
	stream := New(1.1, 2.4, 3.9, 4.4, 5.8)
	sum := Sum(stream)
	assert.Equal(t, 17.6, sum)
}

func TestStreamCount(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	count := Count(stream)
	assert.Equal(t, int64(5), count)
}

func TestStreamForEach(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	var sum int
	ForEach(stream, func(i int) {
		sum += i
	})
	assert.Equal(t, 15, sum)
}

func TestStreamForEach2(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	ForEach(stream, func(i int) {
		fmt.Println(i)
	})
}

func TestStreamCollectToSet(t *testing.T) {
	stream := New(1, 2, 3, 3, 2)
	collected := CollectToSet(stream)
	assert.Equal(t, map[int]struct{}{1: {}, 2: {}, 3: {}}, collected)
}

func TestStreamDistinct(t *testing.T) {
	stream := New(1, 2, 3, 3, 2)
	collected := Collect(Distinct(stream))
	assert.Equal(t, []int{1, 2, 3}, collected)
}

func TestAllMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	allMatch := AllMatch(stream, func(i int) bool {
		return i < 10
	})
	assert.True(t, allMatch)
}
func TestFalseAllMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	allMatch := AllMatch(stream, func(i int) bool {
		return i < 3
	})
	assert.False(t, allMatch)
}

func TestNotAllMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	allMatch := NotAllMatch(stream, func(i int) bool {
		return i < 10
	})
	assert.False(t, allMatch)
}

func TestAnyMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	anyMatch := AnyMatch(stream, func(i int) bool {
		return i == 3
	})
	assert.True(t, anyMatch)
}

func TestNotAnyMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	anyMatch := AnyMatch(stream, func(i int) bool {
		return i == 10
	})
	assert.False(t, anyMatch)
}

func TestDropWhile(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	collected := Collect(DropWhile(stream, func(i int) bool {
		return i < 3
	}))
	assert.Equal(t, []int{3, 4, 5}, collected)
}

func TestTakeWhile(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	collected := Collect(TakeWhile(stream, func(i int) bool {
		return i < 3
	}))
	assert.Equal(t, []int{1, 2}, collected)
}

func TestNoneMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	noneMatch := NoneMatch(stream, func(i int) bool {
		return i == 10
	})
	assert.True(t, noneMatch)
}

func TestFalseNoneMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	noneMatch := NoneMatch(stream, func(i int) bool {
		return i == 2
	})
	assert.False(t, noneMatch)
}

func TestPeek(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	collected := Collect(Peek(stream, func(i int) {
		fmt.Println(i)
	}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, collected)
}

func TestFindFirst(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	first := FindFirst(stream)
	assert.Equal(t, 1, *first)
}

func TestFindFirstEmpty(t *testing.T) {
	stream := New[int]()
	first := FindFirst(stream)
	assert.Nil(t, first)
}

func TestFlatmap(t *testing.T) {
	stream := New([]int{1, 2, 3}, []int{4, 5}, []int{6, 7, 8, 9})
	flatMapped := FlatMap(stream, func(i []int) *Stream[int] {
		return New(i...)
	})
	collected := Collect(flatMapped)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, collected)
}

func TestMin(t *testing.T) {
	stream := New(5, 3, 1, 4, 2)
	minVal := Min(stream)
	assert.Equal(t, 1, *minVal)
}

func TestMax(t *testing.T) {
	stream := New(5, 3, 1, 4, 2)
	maxVal := Max(stream)
	assert.Equal(t, 5, *maxVal)
}

func TestSkip(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	collected := Collect(Skip(stream, 2))
	assert.Equal(t, []int{3, 4, 5}, collected)
}

func TestIfAllMatch(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	IfAllMatch(stream,
		func(i int) bool {
			return i < 10
		},
		func(t int) {
			fmt.Println(t)
		},
	)
}

func TestIfAllMatchNegativeCase(t *testing.T) {
	stream := New(1, 2, 3, 4, 5)
	IfAllMatch(stream,
		func(i int) bool {
			return i > 10
		},
		func(t int) {
			fmt.Println(t)
		},
	)
}
