package ttg

import "fmt"

var Plays [3][3]string

func DrawBoard() {

	Plays[0][1] = "x"
	fmt.Println(Plays)
	CheckPlays()
}

func CheckPlays() {
	p := &Plays
	str := ``

	for i, row := range p {
		for j, col := range row {
			if col == "o" {
				str += "O"
			} else if col == "x" {
				str += "X"
			} else {
				str += " "
			}

			if j == 2 {
				str += "\n"
			} else {
				str += "|"
			}
			fmt.Println(str)
		}

		if i < 2 {
			str += "---+---+---\n"
		}
	}

	fmt.Println(str)
}
