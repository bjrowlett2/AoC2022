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

func CanMove(rock Rock, chamber Chamber, dx, dy int64) bool {
	for coord := range rock.Points {
		coord.X += dx
		coord.Y += dy

		if (coord.X < 0) || (coord.X > 6) {
			return false // Hit the wall.
		}

		if coord.Y <= 0 {
			return false // Hit the ground.
		}

		if _, ok := chamber[coord]; ok {
			return false // Hit another rock.
		}
	}

	return true
}

func Move(rock Rock, chamber Chamber, dx, dy int64) Rock {
	moved := Rock{
		Width:  rock.Width,
		Height: rock.Height,
		Points: make(map[Coord]bool),
	}

	for coord := range rock.Points {
		coord.X += dx
		coord.Y += dy
		moved.Points[coord] = true
	}

	return moved
}

type Key struct {
	RockIndex int64
	JetIndex  int64
}

type State struct {
	I      int64
	Height int64
}

func Simulate(problem *Problem, steps int64) int64 {
	rocks := make([]Rock, 0)
	rocks = append(rocks, NewRock("####"))
	rocks = append(rocks, NewRock(".#.\n###\n.#."))
	rocks = append(rocks, NewRock("..#\n..#\n###"))
	rocks = append(rocks, NewRock("#\n#\n#\n#"))
	rocks = append(rocks, NewRock("##\n##"))

	var numRocks int64 = int64(len(rocks))
	var numJets int64 = int64(len(problem.Jets))

	var height int64 = 0
	chamber := make(Chamber)

	seen := make(map[Key]State)

	var i int64
	var jetIndex int64 = 0
	for i = 0; i < steps; i++ {
		rockIndex := i % numRocks
		template := rocks[rockIndex]

		rock := Rock{
			Width:  template.Width,
			Height: template.Height,
			Points: make(map[Coord]bool),
		}

		for coord := range template.Points {
			coord.X += 2
			coord.Y += height + template.Height + 3
			rock.Points[coord] = true
		}

		for {
			jet := problem.Jets[jetIndex]
			if CanMove(rock, chamber, int64(jet), 0) {
				rock = Move(rock, chamber, int64(jet), 0)
			}

			jetIndex += 1
			jetIndex %= numJets

			if CanMove(rock, chamber, 0, -1) {
				rock = Move(rock, chamber, 0, -1)
			} else {
				for coord := range rock.Points {
					chamber[coord] = true
					height = aoc.Max64(height, coord.Y)
				}

				//
				// Below here is disgusting...
				//

				deleted := make([]Coord, 0)
				for coord := range chamber {
					if coord.Y < (height - 40) {
						deleted = append(deleted, coord)
					}
				}

				for _, coord := range deleted {
					delete(chamber, coord)
				}

				var x int64
				found := true
				for x = 0; x < 7; x++ {
					coord := Coord{X: x, Y: height}
					if _, ok := chamber[coord]; !ok {
						found = false
					}
				}

				if found {
					key := Key{
						RockIndex: rockIndex,
						JetIndex:  jetIndex,
					}

					if state, ok := seen[key]; ok {
						dI := i - state.I
						dH := height - state.Height

						z := (steps - i) / dI

						i += dI * z
						height += dH * z

						added := make([]Coord, 0)
						deleted := make([]Coord, 0)
						for coord := range chamber {
							deleted = append(deleted, coord)

							coord.Y += dH * z
							added = append(added, coord)
						}

						for _, coord := range added {
							chamber[coord] = true
						}

						for _, coord := range deleted {
							delete(chamber, coord)
						}
					}

					seen[key] = State{
						I:      i,
						Height: height,
					}

					//
					// Hopefully, I'll come back and clean that up later...
					//
				}

				break
			}
		}
	}

	return height
}

func (problem *Problem) SolvePart1() error {
	height := Simulate(problem, 2022)
	fmt.Printf("Part 1: %d\n", height)
	return nil
}

func (problem *Problem) SolvePart2() error {
	height := Simulate(problem, 1000000000000)
	fmt.Printf("Part 2: %d\n", height)
	return nil
}
