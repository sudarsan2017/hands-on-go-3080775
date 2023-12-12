// types/constants/begin/main.go
package main

import (
	"fmt"
)

// declare a constant representing pi
const pi = 3.14159

// declare a rune constant for the letter 'a'
const a rune = 'a'
const b int = 10

func main() {
	const b int = 20
	fmt.Printf("b local %T %v\n", b, b)

	// fmt.Printf("a: %c - %T\n", a, a)

	// const a = 20
	// fmt.Printf("a const  %T - %v\n", a, a)

	// print the value and types of pi and a
	// fmt.Printf("pi: %v - %T\n", pi, pi)
	// fmt.Printf("a: %c - %T\n", a, a)

	// use unicode package to confirm that the rune is a letter
	// fmt.Println(unicode.IsLetter(a))

	printGlobalVars()
}
func printGlobalVars() {
	fmt.Printf("b global %T %v\n", b, b)
}
