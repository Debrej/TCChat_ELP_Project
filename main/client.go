package main

import (
	"bufio"
	"net"
	"os"
)

func main() {

	uid := "0"

	ready := make(chan int)

	nickname := Read("Please enter your nickname : ")
	nickname = nickname[:len(nickname)-1]
	msgNickname := "TCCHAT_REGISTER\t" + nickname + "\n"

	conn, errDial := net.Dial("tcp", "192.168.0.161:16000")
	check(errDial)
	_, errWrite := conn.Write([]byte(msgNickname))
	check(errWrite)

	go func() {
		msg, errRead := bufio.NewReader(conn).ReadString('\n')
		check(errRead)

		msgCName, _, msgCParams := ParseClient(msg)
		ClientHandler(msgCName, msgCParams)

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
			ClientHandler(msgCName, msgCParams)
		}
	}()

	for {
		msgPayload := Read(nickname + " : ")
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
