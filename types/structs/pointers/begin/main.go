// types/structs/pointers/begin/main.go
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

// changeName changes the first and last name of the author
func (a *author) changeName(first, last string) {
	a.first = first
	a.last = last
}

func main() {
	a := author{
		first: "Mar1",
		last:  "Twain",
	}

	fmt.Println(a.fullName())

	// call changeName to update name of author
	a.changeName("Mark", "Twain")

	fmt.Println(a.fullName())
}
/* Struct pointers
Selecting transcript lines in this section will navigate to timestamp in the video
- [Instructor] In this lesson, we'll see how we can make changes and keep track of them in our structs. We have an "author struck" with a "first" and "last string" fields, along with a "fullName" method attached to it. We want to add support for changing the name of an author, so let's jump to line 15, to declare it. We'll call this method "changeName." (computer keys clicking) We'll specify the receiver. (computer keys clicking) And we'll provide first, last, both are strings. We're now going to return anything and inside of the function, for the method, we'll change the "first" and assign the incoming "first" parameter and we'll do the same for "last." Okay, now let's jump down to "main," where we'll initialize an "author" and print this up, using the "fullName." Let's now run a program and see what we have so far. Okay, great. Now, we noticed that we have a typo in the name of the author. Thankfully we have our "changeName" method, where we can actually update the name. So we'll call on the "changeName" and provide the updated values. Let's print our "author fullName" again. Hmm. Calling on our "changeName" method did not seem to have the desired effect, did it? That's because we modified a copy of our "author," not the one we initialized. How's that? Well, Go is a pass by value language and our "changeName" method will operate on a copy of the author instance we initialized, unless specified otherwise. To get the desired effect, we must add the "*" operator to the type we're attaching our method to. On line 17, we'll specify "*" in front of the type. Now if we run our program again, we can see that this time, we get the desired effect. Next up, struct embedding. */