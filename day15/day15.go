package day15

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

type Input struct {
	sensor, beacon Point
}

type FoundType uint8

const (
	EMPTY = iota
	SENSOR
	BEACON
	OVERLAPPED
)

var input []Input
var sensor_input_regex = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

const inputType utils.InputType = utils.Input

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 15
	p.InputType = inputType
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	if groups := sensor_input_regex.FindStringSubmatch(val); groups != nil {
		sensor_x, err := strconv.Atoi(groups[1])
		if err != nil {
			return err
		}

		sensor_y, err := strconv.Atoi(groups[2])
		if err != nil {
			return err
		}

		beacon_x, err := strconv.Atoi(groups[3])
		if err != nil {
			return err
		}

		beacon_y, err := strconv.Atoi(groups[4])
		if err != nil {
			return err
		}

		input = append(input, Input{Point{sensor_x, sensor_y}, Point{beacon_x, beacon_y}})

		return nil
	} else {
		return fmt.Errorf("could not parse line: %s", val)
	}
}

func manhatten_dist_xy(p Point, x int, y int) int {
	return utils.Abs(p.x-x) + utils.Abs(p.y-y)
}

func manhatten_dist(p1 Point, p2 Point) int {
	return manhatten_dist_xy(p1, p2.x, p2.y)
}

func part1() (string, error) {
	var locations map[int]FoundType = make(map[int]FoundType, 0)
	smallest_x, largest_x := math.MaxInt, math.MinInt

	add_location := func(x int, t FoundType) {
		if v, found := locations[x]; !found || v == OVERLAPPED {
			locations[x] = t
		}
	}

	check_y := 2000000
	if inputType == utils.Sample {
		check_y = 10
	}

	for _, i := range input {
		dist := manhatten_dist(i.sensor, i.beacon)
		smallest_x, _ = utils.MinMax(smallest_x, i.sensor.x-dist)
		_, largest_x = utils.MinMax(largest_x, i.sensor.x+dist)

		if i.sensor.y == check_y {
			add_location(i.sensor.x, SENSOR)
		}

		if i.beacon.y == check_y {
			add_location(i.beacon.x, BEACON)
		}

		if i.sensor.y-dist <= check_y && check_y <= i.sensor.y+dist {
			for column := i.sensor.x - dist; column <= i.sensor.x+dist; column++ {
				if manhatten_dist(i.sensor, Point{column, check_y}) <= dist {
					add_location(column, OVERLAPPED)
				}
			}
		}
	}

	cannot_be_beacon := 0
	for i := smallest_x; i <= largest_x; i++ {
		v, found := locations[i]
		if found && v == OVERLAPPED {
			cannot_be_beacon++
		}
	}

	return fmt.Sprint(cannot_be_beacon), nil
}

func part2() (string, error) {
	type Range struct {
		start, end int
	}

	x_min, x_max, y_min, y_max := 0, 4000000, 0, 4000000
	if inputType == utils.Sample {
		x_min, x_max, y_min, y_max = 0, 20, 0, 20
	}

	var ranges map[int][]Range = make(map[int][]Range, 0)

	add_range := func(y int, start_x int, end_x int) {
		if y < y_min || y > y_max {
			return
		}

		start_x = utils.Max(start_x, x_min)
		end_x = utils.Min(end_x, x_max)

		if start_x > x_max || end_x < x_min {
			return
		}

		if r, found := ranges[y]; found {
			r = append(r, Range{start_x, end_x})

			utils.Sort(r, func(a *Range, b *Range) bool {
				return a.start <= b.start
			})

			new_ranges := []Range{r[0]}
			for i := 1; i < len(r); i++ {
				last_r := &new_ranges[len(new_ranges)-1]
				if last_r.end+1 >= r[i].start {
					last_r.end = utils.Max(last_r.end, r[i].end)
				} else {
					new_ranges = append(new_ranges, r[i])
				}
			}
			ranges[y] = new_ranges
		} else {
			ranges[y] = []Range{{start_x, end_x}}
		}
	}

	for _, i := range input {
		dist := manhatten_dist(i.sensor, i.beacon)

		for y_dist := 0; y_dist <= dist; y_dist++ {
			x_dist := dist - y_dist
			add_range(i.sensor.y-y_dist, i.sensor.x-x_dist, i.sensor.x+x_dist)
			if y_dist != 0 {
				add_range(i.sensor.y+y_dist, i.sensor.x-x_dist, i.sensor.x+x_dist)
			}
		}
	}

	for y, r := range ranges {
		if len(r) == 2 {
			if r[0].end+2 != r[1].start {
				return "", fmt.Errorf("found a gap in row %d but it was too large", y)
			}

			return fmt.Sprint((r[0].end+1)*4000000 + y), nil
		} else if len(r) == 1 {
			if r[0].start > x_min {
				if r[0].start != 0 {
					return "", fmt.Errorf("found a starting gap in row %d but it was too large", y)
				}
				return fmt.Sprint((r[0].end+1)*4000000 + y), nil
			} else if r[0].end < x_max {
				if r[0].end != x_max {
					return "", fmt.Errorf("found a starting gap in row %d but it was too large", y)
				}
				return fmt.Sprint((r[0].end+1)*4000000 + y), nil
			}
		}
	}

	return "", fmt.Errorf("could not find empy location")
}
