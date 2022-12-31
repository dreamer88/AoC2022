package day3

import (
	"aoc2022/utils"
	"errors"
	"fmt"
)

var input []string

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 3
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	input = append(input, val)
	return nil
}

func getItemScore(c rune) (int, error) {
	c_i := int(c)
	if c_i >= 'A' && c_i <= 'Z' {
		return c_i - 'A' + 27, nil
	} else if c_i >= 'a' && c_i <= 'z' {
		return c_i - 'a' + 1, nil
	}

	return 0, errors.New(fmt.Sprintf("Error, unhandled char '%c'", c))
}

func getBagScore(bag string) (int, error) {
	length := len(bag)
	p1 := utils.HashSetFromString(bag[0 : length/2])
	p2 := utils.HashSetFromString(bag[length/2:])
	intersect := p1.Intersection(p2)

	return getItemScore(intersect.Keys()[0])
}

func part1() (string, error) {
	scores := 0
	for _, bag := range input {
		s, err := getBagScore(bag)
		if err != nil {
			return "", err
		} else {
			scores += s
		}
	}
	return fmt.Sprintf("%d", scores), nil
}

func getBadgeScore(bags []string) (int, error) {
	options := utils.HashSetFromString(bags[0])
	for _, bag := range bags[1:] {
		options = options.Intersection(utils.HashSetFromString(bag))
	}

	return getItemScore(options.Keys()[0])
}

func part2() (string, error) {
	scores := 0
	for _, window := range utils.Windows(input, 3) {
		s, err := getBadgeScore(window)
		if err != nil {
			return "", err
		} else {
			scores += s
		}
	}
	return fmt.Sprintf("%d", scores), nil
}
