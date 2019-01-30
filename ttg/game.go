package ttg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Plays [3][3]string
var Moves = 0
var CurrentPlayer = "x"

var Winner = false

func DrawPrompt() error {
	b := false

	for b == false {
		r, err := makePlay()
		if err != nil {
			return err
		}
		b = r
	}

	return nil
}

func makePlay() (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your move (Ex. 'a1', 'b3'): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	b := regPlay(text)
	return b, nil
}

func regPlay(s string) bool {
	switch s {
	case "a1":
		if b := spaceCheck(Plays[0][0]); b == false {
			return b
		}
		Plays[0][0] = CurrentPlayer
	case "a2":
		if b := spaceCheck(Plays[1][0]); b == false {
			return b
		}
		Plays[1][0] = CurrentPlayer
	case "a3":
		if b := spaceCheck(Plays[2][0]); b == false {
			return b
		}
		Plays[2][0] = CurrentPlayer
	case "b1":
		if b := spaceCheck(Plays[0][1]); b == false {
			return b
		}
		Plays[0][1] = CurrentPlayer
	case "b2":
		if b := spaceCheck(Plays[1][1]); b == false {
			return b
		}
		Plays[1][1] = CurrentPlayer
	case "b3":
		if b := spaceCheck(Plays[2][1]); b == false {
			return b
		}
		Plays[2][1] = CurrentPlayer
	case "c1":
		if b := spaceCheck(Plays[0][2]); b == false {
			return b
		}
		Plays[0][2] = CurrentPlayer
	case "c2":
		if b := spaceCheck(Plays[1][2]); b == false {
			return b
		}
		Plays[1][2] = CurrentPlayer
	case "c3":
		if b := spaceCheck(Plays[2][2]); b == false {
			return b
		}
		Plays[2][2] = CurrentPlayer
	default:
		ShowError("Please select a move inbounds")
		return false
	}

	return true
}

func spaceCheck(s string) bool {
	if s != "" {
		ShowError("That space is occupied. Please select another.")
		return false
	}
	return true
}

func SwitchPlayer() {
	if CurrentPlayer == "x" {
		CurrentPlayer = "o"
	} else {
		CurrentPlayer = "x"
	}
}

func CheckWin() {
	count := 0
	for _, v := range Plays {
		for n := 0; n < 3; n++ {
			if v[n] == CurrentPlayer {
				count++
			}
		}
		if count == 3 {
			setWin()
			return
		}
	}

	// Check columns
	count = 0
	for n := 0; n < 3; n++ {
		for _, v := range Plays {
			if v[n] == CurrentPlayer {
				count++
			}
		}
		if count == 3 {
			setWin()
			return
		}
	}

	// // rows
	// if Plays[0][0] == CurrentPlayer &&
	// 	Plays[0][1] == CurrentPlayer &&
	// 	Plays[0][2] == CurrentPlayer {
	// 	setWin()
	// 	return
	// }

	// if Plays[1][0] == CurrentPlayer &&
	// 	Plays[1][1] == CurrentPlayer &&
	// 	Plays[1][2] == CurrentPlayer {
	// 	setWin()
	// 	return
	// }

	// if Plays[2][0] == CurrentPlayer &&
	// 	Plays[2][1] == CurrentPlayer &&
	// 	Plays[2][2] == CurrentPlayer {
	// 	setWin()
	// 	return
	// }

	// // cols
	// if Plays[0][0] == CurrentPlayer &&
	// 	Plays[1][0] == CurrentPlayer &&
	// 	Plays[2][0] == CurrentPlayer {
	// 	setWin()
	// 	return
	// }

	// if Plays[0][1] == CurrentPlayer &&
	// 	Plays[1][1] == CurrentPlayer &&
	// 	Plays[2][1] == CurrentPlayer {
	// 	setWin()
	// 	return
	// }

	// if Plays[0][2] == CurrentPlayer &&
	// 	Plays[1][2] == CurrentPlayer &&
	// 	Plays[2][2] == CurrentPlayer {
	// 	setWin()
	// 	return
	// }

	// Check diagonals
	// count = 0
	if Plays[0][0] == CurrentPlayer &&
		Plays[1][1] == CurrentPlayer &&
		Plays[2][2] == CurrentPlayer {
		setWin()
		return
	}

	if Plays[0][2] == CurrentPlayer &&
		Plays[1][1] == CurrentPlayer &&
		Plays[2][0] == CurrentPlayer {
		setWin()
		return
	}
}

func setWin() {
	Winner = true
	fmt.Println("WHOOOP")
}

func ShowError(msg string) {
	s := "\nERR: " + msg + "\n"
	fmt.Printf(ColError(s))
}
