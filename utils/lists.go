package utils

import (
	"fmt"
	"math"
	"sort"
)

func IndexInList[T comparable](vals []T, val T) int {
	for i, v := range vals {
		if v == val {
			return i
		}
	}

	return -1
}

func ExistsInList[T comparable](vals []T, val T) bool {
	return IndexInList(vals, val) != -1
}

func FindIndex[T any](vals []T, find_fn func(*T) bool) int {
	for i := 0; i < len(vals); i++ {
		if find_fn(&vals[i]) {
			return i
		}
	}

	return -1
}

func Windows[T any](vals []T, window_size int) [][]T {
	var result [][]T
	size := len(vals)
	num_windows := int(math.Ceil(float64(size) / float64(window_size)))
	for i := 0; i < num_windows; i++ {
		start := i * window_size
		end := Min((i+1)*window_size, size)
		result = append(result, vals[start:end])
	}

	return result
}

func Map[T any, U any](vals []T, fn func(T) U) []U {
	var results = make([]U, len(vals))
	for i, v := range vals {
		results[i] = fn(v)
	}
	return results
}

func Filter[T any](vals []T, fn func(int, *T) bool) []T {
	var results = []T{}

	for i := 0; i < len(vals); i++ {
		if fn(i, &vals[i]) {
			results = append(results, vals[i])
		}
	}

	return results
}

func Any[T any](vals []T, fn func(int, *T) bool) bool {
	for i := 0; i < len(vals); i++ {
		if fn(i, &vals[i]) {
			return true
		}
	}

	return false
}

func All[T any](vals []T, fn func(int, *T) bool) bool {
	for i := 0; i < len(vals); i++ {
		if !fn(i, &vals[i]) {
			return false
		}
	}

	return true
}

func Reduce[T any, U any](vals []T, init U, fn func(T, U) U) U {
	var result = init

	for _, v := range vals {
		result = fn(v, result)
	}

	return result
}

func Reversed[T any](vals []T) []T {
	l := len(vals)
	var results = make([]T, l)

	for i := 0; i < l; i++ {
		results[i] = vals[l-i-1]
	}

	return results
}

func Insert[T any](vals []T, val T, idx int) ([]T, error) {
	l := len(vals)
	if idx < 0 || idx > l {
		return nil, fmt.Errorf("invalid index %d, length %d", idx, l)
	} else if idx == l {
		vals = append(vals, val)
		return vals, nil
	}

	vals = append(vals, vals[l-1])
	for i := l - 1; i > idx; i-- {
		vals[i] = vals[i-1]
	}

	vals[idx] = val

	return vals, nil
}

func Remove[T any](vals []T, idx int) ([]T, error) {
	l := len(vals)
	if idx < 0 || idx >= l {
		return nil, fmt.Errorf("invalid index %d, length %d", idx, l)
	}

	res := make([]T, l-1)
	copy(res, vals[:idx])
	copy(res[idx:], vals[idx+1:])

	return res, nil
}

func Move[T any](vals []T, start_index int, end_index int) ([]T, error) {
	l := len(vals)
	if start_index < 0 || start_index >= l {
		return nil, fmt.Errorf("invalid start index %d, length %d", start_index, l)
	}
	if end_index < 0 || end_index >= l {
		return nil, fmt.Errorf("invalid end index %d, length %d", end_index, l)
	}

	if start_index == end_index {
		return Copy1D(vals), nil
	}

	res := make([]T, l)
	v := vals[start_index]
	min_index, max_index := MinMax(start_index, end_index)

	copy(res, vals[:min_index])
	if start_index < end_index {
		copy(res[start_index:end_index], vals[start_index+1:end_index+1])
	} else {
		copy(res[end_index+1:start_index+1], vals[end_index:start_index])
	}

	copy(res[max_index+1:], vals[max_index+1:])

	res[end_index] = v

	return res, nil
}

func RotateLeft[T any](vals []T, index int, amount int) ([]T, error) {
	end_index := SaneMod(index-amount, len(vals)-1)
	return Move(vals, index, end_index)
}

func RotateRight[T any](vals []T, index int, amount int) ([]T, error) {
	end_index := SaneMod(index+amount, len(vals)-1)
	return Move(vals, index, end_index)
}

func Copy1D[T any](vals []T) []T {
	var results = make([]T, len(vals))
	copy(results, vals)
	return results
}

func Copy2D[T any](vals [][]T) [][]T {
	var results = make([][]T, len(vals))
	for i, v := range vals {
		results[i] = Copy1D(v)
	}
	return results
}

func Copy3D[T any](vals [][][]T) [][][]T {
	var results = make([][][]T, len(vals))
	for i, v := range vals {
		results[i] = Copy2D(v)
	}
	return results
}

func Sort[T any](vals []T, sort_fn func(a *T, b *T) bool) {
	sort.Slice(vals, func(i int, j int) bool { return sort_fn(&vals[i], &vals[j]) })
}

func PrintArray[T any](vals []T, delim string) {
	for i := 0; i < len(vals); i++ {
		if i == 0 {
			print(vals[i])
		} else {
			print(delim, vals[i])
		}
	}
	println("")
}
