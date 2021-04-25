package main

import (
	"fmt"
	"goapi/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const PORT string = ":4000"

	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Word")
	})
	router.HandleFunc("/posts", routes.GetPost).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")

	log.Println("Server Listening on port: ", PORT)
	log.Fatalln(http.ListenAndServe(PORT, router))
}
