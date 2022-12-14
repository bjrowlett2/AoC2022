package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

type Coord struct {
	X int
	Y int
}

type Row []int
type Grid []Row

type Problem struct {
	Map    Grid
	Start  Coord
	Finish Coord
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_12.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	x, y := 0, 0
	problem := Problem{
		Map: make(Grid, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		x = 0
		row := make(Row, 0)

		for _, r := range line {
			if r == 'S' {
				problem.Start = Coord{
					X: x,
					Y: y,
				}

				row = append(row, 0)
			} else if r == 'E' {
				problem.Finish = Coord{
					X: x,
					Y: y,
				}

				row = append(row, 26) // z
			} else {
				height := int(r - 'a')
				row = append(row, height)
			}

			x += 1
		}

		y += 1
		problem.Map = append(problem.Map, row)
	}

	return &problem, nil
}

func Neighbors(grid Grid, c Coord) []Coord {
	neighbors := make([]Coord, 0)

	if c.X > 0 {
		coord := Coord{X: c.X - 1, Y: c.Y}
		neighbors = append(neighbors, coord)
	}

	if c.X < (len(grid[c.Y]) - 1) {
		coord := Coord{X: c.X + 1, Y: c.Y}
		neighbors = append(neighbors, coord)
	}

	if c.Y > 0 {
		coord := Coord{X: c.X, Y: c.Y - 1}
		neighbors = append(neighbors, coord)
	}

	if c.Y < (len(grid) - 1) {
		coord := Coord{X: c.X, Y: c.Y + 1}
		neighbors = append(neighbors, coord)
	}

	return neighbors
}

func ComputePath(grid Grid, s Coord, reverse bool) Grid {
	distances := make(Grid, 0)
	for y := 0; y < len(grid); y++ {
		distances = append(distances, Row{})
		for x := 0; x < len(grid[y]); x++ {
			distances[y] = append(distances[y], math.MaxInt32)
		}
	}

	distances[s.Y][s.X] = 0
	next := make(aoc.Queue[Coord], 0)

	next.Push(s)
	for len(next) != 0 {
		var c Coord
		next.Pop(&c)

		distance := distances[c.Y][c.X] + 1
		for _, n := range Neighbors(grid, c) {
			reachable := (grid[n.Y][n.X] - grid[c.Y][c.X]) <= 1
			if reverse {
				reachable = (grid[n.Y][n.X] - grid[c.Y][c.X]) >= -1
			}

			visited := distances[n.Y][n.X] < math.MaxInt32
			queued := next.Contains(n)

			if reachable && !(visited || queued) {
				next.Push(n)
				distances[n.Y][n.X] = distance
			}
		}
	}

	return distances
}

func (problem *Problem) SolvePart1() error {
	distances := ComputePath(problem.Map, problem.Start, false)
	fmt.Printf("Part 1: %d\n", distances[problem.Finish.Y][problem.Finish.X])
	return nil
}

func (problem *Problem) SolvePart2() error {
	//
	// Optimization:
	//
	// Instead of computing the path starting from each 'a' (ie: multiple passes),
	// if we reverse the problem and compute the path from the end, we can just iterate
	// over the map once and pick the 'a' with the minimum distance.
	//
	distances := ComputePath(problem.Map, problem.Finish, true)

	minimum := math.MaxInt32
	for y := 0; y < len(problem.Map); y++ {
		for x := 0; x < len(problem.Map[y]); x++ {
			if problem.Map[y][x] == 0 { // Zero == 'a'
				steps := distances[y][x]
				if steps < minimum {
					minimum = steps
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", minimum)
	return nil
}
