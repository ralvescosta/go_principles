package main

import (
	"crud/src/frameworks/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	usecases "crud/src/applications/books"
	entities "crud/src/business/entities"
	adapters "crud/src/frameworks/adapters"
	controllers "crud/src/interfaces/controllers"
	middleware "crud/src/interfaces/middleware"
)

const (
	HTTP_PORT = 4444
	HTTP_HOST = "127.0.0.1"
)

// Server ...
type Server struct {
	dbConn *sql.DB
	router *mux.Router
}

//StartHTPServer  ...
func (s *Server) StartHTPServer() {
	_dbConnection := database.DbConnection()
	dbConn, err := _dbConnection.Connect()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	_hMiddleware := middleware.HeadersMiddleware()

	router := mux.NewRouter()
	router.Use(_hMiddleware.Handle)

	s.dbConn = dbConn
	s.router = router

	s.registerRouters()

	server := http.Server{
		Handler:      handlers.CompressHandler(router),
		Addr:         fmt.Sprintf("%s:%d", HTTP_HOST, HTTP_PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server Running on: ", HTTP_HOST, ":", HTTP_PORT)
	log.Fatalln(server.ListenAndServe())
}

func (s *Server) registerRouters() {

	_crudBookUsecase := usecases.Books()
	_createABookController := controllers.CreateABookController(_crudBookUsecase)
	s.router.HandleFunc("/books", adapters.RouteAdapt(_createABookController, &entities.InputCreateBook{})).Methods("POST")
}
