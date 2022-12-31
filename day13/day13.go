package day13

import (
	"aoc2022/utils"
	"fmt"
	"sort"
)

type ValueType uint8

const (
	NONE   ValueType = iota
	NUMBER ValueType = iota
	LIST   ValueType = iota
)

type Result uint8

const (
	CORRECT Result = iota
	EQUAL   Result = iota
	INVALID Result = iota
)

type Value struct {
	valueType ValueType
	number    int
	list      []Value
	line      string
}

var input []Value

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 13
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse_line(val string) Value {
	var val_stack [](*Value)
	var first_value *Value
	number := -1
	finalize_number := func() {
		if number != -1 {
			val_stack[len(val_stack)-1].list = append(val_stack[len(val_stack)-1].list, Value{NUMBER, number, nil, ""})
		}
		number = -1
	}
	for i := 0; i < len(val); i++ {
		c := val[i]
		if c == '[' {
			new_val := Value{
				LIST, 0, make([]Value, 0), "",
			}

			if len(val_stack) > 0 {
				last := val_stack[len(val_stack)-1]
				last.list = append(last.list, new_val)
				val_stack = append(val_stack, &last.list[len(last.list)-1])
			} else {
				new_val.line = val
				val_stack = append(val_stack, &new_val)
				first_value = val_stack[0]
			}
		} else if c == ']' {
			finalize_number()
			val_stack = val_stack[:len(val_stack)-1]
		} else if c == ',' {
			finalize_number()
		} else {
			if number == -1 {
				number = int(c)
			} else {
				number = number*10 + int(c)
			}
		}
	}

	return *first_value
}

func parse(val string) error {
	if len(val) == 0 {
		return nil
	}
	input = append(input, parse_line(val))
	return nil
}

func GetNumResult(left int, right int) Result {
	if left < right {
		return CORRECT
	} else if right < left {
		return INVALID
	} else {
		return EQUAL
	}
}

func IsValid(left *Value, right *Value) Result {
	if left.valueType == right.valueType {
		if left.valueType == NUMBER {
			return GetNumResult(left.number, right.number)
		} else {
			min := len(left.list)
			if len(right.list) < min {
				min = len(right.list)
			}

			for i := 0; i < min; i++ {
				if i >= len(right.list) {
					return INVALID
				}

				result := IsValid(&left.list[i], &right.list[i])
				if result != EQUAL {
					return result
				}
			}

			return GetNumResult(len(left.list), len(right.list))
		}
	} else if left.valueType == NUMBER {
		return IsValid(&Value{LIST, 0, []Value{*left}, ""}, right)
	} else {
		return IsValid(left, &Value{LIST, 0, []Value{*right}, ""})
	}
}

func part1() (string, error) {
	validIndices := 0
	for i, s := range utils.Windows(input, 2) {
		result := IsValid(&s[0], &s[1])
		if result == CORRECT {
			validIndices += (i + 1)
		}
	}
	return fmt.Sprint(validIndices), nil
}

func part2() (string, error) {
	v := parse_line("[[2]]")
	v.number = -1
	values := append(input, v)

	v = parse_line("[[6]]")
	v.number = -1
	values = append(values, v)

	sort.Slice(values, func(i int, j int) bool {
		return IsValid(&values[i], &values[j]) == CORRECT
	})

	result := 1
	for i, v := range values {
		if v.number == -1 {
			result *= (i + 1)
		}
	}

	return fmt.Sprint(result), nil
}
