package main

import (
	"fmt"
	"net"
	"strconv"
)

func ServerRecHandler(msgName string, msgParams map[string]string, conn net.Conn, users map[int]string, conns map[int]net.Conn) (map[int]string, map[int]net.Conn, string) {
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
		str := "new connection from '" + msgParams["nickname"] + "' @ " + net.Addr(conn.RemoteAddr()).String() + "\n"
		fmt.Println(str)
		uidS := strconv.Itoa(i)
		msg = "TCCHAT_WELCOME\tELP_TCCHAT\t" + uidS + "\n"
		msgUserIn := "TCCHAT_USERIN\t" + msgParams["nickname"] + "\n"
		SendBroadcast(msgUserIn, uidS, conns)
		uid = i

	case "TCCHAT_DISCONNECT":
		uid, _ = strconv.Atoi(msgParams["uid"])
		msg = "TCCHAT_USEROUT\t" + users[uid] + "\n"
		SendBroadcast(msg, msgParams["uid"], conns)
		_ = conns[uid].Close()
		delete(conns, uid)
		delete(users, uid)
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
		retString = "TCCHAT_BCAST\t" + msg + "\t" + users[uid] + "\n"
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

func ServerSendHandler(msgName string, msgParams map[string]string, msg string, users map[int]string, conns map[int]net.Conn) {
	switch msgName {
	case "TCCHAT_BCAST":
		SendBroadcast(msg, msgParams["uid"], conns)

	case "TCCHAT_WELCOME":
		msgBcast := "TCCHAT_USERIN\t" + msgParams["nickname"] + "\n"
		SendBroadcast(msgBcast, msgParams["uid"], conns)

	case "TCCHAT_USEROUT":
		SendBroadcast(msg, msgParams["uid"], conns)

	case "TCCHAT_PERSONNAL":
		fmt.Println("Sending to user n°" + msgParams["uid"])
		destUid, _ := strconv.Atoi(msgParams["uid"])
		for user := range users {
			if users[user] == msgParams["nickname"] {
				destUid = user
			}
		}
		_, errWrite := conns[destUid].Write([]byte(msg))
		check(errWrite)
	}
}

func SendBroadcast(msg string, uidS string, conns map[int]net.Conn) {
	uid, _ := strconv.Atoi(uidS)
	for uidCurrent, conn := range conns {
		if uidCurrent != uid {
			fmt.Println("Sending to user n°" + strconv.Itoa(uid))
			_, errWrite := conn.Write([]byte(msg))
			check(errWrite)
		}
	}
}
