package main //go program starts with the package main

//import bloack for external packages
import (
	"fmt" //Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
	"net/http" //Package http provides HTTP client and server implementations.
)

/*
The main func in go is similar to the main method of java. the program starts and ends in the function
*/
func main() {
	fmt.Println("Webserver is Ready and listening in Port 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/*
a simple handler method which prints the response to the web server request
*/
func handler(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "Welcome to GoLang\n")
}