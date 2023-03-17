package main

import "fmt"

func main() {
	var myString string = "Green"

	fmt.Println(myString)
	changeUsingPointer(&myString)
	fmt.Println(myString)

}

func changeUsingPointer(s *string) {
	*s = "Red"
}
