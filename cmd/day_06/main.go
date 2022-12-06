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

type Problem struct {
	Stream string
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_06.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Stream: "",
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		problem.Stream += line
	}

	return &problem, nil
}

func IsStartOfPacket(s string, n, k int) bool {
	set := make(map[byte]bool)
	for j := n - k; j < n; j += 1 {
		c := s[j]
		set[c] = true
	}

	return len(set) == k
}

func (problem *Problem) SolvePart1() error {
	k := 4
	for i := k; i < len(problem.Stream); i += 1 {
		if IsStartOfPacket(problem.Stream, i, k) {
			fmt.Printf("Part 1: %d\n", i)
			break
		}
	}

	return nil
}

func (problem *Problem) SolvePart2() error {
	k := 14
	for i := k; i < len(problem.Stream); i += 1 {
		if IsStartOfPacket(problem.Stream, i, k) {
			fmt.Printf("Part 2: %d\n", i)
			break
		}
	}

	return nil
}
