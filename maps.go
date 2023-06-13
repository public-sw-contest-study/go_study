package main

import "fmt"

func nineth_maps() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string) // to make slices or maps (key:value)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")

	for key, value := range languages {
		fmt.Printf("For Key %v, Value is %v", key, value)
	}
}
