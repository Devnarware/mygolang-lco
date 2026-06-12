package main

import (
	"log"
	"net/http"
)

type config struct{
	addr string
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)

	log.Printf("Server is staring at the port :4000")
	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}

}
