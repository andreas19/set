// Package set implements [Set] types for Go.
package set

type Set[T any] interface {
	// Contains reports whether the element is in the Set.
	Contains(elem T) bool

	// Add adds an element to the Set.
	// Returns true if it was added, false if it already was in the set.
	Add(elem T) bool

	// Update updates the Set with elems.
	Update(elems ...T)

	// Remove removes an element from the Set.
	// Returns true if it was in the set, false otherwise.
	Remove(elem T) bool

	// IsEmpty returns true if Set is an empty set.
	IsEmpty() bool

	// Cardinality returns the number of elements in the Set.
	Cardinality() int

	// Union returns a new Set which is the union of the Set and s2.
	Union(s2 Set[T]) Set[T]

	// Intersection returns a new Set which is the intersection of the Set and s2.
	Intersection(s2 Set[T]) Set[T]

	// Difference returns a new Set which is the set difference of the Set and s2.
	Difference(s2 Set[T]) Set[T]

	// SymDifference returns a new Set which is the symmetric difference of the Set and s2.
	SymDifference(s2 Set[T]) Set[T]

	// IsSubset returns true if the Set is a subset of s2.
	IsSubset(s2 Set[T]) bool

	// IsProperSubset returns true if the Set is a proper subset of s2.
	IsProperSubset(s2 Set[T]) bool

	// Equal returns true if the Set and s2 contain the same elements.
	Equal(s2 Set[T]) bool

	// Clone clones the Set.
	Clone() Set[T]

	// Elements returns a slice with all elements of the Set.
	Elements() []T

	// String returns a string representation of the Set.
	String() string
}
