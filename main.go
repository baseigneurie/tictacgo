package main

import "fmt"

const cellCnt = 9

// This will be the main object that will be passed between the browser and server.
var played = []int{1, 1, 0, 0, 2, 2, 0, 0, 0}

var turn = ""

var combo = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

type moves struct {
	cnt []int
}

func main() {
	turn = "x"
	xMoves, oMoves := moves{}, moves{}
	findMoves(&xMoves, 2)
	findMoves(&oMoves, 4)
	fmt.Println(xMoves, oMoves)
}

func findMoves(m *moves, n int) {
	for _, v := range combo {
		sum := played[v[0]] + played[v[1]] + played[v[2]]
		if sum == n {
			m.cnt = append(m.cnt, findOpen(v))
		}
	}
}

func findOpen(i [3]int) int {
	for _, v := range i {
		if played[v] == 0 {
			return v
		}
	}
	return 0
}
