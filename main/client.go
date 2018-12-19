package main

import (
	"bufio"
	"net"
	"os"
	"strconv"
)

func main() {

	uid := "0"

	ready := make(chan int)

	nickname := Read("Please enter your nickname : ")
	nickname = nickname[:len(nickname)-1]
	msgNickname := "TCCHAT_REGISTER\t" + nickname + "\n"

	conn, errDial := net.Dial("tcp", "192.168.43.10:16000")
	check(errDial)
	_, errWrite := conn.Write([]byte(msgNickname))
	check(errWrite)

	go func() {
		msg, errRead := bufio.NewReader(conn).ReadString('\n')
		check(errRead)

		msgCName, _, msgCParams := ParseClient(msg)
		uidI, _ := strconv.Atoi(uid)
		uid = strconv.Itoa(ClientHandler(msgCName, msgCParams, uidI))

		// CLIENT IS READY
		ready <- 0
	}()

	// WE WAIT UNTIL CLIENT IS READY
	<-ready

	go func() {
		for {
			msg, errRead := bufio.NewReader(conn).ReadString('\n')
			check(errRead)

			msgCName, _, msgCParams := ParseClient(msg)
			uidI, _ := strconv.Atoi(uid)
			ClientHandler(msgCName, msgCParams, uidI)
		}
	}()

	for {
		msgPayload := Read(nickname + " : ")
		if msgPayload != "" {
			msgPayload = msgPayload[:len(msgPayload)-1]
			msgPayload = replacer.Replace(msgPayload)

			msg, disconnect := checkCommand(msgPayload, uid)

			_, errWrite := conn.Write([]byte(msg))
			check(errWrite)

			if disconnect {
				os.Exit(0)
			}
		}
	}
}
