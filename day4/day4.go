package day4

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
)

var input []EPair

var reg = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

type EPair struct {
	first, second *utils.HashSet[int]
}

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 4
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func getElf(start string, end string) *utils.HashSet[int] {
	result := utils.NewHashSet[int]()
	first, _ := strconv.Atoi(start)
	second, _ := strconv.Atoi(end)
	for v := first; v <= second; v++ {
		result.Add(v)
	}
	return result
}

func parse(val string) error {
	match := reg.FindStringSubmatch(val)
	if match == nil {
		return fmt.Errorf("failed to parse line %s", val)
	}

	pair := EPair{getElf(match[1], match[2]), getElf(match[3], match[4])}
	input = append(input, pair)
	return nil
}

func isOverlapping(ep EPair) bool {
	overlap := ep.first.Intersection(ep.second)

	if len(ep.first.Keys()) == len(overlap.Keys()) {
		return true
	} else if len(ep.second.Keys()) == len(overlap.Keys()) {
		return true
	} else {
		return false
	}
}

func part1() (string, error) {
	count := utils.Reduce(input, 0, func(ep EPair, curr int) int {
		if isOverlapping(ep) {
			return curr + 1
		} else {
			return curr
		}
	})
	return fmt.Sprintf("%d", count), nil
}

func anyOverlap(ep EPair) bool {
	return len(ep.first.Intersection(ep.second).Keys()) > 0
}

func part2() (string, error) {
	count := utils.Reduce(input, 0, func(ep EPair, curr int) int {
		if anyOverlap(ep) {
			return curr + 1
		} else {
			return curr
		}
	})
	return fmt.Sprintf("%d", count), nil
}
