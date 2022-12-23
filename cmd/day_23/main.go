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

	if err = problem.SolveBothParts(); err != nil {
		log.Fatal(err)
	}
}

type Coord struct {
	X int
	Y int
}

type Direction int

const (
	DirectionNorth Direction = 0
	DirectionEast  Direction = 1
	DirectionSouth Direction = 2
	DirectionWest  Direction = 3
)

type Problem struct {
	Elves      map[Coord]bool
	Directions []Direction
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_23.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Elves: make(map[Coord]bool),
		Directions: []Direction{
			DirectionNorth,
			DirectionSouth,
			DirectionWest,
			DirectionEast,
		},
	}

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for x, r := range scanner.Text() {
			if r == '#' {
				c := Coord{X: x, Y: y}
				problem.Elves[c] = true
			}
		}

		y += 1
	}

	return &problem, nil
}

func NorthNeighbors(coord Coord) []Coord {
	neighbors := make([]Coord, 3)
	neighbors[0] = Coord{X: coord.X + 0, Y: coord.Y - 1} // N
	neighbors[1] = Coord{X: coord.X + 1, Y: coord.Y - 1} // NE
	neighbors[2] = Coord{X: coord.X - 1, Y: coord.Y - 1} // NW
	return neighbors
}

func EastNeighbors(coord Coord) []Coord {
	neighbors := make([]Coord, 3)
	neighbors[0] = Coord{X: coord.X + 1, Y: coord.Y + 0} // E
	neighbors[1] = Coord{X: coord.X + 1, Y: coord.Y - 1} // NE
	neighbors[2] = Coord{X: coord.X + 1, Y: coord.Y + 1} // SE
	return neighbors
}

func SouthNeighbors(coord Coord) []Coord {
	neighbors := make([]Coord, 3)
	neighbors[0] = Coord{X: coord.X + 0, Y: coord.Y + 1} // S
	neighbors[1] = Coord{X: coord.X + 1, Y: coord.Y + 1} // SE
	neighbors[2] = Coord{X: coord.X - 1, Y: coord.Y + 1} // SW
	return neighbors
}

func WestNeighbors(coord Coord) []Coord {
	neighbors := make([]Coord, 3)
	neighbors[0] = Coord{X: coord.X - 1, Y: coord.Y + 0} // W
	neighbors[1] = Coord{X: coord.X - 1, Y: coord.Y - 1} // NW
	neighbors[2] = Coord{X: coord.X - 1, Y: coord.Y + 1} // SW
	return neighbors
}

func AllNeighbors(coord Coord) []Coord {
	neighbors := make([]Coord, 8)
	neighbors[0] = Coord{X: coord.X + 0, Y: coord.Y - 1} // N
	neighbors[1] = Coord{X: coord.X + 1, Y: coord.Y + 0} // E
	neighbors[2] = Coord{X: coord.X + 0, Y: coord.Y + 1} // S
	neighbors[3] = Coord{X: coord.X - 1, Y: coord.Y + 0} // W
	neighbors[4] = Coord{X: coord.X + 1, Y: coord.Y - 1} // NE
	neighbors[5] = Coord{X: coord.X - 1, Y: coord.Y - 1} // NW
	neighbors[6] = Coord{X: coord.X + 1, Y: coord.Y + 1} // SE
	neighbors[7] = Coord{X: coord.X - 1, Y: coord.Y + 1} // SW
	return neighbors
}

func AreEmpty(problem *Problem, neighbors []Coord) bool {
	for _, neighbor := range neighbors {
		if _, ok := problem.Elves[neighbor]; ok {
			return false
		}
	}

	return true
}

func (problem *Problem) SolveBothParts() error {
	round := 0

	for {
		round += 1

		secondHalf := make([]Coord, 0)
		for coord := range problem.Elves {
			neighbors := AllNeighbors(coord)
			if !AreEmpty(problem, neighbors) {
				secondHalf = append(secondHalf, coord)
			}
		}

		proposed := make(map[Coord][]Coord)
		for _, coord := range secondHalf {
		loop:
			for _, dir := range problem.Directions {
				switch dir {
				case DirectionNorth:
					neighbors := NorthNeighbors(coord)
					if AreEmpty(problem, neighbors) {
						north := neighbors[0]
						proposed[north] = append(proposed[north], coord)
						break loop
					}

				case DirectionEast:
					neighbors := EastNeighbors(coord)
					if AreEmpty(problem, neighbors) {
						east := neighbors[0]
						proposed[east] = append(proposed[east], coord)
						break loop
					}

				case DirectionSouth:
					neighbors := SouthNeighbors(coord)
					if AreEmpty(problem, neighbors) {
						south := neighbors[0]
						proposed[south] = append(proposed[south], coord)
						break loop
					}

				case DirectionWest:
					neighbors := WestNeighbors(coord)
					if AreEmpty(problem, neighbors) {
						west := neighbors[0]
						proposed[west] = append(proposed[west], coord)
						break loop
					}
				}
			}
		}

		moved := 0
		for coord, elves := range proposed {
			if len(elves) == 1 {
				moved += 1
				problem.Elves[coord] = true
				delete(problem.Elves, elves[0])
			}
		}

		if round == 10 {
			minX := math.MaxInt
			maxX := math.MinInt
			minY := math.MaxInt
			maxY := math.MinInt
			for coord := range problem.Elves {
				minX = aoc.Min(minX, coord.X)
				maxX = aoc.Max(maxX, coord.X)
				minY = aoc.Min(minY, coord.Y)
				maxY = aoc.Max(maxY, coord.Y)
			}

			width := maxX - minX + 1
			height := maxY - minY + 1
			emptyTiles := (width * height) - len(problem.Elves)
			fmt.Printf("Part 1: %d\n", emptyTiles)
		}

		if moved == 0 {
			fmt.Printf("Part 2: %d\n", round)
			break
		}

		dir := problem.Directions[0]
		problem.Directions = problem.Directions[1:]
		problem.Directions = append(problem.Directions, dir)
	}

	return nil
}
