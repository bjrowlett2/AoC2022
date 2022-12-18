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

type Jet int64

const (
	JetLeft  Jet = -1
	JetRight Jet = +1
)

type Problem struct {
	Jets []Jet
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_17.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Jets: make([]Jet, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			if r == '<' {
				problem.Jets = append(problem.Jets, JetLeft)
			} else if r == '>' {
				problem.Jets = append(problem.Jets, JetRight)
			}
		}
	}

	return &problem, nil
}

type Coord struct {
	X int64
	Y int64
}

type Rock struct {
	Width  int64
	Height int64
	Points map[Coord]bool
}

func NewRock(shape string) Rock {
	rock := Rock{
		Width:  0,
		Height: 0,
		Points: make(map[Coord]bool),
	}

	coord := Coord{X: 0, Y: 0}
	for _, r := range shape {
		if r == '#' {
			rock.Width = coord.X + 1
			rock.Height = aoc.Abs64(coord.Y) + 1
			rock.Points[coord] = true
		}

		coord.X += 1
		if r == '\n' {
			coord.X = 0
			coord.Y -= 1
		}
	}

	return rock
}

type Chamber map[Coord]bool

func Move(rock Rock, chamber Chamber, jet Jet) (Rock, bool) {
	can := true
	for coord := range rock.Points {
		coord.X += int64(jet)
		if (coord.X < 0) || (coord.X > 6) {
			can = false
			break
		}

		if _, ok := chamber[coord]; ok {
			can = false
			break
		}
	}

	if can {
		moved := Rock{
			Width:  rock.Width,
			Height: rock.Height,
			Points: make(map[Coord]bool),
		}

		for coord := range rock.Points {
			coord.X += int64(jet)
			moved.Points[coord] = true
		}

		return moved, true
	}

	return rock, false
}

func Fall(rock Rock, chamber Chamber) (Rock, bool) {
	can := true
	for coord := range rock.Points {
		coord.Y -= 1
		if coord.Y <= 0 {
			can = false
			break
		}

		if _, ok := chamber[coord]; ok {
			can = false
			break
		}
	}

	if can {
		moved := Rock{
			Width:  rock.Width,
			Height: rock.Height,
			Points: make(map[Coord]bool),
		}

		for coord := range rock.Points {
			coord.Y -= 1
			moved.Points[coord] = true
		}

		return moved, true
	}

	return rock, false
}

func Draw(rock Rock, chamber Chamber) {
	w := 7
	var h int64 = 20

	grid := make([][]byte, 0)

	var y int64 = 0
	for y = 0; y < h; y++ {
		grid = append(grid, make([]byte, 0))
		for x := 0; x < w; x++ {
			grid[y] = append(grid[y], ' ')
		}
	}

	for coord := range chamber {
		grid[h-coord.Y][coord.X] = '#'
	}

	for coord := range rock.Points {
		grid[h-coord.Y][coord.X] = '@'
	}

	for y = 0; y < h; y++ {
		fmt.Print("=")
		fmt.Print(string(grid[y]))
		fmt.Print("=")
		fmt.Println()
	}

	fmt.Println("=========")
}

func (problem *Problem) SolvePart1() error {
	rocks := make([]Rock, 0)
	rocks = append(rocks, NewRock("####"))
	rocks = append(rocks, NewRock(".#.\n###\n.#."))
	rocks = append(rocks, NewRock("..#\n..#\n###"))
	rocks = append(rocks, NewRock("#\n#\n#\n#"))
	rocks = append(rocks, NewRock("##\n##"))

	chamber := make(Chamber)

	jetIndex := 0
	var floor int64 = 0
	for i := 0; i < 2022; i++ {
		rockIndex := i % len(rocks)
		template := rocks[rockIndex]

		rock := Rock{
			Width:  template.Width,
			Height: template.Height,
			Points: make(map[Coord]bool),
		}

		for coord := range template.Points {
			//
			// Each rock appears so that its left edge is two units away
			// from the left wall and its bottom edge is three units above
			// the highest rock in the room (or the floor, if there isn't one).
			//
			coord.X += 2
			coord.Y += floor + template.Height + 3
			rock.Points[coord] = true
		}

		//Draw(rock, chamber)

		for {
			jet := problem.Jets[jetIndex]
			rock, _ = Move(rock, chamber, jet)

			var fell bool
			rock, fell = Fall(rock, chamber)

			jetIndex += 1
			jetIndex %= len(problem.Jets)

			if !fell {
				for coord := range rock.Points {
					chamber[coord] = true
					floor = aoc.Max64(floor, coord.Y)
				}

				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", floor)
	return nil
}

func (problem *Problem) SolvePart2() error {
	rocks := make([]Rock, 0)
	rocks = append(rocks, NewRock("####"))
	rocks = append(rocks, NewRock(".#.\n###\n.#."))
	rocks = append(rocks, NewRock("..#\n..#\n###"))
	rocks = append(rocks, NewRock("#\n#\n#\n#"))
	rocks = append(rocks, NewRock("##\n##"))

	chamber := make(Chamber)

	jetIndex := 0
	var floor int64 = 0
	for i := 0; i < 1000000000000; i++ {
		rockIndex := i % len(rocks)
		template := rocks[rockIndex]

		rock := Rock{
			Width:  template.Width,
			Height: template.Height,
			Points: make(map[Coord]bool),
		}

		for coord := range template.Points {
			//
			// Each rock appears so that its left edge is two units away
			// from the left wall and its bottom edge is three units above
			// the highest rock in the room (or the floor, if there isn't one).
			//
			coord.X += 2
			coord.Y += floor + template.Height + 3
			rock.Points[coord] = true
		}

		//Draw(rock, chamber)

		for {
			jet := problem.Jets[jetIndex]
			rock, _ = Move(rock, chamber, jet)

			var fell bool
			rock, fell = Fall(rock, chamber)

			jetIndex += 1
			jetIndex %= len(problem.Jets)

			if !fell {
				max := make([]int64, 0)
				for i := 0; i < 7; i++ {
					max = append(max, 0)
				}

				for coord := range rock.Points {
					chamber[coord] = true
					floor = aoc.Max64(floor, coord.Y)
					max[coord.Y] = aoc.Max64(max[coord.Y], coord.Y)
				}

				break
			}
		}
	}

	fmt.Printf("Part 2: %d\n", floor)
	return nil
}
