package main

import (
	"bufio"
	"net"
)

func main() {
	conn, errDial := net.Dial("tcp", "golang.org:80")
	check(errDial)
	msg, errRead := bufio.NewReader(conn).ReadString('\n')
	check(errRead)
	msgCName, _, msgCParams := ParseClient(msg)
	ClientHandler(msgCName, msgCParams)
}
