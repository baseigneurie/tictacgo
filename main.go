package main

import (
	"fmt"

	"github.com/baseigneurie/tictacgo/ttg"
)

func main() {
	x, o := ttg.GetMarkers()
	for _, n := range x {
		fmt.Println(n)
	}

	for _, n := range o {
		fmt.Println(n)
	}
}
