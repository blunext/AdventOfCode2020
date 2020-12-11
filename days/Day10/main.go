package Day10

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"sort"
)

func Goooo() {
	fmt.Println("--------- DAY 09 ---------")

	//lines := tools.ConvertIntoInts(tools.ReadFile(("days/Day10/testInput.txt")))
	lines := tools.ConvertIntoInts(tools.ReadFile(("days/Day10/input.txt")))

	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})

	differences := findDifferences(lines)
	a1 := differences[1]
	a3 := differences[3]

	fmt.Printf("part 1: a1 = %d, a2 = %d, * = %d\n", a1, a3, a1*a3)

	// count in excel using tribonacci secuence
	fmt.Printf("part 2: %d", 7*7*7*7*7*2*4*7*7*2*4*4*4*4*4*4*4*7*4*7*4)
}

func findDifferences(lines []int) map[int]int {
	lines = append(lines, lines[len(lines)-1]+3) // last + 3 as a max joilts rate

	differences := make(map[int]int)
	joltage := 0
	for _, v := range lines {
		adapter := v - joltage
		if _, ok := differences[adapter]; ok {
			differences[adapter]++
		} else {
			differences[adapter] = 1
		}
		joltage = v
	}
	return differences
}
