// Package set implements a [Set] type for Go.
package set

import (
	"fmt"
	"strings"
)

// Set type.
type Set[T comparable] struct {
	data map[T]struct{}
}

// New returns a new Set with the given elements.
func New[T comparable](elems ...T) Set[T] {
	data := make(map[T]struct{})
	for _, elem := range elems {
		data[elem] = struct{}{}
	}
	return Set[T]{data: data}
}

// Contains returns true if elem is in the Set.
func (s Set[T]) Contains(elem T) bool {
	_, ok := s.data[elem]
	return ok
}

// Add adds elem to the Set.
// Returns true if it was added, false if it already was in the set.
func (s Set[T]) Add(elem T) bool {
	if _, ok := s.data[elem]; !ok {
		s.data[elem] = struct{}{}
		return true
	}
	return false
}

// Remove removes elem from the Set.
// Returns true if it was in the set, false otherwise.
func (s Set[T]) Remove(elem T) bool {
	if _, ok := s.data[elem]; ok {
		delete(s.data, elem)
		return true
	}
	return false
}

// IsEmpty returns true if Set is an empty set.
func (s Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Cardinality returns the number of elements in the Set.
func (s Set[T]) Cardinality() int {
	return len(s.data)
}

// Union returns a new Set which is the union of s and s2.
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		m[elem] = struct{}{}
	}
	for elem := range s2.data {
		m[elem] = struct{}{}
	}
	return Set[T]{data: m}
}

// Intersection returns a new Set which is the intersection of s and s2.
func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		if _, ok := s2.data[elem]; ok {
			m[elem] = struct{}{}
		}
	}
	return Set[T]{data: m}
}

// Difference returns a new Set which is the set difference of s and s2.
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		if _, ok := s2.data[elem]; !ok {
			m[elem] = struct{}{}
		}
	}
	return Set[T]{data: m}
}

// SymDifference returns a new Set which is the symmetric difference of s and s2.
func (s Set[T]) SymDifference(s2 Set[T]) Set[T] {
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
	return Set[T]{data: m}
}

// IsSubset returns true if s is a subset of s2.
func (s Set[T]) IsSubset(s2 Set[T]) bool {
	for elem := range s.data {
		if _, ok := s2.data[elem]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubset returns true if s is a proper subset of s2.
func (s Set[T]) IsProperSubset(s2 Set[T]) bool {
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
func (s Set[T]) Equal(s2 Set[T]) bool {
	if len(s.data) != len(s2.data) {
		return false
	}
	return s.IsSubset(s2)
}

// Clone clones the Set.
func (s Set[T]) Clone() Set[T] {
	m := make(map[T]struct{})
	for elem := range s.data {
		m[elem] = struct{}{}
	}
	return Set[T]{data: m}
}

// Elements returns a slice with all elements of the Set.
func (s Set[T]) Elements() []T {
	result := make([]T, 0, len(s.data))
	for elem := range s.data {
		result = append(result, elem)
	}
	return result
}

// String returns a string representation of the Set.
func (s Set[T]) String() string {
	sl := make([]string, len(s.data))
	for i, elem := range s.Elements() {
		sl[i] = fmt.Sprintf("%v", elem)
	}
	return fmt.Sprintf("Set{%s}", strings.Join(sl, ", "))
}
