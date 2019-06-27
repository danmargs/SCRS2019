package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func squares(d chan int) {
	for i := 0; i <= 9; i++ {
		d <- i * i
	}

	close(d) // close channel
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	fmt.Println("Canale c Aperto")

	d := make(chan int)
	fmt.Println("Canale d Aperto")

	//primo esempio base
	go greet(c)
	c <- "Pippo"

	close(c)

	//secondo esempio loop
	go squares(d) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for val := range d {

		fmt.Println(val)

	}

	_, ok := <-c
	if !ok {
		fmt.Println("Canale c Chiuso")
	}
	_, ok = <-d
	if !ok {
		fmt.Println("Canale d Chiuso")
	}

	fmt.Println("main() stopped")
}
