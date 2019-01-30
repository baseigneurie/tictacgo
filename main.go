package main

import (
	"fmt"

	"github.com/baseigneurie/tictacgo/ttg"
)

func main() {
	err := gameStart()
	if err != nil {
		fatalError(err)
	}
}

func gameStart() error {
	fmt.Printf("\n")
	for ttg.Winner == false {
		err := ttg.DrawDash()
		if err != nil {
			return err
		}

		err = ttg.DrawBoard()
		if err != nil {
			return err
		}

		err = ttg.DrawPrompt()
		if err != nil {
			ttg.ShowError(err.Error())
			return nil
		}

		ttg.CheckWin()

		ttg.SwitchPlayer()
	}

	return nil
}

func gameEnd() {

}

func fatalError(err error) {
	fmt.Printf("%s%s", ttg.Red("FATAL: "), err.Error())
}
