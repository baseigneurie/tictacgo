package ttg

import (
	"bufio"
	"fmt"
	"os"
)

func DrawPrompt() string {
	t := playerQuestion()
	fmt.Println(t)
	return ""
}

func playerQuestion() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many players?: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	return text
}
