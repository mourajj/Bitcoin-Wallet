package main

import (
	"fmt"
	"log"
	"maxxer/src/config"
	"maxxer/src/router"
	"net/http"
)

func main() {
	config.Load()
	r := router.Gerar()

	fmt.Printf("Listening on Port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
