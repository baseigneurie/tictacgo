package ttg

var Plays [3][3]string
var Moves = 0
var CurrentPlayer string

var GameError string
var Winner = false

func RegPlay(s string) {
	switch s {
	case "a1":
		Plays[0][0] = CurrentPlayer
	case "a2":
		Plays[0][1] = CurrentPlayer
	case "a3":
		Plays[0][2] = CurrentPlayer
	case "b1":
		Plays[1][0] = CurrentPlayer
	case "b2":
		Plays[1][1] = CurrentPlayer
	case "b3":
		Plays[1][2] = CurrentPlayer
	case "c1":
		Plays[2][0] = CurrentPlayer
	case "c2":
		Plays[2][1] = CurrentPlayer
	case "c3":
		Plays[2][2] = CurrentPlayer
	}

	SwitchPlayer()
	Moves++
}

func SwitchPlayer() {
	if CurrentPlayer == "x" {
		CurrentPlayer = "o"
	} else {
		CurrentPlayer = "x"
	}
}
