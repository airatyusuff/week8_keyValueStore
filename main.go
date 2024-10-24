package main

import (
	"fmt"
	"time"
)

type Pair struct {
	key   string
	value string
}

type operation struct {
	name     string
	param    Pair
	response chan string
}

type Request struct {
	command Command
}

var keyValueStore = make(map[string]string)
var requests chan Request = make(chan Request)
var done chan any = make(chan any)

func Start() {
	go monitorRequests()
}

func Stop() {
	shutdown := operation{
		name: "stop",
	}
	requests <- Request{
		command: &StopCommand{op: shutdown},
	}
	<-done
}

func Store(key string, value string) {
	newPair := Pair{
		key:   key,
		value: value,
	}

	store := operation{
		name:  "store",
		param: newPair,
	}

	requests <- Request{command: &StoreCommand{op: store}}
}

func Fetch(key string) {
	fetchop := operation{
		name:     "fetch",
		param:    Pair{key: key},
		response: make(chan string),
	}

	requests <- Request{command: &FetchCommand{op: fetchop}}
	// return <-fetchop.response
	fmt.Printf("Value of %s: %s\n", fetchop.param.key, <-fetchop.response)
}

func findValue(key string) string {
	value, exists := keyValueStore[key]
	if exists {
		return value
	}
	return "Key not found"
}

func monitorRequests() {
	for req := range requests {
		req.command.Execute()
	}

	fmt.Println("All requests processed")
	done <- "done" //send anything to done channel so stop() can unblock
}

func bunchOfOps() {
	go Fetch("last")
	go Store("first", "1")
	go Store("second", "2")
	go Store("third", "3")
	go Store("third", "5")
	go Fetch("third")
	go Fetch("notAKey")
	go Store("last", "56")

	time.Sleep(time.Second)
}

func main() {
	fmt.Print("\n----- Concurrency, with Command pattern for processing requests ----\n\n")
	Start()

	defer Stop()

	bunchOfOps()
}
