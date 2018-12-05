package main

import "net"

func main() {
	ln, err := net.Listen("tcp", ":8080")
	check(err)
	for {
		conn, errAccept := ln.Accept()
		check(errAccept)
		msg := []byte("TCCHAT_WELCOME\tELP_TCCHAT")
		_, errWrite := conn.Write(msg)
		check(errWrite)
	}
}
