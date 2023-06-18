package main

import "fmt"

func thirteen_loop() {
	fmt.Println("Loop in Golang")
	days := []string{"sunday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	// for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	// }

	for i, day := range days {
		// i = index
		fmt.Printf("index is %v and value is %v", i, day)
	}

	value := 1
	for value < 10 {
		if value == 7 {
			goto lco
		}
		if value == 5 {
			break
		}

		if value == 6 {
			value++
			continue // skip 1 phase
		}
		fmt.Println("Value is : ", value)
		value++
	}

lco: // goto
	fmt.Println("Jumping at LearnCodeOnline.in")
}
