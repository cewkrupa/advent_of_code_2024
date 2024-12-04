package main

import (
	"cewkrupa/aoc2024go/util"
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/bitfield/script"
)

func main() {
	util.InitScript(one, two)
}

func one(filename string) {
	input := script.File(filename)
	left, err := input.Column(1).Slice()

	if err != nil {
		panic(err)
	}

	input = script.File(filename)
	right, err := input.Column(2).Slice()
	if err != nil {
		panic(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	if len(left) != len(right) {
		panic("List sizes don't match")
	}

	sum := 0
	for i, lv := range left {
		rv := right[i]

		li, err := strconv.ParseFloat(lv, 32)
		if err != nil {
			panic(err)
		}
		ri, err := strconv.ParseFloat(rv, 32)
		if err != nil {
			panic(err)
		}
		diff := math.Abs(li - ri)
		sum += int(diff)
	}

	fmt.Printf("Total distance: %d\n", sum)

}

func two(filename string) {
	right, err := script.File(filename).Column(2).Freq().Slice()
	if err != nil {
		panic(err)
	}

	rightNumbers, err := script.Slice(right).Column(2).Slice()
	if err != nil {
		panic(err)
	}

	rightFreq, err := script.Slice(right).Column(1).Slice()
	if err != nil {
		panic(err)
	}

	if len(rightNumbers) != len(rightFreq) {
		panic("list sizes don't match")
	}

	left, err := script.File(filename).Column(1).Slice()
	if err != nil {
		panic(err)
	}

	score := 0

	for _, n := range left {
		i := slices.Index(rightNumbers, n)
		if i < 0 {
			// not present
			continue
		}

		freq, err := strconv.Atoi(rightFreq[i])
		if err != nil {
			panic(err)
		}

		nInt, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		score += freq * nInt
	}

	fmt.Printf("score: %d\n", score)
}
