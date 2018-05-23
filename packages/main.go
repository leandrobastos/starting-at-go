package main

import (
	"curso/packages/car"
	"fmt"
)

func main() {
	car := car.Car{"Gol", "Gray"}

	fmt.Println(car.Start())
}
