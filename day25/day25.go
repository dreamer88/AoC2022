package day25

import (
	"aoc2022/utils"
	"fmt"
)

var input []string

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 25
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	input = append(input, val)
	return nil
}

func snafuToInt(v string) (int, error) {
	count := 0

	for i := 0; i < len(v); i++ {
		count *= 5
		switch v[i] {
		case '2':
			count += 2
		case '1':
			count += 1
		case '0':
			count += 0
		case '-':
			count -= 1
		case '=':
			count -= 2
		default:
			return 0, fmt.Errorf("invalid token")
		}
	}

	return count, nil
}

func intToSnafu(v int) (string, error) {
	snafu := ""
	for v > 0 {
		mod := (v+2)%5 - 2
		v = (v - mod) / 5

		switch mod {
		case -2:
			snafu = "=" + snafu
		case -1:
			snafu = "-" + snafu
		case 0:
			snafu = "0" + snafu
		case 1:
			snafu = "1" + snafu
		case 2:
			snafu = "2" + snafu
		default:
			return "", fmt.Errorf("invalid snafu")
		}
	}

	return snafu, nil
}

func part1() (string, error) {
	sum := 0
	for _, v := range input {
		c, err := snafuToInt(v)
		if err != nil {
			return "", err
		}

		sum += c
	}

	snafu, err := intToSnafu(sum)
	return snafu, err
}

func part2() (string, error) {
	return "CONGRATS ON FINISHING", nil
}
