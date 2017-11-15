package account

import (
	"fmt"
	"example/events"
	"example/internal/account"
)

type Dispatcher struct {}

func (Dispatcher) ChangedName(name events.MessageChangedName) {
	account.Load(1)
	fmt.Println("dispatching changedname from account")
}

