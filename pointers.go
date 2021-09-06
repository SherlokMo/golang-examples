package main

import "fmt"

func zerovalue(iValue int) {
	iValue = 0
}

func zeroPtr(prtValue *int) {
	*prtValue = 0
}

func main() {

	i := 5

	// Nothing will happen we're changing the argument value in zeroValue func Scope.
	// i will have a value of 1.
	zerovalue(i)
	fmt.Println(i)

	fmt.Println(&i) // this will print i's location in memory.

	// this will send the address of i and assing it to ptrValue. thus we are doing the change on i adress in memory.
	// i  will = 0
	zeroPtr(&i)
	fmt.Println(i)
}
