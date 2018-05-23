package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"Nome"`
	Year  int    `json:"-"` // remove o campo do serializador
	Color string
}

func main() {
	car := Car{"Gol", 2010, "Gray"}

	result, _ := json.Marshal(car)

	//fmt.Println(result) // imprimi na tabela asc
	fmt.Println(string(result))
}
