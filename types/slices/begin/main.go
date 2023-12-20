// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	names := make([]string, 0)

	// append 3 names to the slice
	names = append(names, "John")
	names = append(names, "Jane")
	names = append(names, "Mary")

	// print the slice
	fmt.Println(names)
	// slice the slice using syntax slice[low:high]
	fmt.Println(names[1:3]) // [Jane Mary]
	fmt.Println(names[1:])  // [Jane Mary]
	fmt.Println(names[:1])  // [John]
	fmt.Println(names[:])   // [John Jane Mary]
}

/*Slices
Selecting transcript lines in this section will navigate to timestamp in the video
- Slices provide a convenient way of working with an underlying array of fixed size. Slices dynamically manage the underlying array for you so you don't have to worry about sizing concerns. Let's see it in action. We'll jump to line six where we'll declare a slice of strings using the make built-in function that allows us to specify a size for the type we're making. We haven't seen the make built-in yet, so let's quickly explain how it works. Make works with slices, maps, and channels to allocate a size in an optional capacity where applicable for these types. More specifically for our case, we're initializing a slice and specifying a size of zero. To fill the slices underlying array, we'll use another built-in function called append and add three values to names The append built-in is where some of the magic of slices happens. Once you run out of room in the underlying array for our slice a new, larger one will be created and have the values of the original transferred to it with room for more, usually double the size of the original array all without us having to do any of that work ourselves. This allows us to have all three names appended to the slice, even though it was initialized with zero. Note how we capture and assign the return value from the call to append, back to the slice we're appending to. We do this so that if the point of location of the slice changes we still have a reference to it. We'll print out the slice and run the program. All right, so far so good. Now let's talk about slicing slices. Slices support a slice operation where we specify a low and high bound, separated by a colon. The high bound is non-inclusive. Some examples: specifying one for the low and three for the high bounds. gets us Jane and Mary. Remember, the upper bound is non-inclusive. Leaving the high bound open gets everything from index one on. Leaving the low bound unspecified gets us everything from index zero up to the non-inclusive high bound. Lastly, just specifying the colon, gets everything from the slice. Let's now run the program, and confirm our output. If we scroll up, we can see for the first, starting from the bottom, we get everything, when we specify colon, we get only John, and when we specify the high bound of one when we specify the low bound of one we get everything from index one and with bounds one and three, we get only Jane and Mary. All right, there's more to slices than we cover here, so look to the README for additional resources. Let's now turn our attention to maps. */
