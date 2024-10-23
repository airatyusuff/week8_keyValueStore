package main

import (
	"fmt"
	"time"
)

type Pair struct {
	key   string
	value string
}

var keyValueStore = make(map[string]string)

type operation struct {
	reqType  string
	param    Pair
	response chan string
}

var requests chan operation = make(chan operation)
var done chan any = make(chan any)

func Start() {
	go monitorRequests()
}

func Stop() {
	shutdown := operation{
		reqType: "stop",
	}
	requests <- shutdown
	<-done
}

func Store(key string, value string) {
	newPair := Pair{
		key:   key,
		value: value,
	}

	requests <- operation{
		reqType:  "store",
		param:    newPair,
		response: nil,
	}
}

func Fetch(key string) {
	fetchop := operation{
		reqType:  "fetch",
		param:    Pair{key: key},
		response: make(chan string),
	}

	requests <- fetchop
	// return <-fetchop.response
	fmt.Println("result of ", fetchop.param.key, ": ", <-fetchop.response)
}

func findValue(key string) string {
	value, exists := keyValueStore[key]
	if exists {
		return value
	}
	return "Key not found"
}

func monitorRequests() {
	for op := range requests {
		switch op.reqType {
		case "store":
			fmt.Printf("Processing: %s request for key %s and value %s\n", op.reqType, op.param.key, op.param.value)
			keyValueStore[op.param.key] = op.param.value

		case "fetch":
			fmt.Println("Processing: ", op.reqType, " request for key ", op.param.key)
			op.response <- findValue(op.param.key)

		case "stop":
			fmt.Println("Shutting down")
			close(requests)
		}
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
	fmt.Println("Helllllllo")
	Start()

	defer Stop()

	bunchOfOps()
}
