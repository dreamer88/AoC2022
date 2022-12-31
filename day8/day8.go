package day8

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

var input [][]int

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 8
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	var new_row []int
	for _, c := range val {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			return err
		} else {
			new_row = append(new_row, i)
		}
	}

	input = append(input, new_row)
	return nil
}

func run_range(visible [][]bool, iterator func() (bool, int, int)) {
	tallest := -1
	for {
		valid, row, col := iterator()
		if !valid {
			return
		}

		val := input[row][col]
		if val > tallest {
			visible[row][col] = true
			tallest = val
		}
	}
}

func part1() (string, error) {
	var visible [][]bool

	rows := len(input)
	cols := len(input[0])

	for i := 0; i < rows; i++ {
		var new_row []bool
		for j := 0; j < cols; j++ {
			new_row = append(new_row, i == 0 || i == (rows-1) || j == 0 || j == (cols-1))
		}
		visible = append(visible, new_row)
	}

	for i := 0; i < rows; i++ {
		col := 0
		run_range(visible, func() (bool, int, int) {
			next := col
			col = col + 1
			return next < cols, i, next
		})

		col = cols - 1
		run_range(visible, func() (bool, int, int) {
			next := col
			col = col - 1
			return next >= 0, i, next
		})
	}

	for j := 0; j < rows; j++ {
		row := 0
		run_range(visible, func() (bool, int, int) {
			next := row
			row = row + 1
			return next < rows, next, j
		})

		row = rows - 1
		run_range(visible, func() (bool, int, int) {
			next := row
			row = row - 1
			return next >= 0, next, j
		})
	}

	visible_count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visible[i][j] {
				visible_count = visible_count + 1
			}
		}
	}

	return fmt.Sprint(visible_count), nil
}

func run_scenic(iterator func() (bool, int, int)) int {
	_, srow, scol := iterator()
	start_height := input[srow][scol]
	view := 0
	for {
		valid, row, col := iterator()
		if !valid {
			return view
		}

		view = view + 1

		val := input[row][col]
		if val >= start_height {
			return view
		}
	}
}

func part2() (string, error) {
	rows := len(input)
	cols := len(input[0])

	highest_score := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			our_score := 1

			row := i
			our_score *= run_scenic(func() (bool, int, int) {
				next := row
				row = row + 1
				return next < rows, next, j
			})

			row = i
			our_score *= run_scenic(func() (bool, int, int) {
				next := row
				row = row - 1
				return next >= 0, next, j
			})

			col := j
			our_score *= run_scenic(func() (bool, int, int) {
				next := col
				col = col + 1
				return next < cols, i, next
			})

			col = j
			our_score *= run_scenic(func() (bool, int, int) {
				next := col
				col = col - 1
				return next >= 0, i, next
			})

			if our_score > highest_score {
				highest_score = our_score
			}
		}
	}

	return fmt.Sprint(highest_score), nil
}
