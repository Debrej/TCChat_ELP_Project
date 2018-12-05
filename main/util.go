package main

import (
	"bufio"
	"fmt"
	"os"
)

func read(str string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(str)
	testString, _ := reader.ReadString('\n')
	return testString
}
