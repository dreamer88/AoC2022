package day18

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

var input []utils.Point3

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 18
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	s := strings.Split(val, ",")

	x, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	y, err := strconv.Atoi(s[1])
	if err != nil {
		return err
	}
	z, err := strconv.Atoi(s[2])
	if err != nil {
		return err
	}

	input = append(input, utils.Point3{x, y, z})
	return nil
}

func part1() (string, error) {
	cache := utils.NewPoint3Set()

	for _, p := range input {
		cache.Add(p)
	}

	surface_area := 0

	for _, p := range input {
		if !cache.Contains(p.Left()) {
			surface_area++
		}
		if !cache.Contains(p.Right()) {
			surface_area++
		}
		if !cache.Contains(p.Up()) {
			surface_area++
		}
		if !cache.Contains(p.Down()) {
			surface_area++
		}
		if !cache.Contains(p.Forward()) {
			surface_area++
		}
		if !cache.Contains(p.Back()) {
			surface_area++
		}
	}

	return fmt.Sprint(surface_area), nil
}

func part2() (string, error) {
	interior_map := utils.NewPoint3Set()
	bounds := utils.EmptyBoundingCube()
	for _, p := range input {
		interior_map.Add(p)
		bounds.UpdateWithPoint(p)
	}

	exterior_map := utils.NewPoint3Set()
	exterior_queue := []utils.Point3{}

	bounds.Min.X--
	bounds.Min.Y--
	bounds.Min.Z--
	bounds.Max.X++
	bounds.Max.Y++
	bounds.Max.Z++

	try_add_to_queue := func(p utils.Point3) {
		if bounds.ContainsPoint(p) {
			if !exterior_map.Contains(p) && !interior_map.Contains(p) {
				exterior_map.Add(p)
				exterior_queue = append(exterior_queue, p)
			}
		}
	}

	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			try_add_to_queue(utils.NewPoint3(x, y, bounds.Min.Z))
			try_add_to_queue(utils.NewPoint3(x, y, bounds.Max.Z))
		}

		for z := bounds.Min.Z; z <= bounds.Max.Z; z++ {
			try_add_to_queue(utils.NewPoint3(x, bounds.Min.Y, z))
			try_add_to_queue(utils.NewPoint3(x, bounds.Max.Y, z))
		}
	}

	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for z := bounds.Min.Z; z <= bounds.Max.Z; z++ {
			try_add_to_queue(utils.NewPoint3(bounds.Min.X, y, z))
			try_add_to_queue(utils.NewPoint3(bounds.Max.X, y, z))
		}
	}

	for len(exterior_queue) > 0 {
		p := exterior_queue[0]
		exterior_queue = exterior_queue[1:]
		try_add_to_queue(p.Left())
		try_add_to_queue(p.Right())
		try_add_to_queue(p.Up())
		try_add_to_queue(p.Down())
		try_add_to_queue(p.Forward())
		try_add_to_queue(p.Back())
	}

	surface_area := 0

	test_point := func(p utils.Point3) bool {
		return exterior_map.Contains(p) && !interior_map.Contains(p)
	}

	for _, p := range input {
		if test_point(p.Left()) {
			surface_area++
		}
		if test_point(p.Right()) {
			surface_area++
		}
		if test_point(p.Up()) {
			surface_area++
		}
		if test_point(p.Down()) {
			surface_area++
		}
		if test_point(p.Forward()) {
			surface_area++
		}
		if test_point(p.Back()) {
			surface_area++
		}
	}

	return fmt.Sprint(surface_area), nil
}
