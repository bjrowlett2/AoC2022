package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bjrowlett2/AoC2022/internal/aoc"
)

func main() {
	var err error

	var problem1 *Problem
	if problem1, err = Load(); err != nil {
		log.Fatal(err)
	}

	if err = problem1.SolvePart1(); err != nil {
		log.Fatal(err)
	}

	var problem2 *Problem
	if problem2, err = Load(); err != nil {
		log.Fatal(err)
	}

	if err = problem2.SolvePart2(); err != nil {
		log.Fatal(err)
	}
}

type Step struct {
	To    int
	From  int
	Count int
}

type Problem struct {
	Steps  []Step
	Stacks []aoc.Stack[rune]
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_05.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Steps:  make([]Step, 0),
		Stacks: nil, // Created below.
	}

	header := true
	headerLines := make(aoc.Stack[string], 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if header {
			if line != "" {
				headerLines = aoc.Push(headerLines, line)
			} else {
				header = false

				var row string
				headerLines, row = aoc.Pop(headerLines)

				n := len(row) - strings.Count(row, " ")
				problem.Stacks = make([]aoc.Stack[rune], n)

				for len(headerLines) != 0 {
					headerLines, row = aoc.Pop(headerLines)
					for i, k := 0, 0; i < len(row); i, k = i+4, k+1 {
						c := strings.TrimSpace(row[i : i+3])

						if c != "" {
							r := rune(c[1])
							problem.Stacks[k] = aoc.Push(problem.Stacks[k], r)
						}
					}
				}
			}
		} else if strings.HasPrefix(line, "move") {
			var to, from, count int
			if _, err = fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to); err != nil {
				return nil, err
			}

			step := Step{
				To:    to - 1,
				From:  from - 1,
				Count: count,
			}

			problem.Steps = append(problem.Steps, step)
		}
	}

	return &problem, nil
}

func (problem *Problem) SolvePart1() error {
	for _, step := range problem.Steps {
		var r rune
		for k := 0; k < step.Count; k += 1 {
			problem.Stacks[step.From], r = aoc.Pop(problem.Stacks[step.From])
			problem.Stacks[step.To] = aoc.Push(problem.Stacks[step.To], r)
		}
	}

	top := ""
	for _, s := range problem.Stacks {
		top += string(aoc.Peek(s))
	}

	fmt.Printf("Part 1: %s\n", top)
	return nil
}

func (problem *Problem) SolvePart2() error {
	for _, step := range problem.Steps {
		var r rune
		t := make(aoc.Stack[rune], 0)
		for k := 0; k < step.Count; k += 1 {
			problem.Stacks[step.From], r = aoc.Pop(problem.Stacks[step.From])
			t = aoc.Push(t, r)
		}

		for k := 0; k < step.Count; k += 1 {
			t, r = aoc.Pop(t)
			problem.Stacks[step.To] = aoc.Push(problem.Stacks[step.To], r)
		}

	}

	top := ""
	for _, s := range problem.Stacks {
		top += string(aoc.Peek(s))
	}

	fmt.Printf("Part 2: %s\n", top)
	return nil
}
