package main

import (
	"strconv"
)

func ServerHandler(msgName string, msgParams map[string]string, users map[int]string) (map[int]string, string) {
	var msg string

	switch msgName {
	case "TCCHAT_REGISTER", "TCCHAT_DISCONNECT":
		users, msg = serverUserHandler(msgName, msgParams, users)

	case "TCCHAT_MESSAGE":
		users, msg = serverMessageHandler(msgParams, users)
	}

	return users, msg
}

func serverUserHandler(msgName string, msgParams map[string]string, users map[int]string) (map[int]string, string) {
	var uid int
	var msg string

	switch msgName {
	case "TCCHAT_REGISTER":
		i := 0
		keyExists := true
		for keyExists {
			if _, ok := users[i]; ok {
				i++
			} else {
				keyExists = false
			}
		}
		users[i] = msgParams["nickname"]
		msg = "TCCHAT_USERIN\t" + msgParams["nickname"]
		uid = i

	case "TCCHAT_DISCONNECT":
		uid, _ = strconv.Atoi(msgParams["uid"])
		msg = "TCCHAT_USEROUT\t" + users[uid]
		delete(users, uid)

	}
	return users, msg
}

func serverMessageHandler(msgParams map[string]string, users map[int]string) (map[int]string, string) {
	isPersonal, msg, dest := checkPersonal(msgParams["msg_payload"])
	var retString string
	uid, _ := strconv.Atoi(msgParams["uid"])
	if isPersonal {
		retString = "TCCHAT_PERSONAL\t" + msg + "\t" + users[uid] + "\t" + dest
	} else {
		retString = "TCCHAT_MESSAGE\t" + msg + "\t" + users[uid]
	}
	return users, retString
}

func checkPersonal(msg string) (bool, string, string) {
	isPersonal := false
	msgRet := ""
	destRet := ""
	if msg[0] == '@' && msg[1] != ' ' {
		isPersonal = true
		i := 0
		for msg[i] != ' ' {
			i++
		}
		destRet = msg[1 : i+1]
		msgRet = msg[i+1:]
	} else {
		msgRet = msg
	}
	return isPersonal, msgRet, destRet
}
