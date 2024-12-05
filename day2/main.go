package main

import (
	"bufio"
	"cewkrupa/aoc2024go/util"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	util.Init(one, two)
}

type Report struct{}

func one(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		safe := true

		// this is just parsing the middle of the array, so it's not getting to the edges
		// we could maybe stop the loop after we know a level isn't safe
		for i := 1; i < len(levels)-1; i++ {
			l := parseLevel(
				levels[i],
				levels[i-1],
				levels[i+1],
			)

			fmt.Println(l)

			absDistPrev := math.Abs(float64(l.distancePrev()))
			absDistNext := math.Abs(float64(l.distanceNext()))
			adjacentRule := absDistPrev >= 1 && absDistPrev <= 3 && absDistNext >= 1 && absDistNext <= 3

			if safe {
				safe = adjacentRule
			}
		}
		fmt.Printf("Safe: %v\n", safe)
	}

	safe := 0
	fmt.Printf("Number of safe reports: %d\n", safe)

}

type ParsedLevel struct {
	lvl  int
	prev int
	next int
}

func (p ParsedLevel) distancePrev() int {
	return p.prev - p.lvl
}
func (p ParsedLevel) distanceNext() int {
	return p.next - p.lvl
}

func (p ParsedLevel) String() string {
	return fmt.Sprintf("\nLvl: %d\nPrev: %d\nNext:%d\n",
		p.lvl,
		p.distancePrev(),
		p.distanceNext(),
	)
}

func parseLevel(lvl, prev, next string) ParsedLevel {
	l, err := strconv.Atoi(lvl)
	util.Err(err)

	p, err := strconv.Atoi(prev)
	util.Err(err)

	n, err := strconv.Atoi(next)
	util.Err(err)

	return ParsedLevel{
		lvl:  l,
		prev: p,
		next: n,
	}
}

func checkReports(rps []Report) int {
	return 0
}

func two(inputFile *os.File) {
	panic("not impelmented")
}
