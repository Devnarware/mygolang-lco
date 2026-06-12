package main

import (
	"fmt"
	"net/http"
	"strconv"
)


func (app *application)home(w http.ResponseWriter, r *http.Request){
	app.logger.Info("Home page visited")
	w.Write([]byte ("hello"))
}



func (app *application)about(w http.ResponseWriter, r *http.Request){
	app.logger.Info("About page visited")
	w.Write([]byte ("About Page"))
}



func (app *application)snippetCreate(w http.ResponseWriter, r *http.Request){
	app.logger.Info("Snippet Created")
	w.Write([]byte ("Snippet Created"))
}



func (app *application)snippetView(w http.ResponseWriter, r *http.Request){

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	msg := fmt.Sprintf("Viewing snippet %d", id)

	app.logger.Info("Viewing snippet", "id", id)
	w.Write([]byte (msg))
}