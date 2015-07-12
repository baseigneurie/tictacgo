package main

import "fmt"

const cellCnt = 9

// 0 - '', 1 - 'X', 2 - 'O'
var plays = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

type grid struct {
	row1 [3]int
	row2 [3]int
	row3 [3]int
	col1 [3]int
	col2 [3]int
	col3 [3]int
}

func main() {
	plays[4] = 1
	plays[5] = 1
	plays[0] = 2
	plays[1] = 2
	var g grid
	loadGrid(&g)
	fmt.Println(g)
}

func loadGrid(g *grid) {
	for i := 0; i < 3; i++ {
		switch {
		case i == 0:
			g.row1 = [3]int{plays[0], plays[1], plays[2]}
			g.col1 = [3]int{plays[0], plays[3], plays[6]}
		case i == 1:
			g.row2 = [3]int{plays[3], plays[4], plays[5]}
			g.col2 = [3]int{plays[1], plays[4], plays[7]}
		case i == 2:
			g.row3 = [3]int{plays[6], plays[7], plays[8]}
			g.col3 = [3]int{plays[2], plays[5], plays[8]}
		}
	}
}
