package account

import (
	"fmt"
	"example/events"
	"example/internal/account"
)

type Dispatcher struct {}

func (Dispatcher) ChangedAccountLogin(message events.MessageAccountLogin) {}

func (Dispatcher) ChangedMachineName(name events.MessageChangedName) {
	//testing cycle importing
	account.Load(1)

	fmt.Println("dispatching changedMachineName")
}

