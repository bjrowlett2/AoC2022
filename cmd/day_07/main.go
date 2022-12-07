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

type File struct {
	Name string
	Size int64
}

type Folder struct {
	Name    string
	Files   []File
	Folders []Folder
}

func (folder *Folder) Size() int64 {
	var total int64 = 0
	for _, s := range folder.Files {
		total += s.Size
	}

	for _, s := range folder.Folders {
		total += s.Size()
	}

	return total
}

func (folder *Folder) Sum(total *int64, limit int64) {
	size := folder.Size()
	if size <= limit {
		*total += size
	}

	for _, s := range folder.Folders {
		s.Sum(total, limit)
	}
}

func (folder *Folder) Delete(results *[]int64, needed int64) {
	size := folder.Size()
	if size >= needed {
		*results = append(*results, size)
	}

	for _, s := range folder.Folders {
		s.Delete(results, needed)
	}
}

func (folder *Folder) AddFile(path []string, name string, size int64) {
	if len(path) == 0 {
		x := File{
			Name: name,
			Size: size,
		}

		folder.Files = append(folder.Files, x)
	} else {
		for i := 0; i < len(folder.Folders); i++ {
			s := &folder.Folders[i]
			if path[0] == s.Name {
				s.AddFile(path[1:], name, size)
			}
		}
	}
}

func (folder *Folder) AddFolder(path []string, name string) {
	if len(path) == 0 {
		x := Folder{
			Name:    name,
			Files:   make([]File, 0),
			Folders: make([]Folder, 0),
		}

		folder.Folders = append(folder.Folders, x)
	} else {
		for i := 0; i < len(folder.Folders); i++ {
			s := &folder.Folders[i]
			if path[0] == s.Name {
				s.AddFolder(path[1:], name)
			}
		}
	}
}

type Problem struct {
	Root Folder
}

func Load() (*Problem, error) {
	var err error

	var file *os.File
	name := "day_07.txt"
	if file, err = os.Open(name); err != nil {
		return nil, err
	}

	defer file.Close()

	problem := Problem{
		Root: Folder{},
	}

	cwd := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ cd /" {
			cwd = []string{}
		} else if line == "$ cd .." {
			n := len(cwd)
			cwd = cwd[:n-1]
		} else if strings.HasPrefix(line, "$ cd ") {
			dest := strings.TrimPrefix(line, "$ cd ")
			cwd = append(cwd, dest)
		} else {
			if line == "$ ls" {
				continue
			}

			if strings.HasPrefix(line, "dir ") {
				name := strings.TrimPrefix(line, "dir ")
				problem.Root.AddFolder(cwd, name)
			} else {
				var size int64
				var name string
				if _, err = fmt.Sscanf(line, "%d %s", &size, &name); err != nil {
					return nil, err
				}

				problem.Root.AddFile(cwd, name, size)
			}
		}
	}

	return &problem, nil
}

func (problem *Problem) SolvePart1() error {
	var total int64 = 0
	problem.Root.Sum(&total, 100000)
	fmt.Printf("Part 1: %d\n", total)
	return nil
}

func (problem *Problem) SolvePart2() error {
	needed := 30000000 - (70000000 - problem.Root.Size())

	results := make([]int64, 0)
	problem.Root.Delete(&results, needed)

	aoc.SortInt64(results)
	fmt.Printf("Part 2: %d\n", results[0])
	return nil
}
