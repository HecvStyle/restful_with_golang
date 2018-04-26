package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Pressing request!")
	w.Write([]byte("OK"))
	log.Println("Finished processing request")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggerRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggerRouter)
}
