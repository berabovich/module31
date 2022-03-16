package main

import (
	"log"
	"module31/internal/app"
	"os"
)

func main() {
	err := app.Run(os.Args)
	//err = app.Run2(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	//	err = app.Run2(os.Args)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
}

//func main() {
//	http.HandleFunc("/", proxy.HandleProxy)
//	log.Fatalln(http.ListenAndServe("localhost:9000", nil))
//}
