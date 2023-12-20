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

/*Maps
Selecting transcript lines in this section will navigate to timestamp in the video
- Go's map is a data type similar to hash tables or dictionaries in other languages. It simply maps keys to values. Let's work through some examples. We'll declare a map of string keys and author values. At this stage, we still have a nil map which we cannot populate quite yet. We must first initialize it. We'll do that with the make built-in function. We don't need to specify a size for the map. Next, we'll fill the map with a handful of author values, using the author initials as the keys. We'll then use the fmt.printf to get a verbose syntactical representation of the map. We'll now run our program, and see what we have so far. Great. As expected, we get the keys and values printed out for each entry in our map. Alternatively, we can use the literal approach to declaring and initializing this map. If we run the program again, we should get the same output, and because we have the, we already know what the type is in this situation that we're initializing the map with, we get all enter here can help us out and tell us when we have redundant task specification. Again, running program again should give us the exact same output. Great. Let's now retrieve a specific value from a map using a known key. We'll comment this out. And run our program now, and we pull out just this value from the map. When working with maps, it's also important to understand what happens when you look up values with a key that is not present in the map in order to avoid subtle bugs. Here we've got a map of authors initialized already. On line 17, we're looking up a non-existent key. Let's run the program to see what we get. No error is thrown. Instead, the set of braces we see printed represents an empty struct, or zero value of the map's declared value type. In some cases, this can produce unexpected logic errors, so it's advisable to use a technique called the comma okay style for map lookups. Let's see that in action. This style of lookup allows us to check the ok value, which is a bullion to determine if the key did indeed exist. From there, we can be certain that the value we're working with is real. Let's see about updating a map next. Naturally, maps can be modified. This example initializes our author's map with a literal as before, and prints it out. Let's jump to line 18, and write the code to update a value with key tm. We'll now uncomment the printing of the map, and run the program and see what we get. And voila, we've updated our map and replaced a value at key tm with James Baldwin. Let's delete a key from the map using the delete built-in function designed specifically for maps. Let's now run our program, after clearing the screen, and voila, we can see in the second output that the key tm was deleted. Okay. In this chapter, we looked at additional commonly used types in Go programs, including structs, arrays, slices, and maps. You absolutely need to wrap your head around these types before becoming proficient with Go. To help with that, let's tackle this chapter's challenge next. */
