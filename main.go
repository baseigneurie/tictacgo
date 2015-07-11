package main

import "fmt"

const cellCnt = 9

var board = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

var avlPlays []int

type grid struct {
	row1 [3]int
	row2 [3]int
	row3 [3]int
	col1 [3]int
	col2 [3]int
	col3 [3]int
}

func main() {
	board[4] = 1
	board[5] = 1
	board[0] = 2
	board[1] = 2
	g := initGrid()
	initGrid()
	checkForRows()
	checkForCols()
	fmt.Println(board)
}

func initGrid() grid {
	g := {
		row1: [3]row1{{0, 0, 0}},
		row2: {0, 0, 0},
		row3: {0, 0, 0},
		col1: {0, 0, 0},
		col2: {0, 0, 0},
		col3: {0, 0, 0}
	}
	return g
}

func checkForRows() {
	for i := 0; i < 3; i++ {
		switch {
		case i == 0:
			x, n := spotSearch(board[0], board[1], board[2])
			if x {
				avlPlays[0] = n
			}
			fmt.Println(i, x, n)
		case i == 1:
			x, n := spotSearch(board[3], board[4], board[5])
			if x {
				avlPlays[1] = n
			}
			fmt.Println(i, x, n)
		case i == 2:
			x, n := spotSearch(board[6], board[7], board[8])
			if x {
				avlPlays[3] = n
			}
			fmt.Println(i, x, n)
		default:
			x, n := false, 0
			fmt.Println(i, x, n)
		}
	}
}

func checkForCols() {
	for i := 0; i < 3; i++ {
		switch {
		case i == 0:
			x, n := spotSearch(board[0], board[3], board[6])
			if x {
				avlPlays[3] = n
			}
			fmt.Println(i, x, n)
		case i == 1:
			x, n := spotSearch(board[1], board[4], board[7])
			if x {
				avlPlays[4] = n
			}
			fmt.Println(i, x, n)
		case i == 2:
			x, n := spotSearch(board[2], board[5], board[8])
			if x {
				avlPlays[5] = n
			}
			fmt.Println(i, x, n)
		default:
			x, n := false, 0
			fmt.Println(i, x, n)
		}
	}
}

func spotSearch(c1, c2, c3 int) (bool, int) {
	n := (c1 + c2 + c3)
	if n == 2 {
		return true, findOpenSpot(c1, c2, c3)
	} else if n == 4 {
		return true, findOpenSpot(c1, c2, c3)
	} else {
		return false, 0
	}
}

func findOpenSpot(a, b, c int) int {
	switch {
	case a == 0:
		return 1
	case b == 0:
		return 2
	case c == 0:
		return 3
	default:
		return 0
	}
}
