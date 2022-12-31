package day16

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Valve struct {
	flow      int
	connected []string
}

var input map[string]Valve = make(map[string]Valve)
var valve_regex = regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w, ]+)`)

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 16
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	if groups := valve_regex.FindStringSubmatch(val); groups != nil {
		valve := groups[1]
		rate, err := strconv.Atoi(groups[2])
		if err != nil {
			return err
		}

		connected := strings.Split(groups[3], ", ")

		input[valve] = Valve{rate, connected}
		return nil
	} else {
		return fmt.Errorf("could not parse line '%s'", val)
	}
}

func part1() (string, error) {
	unopened_valves := utils.NewHashSet[string]()
	for k := range input {
		unopened_valves.Add(k)
	}

	// generate fastest paths
	var fastestPaths map[string]map[string][]string = make(map[string]map[string][]string)
	for _, name := range unopened_valves.Keys() {
		fastestPaths[name] = map[string][]string{}

		unaccessed := unopened_valves.Clone()
		unaccessed.Remove(name)
		current_stack := [][]string{{name}}

		for len(current_stack) > 0 {
			next := current_stack[0]
			current_stack = current_stack[1:]

			next_name := next[len(next)-1]

			fastestPaths[name][next_name] = utils.Copy1D(next)

			for _, other := range input[next_name].connected {
				if unaccessed.Contains(other) {
					unaccessed.Remove(other)
					new_stack := utils.Copy1D(next)
					new_stack = append(new_stack, other)

					current_stack = append(current_stack, new_stack)
				}
			}
		}
	}

	valid_valves := utils.NewHashSet[string]()
	for k := range input {
		if input[k].flow != 0 {
			valid_valves.Add(k)
		}
	}

	type Stack struct {
		current_pos       string
		pressure_released int
		time_left         int
		valves_left       *utils.HashSet[string]
	}

	cache := make(map[string]int)

	finished_stack_max := 0
	valid_remaining := func(s *Stack) bool {
		remaining_valves := s.valves_left.Keys()
		slices.Sort(remaining_valves)
		cache_key := s.current_pos + ";" + strings.Join(remaining_valves, "")

		if v, found := cache[cache_key]; found {
			if v > s.pressure_released {
				return false
			}
		}

		cache[cache_key] = s.pressure_released

		remaining := s.pressure_released
		if finished_stack_max < remaining {
			finished_stack_max = remaining
			return true
		}

		if len(remaining_valves) == 0 {
			return false
		}

		utils.Sort(remaining_valves, func(a *string, b *string) bool {
			return input[*a].flow > input[*b].flow
		})

		time_left := s.time_left - 1
		for _, v := range remaining_valves {
			remaining += input[v].flow * time_left
			time_left--
		}

		return finished_stack_max < remaining
	}

	var dfs func(s Stack)
	dfs = func(s Stack) {
		if s.valves_left.Remove(s.current_pos) {
			s.pressure_released += input[s.current_pos].flow * (s.time_left - 1)
		}

		if !valid_remaining(&s) {
			return
		}

		remaining := s.valves_left.Keys()
		utils.Sort(remaining, func(a *string, b *string) bool {
			return input[*a].flow > input[*b].flow
		})

		for _, new_self := range remaining {
			s_path := fastestPaths[s.current_pos][new_self]

			min_length := len(s_path)
			if min_length < s.time_left {
				dfs(Stack{
					s_path[min_length-1],
					s.pressure_released,
					s.time_left - min_length,
					utils.HashSetFromSlice(remaining),
				})
			}
		}
	}

	// adding one to handle our initial movements
	dfs(Stack{"AA", 0, 31, valid_valves.Clone()})

	return fmt.Sprint(finished_stack_max), nil
}

func part2() (string, error) {
	unopened_valves := utils.NewHashSet[string]()
	for k := range input {
		unopened_valves.Add(k)
	}

	// generate fastest paths
	var fastestPaths map[string]map[string][]string = make(map[string]map[string][]string)
	for _, name := range unopened_valves.Keys() {
		fastestPaths[name] = map[string][]string{}

		unaccessed := unopened_valves.Clone()
		unaccessed.Remove(name)
		current_stack := [][]string{{name}}

		for len(current_stack) > 0 {
			next := current_stack[0]
			current_stack = current_stack[1:]

			next_name := next[len(next)-1]

			fastestPaths[name][next_name] = utils.Copy1D(next)

			for _, other := range input[next_name].connected {
				if unaccessed.Contains(other) {
					unaccessed.Remove(other)
					new_stack := utils.Copy1D(next)
					new_stack = append(new_stack, other)

					current_stack = append(current_stack, new_stack)
				}
			}
		}
	}

	valid_valves := utils.NewHashSet[string]()
	for k := range input {
		if input[k].flow != 0 {
			valid_valves.Add(k)
		}
	}

	type Stack struct {
		current_pos, target_pos, e_current_pos, target_e_pos string
		pressure_released                                    int
		time_left                                            int
		valves_left                                          *utils.HashSet[string]
	}

	cache := make(map[string]int)

	finished_stack_max := 0
	valid_remaining := func(s *Stack) bool {
		cache_key := ""
		if s.current_pos < s.e_current_pos {
			cache_key += s.current_pos + s.e_current_pos
		} else {
			cache_key += s.e_current_pos + s.current_pos
		}

		remaining_valves := s.valves_left.Keys()
		slices.Sort(remaining_valves)
		cache_key += ";" + strings.Join(remaining_valves, "")

		if v, found := cache[cache_key]; found {
			if v > s.pressure_released {
				return false
			}
		}

		cache[cache_key] = s.pressure_released

		remaining := s.pressure_released
		if finished_stack_max < remaining {
			finished_stack_max = remaining
			return true
		}

		if len(remaining_valves) == 0 {
			return false
		}

		utils.Sort(remaining_valves, func(a *string, b *string) bool {
			return input[*a].flow > input[*b].flow
		})

		time_left := s.time_left - 1
		for _, w := range utils.Windows(remaining_valves, 2) {
			for _, v := range w {
				remaining += input[v].flow * time_left
			}
			time_left--
		}

		return finished_stack_max < remaining
	}

	var dfs func(s Stack)
	dfs = func(s Stack) {
		turned_s_valve := false
		turned_e_valve := false
		if s.current_pos == s.target_pos && s.valves_left.Remove(s.current_pos) {
			s.pressure_released += input[s.current_pos].flow * (s.time_left - 1)
			turned_s_valve = true
		}

		if s.e_current_pos == s.target_e_pos && s.valves_left.Remove(s.e_current_pos) {
			s.pressure_released += input[s.e_current_pos].flow * (s.time_left - 1)
			turned_e_valve = true
		}

		if !valid_remaining(&s) {
			return
		}

		remaining := s.valves_left.Keys()
		utils.Sort(remaining, func(a *string, b *string) bool {
			return input[*a].flow > input[*b].flow
		})

		var s_targets []string
		if turned_s_valve {
			s_targets = remaining
		} else {
			s_targets = []string{s.target_pos}
		}

		var e_targets []string
		if turned_e_valve {
			e_targets = remaining
		} else {
			e_targets = []string{s.target_e_pos}
		}

		for _, new_self := range s_targets {
			for _, new_e := range e_targets {
				if new_self == new_e {
					continue
				}

				s_path := fastestPaths[s.current_pos][new_self]
				e_path := fastestPaths[s.e_current_pos][new_e]

				if !turned_s_valve {
					s_path = s_path[1:]
				}
				if !turned_e_valve {
					e_path = e_path[1:]
				}

				min_length := utils.Min(len(s_path), len(e_path))
				if min_length < s.time_left {
					dfs(Stack{
						s_path[min_length-1],
						s_path[len(s_path)-1],
						e_path[min_length-1],
						e_path[len(e_path)-1],
						s.pressure_released,
						s.time_left - min_length,
						utils.HashSetFromSlice(remaining),
					})
				}
			}
		}
	}

	valid_valves_keys := valid_valves.Keys()
	for i, s_pos := range valid_valves_keys {
		for _, e_pos := range valid_valves_keys[i+1:] {
			dfs(Stack{"AA", s_pos, "AA", e_pos, 0, 26, valid_valves.Clone()})
		}
	}

	return fmt.Sprint(finished_stack_max), nil
}
