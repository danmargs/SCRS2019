package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	for index := 0; index < 2000; index++ {
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
		resp2, err2 := http.Get("http://localhost:8080/?where=ws2")
		if err2 != nil {
			// handle error
		}
		defer resp2.Body.Close()

		if resp2.StatusCode == http.StatusOK {
			bodyBytes, _ := ioutil.ReadAll(resp2.Body)
			bodyString := string(bodyBytes)
			fmt.Print(bodyString)
		}
	}

}
