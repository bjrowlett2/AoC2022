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

type Row []rune
type Grid []Row

type Problem struct {
	Map    Grid
	Width  int
	Height int
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_24.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Map: make(Grid, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make(Row, 0)
		for _, r := range scanner.Text() {
			row = append(row, r)
		}

		problem.Map = append(problem.Map, row)
	}

	problem.Width = len(problem.Map[0])
	problem.Height = len(problem.Map)
	return &problem, nil
}

type State struct {
	X    int
	Y    int
	Time int
}

func Wrap(x, limit int) int {
	x %= limit
	if x < 0 {
		x += limit
	}

	return x + 1
}

func CanGo(problem *Problem, x, y, time int) bool {
	if !aoc.Between(0, x, problem.Width-1) {
		return false
	} else if !aoc.Between(0, y, problem.Height-1) {
		return false
	}

	if problem.Map[y][x] == '#' {
		return false
	}

	x1 := Wrap((x-1)-time, problem.Width-2)
	x2 := Wrap((x-1)+time, problem.Width-2)
	y1 := Wrap((y-1)-time, problem.Height-2)
	y2 := Wrap((y-1)+time, problem.Height-2)

	if aoc.Between(0, x1, problem.Width-1) {
		if problem.Map[y][x1] == '>' {
			return false
		}
	}

	if aoc.Between(0, x2, problem.Width-1) {
		if problem.Map[y][x2] == '<' {
			return false
		}
	}

	if aoc.Between(0, y1, problem.Height-1) {
		if problem.Map[y1][x] == 'v' {
			return false
		}
	}

	if aoc.Between(0, y2, problem.Height-1) {
		if problem.Map[y2][x] == '^' {
			return false
		}
	}

	return true
}

func Navigate(problem *Problem, startX, startY, startTime int) int {
	start := State{
		X:    startX,
		Y:    startY,
		Time: startTime,
	}

	minTime := math.MaxInt
	queue := make(aoc.Queue[State], 0)

	queue.Push(start)
	for len(queue) > 0 {
		var state State
		queue.Pop(&state)

		if startY == 0 {
			if state.Y == (problem.Height - 1) {
				minTime = state.Time
				break
			}
		} else {
			if state.Y == 0 {
				minTime = state.Time
				break
			}
		}

		newTime := state.Time + 1
		if CanGo(problem, state.X, state.Y, newTime) {
			next := State{
				X:    state.X,
				Y:    state.Y,
				Time: newTime,
			}

			if !queue.Contains(next) {
				queue.Push(next)
			}
		}

		if CanGo(problem, state.X+1, state.Y, newTime) {
			next := State{
				X:    state.X + 1,
				Y:    state.Y,
				Time: newTime,
			}

			if !queue.Contains(next) {
				queue.Push(next)
			}
		}

		if CanGo(problem, state.X-1, state.Y, newTime) {
			next := State{
				X:    state.X - 1,
				Y:    state.Y,
				Time: newTime,
			}

			if !queue.Contains(next) {
				queue.Push(next)
			}
		}

		if CanGo(problem, state.X, state.Y+1, newTime) {
			next := State{
				X:    state.X,
				Y:    state.Y + 1,
				Time: newTime,
			}

			if !queue.Contains(next) {
				queue.Push(next)
			}
		}

		if CanGo(problem, state.X, state.Y-1, newTime) {
			next := State{
				X:    state.X,
				Y:    state.Y - 1,
				Time: newTime,
			}

			if !queue.Contains(next) {
				queue.Push(next)
			}
		}
	}

	return minTime
}

func FindStart(problem *Problem, x, y int) (int, int) {
	for ; x < problem.Width; x++ {
		if problem.Map[y][x] == '.' {
			break
		}
	}

	return x, y
}

func (problem *Problem) SolveBothParts() error {
	x1, y1 := FindStart(problem, 0, 0)
	part1 := Navigate(problem, x1, y1, 0)
	fmt.Printf("Part 1: %d\n", part1)

	// Going back for snacks...
	x2, y2 := FindStart(problem, 0, problem.Height-1)
	part2 := Navigate(problem, x1, y1, Navigate(problem, x2, y2, part1))
	fmt.Printf("Part 1: %d\n", part2)
	return nil
}
