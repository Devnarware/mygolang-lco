package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)


	log.Print("server is starting")

	err := http.ListenAndServe(":5500", mux)

	if err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello"))
}


