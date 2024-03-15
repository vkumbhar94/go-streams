package streams

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
