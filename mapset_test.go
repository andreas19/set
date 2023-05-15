package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewMapSet(t *testing.T) {
	var tests = []struct {
		args []int
		want map[int]struct{}
	}{
		{[]int{}, map[int]struct{}{}},
		{[]int{1}, map[int]struct{}{1: {}}},
		{[]int{1, 2, 1}, map[int]struct{}{1: {}, 2: {}}},
	}
	for i, test := range tests {
		s := NewMapSet(test.args...)
		if !reflect.DeepEqual(s.data, test.want) {
			t.Errorf("%d: got %v, want %v", i, s.data, test.want)
		}
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		s    *MapSet[int]
		v    int
		want bool
	}{
		{NewMapSet[int](), 1, false},
		{NewMapSet(1), 1, true},
		{NewMapSet(2), 1, false},
		{NewMapSet(2, 3), 1, false},
		{NewMapSet(2, 3), 2, true},
	}
	for i, test := range tests {
		if got := test.s.Contains(test.v); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestAdd(t *testing.T) {
	s := NewMapSet(1)
	if s.Add(1) {
		t.Error("got true, want false")
	}
	if !s.Add(2) {
		t.Error("got false, want true")
	}
	want := map[int]struct{}{1: {}, 2: {}}
	if !reflect.DeepEqual(s.data, want) {
		t.Errorf("got %v, want %v", s.data, want)
	}
}

func TestRemove(t *testing.T) {
	s := NewMapSet(1)
	if s.Remove(2) {
		t.Error("got true, want false")
	}
	if !s.Remove(1) {
		t.Error("got false, want true")
	}
	if s.Remove(1) {
		t.Error("got true, want false")
	}
	if !reflect.DeepEqual(s.data, map[int]struct{}{}) {
		t.Errorf("got %v want empty map", s.data)
	}
}

func TestIsempty(t *testing.T) {
	s := NewMapSet[int]()
	if !s.IsEmpty() {
		t.Error("got false, want true")
	}
	s.Add(1)
	if s.IsEmpty() {
		t.Error("got true, want false")
	}
}

func TestCardinality(t *testing.T) {
	s := NewMapSet[int]()
	if n := s.Cardinality(); n != 0 {
		t.Errorf("got %d, want 0", n)
	}
	s.Add(1)
	if n := s.Cardinality(); n != 1 {
		t.Errorf("got %d, want 1", n)
	}
}

func TestUnion(t *testing.T) {
	var tests = []struct {
		s1, s2 *MapSet[int]
		want   map[int]struct{}
	}{
		{NewMapSet[int](), NewMapSet[int](), map[int]struct{}{}},
		{NewMapSet[int](), NewMapSet(1, 2), map[int]struct{}{1: {}, 2: {}}},
		{NewMapSet(1, 2), NewMapSet(2, 3), map[int]struct{}{1: {}, 2: {}, 3: {}}},
		{NewMapSet(1, 2), NewMapSet(3, 4), map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}},
		{NewMapSet(1, 2), NewMapSet(1, 2), map[int]struct{}{1: {}, 2: {}}},
	}
	for i, test := range tests {
		got1 := test.s1.Union(test.s2).(*MapSet[int]).data
		got2 := test.s2.Union(test.s1).(*MapSet[int]).data
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	var tests = []struct {
		s1, s2 *MapSet[int]
		want   map[int]struct{}
	}{
		{NewMapSet[int](), NewMapSet[int](), map[int]struct{}{}},
		{NewMapSet[int](), NewMapSet(1, 2), map[int]struct{}{}},
		{NewMapSet(1, 2), NewMapSet(2, 3), map[int]struct{}{2: {}}},
		{NewMapSet(1, 2), NewMapSet(3, 4), map[int]struct{}{}},
		{NewMapSet(1, 2), NewMapSet(1, 2), map[int]struct{}{1: {}, 2: {}}},
	}
	for i, test := range tests {
		got1 := test.s1.Intersection(test.s2).(*MapSet[int]).data
		got2 := test.s2.Intersection(test.s1).(*MapSet[int]).data
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestDifference(t *testing.T) {
	var tests = []struct {
		s1, s2 *MapSet[int]
		want   map[int]struct{}
	}{
		{NewMapSet[int](), NewMapSet[int](), map[int]struct{}{}},
		{NewMapSet[int](), NewMapSet(1, 2), map[int]struct{}{}},
		{NewMapSet(1, 2), NewMapSet(2, 3), map[int]struct{}{1: {}}},
		{NewMapSet(1, 2), NewMapSet(3, 4), map[int]struct{}{1: {}, 2: {}}},
		{NewMapSet(1, 2), NewMapSet(1, 2), map[int]struct{}{}},
	}
	for i, test := range tests {
		if got := test.s1.Difference(test.s2).(*MapSet[int]); !reflect.DeepEqual(got.data, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestSymDifference(t *testing.T) {
	var tests = []struct {
		s1, s2 *MapSet[int]
		want   map[int]struct{}
	}{
		{NewMapSet[int](), NewMapSet[int](), map[int]struct{}{}},
		{NewMapSet[int](), NewMapSet(1, 2), map[int]struct{}{1: {}, 2: {}}},
		{NewMapSet(1, 2), NewMapSet(2, 3), map[int]struct{}{1: {}, 3: {}}},
		{NewMapSet(1, 2), NewMapSet(3, 4), map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}},
		{NewMapSet(1, 2), NewMapSet(1, 2), map[int]struct{}{}},
	}
	for i, test := range tests {
		got1 := test.s1.SymDifference(test.s2).(*MapSet[int]).data
		got2 := test.s2.SymDifference(test.s1).(*MapSet[int]).data
		if !(reflect.DeepEqual(got1, test.want) && reflect.DeepEqual(got2, test.want)) {
			t.Errorf("%d: got %#v and %#v, want %#v", i, got1, got2, test.want)
		}
	}
}

func TestIsSubset(t *testing.T) {
	var tests = []struct {
		s1, s2 *MapSet[int]
		want   bool
	}{
		{NewMapSet[int](), NewMapSet[int](), true},
		{NewMapSet[int](), NewMapSet(1), true},
		{NewMapSet(1), NewMapSet(1, 2), true},
		{NewMapSet(1, 2), NewMapSet(1, 2), true},
		{NewMapSet(1, 2), NewMapSet(1), false},
		{NewMapSet(1), NewMapSet[int](), false},
		{NewMapSet(1, 2), NewMapSet(1, 3, 4), false},
		{NewMapSet(1, 2), NewMapSet(1, 2, 4), true},
	}
	for i, test := range tests {
		if got := test.s1.IsSubset(test.s2); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestIsProperSubset(t *testing.T) {
	var tests = []struct {
		s1, s2 *MapSet[int]
		want   bool
	}{
		{NewMapSet[int](), NewMapSet[int](), false},
		{NewMapSet[int](), NewMapSet(1), true},
		{NewMapSet(1), NewMapSet(1, 2), true},
		{NewMapSet(1, 2), NewMapSet(1, 2), false},
		{NewMapSet(1, 2), NewMapSet(1), false},
		{NewMapSet(1), NewMapSet[int](), false},
		{NewMapSet(1, 2), NewMapSet(1, 3, 4), false},
		{NewMapSet(1, 2), NewMapSet(1, 2, 4), true},
	}
	for i, test := range tests {
		if got := test.s1.IsProperSubset(test.s2); got != test.want {
			t.Errorf("%d: got %t, want %t", i, got, test.want)
		}
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		s1, s2 *MapSet[int]
		want   bool
	}{
		{NewMapSet[int](), NewMapSet[int](), true},
		{NewMapSet[int](), NewMapSet(1), false},
		{NewMapSet(1), NewMapSet(1, 2), false},
		{NewMapSet(1, 2), NewMapSet(1, 2), true},
	}
	for i, test := range tests {
		got1 := test.s1.Equal(test.s2)
		got2 := test.s2.Equal(test.s1)
		if got1 != test.want || got2 != test.want {
			t.Errorf("%d: got %t and %t, want %t", i, got1, got2, test.want)
		}
	}
}

func TestClone(t *testing.T) {
	s := NewMapSet(1, 2, 3)
	clone := s.Clone().(*MapSet[int])
	if !reflect.DeepEqual(s.data, clone.data) {
		t.Errorf("got %v, want %v", clone.data, s.data)
	}
}

func TestElements(t *testing.T) {
	var tests = []struct {
		args []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1, 1}, []int{1, 2}},
	}
	for i, test := range tests {
		s := NewMapSet(test.args...)
		sl := s.Elements()
		sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
		if !reflect.DeepEqual(sl, test.want) {
			t.Errorf("%d: got %v, want %v", i, s.Elements(), test.want)
		}
	}
}

func TestString(t *testing.T) {
	s := NewMapSet(1, 2, 3)
	want := "Set{1, 2, 3}"
	if s.String() != want {
		t.Errorf("got %q, want %q", s.String(), want)
	}
}
