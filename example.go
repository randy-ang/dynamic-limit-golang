package main

import (
	"fmt"
	"time"
)

type a struct {
	Val string
}

func test(messages chan a, delay int64) {
	for {
		time.Sleep(time.Duration(delay) * time.Second)
		// process data here
		aVal := a{}
		aVal.Val = "something"
		messages <- aVal // return processed data here
	}
}

func timer(done chan bool, timeout int64) {
	time.Sleep(time.Duration(timeout) * time.Second)
	done <- true
}

func main() {
	// use array to mimic storing data
	aList := []a{}
	messages := make(chan a)
	signals := make(chan bool)
	var timeout int64 = 10
	var msgDelay int64 = 3

	go test(messages, msgDelay)
	go timer(signals, timeout)

outer:
	for {
		select {
		case msg := <-messages:
			fmt.Println("received message & saving data", msg)
			// save data here
			aList = append(aList, msg)
		case <-signals:
			break outer
		}
	}

	fmt.Println("aList length:", len(aList))
	fmt.Println("expected length:", timeout/msgDelay)

}
