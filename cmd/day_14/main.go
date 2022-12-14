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

type Coord struct {
	X int
	Y int
}

type SparseGrid map[Coord]bool

type Problem struct {
	Cave SparseGrid
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_14.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Cave: make(SparseGrid),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		endpoints := make([]Coord, 0)
		for _, point := range strings.Split(line, "->") {
			point = strings.TrimSpace(point)
			coords := strings.Split(point, ",")

			var x int
			if x, err = strconv.Atoi(coords[0]); err != nil {
				return nil, err
			}

			var y int
			if y, err = strconv.Atoi(coords[1]); err != nil {
				return nil, err
			}

			endpoint := Coord{X: x, Y: y}
			endpoints = append(endpoints, endpoint)
		}

		for i := 0; i < len(endpoints)-1; i++ {
			Connect(&problem.Cave, endpoints[i], endpoints[i+1])
		}
	}

	return &problem, nil
}

func Connect(ptr *SparseGrid, a Coord, b Coord) {
	grid := *ptr
	if a.Y == b.Y {
		// Vertical
		min := aoc.Min(a.X, b.X)
		max := aoc.Max(a.X, b.X)
		for i := min; i <= max; i++ {
			c := Coord{X: i, Y: a.Y}
			grid[c] = true
		}
	} else if a.X == b.X {
		// Horizontal
		min := aoc.Min(a.Y, b.Y)
		max := aoc.Max(a.Y, b.Y)
		for i := min; i <= max; i++ {
			c := Coord{X: a.X, Y: i}
			grid[c] = true
		}
	}
}

func Drop(ptr *SparseGrid, c Coord, maxY int, hasFloor bool) bool {
	grid := *ptr

	//
	// Source of the sand is blocked.
	//
	if _, ok := grid[c]; ok {
		return false // Part 2
	}

	for c.Y < maxY {
		down := Coord{
			X: c.X,
			Y: c.Y + 1,
		}

		//
		// There isn't an endless void at the bottom of
		// the scan - there's floor, and you're standing on it!
		//
		if hasFloor && (down.Y == maxY) {
			grid[c] = true
			return true // Part 2
		}

		downLeft := Coord{
			X: c.X - 1,
			Y: c.Y + 1,
		}

		downRight := Coord{
			X: c.X + 1,
			Y: c.Y + 1,
		}

		if _, ok := grid[down]; !ok {
			c = down
		} else if _, ok := grid[downLeft]; !ok {
			c = downLeft
		} else if _, ok := grid[downRight]; !ok {
			c = downRight
		} else {
			//
			// If all three possible destinations are blocked,
			// the unit of sand comes to rest and no longer moves
			//
			grid[c] = true
			return true
		}
	}

	//
	// All further sand flows out the
	// bottom, falling into the endless void.
	//
	return false
}

func (problem *Problem) SolvePart1() error {
	max := 0
	for coord := range problem.Cave {
		if coord.Y > max {
			max = coord.Y
		}
	}

	sand := 0
	c := Coord{X: 500, Y: 0}
	for Drop(&problem.Cave, c, max, false) {
		sand += 1
	}

	fmt.Printf("Part 1: %d\n", sand)
	return nil
}

func (problem *Problem) SolvePart2() error {
	max := 0
	for coord := range problem.Cave {
		if coord.Y > max {
			max = coord.Y
		}
	}

	sand := 0
	c := Coord{X: 500, Y: 0}
	for Drop(&problem.Cave, c, max+2, true) {
		sand += 1
	}

	fmt.Printf("Part 2: %d\n", sand)
	return nil
}
