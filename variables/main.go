package main

import "fmt"

const (
	aa string = "Oiiii"
	bb        = 66
	cc        = 77
)

const xvz int = 1234

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := `Uouuuu
	asdfasf


	asfdasdf
	safdasdfas`

	const xpto = 10

	fmt.Printf("%s \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", xpto)
}
