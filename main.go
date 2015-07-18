package main

import (
	"fmt"
	"math/rand"
	"time"
)

const cellCnt, xCheck, oCheck, xWin, oWin = 9, 2, 10, 3, 15

// This will be the main object that will be passed between the browser and server.
// 0 - "", 1 - X, 5 - O
var played = []int{1, 0, 0, 0, 5, 0, 5, 0, 0}

var turn, mode, comp, player = "", "", "", ""

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
	assignPlayer()
	turn = "o"
	mode = "play"
	xMoves, oMoves := moves{}, moves{}
	if player == turn && mode == "play" {
		if !playRound() {
			// Player wins, game over
		} else {
			endRound()
		}
	} else if comp == turn && mode == "play" {
		//Computer first checks availale moves
		findMoves(&xMoves, xCheck)
		findMoves(&oMoves, oCheck)
		playMove(&xMoves, &oMoves)
		if !playRound() {
			// Computer wins, game over
		} else {
			endRound()
		}
	}
	fmt.Println(xMoves, oMoves)
}

func assignPlayer() {
	player = "x"
	comp = "o"
}

func playRound() bool {
	if turn == "x" {
		r, w := findWin(xWin)
		if r {
			fmt.Println(w)
			mode = "win"
			return false
		}
	} else if turn == "o" {
		r, w := findWin(oWin)
		if r {
			fmt.Println(w)
			mode = "win"
			return false
		}
	}
	return true
}

func playMove(x *moves, o *moves) {
	mark := 0
	num := rand.New(rand.NewSource(time.Now().UnixNano()))
	if comp == "x" {
		mark = 1
		if len(x.cnt) > 0 {
			//Plays a move to win
			i := num.Intn(len(x.cnt))
			played[x.cnt[i]] = mark
		} else if len(o.cnt) > 0 {
			//Plays a move to block
			i := num.Intn(len(o.cnt))
			played[o.cnt[i]] = mark
		} else {
			//Makes no play
			return
		}
	} else if comp == "o" {
		mark = 5
		if len(o.cnt) > 0 {
			//Plays a move to win
			i := num.Intn(len(o.cnt))
			played[o.cnt[i]] = mark
		} else if len(x.cnt) > 0 {
			//Plays a move to block
			i := num.Intn(len(x.cnt))
			played[x.cnt[i]] = mark
		} else {
			//Makes no play
			return
		}
	}

}

func findMoves(m *moves, n int) {
	for _, v := range combo {
		sum := played[v[0]] + played[v[1]] + played[v[2]]
		if sum == n {
			o := (findOpen(v))
			if checkDup(m, o) {
				m.cnt = append(m.cnt, o)
			}
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

func checkDup(m *moves, n int) bool {
	for i := 0; i < len(m.cnt); i++ {
		if n == m.cnt[i] {
			return false
		}
	}
	return true
}

func findWin(n int) (bool, []int) {
	for _, v := range combo {
		sum := played[v[0]] + played[v[1]] + played[v[2]]
		if sum == n {
			win := []int{v[0], v[1], v[2]}
			return true, win
		}
	}
	return false, []int{0, 0, 0}
}

func endRound() {
	// This will end the round and return to the view

}
