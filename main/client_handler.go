package main

import (
	"fmt"
	"strconv"
)

func ClientHandler(msgName string, msgParams map[string]string) {
	switch msgName {
	case "TCCHAT_WELCOME":
		showWelcome(msgParams)

	case "TCCHAT_USERIN":
		showUserIn(msgParams)

	case "TCCHAT_USEROUT":
		showUserOut(msgParams)

	case "TCCHAT_MESSAGE":
		showMsg(msgParams)

	case "TCCHAT_PERSONAL":
		showPrivateMsg(msgParams)
	}

}

func showWelcome(msgParams map[string]string) int {
	serverName := msgParams["server_name"]
	fmt.Println("Welcome to " + serverName)
	uid, _ := strconv.Atoi(msgParams["uid"])
	return uid
}

func showUserIn(msgParams map[string]string) {
	nickname := msgParams["nickname"]
	fmt.Println("A new user arrives, welcome " + nickname)
}

func showUserOut(msgParams map[string]string) {
	nickname := msgParams["nickname"]
	fmt.Println(nickname + " left us... :'(")
}

func showMsg(msgParams map[string]string) {
	nickname := msgParams["src_nickname"]
	msg := msgParams["msg_payload"]
	fmt.Println(nickname + " : " + msg)
}

func showPrivateMsg(msgParams map[string]string) {
	srcNickname := msgParams["src_nickname"]
	destNickname := msgParams["dest_nickname"]
	msg := msgParams["msg_payload"]
	fmt.Println(srcNickname + "@" + destNickname + " : " + msg)
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
