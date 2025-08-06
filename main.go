package main

import (
	"blog/configuration"
	"log"
)

func main() {
	router, err := configuration.CreateRouter()
	if err != nil {
		log.Println("Error creating router")
	} else {
		router.Run()
	}
}
