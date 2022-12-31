package main

import (
	_ "aoc2022/day1"
	_ "aoc2022/day10"
	_ "aoc2022/day11"
	_ "aoc2022/day12"
	_ "aoc2022/day13"
	_ "aoc2022/day14"
	_ "aoc2022/day15"
	_ "aoc2022/day16"
	_ "aoc2022/day17"
	_ "aoc2022/day18"
	_ "aoc2022/day19"
	_ "aoc2022/day2"
	_ "aoc2022/day20"
	_ "aoc2022/day21"
	_ "aoc2022/day22"
	_ "aoc2022/day23"
	_ "aoc2022/day24"
	_ "aoc2022/day25"
	_ "aoc2022/day3"
	_ "aoc2022/day4"
	_ "aoc2022/day5"
	_ "aoc2022/day6"
	_ "aoc2022/day7"
	_ "aoc2022/day8"
	_ "aoc2022/day9"
	"aoc2022/utils"
	"fmt"
)

func main() {
	for puzzle := uint(1); puzzle <= 25; puzzle++ {
		var p = utils.GetPuzzle(puzzle)
		if p != nil {
			err := p.DoPuzzles()
			if err != nil {
				fmt.Print(err.Error())
			}
			println("##########################")
		}
	}
}
