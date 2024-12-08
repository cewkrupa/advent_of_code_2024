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

	scanner := bufio.NewScanner(file)
	r1 := regexp.MustCompile(`mul\((\d+),(\d+)\)|don?\'?t?\(\)`)

	do := "do("
	dont := "don"
	mult := `mul`
	doMatches := 0
	dontMatches := 0
	multMatches := 0

	sum := 0
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()
		instructions := r1.FindAllStringSubmatch(line, -1)

		for _, inst := range instructions {
			switch inst[0][0:3] {
			case do:
				{
					doMatches++
					enabled = true
				}
			case dont:
				{
					dontMatches++
					enabled = false
				}
			case mult:
				{
					if enabled {

						multMatches++
						f, err := strconv.Atoi(inst[1])
						util.Err(err)
						g, err := strconv.Atoi(inst[2])

						sum += f * g
					}
				}
			}

			//fmt.Printf("inst: %v\n", inst)
		}
	}
	fmt.Printf("do: %d\ndon't: %d\nmult: %d\nsum: %d\n", doMatches, dontMatches, multMatches, sum)
}
