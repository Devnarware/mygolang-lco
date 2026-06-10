package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{slug}", snippetView)

	log.Print("Staring the server at the port :5500")

	log.Fatal(http.ListenAndServe(":5500", mux))
}


