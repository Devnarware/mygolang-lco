package main

import (
	"net/http"
)

type config struct{
	addr string
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)


}
