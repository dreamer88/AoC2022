package day1

import (
	"aoc2022/utils"
	"fmt"
	"sort"
	"strconv"
)

var input = []int{0}

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 1
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	i, err := strconv.Atoi(val)
	if err != nil {
		input = append(input, 0)
	} else {
		input[len(input)-1] += i
	}
	return nil
}

func part1() (string, error) {
	max := utils.Reduce(input[1:], input[0], func(v int, r int) int {
		if r > v {
			return r
		} else {
			return v
		}
	})

	return fmt.Sprintf("%d", max), nil
}

func part2() (string, error) {
	sort.Slice(input, func(i int, j int) bool { return input[i] > input[j] })

	max := utils.Reduce(input[0:3], 0, func(v int, r int) int { return v + r })
	return fmt.Sprintf("%d", max), nil
}
