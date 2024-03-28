package streams

import "testing"

func TestNumberStream_Average(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	ns := ToNumberStream(s)
	if ns.Average() != 3 {
		t.Fail()
	}
}
func TestNumberStream_AverageEmpty(t *testing.T) {
	s := New([]int{}...)
	ns := ToNumberStream(s)
	if ns.Average() != 0 {
		t.Fail()
	}
}

func TestNumberStream_Count(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	ns := ToNumberStream(s)
	if ns.Count() != 5 {
		t.Fail()
	}
}

func TestNumberStream_Max(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	ns := ToNumberStream(s)
	if *ns.Max() != 5 {
		t.Fail()
	}
}

func TestNumberStream_Min(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	ns := ToNumberStream(s)
	if *ns.Min() != 1 {
		t.Fail()
	}
}

func TestNumberStream_Sum(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	ns := ToNumberStream(s)
	if ns.Sum() != 15 {
		t.Fail()
	}
}
