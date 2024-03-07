package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/stephan-lopes/golang-clean-arch/adapter/postgres"
	"github.com/stephan-lopes/golang-clean-arch/di"
)

func init() {}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigration()
	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()
	router.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST")
	router.Handle("/product", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")
	log.Printf("Listen on Port: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)

}
