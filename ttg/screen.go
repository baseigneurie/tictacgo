package ttg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func DrawBoard() error {
	var b strings.Builder
	b.WriteString("   A | B | C \n\n")

	for i, row := range Plays {
		_, err := fmt.Fprintf(&b, "%s ", strconv.Itoa(i+1))
		if err != nil {
			return err
		}

		for j, col := range row {
			if col == "o" {
				_, err := b.WriteString(Green(" O "))
				if err != nil {
					return err
				}
			} else if col == "x" {
				_, err := b.WriteString(Blue(" X "))
				if err != nil {
					return err
				}
			} else {
				_, err := b.WriteString("   ")
				if err != nil {
					return err
				}
			}

			if j == 2 {
				_, err := b.WriteString("\n")
				if err != nil {
					return err
				}
			} else {
				_, err := b.WriteString("|")
				if err != nil {
					return err
				}
			}
		}

		if i < 2 {
			_, err := b.WriteString("  ---+---+---\n")
			if err != nil {
				return err
			}
		}
	}

	fmt.Printf(b.String())
	return nil
}

func DrawDash() error {
	var b strings.Builder
	_, err := fmt.Fprintf(&b, "Current Player: %s\n\n", strings.ToUpper(CurrentPlayer))
	if err != nil {
		return err
	}

	fmt.Printf(Yellow(b.String()))
	return nil
}

// **
// Color Presets
// **
func Red(s string) string {
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	return red(s)
}

func Blue(s string) string {
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	return blue(s)
}

func Green(s string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	return green(s)
}

func Yellow(s string) string {
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	return yellow(s)
}

func ColError(s string) string {
	ecol := color.New(color.FgWhite, color.BgRed).SprintFunc()
	return ecol(s)
}
