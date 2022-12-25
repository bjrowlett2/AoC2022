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
}

type Problem struct {
	Numbers []string
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_25.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Numbers: make([]string, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		problem.Numbers = append(problem.Numbers, line)
	}

	return &problem, nil
}

func ToDecimal(snafu string) int64 {
	var decimal int64 = 0
	digits := map[rune]int64{
		'2': 2,
		'1': 1,
		'0': 0,
		'-': -1,
		'=': -2,
	}

	n := int64(len(snafu))
	for i, r := range snafu {
		k := int64(i)
		decimal += digits[r] * aoc.Pow64(5, n-k-1)
	}

	return decimal
}

func ToSnafu(decimal int64) string {
	snafu := ""
	for decimal > 0 {
		digit := decimal % 5
		if digit <= 2 {
			snafu = fmt.Sprintf("%d%s", digit, snafu)
		} else {
			digits := map[int64]rune{
				4: '-',
				3: '=',
			}

			s := string(digits[digit])
			snafu = fmt.Sprintf("%s%s", s, snafu)
			decimal += digit
		}

		decimal /= 5
	}

	return snafu
}

func (problem *Problem) SolvePart1() error {
	var sum int64 = 0
	for _, snafu := range problem.Numbers {
		sum += ToDecimal(snafu)
	}

	snafu := ToSnafu(sum)
	fmt.Printf("Part 1: %s\n", snafu)
	return nil
}
