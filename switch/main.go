package main

import "fmt"

func main() {

	a := "Leandro"
	switch a {
	case "Bob":
		fmt.Println("Hey Bob")
	case "Leandro":
		fmt.Println("Hey Leandro")
	default:
		fmt.Println("I did not find your name!")

	}

}
