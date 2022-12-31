package utils

type HashSet[T comparable] struct {
	data map[T]bool
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{data: make(map[T]bool)}
}

func HashSetFromSlice[T comparable](slice []T) *HashSet[T] {
	result := NewHashSet[T]()
	for _, v := range slice {
		result.Add(v)
	}
	return result
}

func HashSetFromString(str string) *HashSet[rune] {
	result := NewHashSet[rune]()
	for _, v := range str {
		result.Add(v)
	}
	return result
}

func (h *HashSet[T]) Add(val T) {
	h.data[val] = true
}

func (h *HashSet[T]) Remove(val T) bool {
	if h.Contains(val) {
		delete(h.data, val)
		return true
	}

	return false
}

func (h *HashSet[T]) Contains(val T) bool {
	_, found := h.data[val]
	return found
}

func (h *HashSet[T]) Intersection(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for k := range h.data {
		if other.Contains(k) {
			result.Add(k)
		}
	}

	return result
}

func (h *HashSet[T]) Union(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for k := range h.data {
		result.Add(k)
	}

	for k := range other.data {
		result.Add(k)
	}

	return result
}

func (h *HashSet[T]) Keys() []T {
	keys := make([]T, 0, len(h.data))
	for k := range h.data {
		keys = append(keys, k)
	}
	return keys
}

func (h *HashSet[T]) Length() int {
	return len(h.data)
}

func (h *HashSet[T]) Clone() *HashSet[T] {
	return HashSetFromSlice(h.Keys())
}

func (h *HashSet[T]) Equals(other *HashSet[T]) bool {
	if h.Length() != other.Length() {
		return false
	}

	for k := range h.data {
		if !other.Contains(k) {
			return false
		}
	}

	return true
}

func (h *HashSet[T]) IncludedBy(other *HashSet[T]) bool {
	for k := range h.data {
		if !other.Contains(k) {
			return false
		}
	}

	return true
}
