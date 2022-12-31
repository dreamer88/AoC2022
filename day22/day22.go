package day22

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

/*
Input format:
 01
 2
34
5
*/

type Space uint8

const (
	EMPTY Space = iota
	WALKABLE
	WALL
)

type Facing int8

const (
	RIGHT Facing = iota
	DOWN
	LEFT
	UP
	COUNT
)

type Turn int8

const (
	STRAIGHT   Turn = 0
	RIGHT_TURN Turn = 1
	LEFT_TURN  Turn = -1
)

type MovementType uint8

const (
	TURN MovementType = iota
	MOVE
)

type Movement struct {
	moveType MovementType
	turn     Turn
	length   int
}

type Input struct {
	start_point, end_point utils.Point2
	positions              map[utils.Point2]Space
	movements              []Movement
}

type InputPortion uint8

const (
	MAP InputPortion = iota
	MOVEMENTS
)

type CubeFace uint8

const (
	TOP_FACE CubeFace = iota
	RIGHT_FACE
	FRONT_FACE
	BOTTOM_FACE
	LEFT_FACE
	BACK_FACE
)

type Part2MovFn = func(CubeFace, utils.Point2) (Facing, CubeFace, utils.Point2)

var input Input = Input{utils.Point2{X: 0, Y: 0}, utils.Point2{X: 0, Y: 0}, map[utils.Point2]Space{}, []Movement{}}
var curPortion InputPortion = MAP
var curRow int = 0
var inputType utils.InputType = utils.Input

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 22
	p.InputType = inputType
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func getCubeDimension() int {
	if inputType == utils.Sample {
		return 4
	} else {
		return 50
	}
}

func (i Input) getPoint(p utils.Point2) Space {
	if v, found := i.positions[p]; found {
		return v
	} else {
		return EMPTY
	}
}

func getPointMappedToCube(cube_face CubeFace, p utils.Point2) utils.Point2 {
	switch cube_face {
	case FRONT_FACE:
		return utils.Point2{X: p.X + getCubeDimension(), Y: p.Y}

	case RIGHT_FACE:
		return utils.Point2{X: p.X + 2*getCubeDimension(), Y: p.Y}

	case BOTTOM_FACE:
		return utils.Point2{X: p.X + getCubeDimension(), Y: p.Y + getCubeDimension()}

	case BACK_FACE:
		return utils.Point2{X: p.X + getCubeDimension(), Y: p.Y + 2*getCubeDimension()}

	case LEFT_FACE:
		return utils.Point2{X: p.X, Y: p.Y + 2*getCubeDimension()}

	case TOP_FACE:
		return utils.Point2{X: p.X, Y: p.Y + 3*getCubeDimension()}
	}

	panic("shouldn't get here!")
}

func (i Input) getPointOnCube(cube_face CubeFace, p utils.Point2) Space {
	mapped_p := getPointMappedToCube(cube_face, p)
	return i.getPoint(mapped_p)
}

func (i Input) clone() Input {
	new_map := map[utils.Point2]Space{}
	for row, v := range i.positions {
		new_map[row] = v
	}

	return Input{
		i.start_point,
		i.end_point,
		new_map,
		utils.Copy1D(i.movements),
	}
}

func parse(val string) error {
	if curPortion == MAP {
		curRow++
		input.end_point.Y = curRow

		if len(val) == 0 {
			curPortion = MOVEMENTS
		} else {
			input.end_point.X = utils.Max(input.end_point.X, len(val)+1)

			for i := 0; i < len(val); i++ {
				p := utils.Point2{X: i + 1, Y: curRow}
				char := val[i]
				col := i + 1
				if char == '.' {
					if curRow == 1 && input.start_point.Y == 0 {
						input.start_point = utils.Point2{X: col, Y: curRow}
					}

					input.positions[p] = WALKABLE
				} else if char == '#' {
					input.positions[p] = WALL
				}
			}
		}
	} else {
		length := ""
		process_length := func() {
			l, err := strconv.Atoi(length)
			if err != nil {
				panic(err)
			}
			input.movements = append(input.movements, Movement{MOVE, STRAIGHT, l})
			length = ""
		}
		for i := 0; i < len(val); i++ {
			c := val[i]
			if c == 'R' {
				process_length()
				input.movements = append(input.movements, Movement{TURN, RIGHT_TURN, 0})
			} else if c == 'L' {
				process_length()
				input.movements = append(input.movements, Movement{TURN, LEFT_TURN, 0})
			} else {
				length += string(c)
			}
		}
		process_length()
	}

	return nil
}

func part1() (string, error) {
	i := input.clone()

	cur_p := i.start_point
	face := RIGHT

	do_move := func(move_fn func(utils.Point2) utils.Point2, reset_fn func(utils.Point2) utils.Point2) Space {
		next_p := move_fn(cur_p)

		switch i.getPoint(next_p) {
		case WALKABLE:
			cur_p = next_p
			return WALKABLE

		case WALL:
			return WALL

		case EMPTY:
			next_p = reset_fn(next_p)
			v := EMPTY

			for ; v == EMPTY && cur_p != next_p; v = i.getPoint(next_p) {
				next_p = move_fn(next_p)
			}

			if v == WALKABLE {
				cur_p = next_p
			}

			return v
		}

		return EMPTY
	}

	for _, m := range i.movements {
		switch m.moveType {
		case TURN:
			face = Facing(utils.SaneMod(int(m.turn)+int(face), int(COUNT)))

		case MOVE:
			switch face {
			case RIGHT:
				for l := 0; l < m.length; l++ {
					if do_move(func(p utils.Point2) utils.Point2 {
						return p.Right()
					}, func(p utils.Point2) utils.Point2 {
						return utils.Point2{X: 0, Y: p.Y}
					}) == WALL {
						break
					}
				}
			case LEFT:
				for l := 0; l < m.length; l++ {
					if do_move(func(p utils.Point2) utils.Point2 {
						return p.Left()
					}, func(p utils.Point2) utils.Point2 {
						return utils.Point2{X: i.end_point.X + 1, Y: p.Y}
					}) == WALL {
						break
					}
				}
			case DOWN:
				for l := 0; l < m.length; l++ {
					if do_move(func(p utils.Point2) utils.Point2 {
						return p.Down()
					}, func(p utils.Point2) utils.Point2 {
						return utils.Point2{X: p.X, Y: 0}
					}) == WALL {
						break
					}
				}
			case UP:
				for l := 0; l < m.length; l++ {
					if do_move(func(p utils.Point2) utils.Point2 {
						return p.Up()
					}, func(p utils.Point2) utils.Point2 {
						return utils.Point2{X: p.X, Y: i.end_point.Y + 1}
					}) == WALL {
						break
					}
				}
			}
		}
	}

	return fmt.Sprint(1000*cur_p.Y + 4*cur_p.X + int(face)), nil
}

func part2() (string, error) {
	i := input.clone()

	cur_p, cur_cube_face := utils.Point2{X: 1, Y: 1}, FRONT_FACE
	face := RIGHT

	do_move := func(move_fn Part2MovFn) Space {
		next_face, next_cube_face, next_p := move_fn(cur_cube_face, cur_p)

		switch i.getPointOnCube(next_cube_face, next_p) {
		case WALKABLE:
			face, cur_cube_face, cur_p = next_face, next_cube_face, next_p
			return WALKABLE

		case WALL:
			return WALL
		}

		panic("shouldn't get here in part 2!")
	}

	for _, m := range i.movements {
		switch m.moveType {
		case TURN:
			face = Facing(utils.SaneMod(int(m.turn)+int(face), int(COUNT)))

		case MOVE:
		hit_wall:
			for l := 0; l < m.length; l++ {
				switch face {
				case RIGHT:
					if do_move(func(face CubeFace, p utils.Point2) (Facing, CubeFace, utils.Point2) {
						if p.X == getCubeDimension() {
							switch face {
							case FRONT_FACE:
								return RIGHT, RIGHT_FACE, utils.Point2{X: 1, Y: p.Y}
							case RIGHT_FACE:
								return LEFT, BACK_FACE, utils.Point2{X: getCubeDimension(), Y: getCubeDimension() - p.Y + 1}
							case BOTTOM_FACE:
								return UP, RIGHT_FACE, utils.Point2{X: p.Y, Y: getCubeDimension()}
							case BACK_FACE:
								return LEFT, RIGHT_FACE, utils.Point2{X: getCubeDimension(), Y: getCubeDimension() - p.Y + 1}
							case LEFT_FACE:
								return RIGHT, BACK_FACE, utils.Point2{X: 1, Y: p.Y}
							case TOP_FACE:
								return UP, BACK_FACE, utils.Point2{X: p.Y, Y: getCubeDimension()}
							}
						}
						return RIGHT, face, p.Right()
					}) == WALL {
						break hit_wall
					}

				case LEFT:
					if do_move(func(face CubeFace, p utils.Point2) (Facing, CubeFace, utils.Point2) {
						if p.X == 1 {
							switch face {
							case FRONT_FACE:
								return RIGHT, LEFT_FACE, utils.Point2{X: 1, Y: getCubeDimension() - p.Y + 1}
							case RIGHT_FACE:
								return LEFT, FRONT_FACE, utils.Point2{X: getCubeDimension(), Y: p.Y}
							case BOTTOM_FACE:
								return DOWN, LEFT_FACE, utils.Point2{X: p.Y, Y: 1}
							case BACK_FACE:
								return LEFT, LEFT_FACE, utils.Point2{X: getCubeDimension(), Y: p.Y}
							case LEFT_FACE:
								return RIGHT, FRONT_FACE, utils.Point2{X: 1, Y: getCubeDimension() - p.Y + 1}
							case TOP_FACE:
								return DOWN, FRONT_FACE, utils.Point2{X: p.Y, Y: 1}
							}
						}
						return LEFT, face, p.Left()
					}) == WALL {
						break hit_wall
					}

				case DOWN:
					if do_move(func(face CubeFace, p utils.Point2) (Facing, CubeFace, utils.Point2) {
						if p.Y == getCubeDimension() {
							switch face {
							case FRONT_FACE:
								return DOWN, BOTTOM_FACE, utils.Point2{X: p.X, Y: 1}
							case RIGHT_FACE:
								return LEFT, BOTTOM_FACE, utils.Point2{X: getCubeDimension(), Y: p.X}
							case BOTTOM_FACE:
								return DOWN, BACK_FACE, utils.Point2{X: p.X, Y: 1}
							case BACK_FACE:
								return LEFT, TOP_FACE, utils.Point2{X: getCubeDimension(), Y: p.X}
							case LEFT_FACE:
								return DOWN, TOP_FACE, utils.Point2{X: p.X, Y: 1}
							case TOP_FACE:
								return DOWN, RIGHT_FACE, utils.Point2{X: p.X, Y: 1}
							}
						}
						return DOWN, face, p.Down()
					}) == WALL {
						break hit_wall
					}

				case UP:
					if do_move(func(face CubeFace, p utils.Point2) (Facing, CubeFace, utils.Point2) {
						if p.Y == 1 {
							switch face {
							case FRONT_FACE:
								return RIGHT, TOP_FACE, utils.Point2{X: 1, Y: p.X}
							case RIGHT_FACE:
								return UP, TOP_FACE, utils.Point2{X: p.X, Y: getCubeDimension()}
							case BOTTOM_FACE:
								return UP, FRONT_FACE, utils.Point2{X: p.X, Y: getCubeDimension()}
							case BACK_FACE:
								return UP, BOTTOM_FACE, utils.Point2{X: p.X, Y: getCubeDimension()}
							case LEFT_FACE:
								return RIGHT, BOTTOM_FACE, utils.Point2{X: 1, Y: p.X}
							case TOP_FACE:
								return UP, LEFT_FACE, utils.Point2{X: p.X, Y: getCubeDimension()}
							}
						}
						return UP, face, p.Up()
					}) == WALL {
						break hit_wall
					}
				}
			}
		}
	}

	final_p := getPointMappedToCube(cur_cube_face, cur_p)
	return fmt.Sprint(1000*final_p.Y + 4*final_p.X + int(face)), nil
}
