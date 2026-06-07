package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)

	log.Print("Staring the server at the port :5500")

	log.Fatal(http.ListenAndServe(":5500", mux))
}


