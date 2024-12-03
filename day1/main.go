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

	// read input in
	// we can read the input in first and then sort, OR we can split the lists and then sort them
	// sort left list
	// sort right list

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

	// sort the two lists into ascending pairs

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
	// get left list
	// get right list
	// for each in left list
	//
	input := script.File(filename)
	right, err := input.Column(2).Freq().Slice()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", right)
}
