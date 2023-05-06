package main

//Reneil Abner Coleman
//Test 2 Systems Programming
//Dalwin Lewis

import (
	"fmt"
	"log"
	"net/http"
)

func functionone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("functionone has been executed") // preprocessing
		next.ServeHTTP(w, r)                       // call the orginal handler
		log.Print("Executing functionone again")   //post processing after the middleware gets called
	})
}

func finish(w http.ResponseWriter, r *http.Request) {
	log.Print("Finishing up program execution")
	w.Write([]byte("Finished"))
}

//basic example of how to use handler function displays a a message to the screen
func welcome(w http.ResponseWriter, r *http.Request) {
	// Writes a byte slice with the text "This is Middleware"
	// in the response body
	w.Write([]byte("This is Middleware"))
}

//example of a funtion using middleware to log basic info of a HTTP request
//explanation: this funtion logs the network address, protocol version, HTTP method,
//..and request URL to the standard output. This information is available in Go’s http.Request
//..object.
// RemoteAddr: Network address that sent the request (IP:port)
// Proto: Protocol version
// Method: HTTP method
// URL: Request URL
func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG %s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux() //Initialize a new servemux

	finalHandler := http.HandlerFunc(finish)
	mux.Handle("/", functionone(finalHandler)) //register functionone and finalHandler for /
	mux.HandleFunc("/welcome", welcome)

	log.Print("Listening on :3000...") // Run http server with custom servemux at localhost:3000
	err := http.ListenAndServe(":3000", logRequestMiddleware(mux))
	log.Fatal(err) // Exit if any errors occur
}
