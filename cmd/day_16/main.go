package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

type Valve struct {
	FlowRate  int64
	Tunnels   []string
	Distances map[string]int64
}

type Problem struct {
	Valves map[string]Valve
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_16.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Valves: make(map[string]Valve),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ";")
		parts[0] = strings.TrimSpace(parts[0])
		parts[1] = strings.TrimPrefix(parts[1], " tunnel leads to valve ")
		parts[1] = strings.TrimPrefix(parts[1], " tunnels lead to valves ")

		var valve string
		var flowRate int64
		if _, err = fmt.Sscanf(parts[0], "Valve %s has flow rate=%d", &valve, &flowRate); err != nil {
			return nil, err
		}

		problem.Valves[valve] = Valve{
			FlowRate:  flowRate,
			Tunnels:   strings.Split(parts[1], ", "),
			Distances: make(map[string]int64),
		}
	}

	for start, valve := range problem.Valves {
		valve.Distances[start] = 0

		queue := make(aoc.Queue[string], 0)
		visited := make(map[string]bool)

		queue.Push(start)
		for len(queue) != 0 {
			var finish string
			queue.Pop(&finish)

			other := problem.Valves[finish]
			steps := valve.Distances[finish]
			for _, next := range other.Tunnels {
				if !visited[next] {
					queue.Push(next)
					visited[next] = true
					valve.Distances[next] = steps + 1
				}
			}
		}
	}

	return &problem, nil
}

type State struct {
	Valve    string
	Time     int64
	Pressure int64
	Visited  string
}

func TraverseTunnels(problem *Problem, time int64) ([]string, int64) {
	initial := State{
		Valve:    "AA",
		Time:     time,
		Pressure: 0,
		Visited:  "AA",
	}

	var maximumPath string
	var maximumPressure int64 = 0

	queue := make(aoc.Queue[State], 0)

	queue.Push(initial)
	for len(queue) != 0 {
		var state State
		queue.Pop(&state)

		start := state.Valve
		valve := problem.Valves[start]

		if state.Pressure > maximumPressure {
			maximumPath = state.Visited
			maximumPressure = state.Pressure
		}

		for neighbor, distance := range valve.Distances {
			if distance < state.Time {
				if problem.Valves[neighbor].FlowRate != 0 {
					if !strings.Contains(state.Visited, neighbor) {
						newTime := state.Time - valve.Distances[neighbor] - 1
						newPressure := newTime * problem.Valves[neighbor].FlowRate

						next := State{
							Valve:    neighbor,
							Time:     newTime,
							Pressure: state.Pressure + newPressure,
							Visited:  fmt.Sprintf("%s-%s", state.Visited, neighbor),
						}

						queue = append(queue, next)
					}
				}
			}
		}
	}

	return strings.Split(maximumPath, "-"), maximumPressure
}

func (problem *Problem) SolvePart1() error {
	_, pressure := TraverseTunnels(problem, 30)
	fmt.Printf("Part 1: %d\n", pressure)
	return nil
}

func (problem *Problem) SolvePart2() error {
	path, pressure := TraverseTunnels(problem, 26)
	for _, name := range path {
		if valve, ok := problem.Valves[name]; ok {
			valve.FlowRate = 0
			problem.Valves[name] = valve
		}
	}

	_, elephant := TraverseTunnels(problem, 26)
	fmt.Printf("Part 2: %d\n", pressure+elephant)
	return nil
}
