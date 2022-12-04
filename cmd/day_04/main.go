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

type Pair struct {
	A Section
	B Section
}

type Section struct {
	Min int64
	Max int64
}

type Problem struct {
	Pairs []Pair
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_04.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Pairs: make([]Pair, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var minA, maxA int64
		var minB, maxB int64
		if _, err = fmt.Sscanf(line, "%d-%d,%d-%d", &minA, &maxA, &minB, &maxB); err != nil {
			return nil, err
		}

		pair := Pair{
			A: Section{
				Min: minA,
				Max: maxA,
			},
			B: Section{
				Min: minB,
				Max: maxB,
			},
		}

		problem.Pairs = append(problem.Pairs, pair)
	}

	return &problem, nil
}

func Contained(a *Section, b *Section) bool {
	aContained := (b.Min <= a.Min) && (a.Max <= b.Max)
	bContained := (a.Min <= b.Min) && (b.Max <= a.Max)
	return aContained || bContained
}

func Overlapped(a *Section, b *Section) bool {
	return (a.Min <= b.Max) && (b.Min <= a.Max)
}

func (problem *Problem) SolvePart1() error {
	var contained int64 = 0
	for _, pair := range problem.Pairs {
		if Contained(&pair.A, &pair.B) {
			contained += 1
		}
	}

	fmt.Printf("Part 1: %d\n", contained)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var overlapped int64 = 0
	for _, pair := range problem.Pairs {
		if Overlapped(&pair.A, &pair.B) {
			overlapped += 1
		}
	}

	fmt.Printf("Part 2: %d\n", overlapped)
	return nil
}
