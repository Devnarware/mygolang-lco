package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct{
	addr string
}
//2. why we are using the struct method when we can save it in a local variable ??

type application struct{
	logger *slog.Logger
}
//4. creating the logger with the help of struct, so that we can use it anywhere in the package


func main(){

	var cnfg config 
	flag.StringVar(&cnfg.addr, "addr", ":4000", "HTTP network address")
	flag.Parse()
	// it is for the congiguation 
	
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	//3. created a logger inside the main fuction but we can't use it outside the main func
	
	app := &application{
		logger: logger, 
	}
	
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)

	
	logger.Info("Server is staring at", "addr", cnfg.addr)
	err := http.ListenAndServe(cnfg.addr, mux)

	if err != nil {
		logger.Error(err.Error())
	}

	os.Exit(1)

}
