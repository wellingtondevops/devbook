package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := router.Gerar()
	fmt.Println("API listen in port 5003")
	log.Fatal(http.ListenAndServe(":5003", r))

}
