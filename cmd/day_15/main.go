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
	X int64
	Y int64
}

type Sensor struct {
	Position          Coord
	ClosestBeacon     Coord
	ManhattanDistance int64
}

type Problem struct {
	Sensors []Sensor
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_15.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Sensors: make([]Sensor, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		sensor := Coord{}
		beacon := Coord{}
		if _, err = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y); err != nil {
			return nil, err
		}

		next := Sensor{
			Position:          sensor,
			ClosestBeacon:     beacon,
			ManhattanDistance: Distance(sensor, beacon),
		}

		problem.Sensors = append(problem.Sensors, next)
	}

	return &problem, nil
}

func Distance(a, b Coord) int64 {
	return aoc.Abs64(a.X-b.X) + aoc.Abs64(a.Y-b.Y)
}

func (problem *Problem) SolvePart1() error {
	var y int64 = 2000000

	var minimum int64 = math.MaxInt64
	var maximum int64 = -math.MaxInt64
	for _, sensor := range problem.Sensors {
		coord := sensor.Position

		dy := aoc.Abs64(coord.Y - y)
		if dy < sensor.ManhattanDistance {
			dx := sensor.ManhattanDistance - dy
			minimum = aoc.Min64(minimum, coord.X-dx)
			maximum = aoc.Max64(maximum, coord.X+dx)
		}
	}

	x := minimum
	var count int64 = 0

outer:
	for x <= maximum {
		test := Coord{X: x, Y: y}
		for _, sensor := range problem.Sensors {
			coord := sensor.Position
			if Distance(coord, test) <= sensor.ManhattanDistance {
				dy := aoc.Abs64(coord.Y - test.Y)
				dx := sensor.ManhattanDistance - dy
				increment := (coord.X - test.X) + dx + 1

				x += increment
				count += increment
				continue outer
			}
		}

		x += 1
	}

	//
	// Ugly: We over counted on the last
	// iteration because we went past maximum.
	//
	count -= (x - maximum)
	fmt.Printf("Part 1: %d\n", count)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var index int64 = 0

	var width int64 = 4000000
	var height int64 = 4000000

outer:
	for index <= (width * height) {
		var x int64 = index % width
		var y int64 = index / width

		test := Coord{X: x, Y: y}
		for _, sensor := range problem.Sensors {
			coord := sensor.Position
			if Distance(coord, test) <= sensor.ManhattanDistance {
				dy := aoc.Abs64(coord.Y - test.Y)
				dx := sensor.ManhattanDistance - dy
				index += (coord.X - test.X) + dx + 1
				continue outer
			}
		}

		frequency := x*4000000 + y
		fmt.Printf("Part 2: %d\n", frequency)
		break
	}

	return nil
}
