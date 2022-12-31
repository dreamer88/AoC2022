package day24

import (
	"aoc2022/utils"
	"fmt"
)

type Direction uint

const (
	RIGHT Direction = iota
	DOWN
	LEFT
	UP
)

type WindMap = map[utils.Point2][]Direction

var input_rows int = 0
var input_cols int = 0
var input WindMap = WindMap{}

func addWind(m *map[utils.Point2][]Direction, val Direction, p utils.Point2) {
	if _, found := (*m)[p]; found {
		(*m)[p] = append((*m)[p], val)
	} else {
		(*m)[p] = []Direction{val}
	}
}

func hasWind[T any](m *map[utils.Point2]T, p utils.Point2) bool {
	_, found := (*m)[p]
	return found
}

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 24
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	if val[2] != '#' {
		for col := 0; col < len(val)-2; col++ {
			p := utils.Point2{X: col, Y: input_rows}
			switch val[col+1] {
			case '>':
				addWind(&input, RIGHT, p)
			case '<':
				addWind(&input, LEFT, p)
			case '^':
				addWind(&input, UP, p)
			case 'v':
				addWind(&input, DOWN, p)
			}
		}
		input_rows++
	} else {
		input_cols = utils.Max(input_cols, len(val)-2)
	}

	return nil
}

func run_steps(start_p utils.Point2, target_p utils.Point2, wind WindMap) (int, WindMap) {
	bounding_rect := utils.Rect{
		Min: utils.Point2{X: 0, Y: 0},
		Max: utils.Point2{X: input_cols - 1, Y: input_rows - 1},
	}

	num_steps := 1
	states := utils.NewPoint2Set()
	states.Add(start_p)

	for ; ; num_steps++ {
		next_wind := WindMap{}
		for p := range wind {
			for _, wind := range wind[p] {
				switch wind {
				case UP:
					addWind(&next_wind, UP, bounding_rect.WrapPoint(p.Up()))
				case DOWN:
					addWind(&next_wind, DOWN, bounding_rect.WrapPoint(p.Down()))
				case LEFT:
					addWind(&next_wind, LEFT, bounding_rect.WrapPoint(p.Left()))
				case RIGHT:
					addWind(&next_wind, RIGHT, bounding_rect.WrapPoint(p.Right()))
				}
			}
		}

		next_states := utils.NewPoint2Set()
		try_add := func(p utils.Point2) bool {
			if p == target_p {
				return true
			}

			at_start := start_p == p

			if !at_start && !bounding_rect.ContainsPoint(p) {
				return false
			}

			if !hasWind(&next_wind, p) {
				next_states.Add(p)
			}

			return false
		}

		for _, p := range states.Keys() {
			if try_add(p) ||
				try_add(p.Up()) ||
				try_add(p.Down()) ||
				try_add(p.Left()) ||
				try_add(p.Right()) {
				return num_steps, next_wind
			}
		}

		states = next_states
		wind = next_wind
	}
}

func part1() (string, error) {
	start_p := utils.NewPoint2(0, -1)
	target_p := utils.NewPoint2(input_cols-1, input_rows)

	num_steps, _ := run_steps(start_p, target_p, input)

	return fmt.Sprint(num_steps), nil
}

func part2() (string, error) {
	start_p := utils.NewPoint2(0, -1)
	target_p := utils.NewPoint2(input_cols-1, input_rows)
	total_steps := 0

	num_steps, next_wind := run_steps(start_p, target_p, input)
	total_steps += num_steps

	num_steps, next_wind = run_steps(target_p, start_p, next_wind)
	total_steps += num_steps

	num_steps, _ = run_steps(start_p, target_p, next_wind)
	total_steps += num_steps

	return fmt.Sprint(total_steps), nil
}
