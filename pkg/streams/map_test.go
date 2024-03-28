package streams

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMStreamCollect(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"ijk": 3,
		"xyz": 4,
	})
	collected := MCollect(stream)
	assert.True(t, reflect.DeepEqual(collected, map[string]int64{
		"abc": 2,
		"ijk": 3,
		"xyz": 4,
	}))
}

func TestMStreamMap(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"ijk": 3,
		"xyz": 4,
	})
	mapped := Map(stream, func(e MapEntry[string, int64]) MapEntry[string, int64] {
		e.V *= 2
		return e
	})
	collected := MCollect(mapped)
	assert.True(t, reflect.DeepEqual(collected, map[string]int64{
		"abc": 4,
		"ijk": 6,
		"xyz": 8,
	}))
}

func TestMStreamFilter(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"ijk": 3,
		"xyz": 4,
	})
	filtered := Filter(stream, func(e MapEntry[string, int64]) bool {
		return e.V > 2
	})
	collected := MCollect(filtered)
	assert.True(t, reflect.DeepEqual(collected, map[string]int64{
		"ijk": 3,
		"xyz": 4,
	}))
}

func TestMStreamLimit(t *testing.T) {
	stream := MNew(map[string]int64{
		"ijk":   3,
		"xyz":   4,
		"abc":   2,
		"abcd":  2,
		"abcde": 2,
	})
	limited := Limit(stream, 2)
	collected := MCollect(limited)
	// unless sorted, the result is not deterministic
	assert.Len(t, collected, 2)
}

func TestMStreamSorted(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	sorted := MSorted(stream)
	collected := MCollect(sorted)
	assert.True(t, reflect.DeepEqual(collected, map[string]int64{
		"abc": 2,
		"ijk": 3,
		"xyz": 4,
	}))
}

func TestMStreamKeys(t *testing.T) {
	stream := MNewKeys(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	collected := Collect(Sorted(stream, ASC))
	assert.Equal(t, []string{"abc", "ijk", "xyz"}, collected)
}

func TestMStreamValues(t *testing.T) {
	stream := MNewValues(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	collected := Collect(Sorted(stream, ASC))
	assert.Equal(t, []int64{2, 3, 4}, collected)
}

func TestMStreamReduceValues(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	reduced := Reduce(stream, 0, func(ans int64, e MapEntry[string, int64]) int64 {
		return ans + e.V
	})
	assert.Equal(t, int64(9), reduced)
}

func TestMStreamReduceValues2(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	reduced := Reduce(MSorted(stream), "", func(ans string, e MapEntry[string, int64]) string {
		if ans == "" {
			return strconv.FormatInt(e.V, 10)
		}
		return ans + "," + strconv.FormatInt(e.V, 10)
	})
	assert.Equal(t, "2,3,4", reduced)
}

func TestMStreamReduceStringKeysToInt64(t *testing.T) {
	stream := MNew(map[string]int64{
		"abcd": 2,
		"xyz":  4,
		"ijkl": 3,
	})
	reduced := Reduce(stream, 0, func(ans int64, e MapEntry[string, int64]) int64 {
		return ans + int64(len(e.K))
	})
	assert.Equal(t, int64(11), reduced)
}
func TestMStreamReduceStringKeysToString(t *testing.T) {
	stream := MNew(map[string]int64{
		"abcd": 2,
		"xyz":  4,
		"ijkl": 3,
	})

	reduced := Reduce(MSorted(stream), "", func(ans string, e MapEntry[string, int64]) string {
		if ans == "" {
			return e.K
		}
		return ans + "," + e.K
	})
	assert.Equal(t, "abcd,ijkl,xyz", reduced)
}

func TestMStreamSum(t *testing.T) {
	stream := MNewValues(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	sum := Sum(stream)
	assert.Equal(t, int64(9), sum)
}

func TestMStreamCount(t *testing.T) {
	stream := MNewValues(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	count := Count(stream)
	assert.Equal(t, int64(3), count)
}

func TestMStreamSCollectToSet(t *testing.T) {
	stream := MNew(map[string]int64{
		"abcd": 2,
		"xyz":  4,
		"ijkl": 3,
	})
	collected := CollectToSet(stream)
	assert.Equal(t,
		map[MapEntry[string, int64]]struct{}{
			{"abcd", 2}: {},
			{"ijkl", 3}: {},
			{"xyz", 4}:  {},
		},
		collected)
}

func TestMStreamForEachString(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	ForEach(stream, func(e MapEntry[string, int64]) {
		fmt.Println(e)
	})
}

func TestMStreamForEachKeyValue(t *testing.T) {
	stream := MNew(map[string]int64{
		"abc": 2,
		"xyz": 4,
		"ijk": 3,
	})
	ForEach(stream, func(e MapEntry[string, int64]) {
		fmt.Println(e.Key(), e.Value())
	})
}
