package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

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

type ElementType int

const (
	DATA_TYPE_INT ElementType = iota
	DATA_TYPE_LIST
)

type Element struct {
	Type ElementType
	Int  int64
	List []Element
}

func (e Element) Equals(other Element) bool {
	if e.Type == other.Type {
		if e.Type == DATA_TYPE_INT {
			return e.Int == other.Int
		} else if e.Type == DATA_TYPE_LIST {
			m := len(e.List)
			n := len(other.List)

			if m == n {
				equals := true
				for i := 0; i < m; i++ {
					a := e.List[i]
					b := other.List[i]
					equals = equals && a.Equals(b)
				}

				return equals
			}
		}
	}

	return false
}

type Problem struct {
	Packets []Element
}

func Parse(line string) (*Element, error) {
	var err error

	root := Element{
		Type: DATA_TYPE_LIST,
		List: make([]Element, 0),
	}

	read := 0
	active := []*Element{&root}

	for i := 1; i < len(line)-1; i += read {
		read = 1
		x := line[i]

		if x == '[' {
			next := Element{
				Type: DATA_TYPE_LIST,
				List: make([]Element, 0),
			}

			a := len(active)
			current := active[a-1]
			current.List = append(current.List, next)

			b := len(current.List)
			successor := &current.List[b-1]
			active = append(active, successor)
		} else if x == ']' {
			a := len(active)
			active = active[:a-1]
		} else if x != ',' {
			var n int
			var value int64
			if n, err = fmt.Sscanf(line[i:], "%d", &value); err != nil {
				return nil, err
			}

			read = n

			next := Element{
				Type: DATA_TYPE_INT,
				Int:  value,
			}

			a := len(active)
			current := active[a-1]
			current.List = append(current.List, next)
		}
	}

	return active[0], nil
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_13.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Packets: make([]Element, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var left *Element
		if left, err = Parse(line); err != nil {
			return nil, err
		}

		scanner.Scan()
		line = scanner.Text()

		var right *Element
		if right, err = Parse(line); err != nil {
			return nil, err
		}

		scanner.Scan() // Blank line
		problem.Packets = append(problem.Packets, *left)
		problem.Packets = append(problem.Packets, *right)
	}

	return &problem, nil
}

type Result int

const (
	RESULT_ORDERED Result = iota
	RESULT_UNORDERED
	RESULT_UNKNOWN
)

func IsOrdered(left Element, right Element) Result {
	if (left.Type == DATA_TYPE_INT) && (right.Type == DATA_TYPE_INT) {
		if left.Int < right.Int {
			return RESULT_ORDERED
		} else if left.Int > right.Int {
			return RESULT_UNORDERED
		}
	} else if (left.Type == DATA_TYPE_LIST) && (right.Type == DATA_TYPE_LIST) {
		m := len(left.List)
		n := len(right.List)
		min := aoc.Min(m, n)

		for i := 0; i < min; i++ {
			lhs := left.List[i]
			rhs := right.List[i]

			result := IsOrdered(lhs, rhs)
			if result != RESULT_UNKNOWN {
				return result
			}
		}

		if m < n {
			return RESULT_ORDERED
		} else if m > n {
			return RESULT_UNORDERED
		}
	} else {
		if left.Type == DATA_TYPE_INT {
			next := Element{
				Type: DATA_TYPE_INT,
				Int:  left.Int,
			}

			left.Type = DATA_TYPE_LIST
			left.List = append(left.List, next)
		}

		if right.Type == DATA_TYPE_INT {
			next := Element{
				Type: DATA_TYPE_INT,
				Int:  right.Int,
			}

			right.Type = DATA_TYPE_LIST
			right.List = append(right.List, next)
		}

		return IsOrdered(left, right)
	}

	return RESULT_UNKNOWN
}

func (problem *Problem) SolvePart1() error {
	sum := 0
	for i := 0; i < len(problem.Packets); i += 2 {
		idx := i / 2

		left := problem.Packets[i]
		right := problem.Packets[i+1]
		if IsOrdered(left, right) == RESULT_ORDERED {
			sum += idx + 1
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
	return nil
}

func (problem *Problem) SolvePart2() error {
	var err error

	var divider2 *Element
	if divider2, err = Parse("[[2]]"); err != nil {
		return err
	}

	var divider6 *Element
	if divider6, err = Parse("[[6]]"); err != nil {
		return err
	}

	problem.Packets = append(problem.Packets, *divider2)
	problem.Packets = append(problem.Packets, *divider6)

	sort.Slice(problem.Packets, func(i, j int) bool {
		left := problem.Packets[i]
		right := problem.Packets[j]
		result := IsOrdered(left, right)
		return result == RESULT_ORDERED
	})

	decoderKey := 1
	for idx, packet := range problem.Packets {
		if packet.Equals(*divider2) {
			decoderKey *= idx + 1
		} else if packet.Equals(*divider6) {
			decoderKey *= idx + 1
		}
	}

	fmt.Printf("Part 2: %d\n", decoderKey)
	return nil
}
