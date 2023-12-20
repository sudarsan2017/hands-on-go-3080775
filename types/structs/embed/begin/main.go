// types/structs/embed/begin/main.go
package main

import "fmt"

type person struct {
	first string
	last  string
}

// fullName returns the full name of a person
func (p person) fullName() string {
	return p.first + " " + p.last
}

// define author and embed person
type author struct {
	person  person
	penName string
}

// override fullName method for author
func (a author) fullName() string {
	return fmt.Sprintf("%s (%s)", a.person.fullName(), a.penName)
}

func main() {
	// initialize and print a person's full name
	p := person{
		first: "Toni",
		last:  "Morrison",
	}
	fmt.Println(p.fullName())

	// initialize and print an author's full name
	a := author{
		person: person{
			first: "James",
			last:  "Baldwin",
		},
		penName: "Jimmy",
	}
	fmt.Println(a.fullName())
}

/*Struct embedding
Selecting transcript lines in this section will navigate to timestamp in the video
- Go takes the old adage of preferred composition over inheritance to heart, and doesn't support the notion of inheritance at all. What it does though, is provide a seamless way to compose certain types, mainly STRs and interfaces, through what's called embedding. Let's see this in action with STRs. We start out here with a person type, already defined for us along with a method that returns a full name. We'll define a custom type of our own called author. We'll jump to line 17, (Typing) and define the author- type. Since an author is also a person with a first and last name, we'll use go embedding and simply specify the person type as one of the fields for author. We can name this field like any other, or we can leave just a type, and reference it accordingly. Our author type, will have a pen name, so we'll add that as an additional field. And it will be of type string. Down in main, where we already have a person initialized, we'll also initialize an author and compare the outputs. We'll now run the program, and compare the outputs. All right, we see how the person's first and last name fields were promoted and made usable to the author type, but what about our pen name field? Ideally, we'd like to show that in the output for full name, but at the moment, that method promoted as is from person knows nothing about pen names. We'll need to overwrite that method then for our author type. We'll use S print F here. And pass in, the full name as Define an Author, which works just fine, but now, we'll reference the pen name for the second verb that we're specifying in the S Print F function. You can learn more about the font packages, S Sprint F function, on pkg.go.dev. Now that we've overridden the full name for author, let's run the program again and compare the output. And voilà! We're able to get the pin name and our full name output for the author where it's relevant. Remember, embedding also works for interfaces too. More on that later. Next up, erase. */
