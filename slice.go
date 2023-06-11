package main

import (
	"fmt"
	"sort"
)

func eighth_slices() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"} //slice
	fmt.Printf("Type of  fruitList is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList) // Apple Tomato Peach Mango Banana

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) // Tomato Peach Mango Banana, same as python

	highScore := make([]int, 4) // Type + slice
	highScore[0] = 234
	highScore[0] = 123
	highScore[0] = 111
	highScore[0] = 555

	highScore = append(highScore, 32323)
	sort.Ints(highScore) // sort Values

	fmt.Println(highScore)

	// based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "Ruby"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // reactjs, javascript, python, Ruby
}
