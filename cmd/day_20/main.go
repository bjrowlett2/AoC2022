package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var err error

	var problem1 *Problem
	if problem1, err = Load(); err != nil {
		log.Fatal(err)
	}

	if err = problem1.SolvePart1(); err != nil {
		log.Fatal(err)
	}

	var problem2 *Problem
	if problem2, err = Load(); err != nil {
		log.Fatal(err)
	}

	if err = problem2.SolvePart2(); err != nil {
		log.Fatal(err)
	}
}

type Node struct {
	Index    int64
	Number   int64
	Next     *Node
	Previous *Node
}

type Problem struct {
	EncryptedFile []Node
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_20.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		EncryptedFile: make([]Node, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var number int64
		if _, err = fmt.Sscanf(line, "%d", &number); err != nil {
			return nil, err
		}

		// We just store the number for now, we set
		// the index and hook up the linked list below.
		node := Node{
			Number: number,
		}

		problem.EncryptedFile = append(problem.EncryptedFile, node)
	}

	n := int64(len(problem.EncryptedFile))
	for i := range problem.EncryptedFile {
		next := Wrap(int64(i+1), n)
		previous := Wrap(int64(i-1), n)

		problem.EncryptedFile[i].Index = int64(i)
		problem.EncryptedFile[i].Next = &problem.EncryptedFile[next]
		problem.EncryptedFile[i].Previous = &problem.EncryptedFile[previous]
	}

	return &problem, nil
}

func Wrap(i, size int64) int64 {
	result := i % size
	if result < 0 {
		result += size
	}

	return result
}

func Walk(node *Node, moves int64) *Node {
	var i int64
	pointer := node
	for i = 0; i < moves; i++ {
		pointer = pointer.Next
	}

	return pointer
}

func Mix(node *Node, size int64) {
	// We wrap with size-1 because we're removing this
	// node to do the walk, then putting it back in place.
	moves := Wrap(node.Number, size-1)

	if moves > 0 {
		// Remove this node.
		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous

		pointer := Walk(node, moves)

		// Put this node back in place.
		node.Next = pointer.Next
		node.Previous = pointer
		pointer.Next.Previous = node
		pointer.Next = node
	}
}

func (problem *Problem) SolvePart1() error {
	size := int64(len(problem.EncryptedFile))

	for idx := int64(0); idx < size; idx++ {
		Mix(&problem.EncryptedFile[idx], size)
	}

	for _, node := range problem.EncryptedFile {
		if node.Number == 0 {
			node1000 := Walk(&node, 1000)    // 1000
			node2000 := Walk(node1000, 1000) // 2000
			node3000 := Walk(node2000, 1000) // 3000

			groveCoordinates := node1000.Number + node2000.Number + node3000.Number
			fmt.Printf("Part 1: %d\n", groveCoordinates)
			break
		}
	}

	return nil
}

func (problem *Problem) SolvePart2() error {
	var decryptionKey int64 = 811589153
	for i := 0; i < len(problem.EncryptedFile); i++ {
		problem.EncryptedFile[i].Number *= decryptionKey
	}

	size := int64(len(problem.EncryptedFile))

	for iter := 0; iter < 10; iter++ {
		for idx := int64(0); idx < size; idx++ {
			Mix(&problem.EncryptedFile[idx], size)
		}
	}

	for _, node := range problem.EncryptedFile {
		if node.Number == 0 {
			node1000 := Walk(&node, 1000)    // 1000
			node2000 := Walk(node1000, 1000) // 2000
			node3000 := Walk(node2000, 1000) // 3000

			groveCoordinates := node1000.Number + node2000.Number + node3000.Number
			fmt.Printf("Part 1: %d\n", groveCoordinates)
			break
		}
	}

	return nil
}
