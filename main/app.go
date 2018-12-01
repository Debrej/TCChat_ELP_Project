package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	users := make(map[int]string)

	for true {
		/* THIS PART READS THE INPUT FROM THE USER */
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		testString, _ := reader.ReadString('\n')

		///* HERE WE USE ParseServer TO GET THE CORRESPONDING PARAMETERS AND THEIR RESPECTIVE VALUES*/
		msgName, msgParamsName, msgParams := ParseServer(testString)
		//_, msgParamsName, msgParams := ParseClient(testString)

		if msgName != "TCCHAT_MESSAGE" {
			users = serverUserHandler(msgName, msgParams, users)
		}

		for i := 0; i < len(users); i++ {
			if nickname, ok := users[i]; ok {
				fmt.Println(strconv.Itoa(i) + " : " + nickname)
			}
		}

		for i := 0; i < len(msgParams); i++ {
			fmt.Println(msgParamsName[i] + " : " + msgParams[msgParamsName[i]])
		}
	}
}
