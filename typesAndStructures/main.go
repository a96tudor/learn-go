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
