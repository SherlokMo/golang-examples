package main

import "fmt"

func calculate(closure func(a int, b int) int, a int, b int) int {
	return closure(a, b)
}

func main() {

	sum := func(a int, b int) int {
		return a + b
	}

	devide := func(a int, b int) int {
		return a / b
	}

	fmt.Println(calculate(sum, 1, 2))
	fmt.Println(calculate(devide, 2, 2))

}
