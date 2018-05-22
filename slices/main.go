package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)

	//slice = append(slice, 1)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Leandro",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0])   // Hello
	fmt.Println(sliceString[:2])  // Hello World
	fmt.Println(sliceString[1:2]) // World
	fmt.Println(sliceString[2:4]) // Much Better
	fmt.Println(sliceString[2:])  // Todo o restante
}
