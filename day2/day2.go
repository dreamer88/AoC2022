package day2

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

var opp_map = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
}

var base_map = map[string]int{
	"X": 2,
	"Y": 1,
	"Z": 0,
}

var selection = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var scores = map[string]map[string]int{
	"A": {
		"X": 3,
		"Y": 6,
		"Z": 0,
	},
	"B": {
		"X": 0,
		"Y": 3,
		"Z": 6,
	},
	"C": {
		"X": 6,
		"Y": 0,
		"Z": 3,
	},
}

var day2 = map[string]map[string]string{
	"A": {
		"X": "Z",
		"Y": "X",
		"Z": "Y",
	},
	"B": {
		"X": "X",
		"Y": "Y",
		"Z": "Z",
	},
	"C": {
		"X": "Y",
		"Y": "Z",
		"Z": "X",
	},
}

var input [][]string

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 2
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	input = append(input, strings.Split(val, " "))
	return nil
}

func get_score(other string, yours string) int {
	return selection[yours] + scores[other][yours]

}

func part1() (string, error) {
	result := 0

	for _, res := range input {
		result += get_score(res[0], res[1])
	}

	return fmt.Sprintf("%d", result), nil
}

func part2() (string, error) {
	result := 0

	for _, res := range input {
		result += get_score(res[0], day2[res[0]][res[1]])
	}

	return fmt.Sprintf("%d", result), nil
}
