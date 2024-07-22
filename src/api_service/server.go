package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Hello world")
	http.HandleFunc("/helloworld", handleHelloWord)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8484"
	log.Printf("Listening on %s ... ", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleHelloWord(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	fmt.Println("HElloWorld Endpoint triggered ")

	ResponseWriter(writer, "Hello World")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	fmt.Println("Health Endpoint triggered ")
	ResponseWriter(writer, "OK")
}

func ResponseWriter(writer http.ResponseWriter, responseString string) {
	response := []byte(responseString)
	fmt.Println(response)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
