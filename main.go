package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}

func helloHandler(rw http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(rw, "Method is not supported", http.StatusNotFound)
	}

	fmt.Fprintf(rw, "Hello You!")
}

func formHandler(rw http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(rw, "ParseForm() err: %v", err)
		return
	}


	fmt.Fprintf(rw, "POST request successful!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(rw, "Name: %s \n", name)
	fmt.Fprintf(rw, "Address: %s \n ", address)
}