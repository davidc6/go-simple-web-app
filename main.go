package main

import (
	"net/http"
)

func main() {
	// uncomment to try both
	// original()
	alternative()
}

// -------------------------------- A way to do the same --------------------------------

// / - route handler
// w - is a value of type ResponseWriter, used by HTTP handler to create an HTTP response
// r - * is a pointer to Request (how data from request is retrieved)
// e.g. request pointer can provide data from a form request
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home</h1>"))
}

// /about - route handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>About</h1>"))
}

// /contact - route handler
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Contact</h1>"))
}

func original() {
	// provides that ability to handle requests
	// aboutHandler (and the rest) is of type HandleFunc is an adapter that allows to use ordinary functions as HTTP handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	
	// starts an HTTP server on a given address (port) and with a given handler
	// listen on TCP and handles requests on incoming connections
	// when "nil" is supplied then DefaultServeMux is used
	// DefaultServeMux is the default ServeMux which is an HTTP multiplexer (url matching, routing, etc.)
	http.ListenAndServe(":8080", nil)
}

// -------------------------------- Alternative way to do the same --------------------------------

// a collection of fields
type homeHandlerAlt struct {
	name string
}
// ServeHTTP is a method that has a receiver argument named p of type homeHandlerAlt 
// p is not really used inside ServeHTTP method in this example
func (p homeHandlerAlt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home (alternative)</h1>" + "<h2>Hey " + p.name + "!</h2>"))
}

type aboutHandlerAlt struct {}
func (g aboutHandlerAlt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>About (alternative)</h1>"))
}

type contactHandlerAlt struct {}
func (k contactHandlerAlt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Contact (alternative)</h1>"))
}

func alternative() {
	// create a new ServeMux (see above for explanation)
	mux := http.NewServeMux()

	// / - route handler
	// 1st - route, 2nd - create homeHandlerAlt struct
	r := homeHandlerAlt{name: "Gopher"}
	mux.Handle("/", r)
	mux.Handle("/about", aboutHandlerAlt{})
	mux.Handle("/contact", contactHandlerAlt{})

	// set up the server
	server := http.Server{Addr: ":8080", Handler: mux}
	// well, listen and serve
	server.ListenAndServe()
}
