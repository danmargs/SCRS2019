package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//funzione principale del client che avvia la richiesta get al serverweb
func main() {
	var NUM = 1000 //numero richieste

	//inizializza una chiamata bloccante per il main, che deve aspettare che tutte le goroutine finiscano
	var wg sync.WaitGroup
	wg.Add(NUM)

	for i := 0; i < NUM; i++ {
		go richiesta(i, &wg)
	}
	wg.Wait()
	time.Sleep(50 * time.Millisecond)

}

func richiesta(i int, wg *sync.WaitGroup) {
	//richiesta pagina.html al serverweb tramite get su porta 8080
	resp, err := http.Get("http://localhost:8080/pagina.html")
	if err != nil {
		fmt.Println("\n\n errore")
	}
	defer resp.Body.Close()

	//se risposta affermativa allora stampa la risposta ottenuta /pagina.html
	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	fmt.Println("routine num : ", i)

	//la goroutine ha terminato, entra in attesa
	wg.Done()
}
