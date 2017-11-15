package machine

import (
	"example/events"
	"fmt"
	"example/internal/machine"
)

type Dispatcher struct {}

func (Dispatcher) ChangedMachineName(message events.MessageChangedName) {}

func (Dispatcher) ChangedAccountLogin(message events.MessageAccountLogin) {
	//testing cycle importing
	machine.Load(1)
	fmt.Println("dispatching changedAccountLogin event")
}


