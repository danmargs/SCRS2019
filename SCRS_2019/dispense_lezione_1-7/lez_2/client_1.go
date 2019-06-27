package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8080/?where=ws1")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}

	resp1, err1 := http.Get("http://localhost:8080/?why=bho&&where=ws2")
	if err1 != nil {
		// handle error
	}
	defer resp1.Body.Close()

	if resp1.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp1.Body)
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}

}
