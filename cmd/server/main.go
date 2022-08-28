package main

import (
	"log"

	"github.com/kyamauchi1023/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
