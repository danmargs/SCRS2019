package main

import (
	"fmt"
	"log"
	"net/http"
)

func service1(c chan string, r *http.Request) {
	keys, good := r.URL.Query()["where"]
	if !good || len(keys) < 1 {
		//log.Println(r.URL.Query())
		log.Println("Missing 'where' param in URL")
		return
	}
	key := keys[0]
	if key == "ws1" {
		c <- key
	}
}

func service2(c chan string, r *http.Request) {
	keys, good := r.URL.Query()["where"]
	if !good || len(keys) < 1 {
		//log.Println(r.URL.Query())
		log.Println("Missing 'where' param in URL")
		return
	}
	key := keys[0]

	if key == "ws2" {
		c <- key
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	ch := make(chan string)

	go service(ch, r)

}

func main() {
	fmt.Println("proxy started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
