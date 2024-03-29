package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/case2912/go-curd-clean-architecture/driver/database"
	"github.com/case2912/go-curd-clean-architecture/registry"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	registry := registry.NewRegistry()
	registry.SetConfig(database.NewSQLConfig(
		os.Getenv("SQL_NAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("SSL_MODE"),
	))
	router := mux.NewRouter()
	router.Handle("/", registry.NewCreateUserHandler())
	http.Handle("/", router)
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, nil)
}
