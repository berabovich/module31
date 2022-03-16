package main

import (
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter            = 0
	firstInstanceHost  = "http://localhost:8080"
	secondInstanceHost = "http://localhost:8081"
)

func main() {
	http.HandleFunc("/", handleProxy)
	log.Fatalln(http.ListenAndServe(proxyAddr, nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {

	if counter == 0 {
		if _, err := http.NewRequest(r.Method, firstInstanceHost, r.Body); err != nil {
			log.Fatalln(err)
		}

		counter++
		return
	}
	if _, err := http.NewRequest(r.Method, secondInstanceHost, r.Body); err != nil {
		log.Fatalln(err)
	}
	counter--

}
