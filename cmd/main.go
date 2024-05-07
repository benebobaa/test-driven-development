package main

import (
	"fmt"
	"net/http"
	"test-driven-development/pkg/domain"
	"test-driven-development/pkg/sqlstore"
)

// Hardcoded database connection
const postgresDNS = "postgres://root:root@localhost:5432/tdd?sslmode=disable"

func main() {

	db := sqlstore.NewDB(postgresDNS)
	service := domain.NewService(db)
	router := NewRouter(service)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("listening to port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
