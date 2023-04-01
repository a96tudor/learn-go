package main

import "log"

type Owner struct {
	FirstName string
	LastName  string
}

func (user Owner) printName() {
	log.Println(user.FirstName, user.LastName)
}

func createMap() map[string]Owner {
	myMap := make(map[string]Owner)

	myMap["dog"] = Owner{
		FirstName: "Tudor",
		LastName:  "Avram",
	}

	return myMap
}

func createSlice() []int {
	var mySlice1 []string

	mySlice1 = append(mySlice1, "Alice")
	mySlice1 = append(mySlice1, "Thompson")

	mySlice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	return mySlice2
}

func main() {
	myMap := createMap()

	myMap["dog"].printName()

	mySlice := createSlice()
	log.Println(mySlice[2:4])
}
