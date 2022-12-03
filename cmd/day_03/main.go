package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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

type Problem struct {
	Rucksacks []string
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_03.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Rucksacks: make([]string, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		problem.Rucksacks = append(problem.Rucksacks, line)
	}

	return &problem, nil
}

func (problem *Problem) SolvePart1() error {
	var priorities int64 = 0
	for _, rucksack := range problem.Rucksacks {
		half := len(rucksack) / 2

		for _, item := range rucksack[:half] {
			if strings.ContainsRune(rucksack[half:], item) {
				if unicode.IsLower(item) {
					priorities += int64(item) - int64('a') + 1
				} else if unicode.IsUpper(item) {
					priorities += 26 + int64(item) - int64('A') + 1
				}

				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", priorities)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var priorities int64 = 0
	for i := 0; i < len(problem.Rucksacks); i += 3 {
		for _, item := range problem.Rucksacks[i] {
			a := strings.ContainsRune(problem.Rucksacks[i+1], item)
			b := strings.ContainsRune(problem.Rucksacks[i+2], item)

			if a && b {
				if unicode.IsLower(item) {
					priorities += int64(item) - int64('a') + 1
				} else if unicode.IsUpper(item) {
					priorities += 26 + int64(item) - int64('A') + 1
				}

				break
			}
		}
	}

	fmt.Printf("Part 2: %d\n", priorities)
	return nil
}
