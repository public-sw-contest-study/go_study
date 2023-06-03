package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func third_conv() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // input: added "\n" if press enter

	if err != nil {
		// if err exist,
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
