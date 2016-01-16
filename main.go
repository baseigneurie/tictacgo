package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	//maxX, maxY := g.Size()
	if v, err := g.SetView("corner", 0, 0, 10, 10); err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		v.Wrap = true
		fmt.Fprintln(v, "X       X X     X   X   X     X X       X")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.Quit
}

func main() {
	var err error
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.SetLayout(layout)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	err = g.MainLoop()
	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}
