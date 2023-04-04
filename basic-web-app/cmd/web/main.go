package main

import (
	"basic-web-app/pkg/handlers"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(
		fmt.Sprintf("Starting the application at 127.0.0.1%s", portNumber),
	)
	err := http.ListenAndServe(portNumber, nil)
	log.Println(err)
}
