package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//funzione avviata dalla goroutine che inserisce nel canale la pagina della richiesta get
func service(c chan string, r *http.Request) {

	//stampo alcune informazioni base della request
	fmt.Println(r.Method, r.RequestURI, r.Proto)
	//inserimento nel canale della stringa pagina
	c <- r.RequestURI

}

//funzione handler che avvia la goroutine per gestire la richiesta arrivata al serverweb
func handler(w http.ResponseWriter, r *http.Request) {

	//creazione del canale da passare alla goroutine
	ch := make(chan string)

	//creazione goroutine che prende come parametro il canale di comunicazione e la richiesta che deve gestire la funzione service
	go service(ch, r)

	//inserimento in variabile di ciò che il canale ha ottenuto
	risposta := <-ch

	//scrittura del body della risposta con la pagine che è stata richiesta
	io.WriteString(w, risposta)

}

//funzione principale webserver che avvia l'ascolto sulla porta 8080
func main() {

	fmt.Println("ws started su porta 8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
