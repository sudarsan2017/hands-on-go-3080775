// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author

	// initialize the map with make
	authors = make(map[string]author)

	// add authors to the map
	authors["tm"] = author{name: "Toni Morrison"}
	authors["ma"] = author{name: "Marcus Aurelius"}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	fmt.Println(authors["tm"])

	// check when the value for non-existent key
	fmt.Println(authors["jr"])
}

/* Maps
- [Instructor] Go's array declaration syntax is a number inside of square brackets immediately followed by some type T. Let's see what that looks like. We'll jump to line 6. Let's declare an array that can hold up to three integers. We'll print it out. And run the program. As we expect there are three slots available in our array, all initialized to the zero value for the type int. We'll update the first element and print out the array once more. Easy as that. Once an array is declared, however, you cannot resize it. That means if we ran out of a room in our array we'd have to declare a new, larger one, copy the values from the original, and then have space for new values. Thankfully, Go has a more convenient way of working with arrays called slices. Let's look at that next.
*/
