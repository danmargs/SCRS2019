package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler_ws1(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Ciao benvenuto su WS1\n")

}

func main() {
	http.HandleFunc("/", handler_ws1)
	log.Fatal(http.ListenAndServe(":9111", nil))
}
