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

type Cell int

const (
	CellBlank Cell = iota
	CellSolidWall
	CellOpenTile
)

type Row []Cell
type Grid []Row

type Problem struct {
	Map  Grid
	Path string
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_22.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Map: make(Grid, 0),
	}

	y := 0
	maxWidth := 0
	loadingMap := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			loadingMap = false
			continue
		}

		if loadingMap {
			width := len(line)
			maxWidth = aoc.Max(maxWidth, width)

			row := make(Row, width)
			problem.Map = append(problem.Map, row)

			x := 0
			for _, r := range line {
				if r == ' ' {
					problem.Map[y][x] = CellBlank
				} else if r == '#' {
					problem.Map[y][x] = CellSolidWall
				} else if r == '.' {
					problem.Map[y][x] = CellOpenTile
				}

				x += 1
			}

			y += 1
		} else {
			problem.Path = line
		}
	}

	for i := 0; i < len(problem.Map); i++ {
		if len(problem.Map[i]) < maxWidth {
			blanks := make([]Cell, maxWidth-len(problem.Map[i]))
			problem.Map[i] = append(problem.Map[i], blanks...)
		}
	}

	return &problem, nil
}

type Facing int

const (
	FacingUp    Facing = 3
	FacingLeft  Facing = 2
	FacingDown  Facing = 1
	FacingRight Facing = 0
)

func Wrap(x, maximum int) int {
	x = x % maximum
	if x < 0 {
		x += maximum
	}

	return x
}

func Delta(facing Facing) (int, int) {
	dx, dy := 0, 0
	switch facing {
	case FacingUp:
		dx, dy = +0, -1
	case FacingLeft:
		dx, dy = -1, +0
	case FacingDown:
		dx, dy = +0, +1
	case FacingRight:
		dx, dy = +1, +0
	}

	return dx, dy
}

func Walk(grid Grid, x, y, distance int, facing Facing) (int, int) {
	height := len(grid)
	dx, dy := Delta(facing)

	for distance > 0 {
		distance -= 1
		width := len(grid[y])

		newX := Wrap(x+dx, width)
		newY := Wrap(y+dy, height)
		for grid[newY][newX] == CellBlank {
			newX = Wrap(newX+dx, width)
			newY = Wrap(newY+dy, height)
		}

		if grid[newY][newX] == CellSolidWall {
			continue
		}

		x, y = newX, newY
	}

	return x, y
}

func (problem *Problem) SolvePart1() error {
	var err error

	x, y := 0, 0
	for problem.Map[y][x] != CellOpenTile {
		x += 1
	}

	facing := FacingRight
	for i := 0; i < len(problem.Path); i++ {
		if problem.Path[i] == 'L' {
			facing -= 1
			if facing < 0 {
				facing += 4
			}
		} else if problem.Path[i] == 'R' {
			facing += 1
			if facing > 3 {
				facing -= 4
			}
		} else {
			var distance int
			if _, err = fmt.Sscanf(problem.Path[i:], "%d", &distance); err != nil {
				return err
			}

			x, y = Walk(problem.Map, x, y, distance, facing)

			n := len(strconv.Itoa(distance))
			i += n - 1
		}
	}

	row := y + 1
	column := x + 1
	password := 1000*row + 4*column + int(facing)
	fmt.Printf("Part 1: %d\n", password)
	return nil
}

func Between(min, x, max int) bool {
	return (min <= x) && (x < max)
}

func WalkPt2(grid Grid, x, y, distance int, facing Facing) (int, int, Facing) {
	for distance > 0 {
		distance -= 1
		dx, dy := Delta(facing)

		newX := x + dx
		newY := y + dy
		newFacing := facing

		if newX < 0 {
			if Between(100, newY, 150) {
				if facing == FacingLeft {
					newX = 50
					newY = 50 - (newY - 100) - 1
					newFacing = FacingRight
				}
			} else if Between(150, newY, 200) {
				if facing == FacingLeft {
					newX = (newY - 150) + 50
					newY = 0
					newFacing = FacingDown
				}
			}
		} else if Between(0, newX, 50) {
			if Between(0, newY, 50) {
				if facing == FacingLeft {
					newX = 0
					newY = (50 - newY - 1) + 100
					newFacing = FacingRight
				}
			} else if Between(50, newY, 100) {
				if facing == FacingLeft {
					newX = newY - 50
					newY = 100
					newFacing = FacingDown
				} else if facing == FacingUp {
					newY = newX + 50
					newX = 50
					newFacing = FacingRight
				}
			} else if newY >= 200 {
				if facing == FacingDown {
					newX = newX + 100
					newY = 0
					newFacing = FacingDown
				}
			}
		} else if Between(50, newX, 100) {
			if newY < 0 {
				if facing == FacingUp {
					newY = (newX - 50) + 150
					newX = 0
					newFacing = FacingRight
				}
			} else if Between(150, newY, 200) {
				if facing == FacingDown {
					newY = (newX - 50) + 150
					newX = 49
					newFacing = FacingLeft
				} else if facing == FacingRight {
					newX = (newY - 150) + 50
					newY = 149
					newFacing = FacingUp
				}
			}
		} else if Between(100, newX, 150) {
			if newY < 0 {
				if facing == FacingUp {
					newX = newX - 100
					newY = 199
					newFacing = FacingUp
				}
			} else if Between(50, newY, 100) {
				if facing == FacingDown {
					newY = (newX - 100) + 50
					newX = 99
					newFacing = FacingLeft
				} else if facing == FacingRight {
					newX = (newY - 50) + 100
					newY = 49
					newFacing = FacingUp
				}
			} else if Between(100, newY, 150) {
				if facing == FacingRight {
					newY = 50 - (newY - 100) - 1
					newX = 149
					newFacing = FacingLeft
				}
			}
		} else if newX >= 150 {
			if Between(0, newY, 50) {
				if facing == FacingRight {
					newY = (50 - newY - 1) + 100
					newX = 99
					newFacing = FacingLeft
				}
			}
		}

		if grid[newY][newX] == CellSolidWall {
			continue
		}

		x, y = newX, newY
		facing = newFacing
	}

	return x, y, facing
}

func (problem *Problem) SolvePart2() error {
	var err error

	x, y := 0, 0
	for problem.Map[y][x] != CellOpenTile {
		x += 1
	}

	facing := FacingRight
	for i := 0; i < len(problem.Path); i++ {
		if problem.Path[i] == 'L' {
			facing -= 1
			if facing < 0 {
				facing += 4
			}
		} else if problem.Path[i] == 'R' {
			facing += 1
			if facing > 3 {
				facing -= 4
			}
		} else {
			var distance int
			if _, err = fmt.Sscanf(problem.Path[i:], "%d", &distance); err != nil {
				return err
			}

			x, y, facing = WalkPt2(problem.Map, x, y, distance, facing)

			n := len(strconv.Itoa(distance))
			i += n - 1
		}
	}

	row := y + 1
	column := x + 1
	password := 1000*row + 4*column + int(facing)
	fmt.Printf("Part 2: %d\n", password)
	return nil
}
