package main

import "net/http"

func (app *application)routes() *http.ServeMux{

	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	// these are the routes we are gonna use 

	return mux
}