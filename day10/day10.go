package day10

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Op uint8

const (
	NOOP = iota
	ADDX
)

type Inst struct {
	op Op
	v  int
}

var input []Inst
var noop_reg = regexp.MustCompile(`noop`)
var add_reg = regexp.MustCompile(`addx (-?\d+)`)

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 10
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	if noop_reg.MatchString(val) {
		input = append(input, Inst{NOOP, 0})
		return nil
	} else if groups := add_reg.FindStringSubmatch(val); groups != nil {
		v, err := strconv.Atoi(groups[1])
		if err != nil {
			return err
		}

		input = append(input, Inst{ADDX, v})
	} else {
		return fmt.Errorf("unparsed command %s", val)
	}

	return nil
}

func part1() (string, error) {
	x := 1
	count := 0
	current_cycle := 1

	update_cycle := func() {
		current_cycle += 1
		if (current_cycle-20)%40 == 0 {
			add_val := x * current_cycle
			count += add_val
		}
	}

	for _, v := range input {
		switch v.op {
		case NOOP:
			update_cycle()

		case ADDX:
			update_cycle()
			x += v.v
			update_cycle()
		}
	}

	return fmt.Sprint(count), nil
}

func part2() (string, error) {
	x := 0
	current_cycle := 0
	var crt [6][40]byte

	update_cycle := func() {
		pixel := current_cycle % 40
		row := int(math.Floor(float64(current_cycle) / 40.0))
		if row < 6 {
			if pixel >= x && pixel <= (x+2) {
				crt[row][pixel] = '#'
			} else {
				crt[row][pixel] = ' '
			}
		}
		current_cycle += 1
	}

	for _, v := range input {
		switch v.op {
		case NOOP:
			update_cycle()

		case ADDX:
			update_cycle()
			update_cycle()
			x += v.v
		}
	}

	result := ""
	for _, v := range crt {
		result = result + "\n" + string(v[:])
	}

	return result, nil
}
