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

	ch1 := make(chan string)
	ch2 := make(chan string)

	go service1(ch1, r)
	go service2(ch2, r)

	select {
	case request := <-ch1:
		fmt.Println("richiesta ws1", request)
		http.Redirect(w, r, "http://localhost:9111/", 307)
	case request := <-ch2:
		fmt.Println("richiesta ws2", request)
		http.Redirect(w, r, "http://localhost:9112/", 307)
	}

}

func main() {
	fmt.Println("proxy started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
