package Day11

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
)

func Goooo() {
	fmt.Println("--------- DAY 11 ---------")

	//lines := tools.ConvertIntoInts(tools.ReadFile(("days/Day10/testInput.txt")))
	lines := tools.ReadFile(("days/Day11/testInput.txt"))
	//lines := tools.ReadFile(("days/Day11/input.txt"))

	w := newWaitingArea()
	w.addRow()

	initMatrix(lines, w)
	printState(w)
	fmt.Println()

	makeMoves(w)
	printState(w)
	fmt.Println()

	fmt.Println(countOcupied(w))
}

func makeMoves(w *waitingArea) {
	for {
		w.clearMovement()

		w.traverseAll(func(r, c int) {
			w.makeDecision(r, c)
		})

		w.flipMarked()

		//printState(waiting)

		if !w.peopleMoved() {
			break
		}
		//time.Sleep(time.Second)
		//fmt.Println()
	}
}

func initMatrix(lines []string, w *waitingArea) {
	// add 1st row above + 1 floor left + 1 floor right
	for i := 0; i < len(lines[0])+2; i++ {
		w.addSeat(string("."))
	}

	for _, line := range lines {
		w.addRow()
		w.addSeat(string(".")) //left flour
		for _, in := range line {
			w.addSeat(string(in))
		}
		w.addSeat(string(".")) //right flour
	}

	// add last empty row  + 1 floor left + 1 floor right
	w.addRow()
	for i := 0; i < len(lines[0])+2; i++ {
		w.addSeat(string("."))
	}
}

func printState(w *waitingArea) {
	for r := 0; r < w.rowsCount(); r++ {
		for c := 0; c < w.colCount(); c++ {
			fmt.Print(w.printSeat(r, c))
		}
		fmt.Println()
	}
}

func countOcupied(w *waitingArea) int {
	count := 0
	w.traverseAll(func(r, c int) {
		if w.seats[r][c].state == occupied {
			count++
		}
	})
	return count
}
