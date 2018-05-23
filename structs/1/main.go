package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corolla", 2017, "Gray"}
	car2 := Car{"BMW 320i", 2018, "Black"}

	fmt.Println(car1.Name)
	fmt.Println(car2.Name)
	fmt.Println(car1.info())

	sCar := SuperCar{
		Car: Car{
			"Fusca",
			2030,
			"Blue",
		},
		CanFly: true,
	}

	fmt.Println(sCar.info())
}
