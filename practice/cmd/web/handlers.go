package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func snippetView(w http.ResponseWriter, r *http.Request) {

	val, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
	}

	msg := fmt.Sprintln("Viewing Snippet with id:", val)
	w.Write([]byte(msg))
}
