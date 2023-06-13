package main

import "fmt"

func tenth_structs() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "email", true, 16}
	fmt.Println(hitesh)                             // {name email status buol}
	fmt.Printf("hitesh details are: %+v\n", hitesh) // if struct %=v => more details
	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

	// start with Capital letters
}
