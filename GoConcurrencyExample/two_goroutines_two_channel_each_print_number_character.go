package main

import (
	"fmt"
)

func main() {

	chan_n := make(chan bool)
	chan_c := make(chan bool)
	done := make(chan struct{})

	go func() {
		for i := 1; i < 11; i += 2 {
			<-chan_c
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_n <- true
		}
	}()

	go func() {
		char_seq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-chan_n
			fmt.Print(char_seq[i])
			fmt.Print(char_seq[i+1])
			chan_c <- true
		}
		done <- struct{}{}
	}()

	chan_c <- true

	<-done
}
