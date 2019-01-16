package ttg

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var Plays [3][3]string
var boardLayout string

func DrawBoard() {

	Plays[0][0] = "x"
	Plays[0][1] = "x"
	Plays[0][2] = "x"
	Plays[2][2] = "o"

	// boardLayout = "Total Moves: XX   "
	// boardLayout += "Current Player: X\n\n"

	head := "Total Moves: XX   "
	head += "Current Player: X\n\n"

	boardLayout = loadBoard(boardLayout)
	d := color.New(color.FgRed, color.Bold)

	d.Printf(head)
	fmt.Printf(boardLayout)
}

func loadBoard(str string) string {
	p := &Plays
	str += "   A | B | C \n\n"

	for i, row := range p {
		str += strconv.Itoa(i+1) + " "

		for j, col := range row {
			if col == "o" {
				str += " O "
			} else if col == "x" {
				str += " X "
			} else {
				str += "   "
			}

			if j == 2 {
				str += "\n"
			} else {
				str += "|"
			}
		}

		if i < 2 {
			str += "  ---+---+---\n"
		}
	}

	return str
}
