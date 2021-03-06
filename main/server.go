package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	// CREATION DU SERVEUR
	ln, err := net.Listen("tcp", ":16000")
	fmt.Println("Server up\n")
	check(err)

	// CREATION TABLEAU USERS

	users := make(map[int]string)
	conns := make(map[int]net.Conn)

	go func() {
		for {
			conn, errAccept := ln.Accept()
			check(errAccept)

			msg, errRead := bufio.NewReader(conn).ReadString('\n')
			check(errRead)

			msgName, _, msgParams := ParseServer(msg)
			users, conns, msg = ServerRecHandler(msgName, msgParams, conn, users, conns)

			_, errWrite := conn.Write([]byte(msg))
			check(errWrite)

			ServerSendHandler(msgName, msgParams, msg, users, conns)

			go func(conn net.Conn) {
				for {
					msg, errRead := bufio.NewReader(conn).ReadString('\n')
					if errRead != nil {
						_ = conn.Close()
						break
					}
					msgName, _, msgParams := ParseServer(msg)
					users, conns, msg = ServerRecHandler(msgName, msgParams, conn, users, conns)

					fmt.Print("message generated : ")

					fmt.Println(msg)

					ServerSendHandler(msgName, msgParams, msg, users, conns)
				}
			}(conn)
		}
	}()

	for {
		_ = 0
	}
}
