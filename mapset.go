package set

import (
	"fmt"
	"strings"
)

// MapSet type.
type MapSet[T comparable] struct {
	data map[T]struct{}
}

// New returns a new MapSet with the given elements.
func NewMapSet[T comparable](elems ...T) *MapSet[T] {
	data := make(map[T]struct{})
	for _, v := range elems {
		data[v] = struct{}{}
	}
	return &MapSet[T]{data: data}
}

// Contains returns true if the element is in the Set.
func (s *MapSet[T]) Contains(elem T) bool {
	_, ok := s.data[elem]
	return ok
}

// Add adds an element to the Set.
// Returns true if it was added, false if it already was in the set.
func (s *MapSet[T]) Add(elem T) bool {
	if _, ok := s.data[elem]; !ok {
		s.data[elem] = struct{}{}
		return true
	}
	return false
}

// Remove removes an element from the Set.
// Returns true if it was in the set, false otherwise.
func (s *MapSet[T]) Remove(elem T) bool {
	if _, ok := s.data[elem]; ok {
		delete(s.data, elem)
		return true
	}
	return false
}

// IsEmpty returns true if Set is an empty set.
func (s *MapSet[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Cardinality returns the number of elements in the Set.
func (s *MapSet[T]) Cardinality() int {
	return len(s.data)
}

// Union returns a new Set which is the union of s and s2.
func (s *MapSet[T]) Union(s2 *MapSet[T]) *MapSet[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		m[elem] = struct{}{}
	}
	for elem := range s2.data {
		m[elem] = struct{}{}
	}
	return &MapSet[T]{data: m}
}

// Intersection returns a new Set which is the intersection of s and s2.
func (s *MapSet[T]) Intersection(s2 *MapSet[T]) *MapSet[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		if _, ok := s2.data[elem]; ok {
			m[elem] = struct{}{}
		}
	}
	return &MapSet[T]{data: m}
}

// Difference returns a new Set which is the set difference of s and s2.
func (s *MapSet[T]) Difference(s2 *MapSet[T]) *MapSet[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		if _, ok := s2.data[elem]; !ok {
			m[elem] = struct{}{}
		}
	}
	return &MapSet[T]{data: m}
}

// SymDifference returns a new Set which is the symmetric difference of s and s2.
func (s *MapSet[T]) SymDifference(s2 *MapSet[T]) *MapSet[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		if _, ok := s2.data[elem]; !ok {
			m[elem] = struct{}{}
		}
	}
	for elem := range s2.data {
		if _, ok := s.data[elem]; !ok {
			m[elem] = struct{}{}
		}
	}
	return &MapSet[T]{data: m}
}

// IsSubset returns true if s is a subset of s2.
func (s *MapSet[T]) IsSubset(s2 *MapSet[T]) bool {
	for elem := range s.data {
		if _, ok := s2.data[elem]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubset returns true if s is a proper subset of s2.
func (s *MapSet[T]) IsProperSubset(s2 *MapSet[T]) bool {
	if len(s.data) >= len(s2.data) {
		return false
	}
	for elem := range s.data {
		if _, ok := s2.data[elem]; !ok {
			return false
		}
	}
	return true
}

// Equal returns true if s and s2 contain the same elements.
func (s *MapSet[T]) Equal(s2 *MapSet[T]) bool {
	if len(s.data) != len(s2.data) {
		return false
	}
	return s.IsSubset(s2)
}

// Clone clones the Set.
func (s *MapSet[T]) Clone() *MapSet[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		m[elem] = struct{}{}
	}
	return &MapSet[T]{data: m}
}

// Elements returns a slice with all elements of the Set.
func (s *MapSet[T]) Elements() []T {
	result := make([]T, 0, len(s.data))
	for elem := range s.data {
		result = append(result, elem)
	}
	return result
}

// String returns a string representation of the Set.
func (s *MapSet[T]) String() string {
	sl := make([]string, len(s.data))
	for i, elem := range s.Elements() {
		sl[i] = fmt.Sprintf("%v", elem)
	}
	return fmt.Sprintf("Set{%s}", strings.Join(sl, ", "))
}
