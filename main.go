package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", EmptyResponse)
	r.HandleFunc("/game", GameInit)
	http.Handle("/", r)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GameInit(w http.ResponseWriter, req *http.Request) {
	mode = req.URL.Query().Get("mode")
	io.WriteString(w, mode)
}

func EmptyResponse(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Helo World!")
}
