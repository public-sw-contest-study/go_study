package main

import "fmt"

func seventh_array() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	// [Apple Tomato   Peach]
	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List Length is: ", len(fruitList)) // 4

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("vegList is: ", vegList)
}
