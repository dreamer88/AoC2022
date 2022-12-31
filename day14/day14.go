package day14

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Vals uint8

const (
	AIR Vals = iota
	WALL
	SAND
)

var input [][]utils.Point2

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 14
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	var parts []utils.Point2

	for _, s := range strings.Split(val, " -> ") {
		x_y := strings.Split(s, ",")
		x, err := strconv.Atoi(x_y[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(x_y[1])
		if err != nil {
			return err
		}
		parts = append(parts, utils.NewPoint2(x, y))
	}

	input = append(input, parts)

	return nil
}

func convertToMap() (map[utils.Point2]Vals, int) {
	var res map[utils.Point2]Vals = make(map[utils.Point2]Vals, 0)
	lowest := 0
	for _, p := range input {
		for i := 0; i < len(p)-1; i++ {
			rect := utils.ContainingRect(p[i], p[i+1])
			lowest = utils.Max(rect.Max.Y, lowest)
			rect.IteratePoints(func(p utils.Point2) {
				res[p] = WALL
			}, nil)
		}
	}
	return res, lowest
}

func part1() (string, error) {
	m, lowest := convertToMap()
	settled_pieces := 0

	sampleSpace := func(p utils.Point2) Vals {
		if v, found := m[p]; found {
			return v
		}

		return AIR
	}

out:
	for {
		p := utils.NewPoint2(500, 0)
		for {
			p_down := p.Down()
			if sampleSpace(p_down) == AIR {
				p = p_down
				if p.Y > lowest {
					break out
				}
			} else if p_left := p_down.Left(); sampleSpace(p_left) == AIR {
				p = p_left
			} else if p_right := p_down.Right(); sampleSpace(p_right) == AIR {
				p = p_right
			} else {
				m[p] = SAND
				settled_pieces++
				break
			}
		}
	}

	return fmt.Sprint(settled_pieces), nil
}

func part2() (string, error) {
	m, lowest := convertToMap()
	settled_pieces := 0

	sampleSpace := func(p utils.Point2) Vals {
		if p.Y == lowest+2 {
			return WALL
		}

		if v, found := m[p]; found {
			return v
		}

		return AIR
	}

out:
	for {
		p := utils.NewPoint2(500, 0)
		for {
			p_down := p.Down()
			if sampleSpace(p_down) == AIR {
				p = p_down
			} else if p_left := p_down.Left(); sampleSpace(p_left) == AIR {
				p = p_left
			} else if p_right := p_down.Right(); sampleSpace(p_right) == AIR {
				p = p_right
			} else {
				m[p] = SAND
				settled_pieces++
				if p.Y == 0 {
					break out
				} else {
					break
				}
			}
		}
	}

	return fmt.Sprint(settled_pieces), nil
}
