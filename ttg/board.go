package ttg

import "fmt"

type Moves struct {
	A1 string
	A2 string
	A3 string

	B1 string
	B2 string
	B3 string

	C1 string
	C2 string
	C3 string
}

var board [17]string

func DrawBoard() {
	// b := &board
	// n := &NoMarker
	mv := Moves{}

	// str := ""
	// for i, r := range b {
	// 	if i == 6 {
	// 		str += "#"
	// 	}

	// 	if i == 12 {
	// 		str += "#"
	// 	}

	// }

	fmt.Println(str)

}
