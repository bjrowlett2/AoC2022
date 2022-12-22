package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

type MonkeyType int

const (
	MonkeyTypeNumber MonkeyType = iota
	MonkeyTypeFormula
)

type Expression struct {
	Left     string
	Right    string
	Operator string
}

type Monkey struct {
	Name    string
	Type    MonkeyType
	Number  int64
	Formula Expression
}

type Problem struct {
	Monkeys map[string]Monkey
}

func IsExpression(s string) bool {
	return strings.ContainsAny(s, "+-*/")
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_21.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Monkeys: make(map[string]Monkey),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		parts[0] = strings.TrimSpace(parts[0])
		parts[1] = strings.TrimSpace(parts[1])

		monkey := Monkey{
			Name: parts[0],
		}

		if !IsExpression(parts[1]) {
			monkey.Type = MonkeyTypeNumber
			if _, err = fmt.Sscanf(parts[1], "%d", &monkey.Number); err != nil {
				return nil, err
			}
		} else {
			formula := &monkey.Formula
			monkey.Type = MonkeyTypeFormula
			if _, err = fmt.Sscanf(parts[1], "%s %s %s", &formula.Left, &formula.Operator, &formula.Right); err != nil {
				return nil, err
			}
		}

		problem.Monkeys[monkey.Name] = monkey
	}

	return &problem, nil
}

func Yell(problem *Problem, name string) int64 {
	monkey := problem.Monkeys[name]

	var number int64 = 0
	switch monkey.Type {
	case MonkeyTypeNumber:
		return monkey.Number

	case MonkeyTypeFormula:
		formula := monkey.Formula
		left := Yell(problem, formula.Left)
		right := Yell(problem, formula.Right)

		switch formula.Operator {
		case "+":
			number = left + right
		case "-":
			number = left - right
		case "*":
			number = left * right
		case "/":
			number = left / right
		}
	}

	return number
}

func (problem *Problem) SolvePart1() error {
	root := Yell(problem, "root")
	fmt.Printf("Part 1: %d\n", root)
	return nil
}

type Side int

const (
	SideUnknown Side = iota
	SideLeft
	SideRight
)

func FindHuman(problem *Problem, name string) Side {
	monkey := problem.Monkeys[name]

	if monkey.Type == MonkeyTypeFormula {
		formula := monkey.Formula
		if formula.Left == "humn" {
			return SideLeft
		} else if formula.Right == "humn" {
			return SideRight
		}

		if FindHuman(problem, formula.Left) != SideUnknown {
			return SideLeft
		} else if FindHuman(problem, formula.Right) != SideUnknown {
			return SideRight
		}
	}

	return SideUnknown
}

func (problem *Problem) SolvePart2() error {
	// First, you got the wrong job for the monkey named root;
	// specifically, you got the wrong math operation. The correct
	// operation for monkey root should be =.
	root := problem.Monkeys["root"]
	problem.Monkeys["root"] = Monkey{
		Type:   root.Type,
		Number: root.Number,
		Formula: Expression{
			Left:     root.Formula.Left,
			Right:    root.Formula.Right,
			Operator: "=",
		},
	}

	name := "root"
	var answer int64 = 0
	for name != "humn" {
		monkey := problem.Monkeys[name]

		var result int64
		formula := monkey.Formula
		switch FindHuman(problem, name) {
		case SideLeft:
			name = formula.Left
			result = Yell(problem, formula.Right)

			switch formula.Operator {
			case "+":
				answer = answer - result
			case "-":
				answer = answer + result
			case "*":
				answer = answer / result
			case "/":
				answer = answer * result
			case "=":
				answer = result
			}
		case SideRight:
			name = formula.Right
			result = Yell(problem, formula.Left)

			switch formula.Operator {
			case "+":
				answer = answer - result
			case "-":
				answer = result - answer
			case "*":
				answer = answer / result
			case "/":
				answer = result / answer
			case "=":
				answer = result
			}
		}
	}

	fmt.Printf("Part 2: %d\n", answer)
	return nil
}
