package main

import (
	"log"

	"github.com/altinthaqi/jot-fusion/api"
	"github.com/altinthaqi/jot-fusion/db"
)

func main() {
	store, err := db.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
