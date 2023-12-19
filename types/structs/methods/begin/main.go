// types/structs/methods/begin/main.go
package main

import "fmt"

type author struct {
	first string
	last  string
}

// fullName returns the full name of the author
func (a author) fullName() string {
	return a.first + " " + a.last
}

func main() {
	// initialize author
	a := author{
		first: "Marcus",
		last:  "Aurelius",
	}

	// print the author's full name
	fmt.Println(a.fullName())

}

// We can attach functions to our structs like we can for most types in go. We simply specify our struct as what's called the receiver and function declarations to turn these functions into methods on our struct. In this example, we've pre-declared our custom author type and have left room for defining a full name method which should return the first and last name of the author separated by a space. Let's jump to line 10 and start with the definition. To specify a function receiver, we use parenthesis that precede the function name and inside the parenthesis, we specify an optional variable and the name of the type we wish to attach the function to. That's what makes it a method. When specified, the variable is used inside of the function to reference fields that are part of the custom type. In our main, we'll initialize an author value and print out the result of calling the full name method on it. Now let's run our program and see what the output looks like. And there we go. Our full name method reads our struct's fields to produce its return value. We'll see how we can mutate our fields next.
