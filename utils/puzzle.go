package utils

import (
	"fmt"
	"time"
)

type InputType uint8

const (
	Sample InputType = iota
	Input
)

type RunPuzzles uint8

const (
	Day1 RunPuzzles = 1 << iota
	Day2 RunPuzzles = 1 << iota
)

type PartPuzzle func() (string, error)
type ParseFn func() error
type FinalizeFileFn func() error

type IPuzzle interface {
	GetDay() uint

	GetFile() string
	GetSampleFilePath() string
	GetInputFilePath() string

	ParsePart1() error
	ParsePart2() error

	DoPart1() error
	DoPart2() error

	DoPuzzles() error
}

var Puzzles []IPuzzle

func AddPuzzle(p IPuzzle) {
	Puzzles = append(Puzzles, p)
}

func GetPuzzle(puzzle uint) IPuzzle {
	for _, p := range Puzzles {
		if p.GetDay() == puzzle {
			return p
		}
	}

	return nil
}

type Puzzle struct {
	IPuzzle

	Day       uint
	InputType InputType
	Puzzles   RunPuzzles

	Parse1Fn    ParseFn
	Part1Fn     PartPuzzle
	Finalize1Fn FinalizeFileFn

	Parse2Fn    ParseFn
	Part2Fn     PartPuzzle
	Finalize2Fn FinalizeFileFn
}

func (p *Puzzle) GetDay() uint {
	return p.Day
}

func (p *Puzzle) GetFile() string {
	switch p.InputType {
	case Sample:
		return p.GetSampleFilePath()

	default:
		return p.GetInputFilePath()
	}
}

func (p *Puzzle) GetSampleFilePath() string {
	return fmt.Sprintf("day%d/sample.txt", p.Day)
}

func (p *Puzzle) GetInputFilePath() string {
	return fmt.Sprintf("day%d/input.txt", p.Day)
}

func (p *Puzzle) ParsePart1() error {
	if err := p.Parse1Fn(); err != nil {
		return err
	}

	if p.Finalize1Fn != nil {
		return p.Finalize1Fn()
	} else {
		return nil
	}
}

func (p *Puzzle) ParsePart2() error {
	if err := p.Parse2Fn(); err != nil {
		return err
	}

	if p.Finalize2Fn != nil {
		return p.Finalize2Fn()
	} else {
		return nil
	}
}

func (p *Puzzle) DoPart1() error {
	if err := p.ParsePart1(); err != nil {
		return err
	}

	start := time.Now()
	res, err := p.Part1Fn()
	if err == nil {
		fmt.Printf("Day %d (%s), Part 1: %s\n", p.Day, time.Since(start), res)
		return nil
	} else {
		fmt.Printf("Day %d, Part 1 encountered error: \"%s\"\n", p.Day, err.Error())
		return nil
	}
}

func (p *Puzzle) DoPart2() error {
	if err := p.ParsePart2(); err != nil {
		return err
	}

	start := time.Now()
	res, err := p.Part2Fn()
	if err == nil {
		fmt.Printf("Day %d (%s), Part 2: %s\n", p.Day, time.Since(start), res)
		return nil
	} else {
		fmt.Printf("Day %d, Part 2 encountered error: \"%s\"\n", p.Day, err.Error())
		return nil
	}
}

func (p *Puzzle) DoPuzzles() error {
	if p.Puzzles&Day1 != 0 {
		if err := p.DoPart1(); err != nil {
			return err
		}
	}

	if p.Puzzles&Day2 != 0 {
		if err := p.DoPart2(); err != nil {
			return err
		}
	}

	return nil
}

func NewSingleParsePuzzle(parseFn ParseFunction, part1Fn PartPuzzle, part2Fn PartPuzzle, finalizeFn FinalizeFileFn) *Puzzle {
	p := &Puzzle{}

	initialized := false
	sharedParseFn := func() error {
		if initialized {
			return nil
		}

		start := time.Now()
		if err := ProcessByLine(p.GetFile(), parseFn); err != nil {
			return err
		} else {
			fmt.Printf("Parsed in %s\n", time.Since(start))
			initialized = true
			return nil
		}
	}

	p.Parse1Fn = sharedParseFn
	p.Part1Fn = part1Fn
	p.Finalize1Fn = finalizeFn

	p.Parse2Fn = sharedParseFn
	p.Part2Fn = part2Fn
	p.Finalize2Fn = finalizeFn

	return p
}

func NewDualParsePuzzle(parse1Fn ParseFunction, part1Fn PartPuzzle, finalize1Fn FinalizeFileFn, parse2Fn ParseFunction, part2Fn PartPuzzle, finalize2Fn FinalizeFileFn) *Puzzle {
	p := &Puzzle{}

	initialized1 := false
	p.Parse1Fn = func() error {
		if initialized1 {
			return nil
		}

		start := time.Now()
		if err := ProcessByLine(p.GetFile(), parse1Fn); err != nil {
			return err
		} else {
			fmt.Printf("Parsed Part 1 in %s\n", time.Since(start))
			initialized1 = true
			return nil
		}
	}
	p.Part1Fn = part1Fn
	p.Finalize1Fn = finalize1Fn

	initialized2 := false
	p.Parse2Fn = func() error {
		if initialized2 {
			return nil
		}

		start := time.Now()
		if err := ProcessByLine(p.GetFile(), parse2Fn); err != nil {
			return err
		} else {
			fmt.Printf("Parsed Part 2 in %s\n", time.Since(start))
			initialized2 = true
			return nil
		}
	}
	p.Part2Fn = part2Fn
	p.Finalize2Fn = finalize2Fn

	return p
}
