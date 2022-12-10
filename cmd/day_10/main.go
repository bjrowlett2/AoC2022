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

	if err = problem.SolvePart1(); err != nil {
		log.Fatal(err)
	}

	if err = problem.SolvePart2(); err != nil {
		log.Fatal(err)
	}
}

type Problem struct {
	Instructions []string
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_10.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Instructions: make([]string, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		problem.Instructions = append(problem.Instructions, line)
	}

	return &problem, nil
}

func (problem *Problem) SolvePart1() error {
	var err error

	x := 1
	cycle := 0
	strength := 0

	for _, inst := range problem.Instructions {
		temp := 0
		remaining := 0
		if inst == "noop" {
			remaining = 1
		} else {
			var value int
			if _, err = fmt.Sscanf(inst, "addx %d", &value); err != nil {
				return err
			}

			temp = value
			remaining = 2
		}

		for remaining > 0 {
			cycle += 1
			if (cycle-20)%40 == 0 {
				strength += cycle * x
			}

			remaining -= 1
			if remaining == 0 {
				x += temp
			}
		}
	}

	fmt.Printf("Part 1: %d\n", strength)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var err error

	x := 1
	cycle := 0

	screen := make([][]bool, 6)
	for i := 0; i < len(screen); i++ {
		screen[i] = make([]bool, 40)
	}

	for _, inst := range problem.Instructions {
		temp := 0
		remaining := 0
		if inst == "noop" {
			remaining = 1
		} else {
			var value int
			if _, err = fmt.Sscanf(inst, "addx %d", &value); err != nil {
				return err
			}

			temp = value
			remaining = 2
		}

		for remaining > 0 {
			cycle += 1
			idx := cycle - 1

			row := idx / 40
			column := idx % 40
			if aoc.Between(x-1, column, x+1) {
				screen[row][column] = true
			}

			remaining -= 1
			if remaining == 0 {
				x += temp
			}
		}
	}

	for row := 0; row < len(screen); row++ {
		fmt.Print("Part 2: ")
		for column := 0; column < len(screen[row]); column++ {
			if screen[row][column] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	return nil
}
