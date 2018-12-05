package main

import (
	"bufio"
	"fmt"
	"os"
)

func Read(str string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(str)
	testString, _ := reader.ReadString('\n')
	return testString
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
