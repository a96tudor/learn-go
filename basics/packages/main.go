package main

import (
	"basics/packages/helpers"
	"log"
)

func main() {
	var myVar helpers.Person
	myVar.FirstName = "Sam"
	log.Println(myVar.FirstName)
}
