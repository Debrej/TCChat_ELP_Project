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
	msg := ParseServer(testString)
	fmt.Print(msg)
}
