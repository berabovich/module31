package main

import (
	"log"
	"module31/internal/app"
)

func main() {
	//port := os.Args
	port := []string{"", ":8080"}
	err := app.Run(port[1])
	if err != nil {
		log.Fatal(err)
	}

}
