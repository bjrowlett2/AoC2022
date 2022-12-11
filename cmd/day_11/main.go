package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

type Monkey struct {
	Items       []int64
	Operand     string
	Operator    string
	DivisibleBy int64
	IfTrue      int64
	IfFalse     int64
}

type Problem struct {
	Monkeys []Monkey
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_11.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	var idx int
	problem := Problem{
		Monkeys: make([]Monkey, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Monkey ") {
			problem.Monkeys = append(problem.Monkeys, Monkey{})
			if _, err = fmt.Sscanf(line, "Monkey %d", &idx); err != nil {
				return nil, err
			}
		} else {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "Starting items: ") {
				line = strings.TrimPrefix(line, "Starting items: ")
				for _, item := range strings.Split(line, ", ") {
					var value int
					if value, err = strconv.Atoi(item); err != nil {
						return nil, err
					}

					problem.Monkeys[idx].Items = append(problem.Monkeys[idx].Items, int64(value))
				}
			} else if strings.HasPrefix(line, "Operation: ") {
				line = strings.TrimPrefix(line, "Operation: ")

				var operand string
				var operator string
				if _, err = fmt.Sscanf(line, "new = old %s %s", &operator, &operand); err != nil {
					return nil, err
				}

				problem.Monkeys[idx].Operand = operand
				problem.Monkeys[idx].Operator = operator
			} else if strings.HasPrefix(line, "Test: ") {
				line = strings.TrimPrefix(line, "Test: ")

				var div int
				if _, err = fmt.Sscanf(line, "divisible by %d", &div); err != nil {
					return nil, err
				}

				problem.Monkeys[idx].DivisibleBy = int64(div)
			} else if strings.HasPrefix(line, "If true: ") {
				line = strings.TrimPrefix(line, "If true: ")

				var dest int
				if _, err = fmt.Sscanf(line, "throw to monkey %d", &dest); err != nil {
					return nil, err
				}

				problem.Monkeys[idx].IfTrue = int64(dest)
			} else if strings.HasPrefix(line, "If false: ") {
				line = strings.TrimPrefix(line, "If false: ")

				var dest int
				if _, err = fmt.Sscanf(line, "throw to monkey %d", &dest); err != nil {
					return nil, err
				}

				problem.Monkeys[idx].IfFalse = int64(dest)
			}
		}
	}

	return &problem, nil
}

type WorryFunc func(i int, worry int64) int64

func Play(problem *Problem, rounds int, fn WorryFunc) error {
	for r := 0; r < rounds; r++ {
		for i := 0; i < len(problem.Monkeys); i++ {
			monkey := &problem.Monkeys[i]

			for _, item := range monkey.Items {
				var operand int64
				if monkey.Operand == "old" {
					operand = item
				} else {
					value, err := strconv.Atoi(monkey.Operand)

					if err != nil {
						return err
					}

					operand = int64(value)
				}

				var worry int64
				switch monkey.Operator {
				case "+":
					worry = item + operand
				case "*":
					worry = item * operand
				}

				worry = fn(i, worry)

				if (worry % monkey.DivisibleBy) == 0 {
					nextMonkey := &problem.Monkeys[monkey.IfTrue]
					nextMonkey.Items = append(nextMonkey.Items, worry)
				} else {
					nextMonkey := &problem.Monkeys[monkey.IfFalse]
					nextMonkey.Items = append(nextMonkey.Items, worry)
				}
			}

			monkey.Items = []int64{}
		}
	}

	return nil
}

func (problem *Problem) SolvePart1() error {
	n := len(problem.Monkeys)
	inspected := make([]int64, n)

	err := Play(problem, 20, func(i int, worry int64) int64 {
		inspected[i]++
		return worry / 3
	})

	if err != nil {
		return err
	}

	aoc.SortInt64Desc(inspected)
	fmt.Printf("Part 1: %d\n", inspected[0]*inspected[1])
	return nil
}

func (problem *Problem) SolvePart2() error {
	n := len(problem.Monkeys)
	inspected := make([]int64, n)

	var modulo int64 = 1
	for _, monkey := range problem.Monkeys {
		modulo *= monkey.DivisibleBy
	}

	err := Play(problem, 10000, func(i int, worry int64) int64 {
		inspected[i]++
		return worry % modulo
	})

	if err != nil {
		return err
	}

	aoc.SortInt64Desc(inspected)
	fmt.Printf("Part 2: %d\n", inspected[0]*inspected[1])
	return nil
}
