package slice

import "errors"

// Cons creates a new list with head at the front of tail
func Cons[T any](head T, tail []T) []T {
	lst := []T{head}
	return append(lst, tail...)
}

// Head returns slc[0] if the list has at least one element
// in it. Otherwise, returns empty() and a descriptive, non-nil
// error
func Head[T any](slc []T, empty func() T) (T, error) {
	if len(slc) == 0 {
		return empty(), errors.New("Head called on empty list")
	}
	return slc[0], nil
}

// Tail returns slc[1:] only if the list has 2 or more elements in
// it. If it doesn't, returns an empty list and a descriptive, non-nil
// error
func Tail[T any](slc []T) ([]T, error) {
	if len(slc) < 2 {
		return nil, errors.New("Tail called on empty list")
	}
	return slc[1:], nil
}

func minmaxSlice[T any](a, b []T) (smaller, larger []T) {
	if len(a) < len(b) {
		return a, b
	}
	return b, a
}

// Zip combines slc1 and slc2 together. Starting at index 0, every odd number
// index in the returned slice will be the next successive
// element in slc1, and even number will be the next in slc2. If either slice
// is longer than the other, the remainder of the returned slice will just
// have the rest of the elements in the longer slice
func Zip[T any](slc1 []T, slc2 []T) []T {
	ret := make([]T, 0, length)

	smaller, larger := minmaxSlice(slc1, slc2)

	for i := 0; i < len(smaller); i++ {
		ret = append(ret, slc1[i])
		ret = append(ret, slc2[i])
	}

	length := len(slc1) + len(slc2)
	return append(ret, larger[len(larger)-(length-len(ret)):]...)
}
