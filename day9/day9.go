package day9

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Dir uint8

const (
	Up Dir = iota
	Down
	Left
	Right
)

type Move struct {
	dir    Dir
	length int
}

var input []Move
var reg = regexp.MustCompile(`(\w) (\d+)`)

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 9
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	if match := reg.FindStringSubmatch(val); match != nil {
		dir := Up
		switch match[1] {
		case "U":
			break
		case "D":
			dir = Down
		case "L":
			dir = Left
		case "R":
			dir = Right
		default:
			return fmt.Errorf("failed to parse dir %s", match[1])
		}

		count, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}

		input = append(input, Move{dir, count})
		return nil
	} else {
		return fmt.Errorf("failed to parse %s", val)
	}
}

func run_knots(knots []utils.Point2) int {
	seen := utils.NewPoint2Set()
	seen.Add(knots[len(knots)-1])

	var update_knot func(int)
	update_knot = func(idx int) {
		head_pos := &knots[idx-1]
		tail_pos := &knots[idx]

		diff := head_pos.Sub(*tail_pos)
		dist_x, dist_y := utils.Abs(diff.X), utils.Abs(diff.Y)

		if dist_x >= 2 {
			tail_pos.X += diff.X / dist_x
			if dist_y > 0 {
				tail_pos.Y += diff.Y / dist_y
			}
		} else if dist_y >= 2 {
			tail_pos.Y += diff.Y / dist_y
			if dist_x > 0 {
				tail_pos.X += diff.X / dist_x
			}
		}

		next_idx := idx + 1
		if next_idx < len(knots) {
			update_knot(next_idx)
		} else {
			seen.Add(*tail_pos)
		}
	}

	for _, step := range input {
		for i := 0; i < step.length; i++ {
			switch step.dir {
			case Up:
				knots[0] = knots[0].Up()
			case Down:
				knots[0] = knots[0].Down()
			case Left:
				knots[0] = knots[0].Left()
			case Right:
				knots[0] = knots[0].Right()
			}
			update_knot(1)
		}
	}

	return seen.Length()
}

func part1() (string, error) {
	knots := [2]utils.Point2{}
	count := run_knots(knots[:])
	return fmt.Sprint(count), nil
}

func part2() (string, error) {
	knots := [10]utils.Point2{}
	count := run_knots(knots[:])
	return fmt.Sprint(count), nil
}
