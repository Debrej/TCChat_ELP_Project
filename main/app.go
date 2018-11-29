package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	testString, _ := reader.ReadString('\n')

	msgParams := ParseServer(testString)
	for i := 0; i < len(msgParams); i++ {
		fmt.Print(msgParams[i])
	}
}
