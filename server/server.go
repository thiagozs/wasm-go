package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8000", "port")
	flag.Parse()
	log.Printf("Start server on port %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("./public"))))
}
