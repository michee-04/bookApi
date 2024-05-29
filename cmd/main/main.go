package main

import (
	"fmt"

	"github.com/michee-04/dbtest/pkg/routes"
)

const port = ":5000"


func main(){
	// r := mux.NewRouter()
	// http.Handle("/", r)
	// routes.RegisterBookStoreRoutes(r)
	// 	log.Fatal(http.ListenAndServe("localhost:9090", r))


	fmt.Printf("le serveur fonctionne sur http://localhost%s",port)

	routes.RegisterBookStoreRoutes()
}