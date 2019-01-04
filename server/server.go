package main

import (
	"flag"
	"log"
	"net/http"
)

/*
func addMIMEToWasmFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.FileServer(http.Dir("./public/test.wasm"))
}
*/

func main() {
	port := flag.String("port", "8000", "port")
	flag.Parse()
	log.Printf("Start server on port %s\n", *port)
	/*
		mux := http.NewServeMux()
		mux.Handle("/", http.FileServer(http.Dir("./public")))

		mux.HandleFunc("/test.wasm", addMIMEToWasmFiles)
		log.Fatal(http.ListenAndServe(":8000", mux))
	*/
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("./public"))))
}
