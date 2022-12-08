package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

type Row []int
type Grid []Row

type Problem struct {
	Forest Grid
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_08.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Forest: make(Grid, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			problem.Forest = append(problem.Forest, Row{})
		}

		for _, x := range line {
			var height int
			if height, err = strconv.Atoi(string(x)); err != nil {
				return nil, err
			}

			y := len(problem.Forest) - 1
			problem.Forest[y] = append(problem.Forest[y], height)
		}
	}

	return &problem, nil
}

func IsVisible(grid Grid, y, x int) bool {
	target := grid[y][x]

	up := true
	for i := 0; i < y; i++ {
		if grid[i][x] >= target {
			up = false
		}
	}

	left := true
	for i := 0; i < x; i++ {
		if grid[y][i] >= target {
			left = false
		}
	}

	right := true
	for i := x + 1; i < len(grid[y]); i++ {
		if grid[y][i] >= target {
			right = false
		}
	}

	down := true
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] >= target {
			down = false
		}
	}

	return up || left || right || down
}

func (problem *Problem) SolvePart1() error {
	visible := 0
	for y := 0; y < len(problem.Forest); y++ {
		for x := 0; x < len(problem.Forest[y]); x++ {
			if IsVisible(problem.Forest, y, x) {
				visible += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", visible)
	return nil
}

func ScenicScore(grid Grid, y, x int) int {
	target := grid[y][x]

	up := 0
	for j := y - 1; j >= 0; j-- {
		up += 1
		if grid[j][x] >= target {
			break
		}
	}

	left := 0
	for i := x - 1; i >= 0; i-- {
		left += 1
		if grid[y][i] >= target {
			break
		}
	}

	right := 0
	for i := x + 1; i < len(grid[y]); i++ {
		right += 1
		if grid[y][i] >= target {
			break
		}
	}

	down := 0
	for j := y + 1; j < len(grid); j++ {
		down += 1
		if grid[j][x] >= target {
			break
		}
	}

	return up * left * right * down
}

func (problem *Problem) SolvePart2() error {
	highest := 0
	for y := 0; y < len(problem.Forest); y++ {
		for x := 0; x < len(problem.Forest[y]); x++ {
			score := ScenicScore(problem.Forest, y, x)
			if score > highest {
				highest = score
			}
		}
	}

	fmt.Printf("Part 2: %d\n", highest)
	return nil
}
