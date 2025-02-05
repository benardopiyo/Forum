package main

import (
	"flag"
	"fmt"
	"fora/routes"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "Http server port")

	mux := routes.Router()
	server := http.Server{
		Addr: *addr,
		Handler: mux,
	}

	fmt.Printf("Server running on http://localhost%s", *addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
