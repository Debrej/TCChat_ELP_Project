package main

import (
	"bufio"
	"net"
)

func main() {
	conn, errDial := net.Dial("tcp", "golang.org:80")
	check(errDial)
	messages_in := make(chan string)
	messages_out := make(chan string)
	go readServ(conn, messages_in)
	go write(conn, messages_out)

}

func readServ(conn net.Conn, channel chan string) {
	end := false
	reader := bufio.NewReader(conn)

	for !end {
		msg, errRead := reader.ReadString('\n')
		check(errRead)
		msgCName, _, msgCParams := ParseClient(msg)
		ClientHandler(msgCName, msgCParams)
	}
}

func write(conn net.Conn, channel chan string) {
	end := false

}
