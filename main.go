package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What's your name?")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello " + text)
}
