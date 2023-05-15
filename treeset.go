package set

import (
	"fmt"
	"strings"

	"github.com/andreas19/avltree"
	"golang.org/x/exp/constraints"
)

// TreeSet type that implements the [Set] interface.
// It uses an AVL tree to store the elements.
type TreeSet[T any] struct {
	tree *avltree.Tree[T]
	cmp  avltree.Cmp[T]
}

// NewTreeSet returns a new TreeSet with the given elements.
func NewTreeSet[T constraints.Ordered](elems ...T) *TreeSet[T] {
	return (*TreeSet[T])(NewTreeSetFunc(avltree.CmpOrd[T], elems...))
}

// NewTreeSetFunc returns a new TreeSet with the given elements. Function cmp is used
// to compare two elements. It returns 0 if a == b, -1 if a < b, and 1 if a > b.
func NewTreeSetFunc[T any](cmp func(T, T) int, elems ...T) *TreeSet[T] {
	tree := avltree.New(cmp, true)
	for _, elem := range elems {
		tree.Add(elem)
	}
	return &TreeSet[T]{tree: tree, cmp: cmp}
}

// Contains reports whether the element is in the Set.
func (s *TreeSet[T]) Contains(elem T) bool {
	return s.tree.Contains(elem)
}

// Add adds an element to the Set.
// Returns true if it was added, false if it already was in the set.
func (s *TreeSet[T]) Add(elem T) bool {
	return s.tree.Add(elem)
}

// Remove removes an element from the Set.
// Returns true if it was in the set, false otherwise.
func (s *TreeSet[T]) Remove(elem T) bool {
	return s.tree.Del(elem)
}

// IsEmpty returns true if Set is an empty set.
func (s *TreeSet[T]) IsEmpty() bool {
	return s.tree.IsEmpty()
}

// Cardinality returns the number of elements in the Set.
func (s *TreeSet[T]) Cardinality() int {
	return s.tree.Count()
}

// Union returns a new Set which is the union of s and s2.
func (s *TreeSet[T]) Union(s2 Set[T]) Set[T] {
	tree := avltree.New(s.cmp, true)
	s.tree.Each(func(elem T) {
		tree.Add(elem)
	})
	if x, ok := s2.(*TreeSet[T]); ok {
		x.tree.Each(func(elem T) {
			tree.Add(elem)
		})
	} else {
		for _, elem := range s2.Elements() {
			tree.Add(elem)
		}
	}
	return &TreeSet[T]{tree: tree, cmp: s.cmp}
}

// Intersection returns a new Set which is the intersection of s and s2.
func (s *TreeSet[T]) Intersection(s2 Set[T]) Set[T] {
	tree := avltree.New(s.cmp, true)
	s.tree.Each(func(elem T) {
		if s2.Contains(elem) {
			tree.Add(elem)
		}
	})
	return &TreeSet[T]{tree: tree, cmp: s.cmp}
}

// Difference returns a new Set which is the set difference of s and s2.
func (s *TreeSet[T]) Difference(s2 Set[T]) Set[T] {
	tree := avltree.New(s.cmp, true)
	s.tree.Each(func(elem T) {
		if !s2.Contains(elem) {
			tree.Add(elem)
		}
	})
	return &TreeSet[T]{tree: tree, cmp: s.cmp}
}

// SymDifference returns a new Set which is the symmetric difference of s and s2.
func (s *TreeSet[T]) SymDifference(s2 Set[T]) Set[T] {
	tree := avltree.New(s.cmp, true)
	s.tree.Each(func(elem T) {
		if !s2.Contains(elem) {
			tree.Add(elem)
		}
	})
	if x, ok := s2.(*TreeSet[T]); ok {
		x.tree.Each(func(elem T) {
			if !s.Contains(elem) {
				tree.Add(elem)
			}
		})
	} else {
		for _, elem := range s2.Elements() {
			if !s.Contains(elem) {
				tree.Add(elem)
			}
		}
	}
	return &TreeSet[T]{tree: tree, cmp: s.cmp}
}

// IsSubset returns true if s is a subset of s2.
func (s *TreeSet[T]) IsSubset(s2 Set[T]) bool {
	if s.Cardinality() > s2.Cardinality() {
		return false
	}
	b := true
	s.tree.Each(func(elem T) {
		b = b && s2.Contains(elem)
	})
	return b
}

// IsProperSubset returns true if s is a proper subset of s2.
func (s *TreeSet[T]) IsProperSubset(s2 Set[T]) bool {
	if s.Cardinality() >= s2.Cardinality() {
		return false
	}
	return s.IsSubset(s2)
}

// Equal returns true if s and s2 contain the same elements.
func (s *TreeSet[T]) Equal(s2 Set[T]) bool {
	if s.Cardinality() != s2.Cardinality() {
		return false
	}
	return s.IsSubset(s2)
}

// Clone clones the Set.
func (s *TreeSet[T]) Clone() Set[T] {
	tree := avltree.New(s.cmp, true)
	s.tree.Each(func(elem T) {
		tree.Add(elem)
	})
	return &TreeSet[T]{tree: tree, cmp: s.cmp}
}

// Elements returns a slice with all elements of the Set.
func (s *TreeSet[T]) Elements() []T {
	return s.tree.Slice()
}

// String returns a string representation of the Set.
func (s *TreeSet[T]) String() string {
	sl := make([]string, 0, s.tree.Count())
	s.tree.Each(func(elem T) {
		sl = append(sl, fmt.Sprintf("%v", elem))
	})
	return fmt.Sprintf("TreeSet{%s}", strings.Join(sl, ", "))
}
