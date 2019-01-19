package main

import (
	"fmt"

	"github.com/baseigneurie/tictacgo/ttg"
)

func main() {
	dash := ttg.DrawDash()
	board := ttg.DrawBoard()
	p := ttg.DrawPrompt()

	fmt.Printf("\n\n%s%s%s", dash, board, p)
}
