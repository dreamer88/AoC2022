package day11

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Op uint8

const (
	ADD Op = iota
	SUB Op = iota
	DIV Op = iota
	MUL Op = iota
)

type OpValue int

const OldVal OpValue = math.MinInt

type Monkey struct {
	items        []int
	inspections  int
	operation    Op
	leftOpValue  OpValue
	rightOpValue OpValue
	testDivisor  int
	trueIndex    int
	falseIndex   int
}

var input []Monkey

var monkey_reg = regexp.MustCompile(`Monkey \d+`)
var items_reg = regexp.MustCompile(`\s*Starting items: ([\d ,]+)`)
var op_reg = regexp.MustCompile(`\s*Operation: new = ([\w\d]+) ([\+\-\*\/]) ([\w\d]+)`)
var test_reg = regexp.MustCompile(`\s*Test: divisible by (\d+)`)
var trueCond_reg = regexp.MustCompile(`\s*If true: throw to monkey (\d+)`)
var falseCond_reg = regexp.MustCompile(`\s*If false: throw to monkey (\d+)`)

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 11
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func getOpVal(val string) (OpValue, error) {
	if val == "old" {
		return OldVal, nil
	} else {
		v, err := strconv.Atoi(val)
		return OpValue(v), err
	}
}

func parse(val string) error {
	var last_monkey *Monkey = nil
	if len(input) > 0 {
		last_monkey = &input[len(input)-1]
	}
	if monkey_reg.MatchString(val) {
		input = append(input, Monkey{})
	} else if match := items_reg.FindStringSubmatch(val); match != nil {
		items := strings.Split(match[1], ", ")
		for _, v := range items {
			iv, err := strconv.Atoi(v)
			if err != nil {
				return err
			} else {
				last_monkey.items = append(last_monkey.items, iv)
			}
		}
	} else if match := op_reg.FindStringSubmatch(val); match != nil {
		var err error
		last_monkey.leftOpValue, err = getOpVal(match[1])
		if err != nil {
			return err
		}

		last_monkey.rightOpValue, err = getOpVal(match[3])
		if err != nil {
			return err
		}

		switch match[2] {
		case "+":
			last_monkey.operation = ADD
		case "-":
			last_monkey.operation = SUB
		case "*":
			last_monkey.operation = MUL
		case "/":
			last_monkey.operation = DIV

		default:
			return fmt.Errorf("unhandled op %s - %s", match[2], val)
		}
	} else if match := test_reg.FindStringSubmatch(val); match != nil {
		iv, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		} else {
			last_monkey.testDivisor = iv
		}
	} else if match := trueCond_reg.FindStringSubmatch(val); match != nil {
		iv, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		} else {
			last_monkey.trueIndex = iv
		}
	} else if match := falseCond_reg.FindStringSubmatch(val); match != nil {
		iv, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		} else {
			last_monkey.falseIndex = iv
		}
	}

	return nil
}

func doOperation(monkey Monkey, val int) int {
	left_value := monkey.leftOpValue
	if left_value == OldVal {
		left_value = OpValue(val)
	}
	right_value := monkey.rightOpValue
	if right_value == OldVal {
		right_value = OpValue(val)
	}

	switch monkey.operation {
	case ADD:
		return int(left_value) + int(right_value)
	case SUB:
		return int(left_value) - int(right_value)
	case MUL:
		return int(left_value) * int(right_value)
	case DIV:
		return int(left_value) / int(right_value)

	default:
		return 0
	}
}

func doRounds1(monkeys *[]Monkey, rounds int) {
	for r := 1; r <= rounds; r++ {
		for m := 0; m < len(*monkeys); m++ {
			monkey := &(*monkeys)[m]
			for _, item := range monkey.items {
				monkey.inspections++

				item = doOperation(*monkey, item)
				item = int(math.Floor(float64(item) / 3.0))

				if item%monkey.testDivisor == 0 {
					(*monkeys)[monkey.trueIndex].items = append((*monkeys)[monkey.trueIndex].items, item)
				} else {
					(*monkeys)[monkey.falseIndex].items = append((*monkeys)[monkey.falseIndex].items, item)
				}
			}
			monkey.items = make([]int, 0)
		}
	}
}

func part1() (string, error) {
	monkeys := utils.Copy1D(input)
	doRounds1(&monkeys, 20)
	sort.Slice(monkeys, func(i int, j int) bool { return monkeys[i].inspections > monkeys[j].inspections })
	return fmt.Sprint((monkeys[0].inspections * monkeys[1].inspections)), nil
}

func doRounds2(monkeys *[]Monkey, rounds int) {
	fat_modulo := 1
	for _, m := range *monkeys {
		fat_modulo *= m.testDivisor
	}

	for r := 1; r <= rounds; r++ {
		for m := 0; m < len(*monkeys); m++ {
			monkey := &(*monkeys)[m]
			for _, item := range monkey.items {
				monkey.inspections++

				item = doOperation(*monkey, item)
				item = item % fat_modulo

				if item%monkey.testDivisor == 0 {
					(*monkeys)[monkey.trueIndex].items = append((*monkeys)[monkey.trueIndex].items, item)
				} else {
					(*monkeys)[monkey.falseIndex].items = append((*monkeys)[monkey.falseIndex].items, item)
				}
			}
			monkey.items = make([]int, 0)
		}
	}
}

func part2() (string, error) {
	monkeys := utils.Copy1D(input)
	doRounds2(&monkeys, 10000)
	sort.Slice(monkeys, func(i int, j int) bool { return monkeys[i].inspections > monkeys[j].inspections })
	return fmt.Sprint((monkeys[0].inspections * monkeys[1].inspections)), nil
}
