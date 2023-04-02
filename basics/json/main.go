package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool   `json:"has_dog"`
}

func UnmarshallPersons(jsonData string) []Person {
	var unmarshalled []Person

	err := json.Unmarshal(
		[]byte(jsonData), // slice of bytes
		&unmarshalled,
	)

	if err != nil {
		log.Println("Something went wrong:", err)
	}

	return unmarshalled
}

func MarshallPersons(persons []Person) string {
	newJson, err := json.MarshalIndent(persons, "", "  ")

	if err != nil {
		log.Println("Something went wrong:", err)
	}

	return string(newJson)
}

func main() {
	myJSON := `
	[
		{
			"first_name": "Jack",
			"last_name": "Jones",
			"has_dog": true
		},
		{
			"first_name": "Sam",
			"last_name": "Smith",
			"has_dog": false
		}
	]
	`
	var unmarshalled []Person
	unmarshalled = UnmarshallPersons(myJSON)
	log.Println("unmarshalled %v", unmarshalled)

	var p1 Person
	p1.FirstName = "Abby"
	p1.LastName = "Coleman"
	p1.HasDog = true

	var p2 Person
	p2.FirstName = "Michael"
	p2.LastName = "Thompson"
	p2.HasDog = false

	var persons []Person
	persons = append(persons, p1)
	persons = append(persons, p2)

	marshalled := MarshallPersons(persons)
	log.Println(marshalled)
}
