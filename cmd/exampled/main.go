package main

import (
	"example/cmd/server"
	"example/events"
	accEvents "example/internal/events/account"
)
//think about some registrar
func init() {
	events.RegisterDispatcher("changedName", accEvents.Dispatcher{})
}

func main() {
	//TCP listen etc.
	s := server.Grpc{}
	s.HandleSomeRequest()
}
