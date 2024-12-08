package main

import (
	"bufio"
	"cewkrupa/aoc2024go/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	util.Init(one, two)
}

func one(file *os.File) {
	scanner := bufio.NewScanner(file)
	sum := 0
	matches := 0

	r1 := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for scanner.Scan() {
		line := scanner.Text()
		instructions := r1.FindAllStringSubmatch(line, -1)

		for _, inst := range instructions {
			matches++

			//fmt.Printf("inst: %v\n", inst)

			f, err := strconv.Atoi(inst[1])
			util.Err(err)
			g, err := strconv.Atoi(inst[2])

			sum += f * g
		}
	}
	fmt.Printf("matches: %d\nsum: %d\n", matches, sum)
}

func two(file *os.File) {
	panic("not implemented!")
}
