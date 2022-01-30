package main

import (
	"fmt"
	"sync"
	"time"
)

type a struct {
	Val   string
	Index int
}

func test(messages chan a, delay int64, i *int, totalData int) {
	for {

		time.Sleep(time.Duration(delay) * time.Second)
		// process data here
		aVal := a{}
		aVal.Val = "something"
		aVal.Index = *i
		messages <- aVal // return processed data here
		*i++

		if *i == totalData {
			break
		}
	}

	close(messages)
}

func timer(done chan bool, timeout int64) {
	time.Sleep(time.Duration(timeout) * time.Second)
	done <- true
}

func main() {
	startTime := time.Now()
	// use array to mimic storing data
	aList := []a{}
	messages := make(chan a)
	doneSignal := make(chan bool)
	totalMessages := 0
	totalExecution := 0
	var timeout int64 = 14
	var msgDelay int64 = 1
	var saveDelay int64 = 1
	totalData := 5

	go test(messages, msgDelay, &totalMessages, totalData)
	go timer(doneSignal, timeout)
	var wg sync.WaitGroup
outer:
	for {
		var isDone bool
		select {
		case msg, more := <-messages:
			if !more {
				isDone = true
				break
			}
			// save data here
			// goroutine function means if timer expires, it will abruptly cut off the saving process (defer is not called)
			// unless waitgroup is used, then we can wait for all waitgroups to be done before exiting
			wg.Add(1) // used only with goroutine function
			go func() {
				defer func() {
					fmt.Println(msg.Index, " ended. time elapsed:", time.Now().Sub(startTime))
					totalExecution++
					wg.Done() // used only with goroutine function
				}()
				fmt.Println(msg.Index, " started.")
				time.Sleep(time.Duration(saveDelay) * time.Second)
				fmt.Println("received message & saving data", msg)
				aList = append(aList, msg)
			}()

		case done := <-doneSignal:
			isDone = done
		}

		if isDone {
			fmt.Println("break outer")
			// wait for all waitgroups before exiting
			wg.Wait()
			break outer
		}
	}

	timeElapsed := time.Now().Sub(startTime)

	fmt.Println("timeElapsed: ", timeElapsed)
	fmt.Println("aList length:", len(aList))
	fmt.Println("number of messages :", totalMessages)
	fmt.Println("number of execution :", totalExecution)
}
