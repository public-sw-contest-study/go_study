package main

import "fmt"

func sixth_pointer() {
	fmt.Println("Welcome to a class on pointers")
	// use if wanna parse address for get actual value with garantee

	var ptr *int
	fmt.Println("Value of pointer is ", ptr) // nil

	myNumber := 23
	var numPtr = &myNumber

	fmt.Println("Value of actual pointer is ", numPtr)  // address
	fmt.Println("Value of actual pointer is ", *numPtr) // 23

	*numPtr = *numPtr + 2
	fmt.Println("New value is: ", myNumber) // 25

}
