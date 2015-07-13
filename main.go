package main

import "fmt"

const cellCnt = 9

// 0 - '', 1 - 'X', 2 - 'O'
// This will be the main object that will be passed between the browser and server.
var plays = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

type grid struct {
	row1  [3]int
	row2  [3]int
	row3  [3]int
	col1  [3]int
	col2  [3]int
	col3  [3]int
	diag1 [3]int
	diag2 [3]int
}

type openSpot struct {
	row  int
	cell int
}

func main() {
	plays[4] = 1
	plays[5] = 1
	plays[0] = 2
	plays[1] = 2
	var g grid
	loadGrid(&g)
	op := findOpen(&g, 2)
	fmt.Println(g)
	fmt.Println(op)
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
	g.diag1 = [3]int{plays[0], plays[4], plays[8]}
	g.diag1 = [3]int{plays[2], plays[4], plays[6]}
}

func findOpen(g *grid, x int) []openSpot {
	var o []openSpot

	for i := 0; i < 3; i++ {
		switch {
		case i == 0:
			r, v := searchRow(g.row1, x)
			if r {
				o = append(o, openSpot{1, v})
			}
			r, v = searchCol(g.col1, x)
			if r {
				o = append(o, openSpot{1, v})
			}
		case i == 1:
			r, v := searchRow(g.row1, x)
			if r {
				o = append(o, openSpot{2, v})
			}
			r, v = searchCol(g.col1, x)
			if r {
				o = append(o, openSpot{2, v})
			}
		case i == 2:
			r, v := searchRow(g.row1, x)
			if r {
				o = append(o, openSpot{2, v})
			}
			r, v = searchCol(g.col1, x)
			if r {
				o = append(o, openSpot{2, v})
			}
		}
	}
	r, v := searchRow(g.diag1, x)
	if r {
		o = append(o, openSpot{1, v})
	}
	r, v = searchCol(g.diag2, x)
	if r {
		o = append(o, openSpot{1, v})
	}

	return o
}

func searchRow(r [3]int, x int) (bool, int) {
	n, c := 0, 0

	for i, v := range r {
		n += v
		if v == 0 {
			c = i + 1
		}
	}
	if n == x {
		return true, c
	}
	return false, 0
}

func searchCol(r [3]int, x int) (bool, int) {
	n, c := 0, 0

	for i, v := range r {
		n += v
		if v == 0 {
			c = i + 1
		}
	}
	if n == x {
		return true, c
	}
	return false, 0
}

func searchDiag(r [3]int, x int) (bool, int) {
	n, c := 0, 0

	for i, v := range r {
		n += v
		if v == 0 {
			c = i + 1
		}
	}
	if n == x {
		return true, c
	}
	return false, 0
}
