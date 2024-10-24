package main

import "fmt"

type Command interface {
	Execute()
}

type FetchCommand struct {
	op operation
}

type StoreCommand struct {
	op operation
}

type StopCommand struct {
	op operation
}

func (c *FetchCommand) Execute() {
	c.op.fetch()
}

func (c *StoreCommand) Execute() {
	c.op.store()
}

func (c *StopCommand) Execute() {
	c.op.stop()
}

func (op *operation) fetch() {
	fmt.Println("Processing: ", op.name, " request for key: ", op.param.key)
	op.response <- findValue(op.param.key)
}

func (op *operation) store() {
	fmt.Printf("Processing: %s request for key: %s and value: %s\n", op.name, op.param.key, op.param.value)
	keyValueStore[op.param.key] = op.param.value
}

func (op *operation) stop() {
	fmt.Println("Shutting down")
	close(requests)
}
