package utils

import (
	"bufio"
	"math"
	"os"

	"golang.org/x/exp/constraints"
)

type ParseFunction func(string) error

func ProcessByLine(filename string, fn ParseFunction) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // f.Close will run when we're finished.

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := fn(scanner.Text()); err != nil {
			return err
		}
	}

	return nil
}

func Abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func Min[T constraints.Ordered](a T, b T) T {
	if b < a {
		return b
	} else {
		return a
	}
}

func Max[T constraints.Ordered](a T, b T) T {
	if b < a {
		return a
	} else {
		return b
	}
}

func MinMax[T constraints.Ordered](a T, b T) (T, T) {
	if b < a {
		return b, a
	} else {
		return a, b
	}
}

func Clamp[T constraints.Ordered](v T, start T, end T) T {
	if v < start {
		return start
	} else if v > end {
		return end
	} else {
		return v
	}
}

func Ceil[T constraints.Integer](num T, denom T) T {
	return T(math.Ceil(float64(num) / float64(denom)))
}

func SaneMod[T constraints.Integer](x T, l T) T {
	v := x % l
	if v < 0 {
		return v + l
	} else {
		return v
	}
}
