package day7

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
)

var down_dir_regex = regexp.MustCompile(`\$ cd (\w+)`)
var up_dir_regex = regexp.MustCompile(`\$ cd \.\.`)
var root_dir_regex = regexp.MustCompile(`\$ cd /`)
var ls_regex = regexp.MustCompile(`\$ ls`)
var file_regex = regexp.MustCompile(`(\d+) ([\w\.]+)`)
var dir_regex = regexp.MustCompile(`dir (\w+)`)

var input []string

func init() {
	p := utils.NewSingleParsePuzzle(parse, part1, part2, nil)

	p.Day = 7
	p.InputType = utils.Input
	p.Puzzles = utils.Day1 | utils.Day2

	utils.AddPuzzle(p)
}

func parse(val string) error {
	input = append(input, val)
	return nil
}

type File struct {
	size int
}

type Directory struct {
	name        string
	directories map[string]Directory
	files       map[string]File
}

func (d Directory) GetDirectoriesBelowSize(test_size int) (int, int) {
	size := 0
	sub_dirs_below_size := 0
	for _, sub_dir := range d.directories {
		sub_size, below_size := sub_dir.GetDirectoriesBelowSize(test_size)
		size += sub_size
		sub_dirs_below_size += below_size
	}

	for _, file := range d.files {
		size += file.size
	}

	if size <= test_size {
		sub_dirs_below_size += size
	}
	return size, sub_dirs_below_size
}

func (d Directory) GetDirSize() int {
	size := 0
	for _, sub_dir := range d.directories {
		size += sub_dir.GetDirSize()
	}

	for _, file := range d.files {
		size += file.size
	}
	return size
}

func part1() (string, error) {
	cmds := utils.Copy1D(input)

	rootDirectory := Directory{"/", make(map[string]Directory), make(map[string]File)}
	directories := []Directory{rootDirectory}

	for _, line := range cmds {
		if root_dir_regex.MatchString(line) {
			directories = []Directory{rootDirectory}
		} else if up_dir_regex.MatchString(line) {
			directories = directories[:len(directories)-1]
		} else if down_dir := down_dir_regex.FindStringSubmatch(line); down_dir != nil {
			dir_name := down_dir[1]
			current := directories[len(directories)-1]
			if v, found := current.directories[dir_name]; found {
				directories = append(directories, v)
			} else {
				newDir := Directory{dir_name, make(map[string]Directory), make(map[string]File)}
				current.directories[dir_name] = newDir
				directories = append(directories, newDir)
			}
		} else if ls_regex.MatchString(line) {
			// do nothing lul
		} else if file_match := file_regex.FindStringSubmatch(line); file_match != nil {
			file_size, file_err := strconv.Atoi(file_match[1])
			if file_err != nil {
				return "", file_err
			}

			file_name := file_match[2]
			current := directories[len(directories)-1]
			current.files[file_name] = File{size: file_size}
		} else if dir_match := dir_regex.FindStringSubmatch(line); dir_match != nil {
			// do nothing lul
		} else {
			return "", fmt.Errorf("unhandled line %s", line)
		}
	}

	_, sub_dirs_below_size := rootDirectory.GetDirectoriesBelowSize(100000)
	return fmt.Sprint(sub_dirs_below_size), nil
}

func part2() (string, error) {
	cmds := utils.Copy1D(input)

	rootDirectory := Directory{"/", make(map[string]Directory), make(map[string]File)}
	directories := []Directory{rootDirectory}
	var all_directories []Directory

	for _, line := range cmds {
		if root_dir_regex.MatchString(line) {
			directories = []Directory{rootDirectory}
		} else if up_dir_regex.MatchString(line) {
			directories = directories[:len(directories)-1]
		} else if down_dir := down_dir_regex.FindStringSubmatch(line); down_dir != nil {
			dir_name := down_dir[1]
			current := directories[len(directories)-1]
			if v, found := current.directories[dir_name]; found {
				directories = append(directories, v)
			} else {
				newDir := Directory{dir_name, make(map[string]Directory), make(map[string]File)}
				all_directories = append(all_directories, newDir)
				current.directories[dir_name] = newDir
				directories = append(directories, newDir)
			}
		} else if ls_regex.MatchString(line) {
			// do nothing lul
		} else if file_match := file_regex.FindStringSubmatch(line); file_match != nil {
			file_size, file_err := strconv.Atoi(file_match[1])
			if file_err != nil {
				return "", file_err
			}

			file_name := file_match[2]
			current := directories[len(directories)-1]
			current.files[file_name] = File{size: file_size}
		} else if dir_match := dir_regex.FindStringSubmatch(line); dir_match != nil {
			dir_name := dir_match[1]
			current := directories[len(directories)-1]
			if v, found := current.directories[dir_name]; found {
				directories = append(directories, v)
			} else {
				newDir := Directory{dir_name, make(map[string]Directory), make(map[string]File)}
				all_directories = append(all_directories, newDir)
				current.directories[dir_name] = newDir
			}
		} else {
			return "", fmt.Errorf("unhandled line %s", line)
		}
	}

	best_size := 70000000
	curr_size := rootDirectory.GetDirSize()
	required_size := curr_size - 40000000
	for _, dir := range all_directories {
		size := dir.GetDirSize()
		if size >= required_size && size < best_size {
			best_size = size
		}
	}

	return fmt.Sprint(best_size), nil
}
