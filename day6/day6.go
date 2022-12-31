package day6

import (
	"aoc2022/utils"
	"fmt"
)

var input string

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 6
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	input = val
	return nil
}

func findUnique(v string, count int) int {
	for i := count; i <= len(input); i++ {
		start := i - count
		vals := utils.HashSetFromString(v[start:i])
		if len(vals.Keys()) == count {
			return i
		}
	}
	return -1
}

func part1() (string, error) {
	return fmt.Sprintf("%d", findUnique(input, 4)), nil
}

func part2() (string, error) {
	return fmt.Sprintf("%d", findUnique(input, 14)), nil
}
