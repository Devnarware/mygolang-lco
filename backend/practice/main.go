package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/book", bookHandler)


	fmt.Println("Server is starting on the port :5050  .....")
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}


func bookHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet{
		newBook := Book{"Atomic habit", "James clear", 369}

		w.Header().Set("Content-Type", "application/json")

		// Ecoding it into the json
		 err := json.NewEncoder(w).Encode(newBook)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
	}else if r.Method == "POST" {
		var book Book 
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, "incorrect JSON data", http.StatusBadRequest)
			return 
		}

		fmt.Printf("The book %s is written by %s\n", book.Name, book.Author)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Book data recieved successfully"))
	}else{
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}