package main

import (
	"log"
	"net/http"

	"usermanagement/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	routes.UserManagementRoutes(r)
	http.Handle("/", r)
	log.Println("starting")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
	log.Println("started")

}
