package main

import (
	_ "embed"
	"log"

	"github.com/Omar-Arabi1/GoShorten/handlers"
)

const addr = ":8080"
const dbName = "urls.db"

//go:embed schema.sql
var schema string

func main() {
	router, err := handlers.New(dbName, schema)
	if err != nil {
		// the reason we are aware that the error must be a db error is because that is the only possible error
		log.Fatalf("Encountered an error while setting up the db connection: %s\n", err)
	}
	router.Run(addr)
}
