package main

import (
	"bufio"
	"fmt"
	"os"
)

func second_userinput() {
	input := "Welcome to user input"
	fmt.Println(input)

	// buffer, which can read input output, store something

	reader := bufio.NewReader(os.Stdin) // pointer
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok || err ok => treat errors like variables

	inputData, _ := reader.ReadString('\n') // success || error
	fmt.Println("Thanks for rating, ", inputData)
	fmt.Printf("Type of this rating is %T ", inputData)
}
