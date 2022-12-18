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

type Coord3d struct {
	X int64
	Y int64
	Z int64
}

type Problem struct {
	Droplet map[Coord3d]bool
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_18.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Droplet: make(map[Coord3d]bool),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		coord := Coord3d{}
		if _, err = fmt.Sscanf(line, "%d,%d,%d", &coord.X, &coord.Y, &coord.Z); err != nil {
			return nil, err
		}

		problem.Droplet[coord] = true
	}

	return &problem, nil
}

func Neighbors(coord Coord3d) []Coord3d {
	neighbors := make([]Coord3d, 6)
	neighbors[0] = Coord3d{X: coord.X - 1, Y: coord.Y, Z: coord.Z}
	neighbors[1] = Coord3d{X: coord.X + 1, Y: coord.Y, Z: coord.Z}
	neighbors[2] = Coord3d{X: coord.X, Y: coord.Y - 1, Z: coord.Z}
	neighbors[3] = Coord3d{X: coord.X, Y: coord.Y + 1, Z: coord.Z}
	neighbors[4] = Coord3d{X: coord.X, Y: coord.Y, Z: coord.Z - 1}
	neighbors[5] = Coord3d{X: coord.X, Y: coord.Y, Z: coord.Z + 1}
	return neighbors
}

func (problem *Problem) SolvePart1() error {
	var surfaceArea int64 = 0
	for coord := range problem.Droplet {
		for _, neighbor := range Neighbors(coord) {
			if _, ok := problem.Droplet[neighbor]; !ok {
				surfaceArea += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", surfaceArea)
	return nil
}

type Bounds struct {
	MinX int64
	MaxX int64
	MinY int64
	MaxY int64
	MinZ int64
	MaxZ int64
}

type BoundingBox map[Coord3d]bool

func DisplaceAir(problem *Problem, outside BoundingBox, bounds Bounds) {
	queue := make(aoc.Queue[Coord3d], 0)

	start := Coord3d{
		X: bounds.MinX,
		Y: bounds.MinY,
		Z: bounds.MinZ,
	}

	queue.Push(start)
	for len(queue) != 0 {
		var coord Coord3d
		queue.Pop(&coord)

		outside[coord] = true
		for _, neighbor := range Neighbors(coord) {
			if queue.Contains(neighbor) {
				continue // We're in the queue.
			}

			if _, ok := outside[neighbor]; ok {
				continue // We've already visited it.
			}

			if _, ok := problem.Droplet[neighbor]; ok {
				continue // We're inside the droplet now.
			}

			if (neighbor.X < bounds.MinX) || (neighbor.X > bounds.MaxX) {
				continue // We're outside the bounding box now.
			}

			if (neighbor.Y < bounds.MinY) || (neighbor.Y > bounds.MaxY) {
				continue // We're outside the bounding box now.
			}

			if (neighbor.Z < bounds.MinZ) || (neighbor.Z > bounds.MaxZ) {
				continue // We're outside the bounding box now.
			}

			queue.Push(neighbor)
		}
	}
}

func (problem *Problem) SolvePart2() error {
	bounds := Bounds{
		MinX: math.MaxInt64,
		MaxX: math.MinInt64,
		MinY: math.MaxInt64,
		MaxY: math.MinInt64,
		MinZ: math.MaxInt64,
		MaxZ: math.MinInt64,
	}

	for coord := range problem.Droplet {
		bounds.MinX = aoc.Min64(bounds.MinX, coord.X-1)
		bounds.MaxX = aoc.Max64(bounds.MaxX, coord.X+1)
		bounds.MinY = aoc.Min64(bounds.MinY, coord.Y-1)
		bounds.MaxY = aoc.Max64(bounds.MaxY, coord.Y+1)
		bounds.MinZ = aoc.Min64(bounds.MinZ, coord.Z-1)
		bounds.MaxZ = aoc.Max64(bounds.MaxZ, coord.Z+1)
	}

	outside := make(map[Coord3d]bool)
	DisplaceAir(problem, outside, bounds)

	var surfaceArea int64 = 0
	for coord := range problem.Droplet {
		for _, neighbor := range Neighbors(coord) {
			if _, ok := outside[neighbor]; ok {
				if _, ok := problem.Droplet[neighbor]; !ok {
					surfaceArea += 1
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", surfaceArea)
	return nil
}
