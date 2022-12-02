package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Round struct {
	Me   byte
	Them byte
}

type Problem struct {
	Rounds []Round
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_02.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Rounds: make([]Round, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		round := Round{
			Me:   line[2],
			Them: line[0],
		}

		problem.Rounds = append(problem.Rounds, round)
	}

	return &problem, nil
}

func (problem *Problem) SolvePart1() error {
	var score int64 = 0
	for _, round := range problem.Rounds {
		switch round.Me {
		case 'X': // Rock
			score += 1
			switch round.Them {
			case 'A': // Rock
				score += 3 // Draw
			case 'B': // Paper
				score += 0 // Lose
			case 'C': // Scissors
				score += 6 // Win
			}
		case 'Y': // Paper
			score += 2
			switch round.Them {
			case 'A': // Rock
				score += 6 // Win
			case 'B': // Paper
				score += 3 // Draw
			case 'C': // Scissors
				score += 0 // Lose
			}
		case 'Z': // Scissors
			score += 3
			switch round.Them {
			case 'A': // Rock
				score += 0 // Lose
			case 'B': // Paper
				score += 6 // Win
			case 'C': // Scissors
				score += 3 // Draw
			}
		}
	}

	fmt.Printf("Part 1: %d\n", score)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var score int64 = 0
	for _, round := range problem.Rounds {
		switch round.Me {
		case 'X': // Lose
			score += 0
			switch round.Them {
			case 'A': // Rock
				score += 3 // Scissors
			case 'B': // Paper
				score += 1 // Rock
			case 'C': // Scissors
				score += 2 // Paper
			}
		case 'Y': // Draw
			score += 3
			switch round.Them {
			case 'A': // Rock
				score += 1 // Rock
			case 'B': // Paper
				score += 2 // Paper
			case 'C': // Scissors
				score += 3 // Scissors
			}
		case 'Z': // Win
			score += 6
			switch round.Them {
			case 'A': // Rock
				score += 2 // Paper
			case 'B': // Paper
				score += 3 // Scissors
			case 'C': // Scissors
				score += 1 // Rock
			}
		}
	}

	fmt.Printf("Part 2: %d\n", score)
	return nil
}
