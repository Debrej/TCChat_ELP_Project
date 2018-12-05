package main

import (
	"bufio"
	"net"
)

func main() {

	nickname := Read("Please enter your nickname : ")
	msg_nickname := "TCCHAT_REGISTER\t" + nickname + "\n"

	conn, errDial := net.Dial("tcp", "192.168.43.10:16000")
	check(errDial)
	_, errWrite := conn.Write([]byte(msg_nickname))
	check(errWrite)

	msg, errRead := bufio.NewReader(conn).ReadString('\n')
	check(errRead)

	msgCName, _, msgCParams := ParseClient(msg)
	ClientHandler(msgCName, msgCParams)
}
