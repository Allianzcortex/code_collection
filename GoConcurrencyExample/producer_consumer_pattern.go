// This is the demo of basic producer and consumer explanation
// THe channel can send and receive data

// demo result is:
// send value:  0
// send value:  1
// receive value:  0
// receive value:  1
// send value:  2
// send value:  3
// receive value:  2
// receive value:  3

package main

import "fmt"

func consumer(data chan int, done chan bool) {
	for value := range data {
		fmt.Println("receive value: ", value)
	}

	done <- true
}

func producer(data chan int) {
	for value := 0; value < 4; value++ {
		fmt.Println("send value: ", value)
		data <- value
	}

	close(data)
}

func main() {
	data := make(chan int)               // unbuffered
	var done chan bool = make(chan bool) // another define way

	go consumer(data, done)
	go producer(data)

	<-done
}
