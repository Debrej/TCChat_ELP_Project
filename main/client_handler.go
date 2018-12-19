package main

import (
	"fmt"
	"os"
	"strconv"
)

func ClientHandler(msgName string, msgParams map[string]string, uid int, f *os.File) int {
	switch msgName {
	case "TCCHAT_WELCOME":
		uid = showWelcome(msgParams, f)

	case "TCCHAT_USERIN":
		showUserIn(msgParams, f)

	case "TCCHAT_USEROUT":
		showUserOut(msgParams, f)

	case "TCCHAT_MESSAGE":
		showMsg(msgParams, f)

	case "TCCHAT_PERSONAL":
		showPrivateMsg(msgParams, f)
	}
	return uid
}

func showWelcome(msgParams map[string]string, f *os.File) int {
	serverName := msgParams["server_name"]
	_, _ = f.Write([]byte("Welcome to " + serverName))
	uid, _ := strconv.Atoi(msgParams["uid"])
	return uid
}

func showUserIn(msgParams map[string]string, f *os.File) {
	nickname := msgParams["nickname"]
	_, _ = f.Write([]byte("\nA new user arrives, welcome " + nickname))
}

func showUserOut(msgParams map[string]string, f *os.File) {
	nickname := msgParams["nickname"]
	_, _ = f.Write([]byte("\n" + nickname + " left us... :'("))
}

func showMsg(msgParams map[string]string, f *os.File) {
	nickname := msgParams["src_nickname"]
	msg := msgParams["msg_payload"]
	_, _ = f.Write([]byte("\n" + nickname + " : " + msg))
}

func showPrivateMsg(msgParams map[string]string, f *os.File) {
	srcNickname := msgParams["src_nickname"]
	destNickname := msgParams["dest_nickname"]
	msg := msgParams["msg_payload"]
	_, _ = f.Write([]byte("\n" + srcNickname + "@" + destNickname + " : " + msg))
}

func checkCommand(msgPayload string, uid string) (string, bool) {
	msg := ""
	disconnect := false

	if msgPayload[0] == '/' {
		if len(msgPayload) >= 11 && msgPayload[1:11] == "disconnect" {
			msg = "TCCHAT_DISCONNECT\t" + uid + "\n"
			fmt.Println("Goodbye")
			disconnect = true
		}
	} else {
		msg = "TCCHAT_MESSAGE\t" + msgPayload + "\t" + uid + "\n"
	}

	return msg, disconnect
}
