package main

import (
	"log"
	"module31/internal/app"
	"os"
)

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
