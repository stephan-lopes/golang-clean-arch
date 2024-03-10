package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	_ "github.com/stephan-lopes/golang-clean-arch/adapter/http/docs"
	"github.com/stephan-lopes/golang-clean-arch/adapter/postgres"
	"github.com/stephan-lopes/golang-clean-arch/di" // swagger embed files
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Keven Lopes
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigration()
	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST")
	router.Handle("/product", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods(http.MethodGet)

	port := viper.GetString("server.port")
	log.Printf("Listen on Port: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)

}
