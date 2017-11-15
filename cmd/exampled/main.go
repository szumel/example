package main

import (
	"example/cmd/server"
	"example/events"
	accEvents "example/internal/events/account"
)

func init() {
	events.RegisterDispatcher("changedName", accEvents.Dispatcher{})
}

func main() {
	//TCP listen etc.
	s := server.Grpc{}
	s.HandleSomeRequest()
}
