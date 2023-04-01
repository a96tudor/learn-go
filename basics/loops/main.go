package main

import "log"

func forLoop() {
	for i := 0; i <= 5; i++ {
		log.Println("i is", i)
	}

	animals := []string{"horse", "cow", "dog", "cat"}
	for i, animal := range animals {
		log.Println("Animal number ", i, "is", animal)
	}

	animals2 := []string{"horse", "cow", "dog", "cat"}
	for _, animal := range animals2 {
		log.Println(animal)
	}

	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Tenko"
	for key, value := range animalsMap {
		log.Println(key, value)
	}
}

func main() {

}
