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

func parseJSON(str string) []Message {
	decoderJson := json.NewDecoder(strings.NewReader(str))
	var mA MessageArray
	errA := decoderJson.Decode(&mA)
	check(errA)
	return mA.Messages
}

func ParseServer(str string) (string, []string, map[string]string) {
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
	var retParamNames []string
	retParamValues := make(map[string]string)
	for !found && !(i > len(messages)) {
		if messages[i].Name == msgName {
			found = true
			retParamNames = messages[i].Parameters
		} else {
			i++
		}
	}
	for i := 0; i < len(retParamNames); i++ {
		retParamValues[retParamNames[i]] = aStr[i+1]
	}
	return msgName, retParamNames, retParamValues
}

func ParseClient(str string) (string, []string, map[string]string) {
	absPath, _ := filepath.Abs("./main/ClientMessages.json")
	messagesArrayBytes, err := ioutil.ReadFile(absPath)
	check(err)
	messagesArray := string(messagesArrayBytes[:])
	messages := parseJSON(messagesArray)
	lenStr := len(str)
	tmpStr := str[:lenStr]
	aStr := strings.Split(tmpStr, "\t")
	msgName := aStr[0]
	found := false
	i := 0
	var retParamNames []string
	retParamValues := make(map[string]string)
	for !found && !(i > len(messages)) {
		if messages[i].Name == msgName {
			found = true
			retParamNames = messages[i].Parameters
		} else {
			i++
		}
	}
	for i := 0; i < len(retParamNames); i++ {
		retParamValues[retParamNames[i]] = aStr[i+1]
	}
	return msgName, retParamNames, retParamValues
}
