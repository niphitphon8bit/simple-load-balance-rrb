package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var serverList string
	var mode string
	var port int
	var name string

	flag.StringVar(&serverList, "backends", "", "Load balanced backends, use commas to separate")
	flag.StringVar(&mode, "mode", "backend", "set mode to load balanced (default backend)")
	flag.StringVar(&name, "name", "", "set name")
	flag.IntVar(&port, "port", 3030, "Port to serve")
	flag.Parse()

	if mode == "backend" {

		http.HandleFunc("/", handler)
		http.HandleFunc("/health", healthCheckHandler)

		addr := fmt.Sprintf(":%d", port) // Use %d for decimal integers
		fmt.Printf("Server starting on %s\n", addr)

		if err := http.ListenAndServe(addr, nil); err != nil {
			panic(err)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
}
