package main

import (
	"log"
	"time"
)

type User struct {
	firstName string
	lastName  string
	age       int
	birthday  time.Time
}

// Associating a function with a struct
func (m *User) printFirstName() string {
	log.Println(m.firstName)
	return m.firstName
}

func main() {
	user := User{
		firstName: "Trevor",
		lastName:  "Johnson",
		age:       25,
	}

	log.Println(user.birthday)
}

// Private to this package
func something() {

}

// Visible outside of the package
func Something() {

}
