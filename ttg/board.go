package ttg

import "fmt"

var moves [3][3]string

var board [17]string

type Quadrant struct {
	Row [5]string
	Col [5]string
}

func DrawBoard() {
	b := &board
	n := &NoMarker
	str := ""
	for i, r := range b {
		if i == 6 {
			str += "#"
		}

		if i == 12 {
			str += "#"
		}

	}

	fmt.Println(str)

}
