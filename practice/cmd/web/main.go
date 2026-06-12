package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct{
	addr string
}

func main(){

	var cnfg config 
	flag.StringVar(&cnfg.addr, "addr", ":4000", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)

	log.Printf("Server is staring at the port %s", cnfg.addr)
	err := http.ListenAndServe(cnfg.addr, mux)

	if err != nil {
		log.Fatal(err)
	}

}
