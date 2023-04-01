package main

import "log"

func ifs() {
	var isTrue bool
	isTrue = false
	if isTrue {
		log.Println("isTrue is true")
	} else {
		log.Println("isTrue is false")
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("cat is cat")
	}

	number := 100
	if number > 99 && !isTrue {
		log.Println("number is greater than 99 and isTrue is set to false")
	} else if number < 100 && isTrue {
		log.Println("number is smaller than 100 and isTrue is set to true")
	} else {
		log.Println("Something else")
	}
}

func switchStatement() {
	myVar := "horse"
	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")

	case "dog":
		log.Println("myVar is set to dog")

	default:
		log.Println("myVar is something else")
	}
}

func main() {

}
