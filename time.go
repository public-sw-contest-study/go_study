package main

import (
	"fmt"
	"time"
)

func fourth_time() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // Basic time for standard formatting, no exception

	createdDate := time.Date(2002, time.May, 4, 23, 50, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
