package main

// https://www.youtube.com/watch?v=9fYqg6uo-UU
import "fmt"

const LoginToken string = "test" // Public if use Capltal to first letter

func first_variables() {
	var userName string = "lavi02"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 // 0 ~ 2^8 - 1, int8 = -(2^8 - 1) ~ 2^8 - 1
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4545 // Decimal
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var implType = "test"
	fmt.Println(implType)

	numberOfUser := 300 // no var style, cannot use global outSide
	fmt.Println(numberOfUser)
}
