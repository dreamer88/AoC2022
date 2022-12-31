package day12

import (
	"aoc2022/utils"
	"fmt"
)

type Map struct {
	start_x, start_y, end_x, end_y int
	heights                        [][]int
}

type MapIter struct {
	x, y, steps int
}

var input = Map{
	0, 0, 0, 0,
	make([][]int, 0),
}

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 12
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	var row []int

	for i := 0; i < len(val); i++ {
		c := val[i]
		if c == 'S' {
			input.start_x = len(row)
			input.start_y = len(input.heights)
			row = append(row, int('a'))
		} else if c == 'E' {
			input.end_x = len(row)
			input.end_y = len(input.heights)
			row = append(row, int('z'))
		} else {
			row = append(row, int(c))
		}
	}

	input.heights = append(input.heights, row)
	return nil
}

func part1() (string, error) {
	seen := utils.NewHashSet[string]()

	rows := len(input.heights)
	cols := len(input.heights[0])

	curr_positions := []MapIter{{input.start_x, input.start_y, 0}}

	try_add_pos := func(curr_x int, curr_y int, end_x int, end_y int, steps int) {
		if end_x >= 0 && end_x < cols && end_y >= 0 && end_y < rows {
			curr_height := input.heights[curr_y][curr_x]
			next_height := input.heights[end_y][end_x]
			if next_height <= curr_height+1 {
				set := fmt.Sprint(end_x, ",", end_y)
				if !seen.Contains(set) {
					seen.Add(set)
					curr_positions = append(curr_positions, MapIter{end_x, end_y, steps})
				}
			}
		}
	}

	for len(curr_positions) > 0 {
		curr_iter := curr_positions[0]
		if curr_iter.x == input.end_x && curr_iter.y == input.end_y {
			return fmt.Sprint(curr_iter.steps), nil
		}
		curr_positions = curr_positions[1:]

		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x-1, curr_iter.y, curr_iter.steps+1)
		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x+1, curr_iter.y, curr_iter.steps+1)
		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x, curr_iter.y-1, curr_iter.steps+1)
		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x, curr_iter.y+1, curr_iter.steps+1)
	}
	return "", fmt.Errorf("did not find path")
}

func part2() (string, error) {
	seen := utils.NewHashSet[string]()

	rows := len(input.heights)
	cols := len(input.heights[0])

	curr_positions := []MapIter{{input.end_x, input.end_y, 0}}

	try_add_pos := func(curr_x int, curr_y int, end_x int, end_y int, steps int) {
		if end_x >= 0 && end_x < cols && end_y >= 0 && end_y < rows {
			curr_height := input.heights[curr_y][curr_x]
			next_height := input.heights[end_y][end_x]
			if curr_height <= next_height+1 {
				set := fmt.Sprint(end_x, ",", end_y)
				if !seen.Contains(set) {
					seen.Add(set)
					curr_positions = append(curr_positions, MapIter{end_x, end_y, steps})
				}
			}
		}
	}

	for len(curr_positions) > 0 {
		curr_iter := curr_positions[0]
		if input.heights[curr_iter.y][curr_iter.x] == 'a' {
			return fmt.Sprint(curr_iter.steps), nil
		}
		curr_positions = curr_positions[1:]

		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x-1, curr_iter.y, curr_iter.steps+1)
		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x+1, curr_iter.y, curr_iter.steps+1)
		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x, curr_iter.y-1, curr_iter.steps+1)
		try_add_pos(curr_iter.x, curr_iter.y, curr_iter.x, curr_iter.y+1, curr_iter.steps+1)
	}
	return "", fmt.Errorf("did not find path")
}
