package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handles form submission
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err) // ✅ Use fmt.Fprintf for client response
		return
	}

	fmt.Fprintf(w, "POST request Successful\n") // ✅ Use fmt.Fprintf instead of fmt.Printf

	name := r.FormValue("name")
	address := r.FormValue("address")

	// ✅ Print output to both terminal & client response
	fmt.Println("Received Form Data:")
	fmt.Println("Name:", name)
	fmt.Println("Address:", address)

	// ✅ Send response to client
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

// Handles /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // ✅ Capital "F" in FileServer
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)  // ✅ Correct function name
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080...")

	// ✅ Use "nil" instead of "null"
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
