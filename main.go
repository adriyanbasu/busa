package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server running on port 8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
