package main

import (
	"example/cmd/server"
	"example/events"
	accEvents "example/internal/events/account"
	machineEvents "example/internal/events/machine"
)
//think about some registrar
func init() {
	events.RegisterDispatcher("changedMachineName", accEvents.Dispatcher{})
	events.RegisterDispatcher("changedAccountLogin", machineEvents.Dispatcher{})
}

func main() {
	//TCP listen etc.
	s := server.Grpc{}
	s.HandleSomeRequest()
}
