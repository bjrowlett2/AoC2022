package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

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

type Motion struct {
	Dx    int
	Dy    int
	Total int
}

type Problem struct {
	Motions []Motion
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_09.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Motions: make([]Motion, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var amount int
		if amount, err = strconv.Atoi(line[2:]); err != nil {
			return nil, err
		}

		switch line[0] {
		case 'U':
			motion := Motion{Dx: 0, Dy: +1, Total: amount}
			problem.Motions = append(problem.Motions, motion)
		case 'D':
			motion := Motion{Dx: 0, Dy: -1, Total: amount}
			problem.Motions = append(problem.Motions, motion)
		case 'L':
			motion := Motion{Dx: -1, Dy: 0, Total: amount}
			problem.Motions = append(problem.Motions, motion)
		case 'R':
			motion := Motion{Dx: +1, Dy: 0, Total: amount}
			problem.Motions = append(problem.Motions, motion)
		}
	}

	return &problem, nil
}

type Coord struct {
	X int
	Y int
}

func Follow(head Coord, tail *Coord) {
	absX := aoc.Abs(head.X - tail.X)
	absY := aoc.Abs(head.Y - tail.Y)

	signX := aoc.Sign(head.X - tail.X)
	signY := aoc.Sign(head.Y - tail.Y)

	// If the head is ever two steps directly up, down, left,
	// or right from the tail, the tail must also move one step
	// in that direction so it remains close enough:
	if (absX == 2) && (absY == 0) {
		tail.X += signX
	} else if (absX == 0) && (absY == 2) {
		tail.Y += signY
	}

	// Otherwise, if the head and tail aren't touching and
	// aren't in the same row or column, the tail always moves
	// one step diagonally to keep up:
	if (absX == 2) && (absY == 1) {
		tail.X += signX
		tail.Y += signY
	} else if (absX == 1) && (absY == 2) {
		tail.X += signX
		tail.Y += signY
	}

	// However, be careful: more types of motion are possible
	// than before, so you might want to visually compare your
	// simulated rope to the one above.
	if (absX == 2) && (absY == 2) {
		tail.X += signX
		tail.Y += signY
	}
}

func (problem *Problem) SolvePart1() error {
	head := Coord{X: 0, Y: 0}
	tail := Coord{X: 0, Y: 0}

	visited := make(aoc.Set[Coord])
	for _, m := range problem.Motions {
		for i := 0; i < m.Total; i++ {
			head.X += m.Dx
			head.Y += m.Dy
			Follow(head, &tail)

			visited.Add(tail)
		}
	}

	total := 0
	for _, v := range visited {
		if v {
			total += 1
		}
	}

	fmt.Printf("Part 1: %d\n", total)
	return nil
}

func (problem *Problem) SolvePart2() error {
	knots := make([]Coord, 10)
	visited := make(aoc.Set[Coord])

	n := len(knots)
	head := &knots[0]
	for _, m := range problem.Motions {
		for i := 0; i < m.Total; i++ {
			head.X += m.Dx
			head.Y += m.Dy

			for k := 1; k < len(knots); k++ {
				Follow(knots[k-1], &knots[k])
			}

			visited.Add(knots[n-1])
		}
	}

	total := 0
	for _, v := range visited {
		if v {
			total += 1
		}
	}

	fmt.Printf("Part 2: %d\n", total)
	return nil
}
