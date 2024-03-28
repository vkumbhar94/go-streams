package streams

import (
	"fmt"
	"reflect"
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

func TestFilter(t *testing.T) {
	collected := New([]int{1, 2, 3, 4, 5}...).Filter(func(i int) bool {
		return i > 2
	}).Collect()
	if !reflect.DeepEqual(collected, []int{3, 4, 5}) {
		t.Errorf("Expected [3, 4, 5] but got %v", collected)
	}
}

func TestLimit(t *testing.T) {
	collected := New([]int{1, 2, 3, 4, 5}...).Limit(3).Collect()
	if !reflect.DeepEqual(collected, []int{1, 2, 3}) {
		t.Errorf("Expected [1, 2, 3] but got %v", collected)
	}
}

func TestForEach(t *testing.T) {
	var sum int
	New([]int{1, 2, 3, 4, 5}...).ForEach(func(i int) {
		sum += i
	})
	if sum != 15 {
		t.Errorf("Expected 15 but got %v", sum)
	}
}

func TestMethodAllMatch(t *testing.T) {
	allMatch := New([]int{1, 2, 3, 4, 5}...).AllMatch(func(i int) bool {
		return i < 10
	})
	if !allMatch {
		t.Errorf("Expected true but got %v", allMatch)
	}
}

func TestMethodFalseAllMatch(t *testing.T) {
	allMatch := New([]int{1, 2, 3, 4, 5}...).AllMatch(func(i int) bool {
		return i < 3
	})
	if allMatch {
		t.Errorf("Expected false but got %v", allMatch)
	}
}

func TestMethodNotAllMatch(t *testing.T) {
	allMatch := New([]int{1, 2, 3, 4, 5}...).NotAllMatch(func(i int) bool {
		return i < 10
	})
	if allMatch {
		t.Errorf("Expected false but got %v", allMatch)
	}
}

func TestMethodAnyMatch(t *testing.T) {
	anyMatch := New([]int{1, 2, 3, 4, 5}...).AnyMatch(func(i int) bool {
		return i == 3
	})
	if !anyMatch {
		t.Errorf("Expected true but got %v", anyMatch)
	}
}
func TestMethodFalseAnyMatch(t *testing.T) {
	anyMatch := New([]int{1, 2, 3, 4, 5}...).AnyMatch(func(i int) bool {
		return i == 10
	})
	if anyMatch {
		t.Errorf("Expected false but got %v", anyMatch)
	}
}

func TestMethodNoneMatch(t *testing.T) {
	noneMatch := New([]int{1, 2, 3, 4, 5}...).NoneMatch(func(i int) bool {
		return i == 6
	})
	if !noneMatch {
		t.Errorf("Expected true but got %v", noneMatch)
	}
}

func TestMethodDropWhile(t *testing.T) {
	collected := New([]int{1, 2, 3, 4, 5}...).DropWhile(func(i int) bool {
		return i < 3
	}).Collect()
	if !reflect.DeepEqual(collected, []int{3, 4, 5}) {
		t.Errorf("Expected [3, 4, 5] but got %v", collected)
	}
}

func TestMethodTakeWhile(t *testing.T) {
	collected := New([]int{1, 2, 3, 4, 5}...).TakeWhile(func(i int) bool {
		return i < 3
	}).Collect()
	if !reflect.DeepEqual(collected, []int{1, 2}) {
		t.Errorf("Expected [1, 2] but got %v", collected)
	}
}

func TestMethodPeek(t *testing.T) {
	var sum int
	collected := New([]int{1, 2, 3, 4, 5}...).Peek(func(i int) {
		sum += i
	}).Collect()
	if !reflect.DeepEqual(collected, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected [1, 2, 3, 4, 5] but got %v", collected)
	}
	if sum != 15 {
		t.Errorf("Expected 15 but got %v", sum)
	}
}

func TestMethodFindFirst(t *testing.T) {
	first := New([]int{1, 2, 3, 4, 5}...).FindFirst()
	if *first != 1 {
		t.Errorf("Expected 1 but got %v", first)
	}
}

func TestMethodFindFirstNil(t *testing.T) {
	first := New([]int{}...).FindFirst()
	if first != nil {
		t.Errorf("Expected nil but got %v", first)
	}
}

func TestMethodFindFirstOr(t *testing.T) {
	first := New([]int{1, 2, 3}...).FindFirstOr().Or(10)
	assert.Equal(t, 1, first)
}
func TestMethodFindFirstOrEmpty(t *testing.T) {
	first := New([]int{}...).FindFirstOr().Or(2)
	assert.Equal(t, 2, first)
}

func TestMethodSkip(t *testing.T) {
	collected := New([]int{1, 2, 3, 4, 5}...).Skip(3).Collect()
	if !reflect.DeepEqual(collected, []int{4, 5}) {
		t.Errorf("Expected [4, 5] but got %v", collected)
	}
}

func TestMethodIfAllMatch(t *testing.T) {
	var sum int
	New([]int{1, 2, 3, 4, 5}...).IfAllMatch(func(i int) bool {
		return i < 10
	}, func(i int) {
		sum += i * 3
	})
	assert.Equal(t, 45, sum)
}

func TestMethodIfAllMatchElse(t *testing.T) {
	var sum int
	New([]int{1, 2, 3, 4, 5}...).IfAllMatch(func(i int) bool {
		return i < 3
	}, func(i int) {
		sum += i
	}).Else(func(i int) {
		sum += i * 2
	})
	assert.Equal(t, 30, sum)
}

func TestMethodMap(t *testing.T) {
	collected := New([]int{1, 2, 3, 4, 5}...).Map(func(i int) int {
		return i * 2
	}).Collect()
	if !reflect.DeepEqual(collected, []int{2, 4, 6, 8, 10}) {
		t.Errorf("Expected [2, 4, 6, 8, 10] but got %v", collected)
	}
}

func TestMethodCollectToSet(t *testing.T) {
	collected := ToComparableStream(New([]int{1, 2, 3, 4, 5}...)).CollectToSet()
	if !reflect.DeepEqual(collected, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}}) {
		t.Errorf("Expected [1, 2, 3, 4, 5] but got %v", collected)
	}
}

func TestMethodDistinct(t *testing.T) {
	collected := ToComparableStream(New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}...)).Distinct().Collect()
	if !reflect.DeepEqual(collected, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected [1, 2, 3, 4, 5] but got %v", collected)
	}

}
func TestMethodDistinctAndThen(t *testing.T) {
	collected := ToComparableStream(New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}...)).DistinctAndThen().CollectToSet()
	if !reflect.DeepEqual(collected, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}}) {
		t.Errorf("Expected [2, 4, 6, 8, 10] but got %v", collected)
	}
}

func TestMethodSorted(t *testing.T) {
	collected := ToOrderedStream(New([]int{5, 3, 1, 4, 2}...)).Sorted(ASC).Collect()
	if !reflect.DeepEqual(collected, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected [1, 2, 3, 4, 5] but got %v", collected)
	}
}
func TestMethodSortedDesc(t *testing.T) {
	collected := ToOrderedStream(New([]int{5, 3, 1, 4, 2}...)).Sorted(DESC).Collect()
	if !reflect.DeepEqual(collected, []int{5, 4, 3, 2, 1}) {
		t.Errorf("Expected [1, 2, 3, 4, 5] but got %v", collected)
	}
}

func TestMethodReduce(t *testing.T) {
	reduced := New([]int{1, 2, 3, 4, 5}...).Reduce(0, func(ans, i int) int {
		return ans + i
	})
	if reduced != 15 {
		t.Errorf("Expected 15 but got %v", reduced)
	}
}

func TestMethodSum(t *testing.T) {
	sum := ToNumberStream(New([]int{1, 2, 3, 4, 5}...)).Sum()
	if sum != 15 {
		t.Errorf("Expected 15 but got %v", sum)
	}
}

func TestMethodCount(t *testing.T) {
	count := New([]int{1, 2, 3, 4, 5}...).Count()
	if count != 5 {
		t.Errorf("Expected 5 but got %v", count)
	}
}

func TestMSorted(t *testing.T) {
	collected := ToMOrderedStream(MNew(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})).Sorted(ASC).Collect()
	assert.True(t, reflect.DeepEqual(collected, []MapEntry[string, int64]{
		{"abc", 2},
		{"ijk", 3},
		{"xyz", 4},
	}))
}
func TestMSortedDesc(t *testing.T) {
	collected := ToMOrderedStream(MNew(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})).Sorted(DESC).Collect()
	assert.True(t, reflect.DeepEqual(collected, []MapEntry[string, int64]{
		{"xyz", 4},
		{"ijk", 3},
		{"abc", 2},
	}))
}

func TestMethodReverse(t *testing.T) {
	collected := New([]int{1, 2, 3, 4, 5}...).Reverse().Collect()
	assert.Equal(t, []int{5, 4, 3, 2, 1}, collected)
}
