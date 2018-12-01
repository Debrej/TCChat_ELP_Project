package main

import (
	"strconv"
)

func serverUserHandler(msgName string, paramValues map[string]string, users map[int]string) map[int]string {
	var uid int

	switch msgName {
	case "TCCHAT_REGISTER":
		i := 0
		keyExists := true
		for keyExists {
			if _, ok := users[i]; ok {
				i++
			} else {
				break
			}
		}
		users[i] = paramValues["nickname"]
		uid = i

	case "TCCHAT_DISCONNECT":
		uid, _ = strconv.Atoi(paramValues["uid"])
		delete(users, uid)
	}
	return users
}

func serverMessageHandler(paramValues map[string]string, users map[int]string) {

}
