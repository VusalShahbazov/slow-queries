package main

import (
	"log"

	"github.com/VusalShahbazov/slow-query-logs/internal/bootstrap/show_queries"
)

func main() {

	cnf := show_queries.NewConfig()

	srv := show_queries.NewServer(cnf)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
