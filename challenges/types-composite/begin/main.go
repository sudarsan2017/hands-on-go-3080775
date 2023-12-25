// challenges/types-composite/begin/main.go
package main

import (
	"fmt"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthorName(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	lib := library{}
	fmt.Println(lib)

	// add 2 books to the library by the same author
	jb := author{name: "James Baldwin"}

	// add 1 book to the library by a different author
	lib.addBook(book{
		title:  "The Fire Next time",
		author: jb,
	})
	lib.addBook(book{
		title:  "Go Tell It is on the Mountain",
		author: jb,
	})

	lib.addBook(book{
		title: "Meditations",
		author: author{
			name: "Marcus Aurelius",
		},
	})

	// dump the library with spew
	//spew.Dump(lib)

	// lookup books by known author in the library
	books := lib.lookupByAuthorName(jb.name)
	//	spew.Dump(books)

	// print out the first book's title and its author's name
	if len(books) != 0 {
		b := books[0]
		fmt.Println(b.title, "by", b.author.name)
	}

}
