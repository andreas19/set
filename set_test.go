package set

import (
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		args []int
		want map[int]struct{}
	}{
		{[]int{}, map[int]struct{}{}},
		{[]int{1}, map[int]struct{}{1: {}}},
		{[]int{1, 2, 1}, map[int]struct{}{1: {}, 2: {}}},
	}
	for _, test := range tests {
		s := New(test.args...)
		if !mapEqual(s.data, test.want) {
			t.Errorf("got %v, want %v", s.data, test.want)
		}
	}
}

func TestContains(t *testing.T) {
	s := New(1)
	if !s.Contains(1) {
		t.Error("got false, want true")
	}
	if s.Contains(2) {
		t.Error("got true, want false")
	}
}

func TestAdd(t *testing.T) {
	s := New(1)
	if s.Add(1) {
		t.Error("got true, want false")
	}
	if !s.Add(2) {
		t.Error("got false, want true")
	}
	want := map[int]struct{}{1: {}, 2: {}}
	if !mapEqual(s.data, want) {
		t.Errorf("got %v, want %v", s.data, want)
	}
}

func TestRemove(t *testing.T) {
	s := New(1)
	if s.Remove(2) {
		t.Error("got true, want false")
	}
	if !s.Remove(1) {
		t.Error("got false, want true")
	}
	if s.Remove(1) {
		t.Error("got true, want false")
	}
	if !mapEqual(s.data, map[int]struct{}{}) {
		t.Errorf("got %v want empty map", s.data)
	}
}

func TestIsempty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Error("got false, want true")
	}
	s.Add(1)
	if s.IsEmpty() {
		t.Error("got true, want false")
	}
}

func TestCardinality(t *testing.T) {
	s := New[int]()
	if n := s.Cardinality(); n != 0 {
		t.Errorf("got %d, want 0", n)
	}
	s.Add(1)
	if n := s.Cardinality(); n != 1 {
		t.Errorf("got %d, want 1", n)
	}
}

func TestUnion(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	want := map[int]struct{}{1: {}, 2: {}, 3: {}}
	union := s1.Union(s2)
	if !mapEqual(union.data, want) {
		t.Errorf("got %v, want %v", union.data, want)
	}
}

func TestIntersection(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	want := map[int]struct{}{2: {}}
	intersection := s1.Intersection(s2)
	if !mapEqual(intersection.data, want) {
		t.Errorf("got %v, want %v", intersection.data, want)
	}
}

func TestDifference(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	want := map[int]struct{}{1: {}}
	diff := s1.Difference(s2)
	if !mapEqual(diff.data, want) {
		t.Errorf("got %v, want %v", diff.data, want)
	}
}

func TestIsSubset(t *testing.T) {
	s := New(1, 2, 3)
	s1 := New(1, 2)
	s2 := New(3, 4)
	if !s1.IsSubset(s) {
		t.Error("got false, want true")
	}
	if s2.IsSubset(s) {
		t.Error("got true, want false")
	}
}

func TestEqual(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(1, 2)
	s3 := New(3, 4)
	s4 := New(0)
	if !(s1.Equal(s2) && s2.Equal(s1)) {
		t.Error("got false, want true")
	}
	if !s1.Equal(s2) {
		t.Error("got false, want true")
	}
	if s1.Equal(s3) {
		t.Error("got true, want false")
	}
	if s3.Equal(s4) {
		t.Error("got true, want false")
	}
}

func TestClone(t *testing.T) {
	s := New(1, 2, 3)
	clone := s.Clone()
	if !mapEqual(s.data, clone.data) {
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
	for _, test := range tests {
		s := New(test.args...)
		sl := s.Elements()
		sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
		if !sliceEqual(sl, test.want) {
			t.Errorf("got %v, want %v", s.Elements(), test.want)
		}
	}
}

func TestString(t *testing.T) {
	s := New(1)
	want := "Set{1}"
	if s.String() != want {
		t.Errorf("got %q, want %q", s.String(), want)
	}
}

func mapEqual(m1, m2 map[int]struct{}) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || (v1 != v2) {
			return false
		}
	}
	return true
}

func sliceEqual(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, v := range sl1 {
		if v != sl2[i] {
			return false
		}
	}
	return true
}
