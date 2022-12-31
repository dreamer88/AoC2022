package day20

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

type Val struct {
	v, index int
}

var input []Val

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 20
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	} else {
		input = append(input, Val{i, len(input)})
		return nil
	}
}

func run_step(vals []Val) ([]Val, error) {
	var err error

	for i := 0; i < len(vals); i++ {
		start_index := utils.FindIndex(vals, func(v *Val) bool {
			return v.index == i
		})

		o := vals[start_index]
		if o.v < 0 {
			vals, err = utils.RotateLeft(vals, start_index, -o.v)
			if err != nil {
				return vals, err
			}
		} else {
			vals, err = utils.RotateRight(vals, start_index, o.v)
			if err != nil {
				return vals, err
			}
		}
	}

	return vals, nil
}

func calc_score(vals []Val) int {
	zero_index := utils.FindIndex(vals, func(v *Val) bool {
		return v.v == 0
	})

	v1 := vals[(zero_index+1000)%len(vals)]
	v2 := vals[(zero_index+2000)%len(vals)]
	v3 := vals[(zero_index+3000)%len(vals)]

	return v1.v + v2.v + v3.v
}

func part1() (string, error) {
	new_array := utils.Copy1D(input)

	ans, err := run_step(new_array)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(calc_score(ans)), nil
}

func part2() (string, error) {
	new_array := utils.Map(input, func(v Val) Val {
		return Val{
			v.v * 811589153,
			v.index,
		}
	})

	var err error

	for r := 0; r < 10; r++ {
		new_array, err = run_step(new_array)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprint(calc_score(new_array)), nil
}
