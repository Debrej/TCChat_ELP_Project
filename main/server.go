package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	// CREATION DU SERVEUR
	ln, err := net.Listen("tcp", ":16000")
	fmt.Println("Server up")
	check(err)

	// CREATION TABLEAU USERS

	users := make(map[int]string)
	conns := make(map[int]net.Conn)

	//go func() {
	for {
		conn, errAccept := ln.Accept()
		check(errAccept)

		msg, errRead := bufio.NewReader(conn).ReadString('\n')
		check(errRead)
		msgName, _, msgParams := ParseServer(msg)
		fmt.Println("new connection from '" + msgParams["nickname"] + "' @ " + net.Addr(conn.RemoteAddr()).String() + "\n\n")
		users, conns, msg = ServerHandler(msgName, msgParams, conn, users, conns)

		fmt.Println(users)
		fmt.Println(conns)

		_, errWrite := conn.Write([]byte(msg))
		check(errWrite)
	}
	//}()
}
