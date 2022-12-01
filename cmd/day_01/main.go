package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/bjrowlett2/AoC2022/internal/aoc"
)

func main() {
	var err error

	var problem *Problem
	if problem, err = Load(); err != nil {
		log.Fatal(err)
	}

	aoc.SortInt64Desc(problem.Calories)

	if err = problem.SolvePart1(); err != nil {
		log.Fatal(err)
	}

	if err = problem.SolvePart2(); err != nil {
		log.Fatal(err)
	}
}

type Problem struct {
	Calories []int64
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_01.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Calories: make([]int64, 0),
	}

	var total int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			problem.Calories = append(problem.Calories, total)
			total = 0
		} else {
			var calories int64
			if _, err = fmt.Sscanf(line, "%d", &calories); err != nil {
				return nil, err
			}

			total += calories
		}
	}

	return &problem, nil
}

func (problem *Problem) SolvePart1() error {
	fmt.Printf("Part 1: %d\n", problem.Calories[0])
	return nil
}

func (problem *Problem) SolvePart2() error {
	fmt.Printf("Part 2: %d\n", problem.Calories[0]+problem.Calories[1]+problem.Calories[2])
	return nil
}
