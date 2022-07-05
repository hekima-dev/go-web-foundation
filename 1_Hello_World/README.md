### Introduction
Go is a battery included programming language and has a webserver already built in. The `net/http` package from the standard library contains all functionalities about the HTTP protocol. This includes (among many other things) an HTTP client and an HTTP server. In this example you will figure out how simple it is, to create a webserver that you can view in your browser.

### Registering a Request Handler
First, create a Handler which receives all incomming HTTP connections from browsers, HTTP clients or API requests. A handler in Go is a function with this signature:

```Go
    func (Response http.ResponseWriter, Request *http.Request)
```

The function receives two parameters:

An `http.ResponseWriter` which is where you write your text/html response to.
An `http.Request` which contains all information about this HTTP request including things like the URL or header fields.
Registering a request handler to the default HTTP Server is as simple as this:

```Go

    http.HandleFunc("/", func (Response http.ResponseWriter, Request *http.Request) {
        fmt.Fprintf(Response, "Hello World")
    })

```

### Listen for HTTP Connections
The request handler alone can not accept any HTTP connections from the outside. An HTTP server has to listen on a port to pass connections on to the request handler.

The following code will start Go’s default HTTP server and listen for connections on port 5000. You can navigate your browser to http://localhost:5000/ and see your server handing your request:

```Go

    Port := ":5000"
	Server := http.Server{
		Addr:           Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        nil,
		MaxHeaderBytes: 1 << 20,
	}

	Server.ListenAndServe()

```

### The Code (for copy/paste)

```Go

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

```