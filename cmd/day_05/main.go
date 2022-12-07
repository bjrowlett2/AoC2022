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
	Stacks []aoc.Stack[byte]
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
				headerLines.Push(line)
			} else {
				header = false

				var row string
				headerLines.Pop(&row)

				n := len(row) - strings.Count(row, " ")
				problem.Stacks = make([]aoc.Stack[byte], n)

				for headerLines.Pop(&row) {
					for i, k := 0, 0; i < len(row); i, k = i+4, k+1 {
						c := strings.TrimSpace(row[i : i+3])

						if c != "" {
							b := byte(c[1])
							problem.Stacks[k].Push(b)
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
	var b byte
	for _, step := range problem.Steps {
		for k := 0; k < step.Count; k += 1 {
			if problem.Stacks[step.From].Pop(&b) {
				problem.Stacks[step.To].Push(b)
			}
		}
	}

	top := ""
	for _, s := range problem.Stacks {
		if s.Peek(&b) {
			top += string(b)
		}
	}

	fmt.Printf("Part 1: %s\n", top)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var b byte
	for _, step := range problem.Steps {
		temp := make(aoc.Stack[byte], 0)
		for k := 0; k < step.Count; k += 1 {
			if problem.Stacks[step.From].Pop(&b) {
				temp.Push(b)
			}
		}

		for k := 0; k < step.Count; k += 1 {
			if temp.Pop(&b) {
				problem.Stacks[step.To].Push(b)
			}
		}
	}

	top := ""
	for _, s := range problem.Stacks {
		if s.Peek(&b) {
			top += string(b)
		}
	}

	fmt.Printf("Part 2: %s\n", top)
	return nil
}
