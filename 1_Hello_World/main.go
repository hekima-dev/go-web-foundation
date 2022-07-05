/* defining package name */
package main

/* importing other go package */
import (
	"fmt"
	"net/http"
	"time"
)

/* creating main function */
func main() {

	/* creating http handler function */
	http.HandleFunc("/", HelloWorldHandler)

	/* defining server information */
	Port := ":5000"
	Server := http.Server{
		Addr:           Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        nil,
		MaxHeaderBytes: 1 << 20,
	}

	/* printing message to the terminal */
	fmt.Println("Starting server on http://localhost" + Port)

	/* starting server on port 5000 */
	Server.ListenAndServe()

}

/* hello world handler function */
func HelloWorldHandler(Response http.ResponseWriter, Request *http.Request) {

	/* response to the client with hello world */
	fmt.Fprintf(Response, "<h1> Hello World </h1>")

}
