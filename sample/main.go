package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    var age int
    fmt.Print("Enter your age: ")
    fmt.Scanf("%d", &age)

    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}