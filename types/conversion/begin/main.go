// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	a := 42.5
	b := uint(a)
	fmt.Printf("%v = %T\n %v = %T\n", a, a, b, b)
}
