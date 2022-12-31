package day21

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Op uint8

const (
	VAL Op = iota
	ADD Op = iota
	SUB Op = iota
	MUL Op = iota
	DIV Op = iota

	UNKNOWN Op = iota
	EQUAL   Op = iota
)

type Monkey struct {
	op                        Op
	val                       int
	first_other, second_other string
}

var val_monkey = regexp.MustCompile(`(\w+): (-?\d+)`)
var op_monkey = regexp.MustCompile(`(\w+): (\w+) ([/\*\-\+]) (\w+)`)

var input map[string]Monkey = make(map[string]Monkey)

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 21
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	if match := val_monkey.FindStringSubmatch(val); match != nil {
		name := match[1]
		val, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}

		input[name] = Monkey{VAL, val, "", ""}
	} else if match := op_monkey.FindStringSubmatch(val); match != nil {
		name := match[1]
		first_other := match[2]
		second_other := match[4]

		switch match[3] {
		case "+":
			input[name] = Monkey{ADD, 0, first_other, second_other}

		case "-":
			input[name] = Monkey{SUB, 0, first_other, second_other}

		case "*":
			input[name] = Monkey{MUL, 0, first_other, second_other}

		case "/":
			input[name] = Monkey{DIV, 0, first_other, second_other}

		default:
			return fmt.Errorf("cannot process operator \"%s\" in \"%s\"", match[3], val)
		}
	} else {
		return fmt.Errorf("could not parse line \"%s\"", val)
	}

	return nil
}

func copy_monkeys() map[string]Monkey {
	var res map[string]Monkey = make(map[string]Monkey)
	for key := range input {
		res[key] = input[key]
	}
	return res
}

func get_monkey_value(monkey string, monkeys *map[string]Monkey) (int, error) {
	if v, found := (*monkeys)[monkey]; !found {
		return 0, fmt.Errorf("could not find monkey %s", monkey)
	} else {
		if v.op == VAL {
			return v.val, nil
		} else {
			first_val, err := get_monkey_value(v.first_other, monkeys)
			if err != nil {
				return 0, err
			}

			second_val, err := get_monkey_value(v.second_other, monkeys)
			if err != nil {
				return 0, err
			}

			new_val := 0
			switch v.op {
			case VAL:
				return v.val, nil

			case ADD:
				new_val = first_val + second_val

			case SUB:
				new_val = first_val - second_val

			case MUL:
				new_val = first_val * second_val

			case DIV:
				new_val = first_val / second_val

			default:
				return 0, fmt.Errorf("invalid run op %d", v.op)
			}

			(*monkeys)[monkey] = Monkey{VAL, new_val, "", ""}
			return new_val, nil
		}
	}
}

func part1() (string, error) {
	monkeys := copy_monkeys()

	result, err := get_monkey_value("root", &monkeys)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(result), nil
}

type ResultPair struct {
	left, right Result
}

type Result struct {
	op          Op
	v           int
	result      *ResultPair
	has_unknown bool
}

func get_monkey_value_two(monkey string, monkeys *map[string]Monkey) (Result, error) {
	if v, found := (*monkeys)[monkey]; !found {
		return Result{}, fmt.Errorf("could not find monkey %s", monkey)
	} else {
		if v.op == VAL {
			return Result{VAL, v.val, &ResultPair{}, false}, nil
		} else if v.op == UNKNOWN {
			return Result{UNKNOWN, 0, &ResultPair{}, true}, nil
		} else {
			first_val, err := get_monkey_value_two(v.first_other, monkeys)
			if err != nil {
				return Result{}, err
			}

			second_val, err := get_monkey_value_two(v.second_other, monkeys)
			if err != nil {
				return Result{}, err
			}

			new_val := 0

			if first_val.has_unknown || second_val.has_unknown {
				return Result{v.op, 0, &ResultPair{first_val, second_val}, true}, nil
			}

			switch v.op {
			case VAL:
				return Result{VAL, v.val, &ResultPair{}, false}, nil

			case ADD:
				new_val = first_val.v + second_val.v

			case SUB:
				new_val = first_val.v - second_val.v

			case MUL:
				new_val = first_val.v * second_val.v

			case DIV:
				new_val = first_val.v / second_val.v

			default:
				return Result{}, fmt.Errorf("invalid run op %d", v.op)
			}

			(*monkeys)[monkey] = Monkey{VAL, new_val, "", ""}
			return Result{VAL, new_val, &ResultPair{}, false}, nil
		}
	}
}

func resolve_unknown(r Result, passed_value int) (int, error) {
	if r.result.left.has_unknown {
		switch r.op {
		case UNKNOWN:
			return passed_value, nil

		case EQUAL:
			passed_value = r.result.right.v

		case ADD:
			passed_value -= r.result.right.v

		case SUB:
			passed_value += r.result.right.v

		case MUL:
			passed_value /= r.result.right.v

		case DIV:
			passed_value *= r.result.right.v

		default:
			return 0, fmt.Errorf("unhandled op %d", r.op)
		}

		return resolve_unknown(r.result.left, passed_value)
	} else {
		switch r.op {
		case UNKNOWN:
			return passed_value, nil

		case EQUAL:
			passed_value = r.result.left.v

		case ADD:
			passed_value -= r.result.left.v

		case SUB:
			passed_value = r.result.left.v - passed_value

		case MUL:
			passed_value /= r.result.left.v

		case DIV:
			passed_value = r.result.left.v / passed_value

		default:
			return 0, fmt.Errorf("unhandled op %d", r.op)
		}

		return resolve_unknown(r.result.right, passed_value)
	}
}

func part2() (string, error) {
	monkeys := copy_monkeys()

	r := monkeys["root"]
	monkeys["root"] = Monkey{EQUAL, 0, r.first_other, r.second_other}

	monkeys["humn"] = Monkey{UNKNOWN, 0, "", ""}

	result, err := get_monkey_value_two("root", &monkeys)
	if err != nil {
		return "", err
	}

	val, err := resolve_unknown(result, 0)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(val), nil
}
