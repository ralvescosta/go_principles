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
	repositories "crud/src/frameworks/repositories"
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
	/*
	* Database Connection
	 */
	_Database := database.DbConnection()
	dbConn, err := _Database.Connect()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	s.dbConn = dbConn

	/*
	* Mux Router Handler
	 */
	router := mux.NewRouter()

	_cors := middleware.Cors()
	router.Use(adapters.MiddlewareAdapt(_cors))

	s.router = router

	/*
	* Routers
	 */
	s.registerRouters()

	/*
	* HTTP Server Configuration
	 */
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

	_crudBooksRepository := repositories.BooksRepository(s.dbConn)
	_crudBookUsecase := usecases.Books(_crudBooksRepository)

	_createABookController := controllers.CreateABookController(_crudBookUsecase)
	s.router.HandleFunc("/books", adapters.RouteAdapt(_createABookController, &entities.InputCreateBook{})).Methods("POST")

	_indexBookController := controllers.IndexBookController(_crudBookUsecase)
	s.router.HandleFunc("/books/{id}", adapters.RouteAdapt(_indexBookController, nil)).Methods("GET")

	_showBookController := controllers.ShowBookController(_crudBookUsecase)
	s.router.HandleFunc("/books", adapters.RouteAdapt(_showBookController, nil)).Methods("GET")

	_updateBookController := controllers.UpdateBookController(_crudBookUsecase)
	s.router.HandleFunc("/books/{id}", adapters.RouteAdapt(_updateBookController, &entities.InputCreateBook{})).Methods("PUT")

	_deleteBookController := controllers.DeleteABookController(_crudBookUsecase)
	s.router.HandleFunc("/books/{id}", adapters.RouteAdapt(_deleteBookController, nil)).Methods("DELETE")
}
