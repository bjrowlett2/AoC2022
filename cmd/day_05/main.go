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

type Rearrangement struct {
	To     int64
	From   int64
	Crates int64
}

type Problem struct {
	Stacks    []aoc.Stack[rune]
	Procedure []Rearrangement
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
		Stacks:    nil, // Created below.
		Procedure: make([]Rearrangement, 0),
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
			var to, from, crates int64
			if _, err = fmt.Sscanf(line, "move %d from %d to %d", &crates, &from, &to); err != nil {
				return nil, err
			}

			rearrangement := Rearrangement{
				To:     to - 1,
				From:   from - 1,
				Crates: crates,
			}

			problem.Procedure = append(problem.Procedure, rearrangement)
		}
	}

	return &problem, nil
}

func (problem *Problem) SolvePart1() error {
	for _, rearrangement := range problem.Procedure {
		var r rune
		for k := 0; k < int(rearrangement.Crates); k += 1 {
			problem.Stacks[rearrangement.From], r = aoc.Pop(problem.Stacks[rearrangement.From])
			problem.Stacks[rearrangement.To] = aoc.Push(problem.Stacks[rearrangement.To], r)
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
	for _, rearrangement := range problem.Procedure {
		var r rune
		t := make(aoc.Stack[rune], 0)
		for k := 0; k < int(rearrangement.Crates); k += 1 {
			problem.Stacks[rearrangement.From], r = aoc.Pop(problem.Stacks[rearrangement.From])
			t = aoc.Push(t, r)
		}

		for k := 0; k < int(rearrangement.Crates); k += 1 {
			t, r = aoc.Pop(t)
			problem.Stacks[rearrangement.To] = aoc.Push(problem.Stacks[rearrangement.To], r)
		}

	}

	top := ""
	for _, s := range problem.Stacks {
		top += string(aoc.Peek(s))
	}

	fmt.Printf("Part 1: %s\n", top)
	return nil
}
