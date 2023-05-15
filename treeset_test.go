package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewTreeSet(t *testing.T) {
	var tests = []struct {
		args, want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 1}, []int{1, 2}},
	}
	for i, test := range tests {
		s := NewTreeSet(test.args...)
		if got := s.Elements(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestContainsTS(t *testing.T) {
	var tests = []struct {
		s    *TreeSet[int]
		v    int
		want bool
	}{
		{NewTreeSet[int](), 1, false},
		{NewTreeSet(1), 1, true},
		{NewTreeSet(2), 1, false},
		{NewTreeSet(2, 3), 1, false},
		{NewTreeSet(2, 3), 2, true},
	}
	for i, test := range tests {
		if got := test.s.Contains(test.v); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestAddTS(t *testing.T) {
	s := NewTreeSet(1)
	if s.Add(1) {
		t.Error("got true, want false")
	}
	if !s.Add(2) {
		t.Error("got false, want true")
	}
	want := []int{1, 2}
	if got := s.Elements(); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRemoveTS(t *testing.T) {
	s := NewTreeSet(1)
	if s.Remove(2) {
		t.Error("got true, want false")
	}
	if !s.Remove(1) {
		t.Error("got false, want true")
	}
	if s.Remove(1) {
		t.Error("got true, want false")
	}
	if got := s.Elements(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("got %v want empty map", got)
	}
}

func TestIsEmptyTS(t *testing.T) {
	s := NewTreeSet[int]()
	if !s.IsEmpty() {
		t.Error("got false, want true")
	}
	s.Add(1)
	if s.IsEmpty() {
		t.Error("got true, want false")
	}
}

func TestCardinalityTS(t *testing.T) {
	s := NewTreeSet[int]()
	if n := s.Cardinality(); n != 0 {
		t.Errorf("got %d, want 0", n)
	}
	s.Add(1)
	if n := s.Cardinality(); n != 1 {
		t.Errorf("got %d, want 1", n)
	}
}

func TestUnionTS(t *testing.T) {
	var tests = []struct {
		s1, s2 *TreeSet[int]
		want   []int
	}{
		{NewTreeSet[int](), NewTreeSet[int](), []int{}},
		{NewTreeSet[int](), NewTreeSet(1, 2), []int{1, 2}},
		{NewTreeSet(1, 2), NewTreeSet(2, 3), []int{1, 2, 3}},
		{NewTreeSet(1, 2), NewTreeSet(3, 4), []int{1, 2, 3, 4}},
		{NewTreeSet(1, 2), NewTreeSet(1, 2), []int{1, 2}},
	}
	for i, test := range tests {
		got1 := test.s1.Union(test.s2).(*TreeSet[int]).Elements()
		got2 := test.s2.Union(test.s1).(*TreeSet[int]).Elements()
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestUnion2TS(t *testing.T) {
	var tests = []struct {
		s1, s2 Set[int]
		want   []int
	}{
		{NewTreeSet[int](), NewMapSet[int](), []int{}},
		{NewTreeSet[int](), NewMapSet(1, 2), []int{1, 2}},
		{NewTreeSet(1, 2), NewMapSet(2, 3), []int{1, 2, 3}},
		{NewTreeSet(1, 2), NewMapSet(3, 4), []int{1, 2, 3, 4}},
		{NewTreeSet(1, 2), NewMapSet(1, 2), []int{1, 2}},
	}
	for i, test := range tests {
		got1 := test.s1.Union(test.s2).Elements()
		got2 := test.s2.Union(test.s1).Elements()
		sort.Ints(got2)
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestIntersectionTS(t *testing.T) {
	var tests = []struct {
		s1, s2 *TreeSet[int]
		want   []int
	}{
		{NewTreeSet[int](), NewTreeSet[int](), []int{}},
		{NewTreeSet[int](), NewTreeSet(1, 2), []int{}},
		{NewTreeSet(1, 2), NewTreeSet(2, 3), []int{2}},
		{NewTreeSet(1, 2), NewTreeSet(3, 4), []int{}},
		{NewTreeSet(1, 2), NewTreeSet(1, 2), []int{1, 2}},
	}
	for i, test := range tests {
		got1 := test.s1.Intersection(test.s2).Elements()
		got2 := test.s2.Intersection(test.s1).Elements()
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestDifferenceTS(t *testing.T) {
	var tests = []struct {
		s1, s2 *TreeSet[int]
		want   []int
	}{
		{NewTreeSet[int](), NewTreeSet[int](), []int{}},
		{NewTreeSet[int](), NewTreeSet(1, 2), []int{}},
		{NewTreeSet(1, 2), NewTreeSet(2, 3), []int{1}},
		{NewTreeSet(1, 2), NewTreeSet(3, 4), []int{1, 2}},
		{NewTreeSet(1, 2), NewTreeSet(1, 2), []int{}},
	}
	for i, test := range tests {
		if got := test.s1.Difference(test.s2).Elements(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestSymDifferenceTS(t *testing.T) {
	var tests = []struct {
		s1, s2 *TreeSet[int]
		want   []int
	}{
		{NewTreeSet[int](), NewTreeSet[int](), []int{}},
		{NewTreeSet[int](), NewTreeSet(1, 2), []int{1, 2}},
		{NewTreeSet(1, 2), NewTreeSet(2, 3), []int{1, 3}},
		{NewTreeSet(1, 2), NewTreeSet(3, 4), []int{1, 2, 3, 4}},
		{NewTreeSet(1, 2), NewTreeSet(1, 2), []int{}},
	}
	for i, test := range tests {
		got1 := test.s1.SymDifference(test.s2).Elements()
		got2 := test.s2.SymDifference(test.s1).Elements()
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestSymDifference2TS(t *testing.T) {
	var tests = []struct {
		s1, s2 Set[int]
		want   []int
	}{
		{NewTreeSet[int](), NewMapSet[int](), []int{}},
		{NewTreeSet[int](), NewMapSet(1, 2), []int{1, 2}},
		{NewTreeSet(1, 2), NewMapSet(2, 3), []int{1, 3}},
		{NewTreeSet(1, 2), NewMapSet(3, 4), []int{1, 2, 3, 4}},
		{NewTreeSet(1, 2), NewMapSet(1, 2), []int{}},
	}
	for i, test := range tests {
		got1 := test.s1.SymDifference(test.s2).Elements()
		got2 := test.s2.SymDifference(test.s1).Elements()
		sort.Ints(got2)
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestIsSubsetTS(t *testing.T) {
	var tests = []struct {
		s1, s2 *TreeSet[int]
		want   bool
	}{
		{NewTreeSet[int](), NewTreeSet[int](), true},
		{NewTreeSet[int](), NewTreeSet(1), true},
		{NewTreeSet(1), NewTreeSet(1, 2), true},
		{NewTreeSet(1, 2), NewTreeSet(1, 2), true},
		{NewTreeSet(1, 2), NewTreeSet(1), false},
		{NewTreeSet(1), NewTreeSet[int](), false},
		{NewTreeSet(1, 2), NewTreeSet(1, 3, 4), false},
		{NewTreeSet(1, 2), NewTreeSet(1, 2, 4), true},
	}
	for i, test := range tests {
		if got := test.s1.IsSubset(test.s2); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestIsProperSubsetTS(t *testing.T) {
	var tests = []struct {
		s1, s2 *TreeSet[int]
		want   bool
	}{
		{NewTreeSet[int](), NewTreeSet[int](), false},
		{NewTreeSet[int](), NewTreeSet(1), true},
		{NewTreeSet(1), NewTreeSet(1, 2), true},
		{NewTreeSet(1, 2), NewTreeSet(1, 2), false},
		{NewTreeSet(1, 2), NewTreeSet(1), false},
		{NewTreeSet(1), NewTreeSet[int](), false},
		{NewTreeSet(1, 2), NewTreeSet(1, 3, 4), false},
		{NewTreeSet(1, 2), NewTreeSet(1, 2, 4), true},
	}
	for i, test := range tests {
		if got := test.s1.IsProperSubset(test.s2); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestEqualTS(t *testing.T) {
	var tests = []struct {
		s1, s2 *TreeSet[int]
		want   bool
	}{
		{NewTreeSet[int](), NewTreeSet[int](), true},
		{NewTreeSet[int](), NewTreeSet(1), false},
		{NewTreeSet(1), NewTreeSet(1, 2), false},
		{NewTreeSet(1, 2), NewTreeSet(1, 2), true},
	}
	for i, test := range tests {
		got1 := test.s1.Equal(test.s2)
		got2 := test.s2.Equal(test.s1)
		if got1 != test.want || got2 != test.want {
			t.Errorf("%d: got %t and %t, want %t", i, got1, got2, test.want)
		}
	}
}

func TestCloneTS(t *testing.T) {
	s := NewTreeSet(1, 2, 3)
	clone := s.Clone()
	want := []int{1, 2, 3}
	if got := clone.Elements(); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestElementsTS(t *testing.T) {
	var tests = []struct {
		args []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1, 1}, []int{1, 2}},
	}
	for i, test := range tests {
		s := NewTreeSet(test.args...)
		sl := s.Elements()
		if !reflect.DeepEqual(sl, test.want) {
			t.Errorf("%d: got %v, want %v", i, s.Elements(), test.want)
		}
	}
}

func TestStringTS(t *testing.T) {
	s := NewTreeSet(1, 2)
	want := "TreeSet{1, 2}"
	if got := s.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
