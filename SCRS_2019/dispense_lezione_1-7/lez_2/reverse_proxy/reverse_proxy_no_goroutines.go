package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	keys, good := r.URL.Query()["where"]
	if !good || len(keys) < 1 {
		//log.Println(r.URL.Query())
		log.Println("Missing 'where' param in URL")
		return
	}
	key := keys[0]
	fmt.Print("pippo.. ", keys)
	log.Println("Utente to ws")
	if key == "ws1" {
		http.Redirect(w, r, "http://localhost:9111/", 307)
	} else if key == "ws2" {
		http.Redirect(w, r, "http://localhost:9112/", 307)
	}

}

func main() {
	fmt.Println("proxy started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
