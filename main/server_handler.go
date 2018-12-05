package main

import (
	"fmt"
	"net"
	"strconv"
)

func ServerHandler(msgName string, msgParams map[string]string, conn net.Conn, users map[int]string, conns map[int]net.Conn) (map[int]string, map[int]net.Conn, string) {
	var msg string

	switch msgName {
	case "TCCHAT_REGISTER", "TCCHAT_DISCONNECT":
		users, conns, msg = serverUserHandler(msgName, msgParams, conn, users, conns)

	case "TCCHAT_MESSAGE":
		users, msg = serverMessageHandler(msgParams, users)
	}

	return users, conns, msg
}

func serverUserHandler(msgName string, msgParams map[string]string, conn net.Conn, users map[int]string, conns map[int]net.Conn) (map[int]string, map[int]net.Conn, string) {
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
		conns[i] = conn
		msg = "TCCHAT_WELCOME\tELP_TCCHAT\t" + strconv.Itoa(i) + "\n"
		fmt.Println(msg)
		uid = i

	case "TCCHAT_DISCONNECT":
		uid, _ = strconv.Atoi(msgParams["uid"])
		msg = "TCCHAT_USEROUT\t" + users[uid] + "\n"
		delete(users, uid)
		delete(conns, uid)

	}
	return users, conns, msg
}

func serverMessageHandler(msgParams map[string]string, users map[int]string) (map[int]string, string) {
	isPersonal, msg, dest := checkPersonal(msgParams["msg_payload"])
	var retString string
	uid, _ := strconv.Atoi(msgParams["uid"])
	if isPersonal {
		retString = "TCCHAT_PERSONAL\t" + msg + "\t" + users[uid] + "\t" + dest + "\n"
	} else {
		retString = "TCCHAT_MESSAGE\t" + msg + "\t" + users[uid] + "\n"
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
