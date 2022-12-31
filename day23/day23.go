package day23

import (
	"aoc2022/utils"
	"fmt"
)

type OpFunc func(utils.Point2Set, utils.Point2) (bool, utils.Point2)

var input utils.Point2Set = utils.NewPoint2Set()
var row int = 0

var operations []OpFunc = []OpFunc{
	func(last_map utils.Point2Set, p utils.Point2) (bool, utils.Point2) {
		n := p.Up()
		if !last_map.Contains(n) && !last_map.Contains(n.Left()) && !last_map.Contains(n.Right()) {
			return true, n
		}
		return false, p
	},
	func(last_map utils.Point2Set, p utils.Point2) (bool, utils.Point2) {
		s := p.Down()
		if !last_map.Contains(s) && !last_map.Contains(s.Left()) && !last_map.Contains(s.Right()) {
			return true, s
		}
		return false, p
	},
	func(last_map utils.Point2Set, p utils.Point2) (bool, utils.Point2) {
		w := p.Left()
		if !last_map.Contains(w) && !last_map.Contains(w.Up()) && !last_map.Contains(w.Down()) {
			return true, w
		}
		return false, p
	},
	func(last_map utils.Point2Set, p utils.Point2) (bool, utils.Point2) {
		e := p.Right()
		if !last_map.Contains(e) && !last_map.Contains(e.Up()) && !last_map.Contains(e.Down()) {
			return true, e
		}
		return false, p
	},
}

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 23
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	for i := 0; i < len(val); i++ {
		if val[i] == '#' {
			input.Add(utils.NewPoint2(i, row))
		}
	}

	row++
	return nil
}

func run_step(starting_map utils.Point2Set, operations []OpFunc) (utils.Point2Set, bool) {
	any_moved := false
	var try_map map[utils.Point2][]utils.Point2 = make(map[utils.Point2][]utils.Point2)
	for _, p := range starting_map.Keys() {
		new_p := p

		n := p.Up()
		e := p.Right()
		s := p.Down()
		w := p.Left()

		if starting_map.Contains(n) || starting_map.Contains(s) ||
			starting_map.Contains(e) || starting_map.Contains(w) ||
			starting_map.Contains(n.Left()) || starting_map.Contains(n.Right()) ||
			starting_map.Contains(s.Left()) || starting_map.Contains(s.Right()) {

			for _, op := range operations {
				handled, op_p := op(starting_map, p)
				if handled {
					any_moved = true
					new_p = op_p
					break
				}
			}
		}

		if _, found := try_map[new_p]; found {
			try_map[new_p] = append(try_map[new_p], p)
		} else {
			try_map[new_p] = []utils.Point2{p}
		}
	}

	new_map := utils.NewPoint2Set()
	for p := range try_map {
		points := try_map[p]
		if len(points) == 1 {
			new_map.Add(p)
		} else {
			for _, old_p := range points {
				new_map.Add(old_p)
			}
		}
	}

	return new_map, any_moved
}

func part1() (string, error) {
	last_map := input
	ops := utils.Copy1D(operations)

	for round := 1; round <= 10; round++ {
		new_map, _ := run_step(last_map, ops)

		last_map = new_map

		ops = []OpFunc{
			ops[1],
			ops[2],
			ops[3],
			ops[0],
		}
	}

	bounding_rect := utils.EmptyBoundingRect()
	for _, p := range last_map.Keys() {
		bounding_rect.UpdateWithPoint(p)
	}

	empty_spaces := 0
	bounding_rect.IteratePoints(func(p utils.Point2) {
		if !last_map.Contains(p) {
			empty_spaces++
		}
	}, nil)

	return fmt.Sprint(empty_spaces), nil
}

func part2() (string, error) {
	last_map := input
	ops := utils.Copy1D(operations)

	round := 1
	for ; ; round++ {
		new_map, moved := run_step(last_map, ops)
		if !moved {
			break
		}

		last_map = new_map
		ops = []OpFunc{
			ops[1],
			ops[2],
			ops[3],
			ops[0],
		}
	}

	return fmt.Sprint(round), nil
}
