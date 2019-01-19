package ttg

import (
	"strconv"

	"github.com/fatih/color"
)

func DrawBoard() string {

	p := &Plays
	p[0][0] = "x"
	p[0][1] = "x"
	p[0][2] = "x"
	p[2][2] = "o"

	return loadBoard()
}

func loadBoard() string {
	p := &Plays

	str := "   A | B | C \n\n"

	for i, row := range p {
		str += strconv.Itoa(i+1) + " "

		for j, col := range row {
			if col == "o" {
				str += Green(" O ")
			} else if col == "x" {
				str += Blue(" X ")
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

func Red(s string) string {
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	return red(s)
}

func Blue(s string) string {
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	return blue(s)
}

func Green(s string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	return green(s)
}

func Yellow(s string) string {
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	return yellow(s)
}

func ColError(s string) string {
	ecol := color.New(color.FgWhite, color.BgRed).SprintFunc()
	return ecol(s)
}
