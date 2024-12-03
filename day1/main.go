package main

import (
	"cewkrupa/aoc2024go/util"
	"fmt"
	"io"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	util.InitScript(one, two)
}

func one(filename string) {

	input := script.File(filename)

	var leftList []int
	var rightList []int

	out, err := input.First(10).FilterScan(func(line string, w io.Writer) {
		split := strings.Split(line)
		fmt.Fprintf(w, "scanned line: %q\n", line)
	}).String()

	fmt.Println(out)
	if err != nil {
		panic(err)
	}

	// regex out the two matches

	// sort the two lists into ascending pairs
	// keep running sum of differences
	// for each pair
	//   add absolute difference of pairs to sum
	//

}

func two(filename string) {
	panic("not implemented")
}
