package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type MessageArray struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseJSON(str string) []Message {
	decoderJson := json.NewDecoder(strings.NewReader(str))
	var mA MessageArray
	errA := decoderJson.Decode(&mA)
	check(errA)
	return mA.Messages
}

func ParseServer(str string) []string {
	absPath, _ := filepath.Abs("./main/ServerMessages.json")
	messagesArrayBytes, err := ioutil.ReadFile(absPath)
	check(err)
	messagesArray := string(messagesArrayBytes[:])
	messages := parseJSON(messagesArray)
	lenStr := len(str) - 1
	tmpStr := str[:lenStr]
	aStr := strings.Split(tmpStr, "\t")
	msgName := aStr[0]
	found := false
	i := 0
	var retParam []string
	for !found {
		if messages[i].Name == msgName {
			found = true
			retParam = messages[i].Parameters
		} else {
			i++
		}
	}
	return retParam
}
