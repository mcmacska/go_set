package set

type Set[T comparable] map[T]bool

// adds an element to the set
func (s Set[T]) Add(value T) {
	s[value] = true
}

// removes an element from the set
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// checks if an element exists in the set
func (s Set[T]) Contains(value T) bool {
	return s[value]
}

// gets size of the set
func (s Set[T]) Size() int {
	return len(s)
}

// merges another set into the current set
func (s Set[T]) Merge(other Set[T]) {
	for key := range other {
		s[key] = true
	}
}

// returns a new set containing elements of original set that are not in the other
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := Set[T]{}
	for key := range s {
		if !other[key] {
			result[key] = true
		}
	}
	return result
}

// returns a new set containing n number of elements
func (s Set[T]) FirstN(n int) Set[T] {
	result := Set[T]{}
	count := 0
	for key := range s {
		if count >= n {
			break
		}
		result[key] = true
		count++
	}
	return result
}

// returns a new set containing elements of both sets
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := Set[T]{}
	for key := range s {
		if other[key] {
			result[key] = true
		}
	}
	return result
}

// converts set to a list
func (s Set[T]) ToList() []T {
	var list []T
	for key := range s {
		list = append(list, key)
	}
	return list
}
