package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name  string
	Color string
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "Swissy",
	}
	gorilla := Gorilla{
		Name:  "Sammy",
		Color: "brown",
	}

	printInfo(&dog)
	printInfo(&gorilla)
}

func printInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
