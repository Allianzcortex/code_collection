// This is an advanced producer and consumer explanation
// You use WaitGroup to control the grountine
// Certaily you can also use channel,it still workds(D

// demo result is:
// send value:  0
// send value:  4
// receive value:  0
// temporary result is:  0
// receive value:  4
// temporary result is:  4
// send value:  5
// send value:  1
// receive value:  5
// temporary result is:  9
// receive value:  1
// temporary result is:  10
// send value:  2
// send value:  6
// receive value:  2
// temporary result is:  12
// receive value:  6
// temporary result is:  18
// send value:  7
// receive value:  7
// temporary result is:  25
// send value:  3
// receive value:  3
// temporary result is:  28
// consumer is done
// 28

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var sum = 7
var wg sync.WaitGroup

func consumer(data chan int, result *int) {
	// here ,you cannot use defer wg.done()
	// or the result will be :fatal error: all goroutines are asleep - deadlock!
	// it is simply because the value will always wait for the value sent from data(channel)
	// you must specific the acact stop position
	var i int64 = 0
	for value := range data {
		fmt.Println("receive value: ", value)
		atomic.AddInt64(&i, 1)
		*result += value
		fmt.Println("temporary result is: ", *result)
		if int(atomic.LoadInt64(&i)) == sum+1 {
			fmt.Println("consumer is done")
			wg.Done()
			// you can also use channel to control,just uncoment the next line
			// and add (done chan bool) to this function's paramater
			// done<-true
		}

	}

}

func producer(data chan int, start int) {
	defer wg.Done()
	for value := start; value < start+4; value++ {
		time.Sleep(1)
		fmt.Println("send value: ", value)
		data <- value
	}

}

func main() {
	data := make(chan int) // unbuffered
	// uncomment next line if you want to use channel
	// var done chan bool = make(chan bool) // another define way
	x := 0
	wg.Add(3)
	go consumer(data, &x)
	go producer(data, 0)
	go producer(data, 4)
	wg.Wait()
	// channel still works,uncomment next line and you can also get the correct conclusion
	// <-done

	fmt.Println(x)
}
