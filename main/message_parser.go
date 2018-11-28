package main

import (
	"strings"
)

func ParseServer(str string) string {
	lenStr := len(str) - 1
	tmpStr := str[:lenStr]
	aStr := strings.Split(tmpStr, "\t")
	return aStr[0]
}
