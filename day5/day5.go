package day5

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Move struct {
	count, from, to int
}

var moveParse = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
var parsingMoves bool = false
var input [][]byte

var moves []Move

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 5
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	if val == "" {
		parsingMoves = true
		return nil
	} else if !parsingMoves {
		var row []byte
		any_valid := false
		for idx := 1; idx < len(val); idx += 4 {
			if val[idx-1] == '[' {
				any_valid = true
			}

			row = append(row, val[idx])
		}

		if any_valid {
			input = append(input, row)
		}
	} else {
		match := moveParse.FindStringSubmatch(val)
		if match == nil {
			return fmt.Errorf("failed to parse line %s", val)
		}

		count, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		}

		from, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}

		to, err := strconv.Atoi(match[3])
		if err != nil {
			return err
		}

		moves = append(moves, Move{count, from - 1, to - 1})
	}

	return nil
}

func getFirstEmptySpace(column int, crates *[][]byte) int {
	for row := 0; row < len(*crates); row++ {
		if (*crates)[row][column] == ' ' {
			return row
		}
	}

	return len(*crates)
}

func doMove(move Move, crates *[][]byte) {
	current_to_row := getFirstEmptySpace(move.to, crates)
	current_from_row := getFirstEmptySpace(move.from, crates) - 1

	for m := 0; m < move.count; m++ {
		for current_to_row >= len(*crates) {
			var new_row []byte
			for range (*crates)[0] {
				new_row = append(new_row, ' ')
			}
			(*crates) = append(*crates, new_row)
		}

		(*crates)[current_to_row][move.to] = (*crates)[current_from_row][move.from]
		(*crates)[current_from_row][move.from] = ' '
		current_to_row++
		current_from_row--
	}
}

func getTops(crates *[][]byte) []byte {
	var result []byte

	for column := 0; column < len((*crates)[0]); column++ {
		idx := getFirstEmptySpace(column, crates)
		result = append(result, (*crates)[idx-1][column])
	}

	return result
}

func part1() (string, error) {
	var crates [][]byte
	for row := len(input) - 1; row >= 0; row-- {
		cpy := make([]byte, len(input[row]))
		copy(cpy, input[row])
		crates = append(crates, cpy)
	}

	for _, move := range moves {
		doMove(move, &crates)
	}

	results := getTops(&crates)

	return string(results), nil
}

func doMove2(move Move, crates *[][]byte) {
	current_to_row := getFirstEmptySpace(move.to, crates)
	current_from_row := getFirstEmptySpace(move.from, crates) - move.count

	for m := 0; m < move.count; m++ {
		for current_to_row >= len(*crates) {
			var new_row []byte
			for range (*crates)[0] {
				new_row = append(new_row, ' ')
			}
			(*crates) = append(*crates, new_row)
		}

		(*crates)[current_to_row][move.to] = (*crates)[current_from_row][move.from]
		(*crates)[current_from_row][move.from] = ' '
		current_to_row++
		current_from_row++
	}
}

func part2() (string, error) {
	var crates [][]byte
	for row := len(input) - 1; row >= 0; row-- {
		cpy := make([]byte, len(input[row]))
		copy(cpy, input[row])
		crates = append(crates, cpy)
	}

	for _, move := range moves {
		doMove2(move, &crates)
	}

	results := getTops(&crates)

	return string(results), nil
}
