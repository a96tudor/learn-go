package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// Variable declaration
	var whatToSay string
	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	var i int
	i = 7
	fmt.Println("i is set to ", i)

	something, somethingElse := saySomething()
	fmt.Println(something, somethingElse)

}

func saySomething() (string, string) {
	return "something", "else"
}
