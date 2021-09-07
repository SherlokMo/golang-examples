package main

import "fmt"

type rect struct {
	width  int
	height int
}

//pointer passing
func (r *rect) area() int {
	return r.height * r.height
}

//value passing
func (r rect) perim() int {

	return 2 * r.height * 2 * r.width
}

func main() {

	r := rect{5, 5}

	fmt.Println(r.area())
	fmt.Println(r.perim())
}
