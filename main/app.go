package main

func main() {

	users := make(map[int]string)

	for true {
		var msg string
		/* THIS PART READS THE INPUT FROM THE USER */
		testString := Read("Enter text: ")

		/* HERE WE USE ParseServer TO GET THE CORRESPONDING PARAMETERS AND THEIR RESPECTIVE VALUES*/
		//msgName, msgParamsName, msgParams := ParseServer(testString)
		msgName, _, msgParams := ParseServer(testString)

		users, msg = ServerHandler(msgName, msgParams, users)

		msgCName, _, msgCParams := ParseClient(msg)
		ClientHandler(msgCName, msgCParams)

		/*
			for i := 0; i < len(msgParams); i++ {
				fmt.Println(msgParamsName[i] + " : " + msgParams[msgParamsName[i]])
			}
		*/
	}
}
