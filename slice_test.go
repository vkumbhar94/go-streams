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
