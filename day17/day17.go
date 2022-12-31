package day17

import (
	"aoc2022/utils"
	"fmt"
	"math"
)

type Space uint8

const (
	NONE Space = iota
	ROCK Space = iota
)

var input string

var stones [][][]Space = [][][]Space{
	{
		{ROCK, ROCK, ROCK, ROCK},
	},

	{
		{NONE, ROCK, NONE},
		{ROCK, ROCK, ROCK},
		{NONE, ROCK, NONE},
	},

	{
		{NONE, NONE, ROCK},
		{NONE, NONE, ROCK},
		{ROCK, ROCK, ROCK},
	},

	{
		{ROCK},
		{ROCK},
		{ROCK},
		{ROCK},
	},

	{
		{ROCK, ROCK},
		{ROCK, ROCK},
	},
}

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 17
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	input = val
	return nil
}

func part1() (string, error) {
	const tunnel_width int = 7
	var rows [][tunnel_width]bool = [][tunnel_width]bool{}
	tunnel_point_empty := func(x int, y int) bool {
		if y < 0 || x < 0 || x >= tunnel_width {
			return false
		} else if y >= len(rows) {
			return true
		} else {
			return !rows[y][x]
		}
	}

	rock_fits := func(rock *[][]Space, x int, bottom_y int) bool {
		rock_height := len(*rock)
		rock_length := len((*rock)[0])
		for row := 0; row < rock_height; row++ {
			rock_row := rock_height - row - 1
			check_height := bottom_y + row

			for rock_col := 0; rock_col < rock_length; rock_col++ {
				if (*rock)[rock_row][rock_col] == ROCK {
					if !tunnel_point_empty(x+rock_col, check_height) {
						return false
					}
				}
			}
		}

		return true
	}

	wind_index := 0
	for r := 0; r < 2022; r++ {
		rock := stones[r%len(stones)]
		rock_height := len(rock)
		rock_length := len(rock[0])

		rock_x := 2
		rock_bottom_y := len(rows) + 3

	rock_loop:
		for {
			wind_dir := input[wind_index]
			wind_index = (wind_index + 1) % len(input)

			switch wind_dir {
			case '<':
				new_x := rock_x - 1
				if rock_fits(&rock, new_x, rock_bottom_y) {
					rock_x = new_x
				}
			case '>':
				new_x := utils.Min(tunnel_width-rock_length, rock_x+1)
				if rock_fits(&rock, new_x, rock_bottom_y) {
					rock_x = new_x
				}
			default:
				return "", fmt.Errorf("invalid wind direction %b", wind_dir)
			}

			if rock_fits(&rock, rock_x, rock_bottom_y-1) {
				rock_bottom_y--
			} else {
				break rock_loop
			}
		}

		// add space for our new rock
		for i := len(rows); i < rock_bottom_y+rock_height; i++ {
			rows = append(rows, [tunnel_width]bool{})
		}

		// add our new rock
		for rock_row := 0; rock_row < rock_height; rock_row++ {
			for rock_col := 0; rock_col < rock_length; rock_col++ {
				if rock[rock_row][rock_col] == ROCK {
					rows[(rock_height-rock_row)+rock_bottom_y-1][rock_x+rock_col] = true
				}
			}
		}
	}

	return fmt.Sprint(len(rows)), nil
}

func part2() (string, error) {
	const rocks_wanted int = 1000000000000
	const tunnel_width int = 7
	var rows [][tunnel_width]bool = [][tunnel_width]bool{}
	tunnel_point_empty := func(x int, y int) bool {
		if y < 0 || x < 0 || x >= tunnel_width {
			return false
		} else if y >= len(rows) {
			return true
		} else {
			return !rows[y][x]
		}
	}

	rock_fits := func(rock *[][]Space, x int, bottom_y int) bool {
		rock_height := len(*rock)
		rock_length := len((*rock)[0])
		for row := 0; row < rock_height; row++ {
			rock_row := rock_height - row - 1
			check_height := bottom_y + row

			for rock_col := 0; rock_col < rock_length; rock_col++ {
				if (*rock)[rock_row][rock_col] == ROCK {
					if !tunnel_point_empty(x+rock_col, check_height) {
						return false
					}
				}
			}
		}

		return true
	}

	type CacheObject struct {
		row, start_height, end_height int
	}

	r := 0
	var previous map[string]int = make(map[string]int)
	var previous_rounds []CacheObject

	calc_cache_key := func(rock_index int, wind_index int) string {
		var has_value [tunnel_width]bool
		cache_index := fmt.Sprint(rock_index, wind_index)

		for row := len(rows) - 1; row >= 0; row-- {
			for col := 0; col < tunnel_width; col++ {
				if tunnel_point_empty(col, row) {
					cache_index += "."
				} else {
					cache_index += "#"
					has_value[col] = true
				}
			}

			if utils.All(has_value[:], func(_ int, b *bool) bool { return *b }) {
				break
			}
		}

		return cache_index
	}

	wind_index := 0

round_loop:
	for ; r < rocks_wanted; r++ {
		rock_index := r % len(stones)
		cache_index := calc_cache_key(rock_index, wind_index)

		if v, found := previous[cache_index]; found {
			r = v
			break round_loop
		} else {
			previous[cache_index] = r
		}

		start_height := len(rows)

		rock := stones[rock_index]
		rock_height := len(rock)
		rock_length := len(rock[0])

		rock_x := 2
		rock_bottom_y := len(rows) + 3

	rock_loop:
		for {
			wind_dir := input[wind_index]
			wind_index = (wind_index + 1) % len(input)

			switch wind_dir {
			case '<':
				new_x := rock_x - 1
				if rock_fits(&rock, new_x, rock_bottom_y) {
					rock_x = new_x
				}
			case '>':
				new_x := utils.Min(tunnel_width-rock_length, rock_x+1)
				if rock_fits(&rock, new_x, rock_bottom_y) {
					rock_x = new_x
				}
			default:
				return "", fmt.Errorf("invalid wind direction %b", wind_dir)
			}

			if rock_fits(&rock, rock_x, rock_bottom_y-1) {
				rock_bottom_y--
			} else {
				break rock_loop
			}
		}

		// add space for our new rock
		for i := len(rows); i < rock_bottom_y+rock_height; i++ {
			rows = append(rows, [tunnel_width]bool{})
		}

		// add our new rock
		for rock_row := 0; rock_row < rock_height; rock_row++ {
			for rock_col := 0; rock_col < rock_length; rock_col++ {
				if rock[rock_row][rock_col] == ROCK {
					rows[(rock_height-rock_row)+rock_bottom_y-1][rock_x+rock_col] = true
				}
			}
		}

		previous_rounds = append(previous_rounds, CacheObject{r, start_height, len(rows)})
	}

	cache_start := r
	cache_end := len(previous_rounds) - 1
	cache_length := cache_end - cache_start + 1
	per_cache_cycle := previous_rounds[cache_end].end_height - previous_rounds[cache_start].start_height

	round_left := rocks_wanted - (r + 1)
	num_copied := int(math.Floor(float64(round_left) / float64(cache_length)))

	length := previous_rounds[cache_start].start_height + num_copied*per_cache_cycle
	r += cache_length * num_copied

	remaining_start_index := (r-cache_start)%cache_length + cache_start
	remaining_start := previous_rounds[remaining_start_index].start_height

	remaining_end_index := (rocks_wanted-cache_start-1)%cache_length + cache_start
	remaining_end := previous_rounds[remaining_end_index].end_height

	remainder := remaining_end - remaining_start
	length += remainder

	r += (remaining_end_index - remaining_start_index)

	return fmt.Sprint(length), nil
}
