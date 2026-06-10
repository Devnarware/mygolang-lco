package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("../../ui/html/pages/home.html")

	if err != nil {
		log.Fatal(err)
	}

	err = ts.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}
}	



func snippetView(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(r.PathValue("slug"))

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	msg := fmt.Sprintln("Viewing Snippet with id: ", id)
	w.Write([]byte(msg))
}
