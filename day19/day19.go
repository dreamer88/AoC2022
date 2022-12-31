package day19

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Resource int8

const (
	RES_START Resource = iota
	ORE       Resource = iota - 1
	CLAY
	OBSIDIAN
	GEODE
	RES_COUNT
)

type Purchaseable uint8

const (
	NOT        Purchaseable = iota
	NEW        Purchaseable = iota
	PREVIOUSLY Purchaseable = iota
)

type Cost struct {
	count    int
	resource Resource
}

type Blueprint struct {
	idx           int
	ore_cost      []Cost
	clay_cost     []Cost
	obsidian_cost []Cost
	geode_cost    []Cost
}

type State struct {
	robots, resources [RES_COUNT]int
}

var input []Blueprint

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 19
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	s := strings.Split(val, ": ")
	bp := strings.Split(s[0], " ")
	bp_idx, err := strconv.Atoi(bp[1])
	if err != nil {
		return err
	}

	getType := func(s string) Resource {
		switch s {
		case "ore":
			return ORE

		case "clay":
			return CLAY

		case "obsidian":
			return OBSIDIAN

		default:
			panic(fmt.Sprint("can't parse ", s))
		}
	}

	parseRobot := func(s string) []Cost {
		var res []Cost
		s = strings.Trim(s, ".")

		p := strings.Split(s, "costs ")

		re := strings.Split(p[1], " and ")
		for _, r := range re {
			v := strings.Split(r, " ")

			t := getType(v[1])
			i, _ := strconv.Atoi(v[0])
			res = append(res, Cost{i, t})
		}

		return res
	}

	robots := strings.Split(s[1], ". ")

	blue := Blueprint{bp_idx, parseRobot(robots[0]), parseRobot(robots[1]), parseRobot(robots[2]), parseRobot(robots[3])}
	input = append(input, blue)

	return nil
}

func copy_state(s *State) State {
	return State{
		[RES_COUNT]int{
			s.robots[ORE],
			s.robots[CLAY],
			s.robots[OBSIDIAN],
			s.robots[GEODE],
		}, [RES_COUNT]int{
			s.resources[ORE],
			s.resources[CLAY],
			s.resources[OBSIDIAN],
			s.resources[GEODE],
		},
	}
}

func run_dfs(blueprints []Blueprint, minutes int) []int {
	qualities := []int{}

	for _, bp := range blueprints {
		get_costs := func(r Resource) []Cost {
			switch r {
			case ORE:
				return bp.ore_cost
			case CLAY:
				return bp.clay_cost
			case OBSIDIAN:
				return bp.obsidian_cost
			case GEODE:
				return bp.geode_cost
			}

			return []Cost{}
		}

		base_state := State{
			[RES_COUNT]int{1, 0, 0, 0},
			[RES_COUNT]int{0, 0, 0, 0},
		}

		turns_until_robot := func(s *State, r Resource) int {
			max_turns := 0
			for _, c := range get_costs(r) {
				if s.robots[c.resource] == 0 {
					return -1
				}

				max_turns = utils.Max(max_turns, utils.Ceil(c.count-s.resources[c.resource], s.robots[c.resource]))
			}

			return max_turns
		}

		max_geodes := 0

		var dfs func(s State, minutes int)
		dfs = func(s State, minutes int) {
			remaining_geodes := s.resources[GEODE] + s.robots[GEODE]*minutes
			max_geodes = utils.Max(max_geodes, remaining_geodes)

			if remaining_geodes+(minutes*(minutes-1))/2 <= max_geodes {
				return
			}

			for res := GEODE; res >= ORE; res-- {
				turns_to_next := turns_until_robot(&s, res)
				if turns_to_next >= 0 && turns_to_next < minutes {
					new_s := copy_state(&s)
					for _, c := range get_costs(res) {
						new_s.resources[c.resource] -= c.count
					}

					for r := RES_START; r < RES_COUNT; r++ {
						new_s.resources[r] += new_s.robots[r] * (turns_to_next + 1)
					}

					new_s.robots[res]++
					dfs(new_s, minutes-turns_to_next-1)
				}
			}
		}

		dfs(base_state, minutes)

		qualities = append(qualities, max_geodes)
	}

	return qualities
}

func part1() (string, error) {
	qualities := run_dfs(input, 24)
	bp := 0

	qualities = utils.Map(qualities, func(v int) int {
		bp++
		return v * bp
	})
	total_quality := utils.Reduce(qualities, 0, func(q int, s int) int {
		return q + s
	})
	return fmt.Sprint(total_quality), nil
}

func part2() (string, error) {
	qualities := run_dfs(input[:3], 32)
	total_quality := utils.Reduce(qualities, 1, func(q int, s int) int {
		return q * s
	})
	return fmt.Sprint(total_quality), nil
}
