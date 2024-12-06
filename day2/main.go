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

func one(file *os.File) {
	scanner := bufio.NewScanner(file)
	numSafe := 0

	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		increasing := false
		decreasing := false
		safe := true

		// this is just parsing the middle of the array, so it's not getting to the edges
		// we could maybe stop the loop after we know a level isn't safe
		for i := 1; i < len(levels); i++ {
			l := parseLevel(
				levels[i],
				levels[i-1],
			)
			absDistPrev := math.Abs(float64(l.distancePrev()))
			adjacentRule := absDistPrev >= 1 && absDistPrev <= 3

			if l.distancePrev() > 0 {
				increasing = true
			} else if l.distancePrev() < 0 {
				decreasing = true
			}

			if safe {
				safe = adjacentRule && increasing != decreasing
			}
		}
		if safe {
			numSafe++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", numSafe)

}

type ParsedLevel struct {
	lvl  int
	prev int
}

func (p ParsedLevel) distancePrev() int {
	return p.prev - p.lvl
}

func (p ParsedLevel) String() string {
	return fmt.Sprintf("\nLvl: %d\nPrev: %d\n",
		p.lvl,
		p.distancePrev(),
	)
}

func parseLevel(lvl, prev string) ParsedLevel {
	l, err := strconv.Atoi(lvl)
	util.Err(err)

	p, err := strconv.Atoi(prev)
	util.Err(err)

	return ParsedLevel{
		lvl:  l,
		prev: p,
	}
}

func two(file *os.File) {
	scanner := bufio.NewScanner(file)
	numSafe := 0

	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")

		firstCheck := checkLevels(levels)
		if firstCheck {
			numSafe++
			continue
		}

		for i := 0; i < len(levels); i++ {
			pre := levels[:i]
			post := levels[i+1:]
			subset := []string{}
			subset = append(subset, pre...)
			subset = append(subset, post...)

			check := checkLevels(subset)
			if check {
				numSafe++
				break
			}
		}
	}

	fmt.Printf("Number of safe reports: %d\n", numSafe)
}

func checkLevels(levels []string) bool {
	increasing := false
	decreasing := false
	safe := true
	for i := 1; i < len(levels); i++ {
		l := parseLevel(
			levels[i],
			levels[i-1],
		)
		absDistPrev := math.Abs(float64(l.distancePrev()))
		adjacentRule := absDistPrev >= 1 && absDistPrev <= 3

		if l.distancePrev() > 0 {
			increasing = true
		} else if l.distancePrev() < 0 {
			decreasing = true
		}

		if safe {
			newSafe := adjacentRule && increasing != decreasing

			safe = newSafe
		}
	}
	return safe
}
