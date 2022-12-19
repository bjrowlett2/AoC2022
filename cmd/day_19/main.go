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

type Blueprint struct {
	OreRobotCost      int64   // Costs x ore.
	ClayRobotCost     int64   // Costs x ore.
	ObsidianRobotCost []int64 // Costs x ore and y clay.
	GeodeRobotCost    []int64 // Costs x ore and y obsidian.
}

type Problem struct {
	Blueprints []Blueprint
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_19.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Blueprints: make([]Blueprint, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		colon := strings.Split(line, ":")
		period := strings.Split(colon[1], ".")

		blueprint := Blueprint{
			OreRobotCost:      0,
			ClayRobotCost:     0,
			ObsidianRobotCost: make([]int64, 2),
			GeodeRobotCost:    make([]int64, 2),
		}

		fmt.Sscanf(period[0], " Each ore robot costs %d ore", &blueprint.OreRobotCost)
		fmt.Sscanf(period[1], " Each clay robot costs %d ore", &blueprint.ClayRobotCost)
		fmt.Sscanf(period[2], " Each obsidian robot costs %d ore and %d clay", &blueprint.ObsidianRobotCost[0], &blueprint.ObsidianRobotCost[1])
		fmt.Sscanf(period[3], " Each geode robot costs %d ore and %d obsidian", &blueprint.GeodeRobotCost[0], &blueprint.GeodeRobotCost[1])

		problem.Blueprints = append(problem.Blueprints, blueprint)
	}

	return &problem, nil
}

type State struct {
	Time int64

	Ore      int64
	Clay     int64
	Obsidian int64
	Geodes   int64

	OreRobots      int64
	ClayRobots     int64
	ObsidianRobots int64
	GeodeRobots    int64
}

func OpenGeodes(blueprint Blueprint, time int64) int64 {
	initial := State{
		Ore:  0,
		Time: time,
		// Fortunately, you have exactly one ore-collecting robot in
		// your pack that you can use to kickstart the whole operation.
		OreRobots: 1,
	}

	maxOreCost := blueprint.OreRobotCost
	maxOreCost = aoc.Max64(maxOreCost, blueprint.ClayRobotCost)
	maxOreCost = aoc.Max64(maxOreCost, blueprint.ObsidianRobotCost[0])
	maxOreCost = aoc.Max64(maxOreCost, blueprint.GeodeRobotCost[0])
	maxClayCost := blueprint.ObsidianRobotCost[1]
	maxObsidianCost := blueprint.GeodeRobotCost[1]

	var maximumGeodes int64 = 0
	queue := make(aoc.Queue[State], 0)

	queue.Push(initial)
	for len(queue) != 0 {
		var state State
		queue.Pop(&state)

		if state.Geodes < maximumGeodes {
			continue
		}

		maximumGeodes = aoc.Max64(maximumGeodes, state.Geodes)

		if state.Time > 0 {
			base := State{
				Time: state.Time - 1,

				Ore:      state.Ore + state.OreRobots,
				Clay:     state.Clay + state.ClayRobots,
				Obsidian: state.Obsidian + state.ObsidianRobots,
				Geodes:   state.Geodes + state.GeodeRobots,

				OreRobots:      state.OreRobots,
				ClayRobots:     state.ClayRobots,
				ObsidianRobots: state.ObsidianRobots,
				GeodeRobots:    state.GeodeRobots,
			}

			// Always build another geode robot if we can.
			if state.Ore >= blueprint.GeodeRobotCost[0] {
				if state.Obsidian >= blueprint.GeodeRobotCost[1] {
					next := base
					next.Ore -= blueprint.GeodeRobotCost[0]
					next.Obsidian -= blueprint.GeodeRobotCost[1]
					next.GeodeRobots += 1
					queue.Push(next)
					continue
				}
			}

			// Always build another obsidian robot if we can.
			if state.Ore >= blueprint.ObsidianRobotCost[0] {
				if state.Clay >= blueprint.ObsidianRobotCost[1] {
					// But only if we'd benefit from more obsidian per minute.
					if state.ObsidianRobots < maxObsidianCost {
						next := base
						next.Ore -= blueprint.ObsidianRobotCost[0]
						next.Clay -= blueprint.ObsidianRobotCost[1]
						next.ObsidianRobots += 1
						queue.Push(next)
						continue
					}
				}
			}

			// We could choose to build another ore robot.
			if state.Ore >= blueprint.OreRobotCost {
				// But only if we'd benefit from more ore per minute.
				if state.OreRobots < maxOreCost {
					next := base
					next.Ore -= blueprint.OreRobotCost
					next.OreRobots += 1
					queue.Push(next)
				}
			}

			// We could choose to build another clay robot.
			if state.Ore >= blueprint.ClayRobotCost {
				// But only if we'd benefit from more clay per minute.
				if state.OreRobots < maxClayCost {
					next := base
					next.Ore -= blueprint.ClayRobotCost
					next.ClayRobots += 1
					queue.Push(next)
				}
			}

			// We could choose to do nothing.
			queue.Push(base)
		}
	}

	return maximumGeodes
}

func (problem *Problem) SolvePart1() error {
	var sum int64 = 0
	for i, blueprint := range problem.Blueprints {
		quality := int64(i+1) * OpenGeodes(blueprint, 24)
		sum += quality
	}

	fmt.Printf("Part 1: %d\n", sum)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var product int64 = 1
	for i := 0; i < 3; i++ {
		blueprint := problem.Blueprints[i]
		product *= OpenGeodes(blueprint, 32)
	}

	fmt.Printf("Part 2: %d\n", product)
	return nil
}
